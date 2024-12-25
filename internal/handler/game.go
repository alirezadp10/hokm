package handler

import (
    "context"
    "errors"
    "fmt"
    "github.com/alirezadp10/hokm/internal/database/redis"
    "github.com/alirezadp10/hokm/internal/database/sqlite"
    "github.com/alirezadp10/hokm/internal/service/cards_service"
    "github.com/alirezadp10/hokm/internal/service/game_service"
    "github.com/alirezadp10/hokm/internal/service/players_service"
    "github.com/alirezadp10/hokm/internal/service/points_service"
    "github.com/alirezadp10/hokm/internal/utils/my_bool"
    "github.com/alirezadp10/hokm/internal/utils/my_slice"
    "github.com/alirezadp10/hokm/internal/utils/trans"
    "github.com/google/uuid"
    "github.com/labstack/echo/v4"
    "github.com/redis/rueidis"
    "net/http"
    "strconv"
    "strings"
    "time"
)

func (h *Handler) GetGameId(c echo.Context) error {
    username := c.Get("username").(string)

    var gameId string

    if gid, ok := sqlite.DoesPlayerHaveAnActiveGame(h.sqlite, username); ok {
        return c.JSON(http.StatusForbidden, map[string]interface{}{
            "message": trans.Get("You have already an active game."),
            "gameId":  *gid,
        })
    }

    gameId = uuid.New().String()
    go game_service.Matchmaking(c.Request().Context(), h.redis, username, gameId)

    err := redis.Subscribe(c.Request().Context(), h.redis, "game_creation", func(msg rueidis.PubSubMessage) {
        message := strings.Split(msg.Message, "|")
        players := strings.Split(message[0], ",")
        if my_slice.Has(players, username) {
            gameId = message[1]
            _, err := sqlite.AddPlayerToGame(h.sqlite, username, gameId)
            if err != nil {
                fmt.Println(err)
            }
            redis.Unsubscribe(c.Request().Context(), h.redis, "game_creation")
        }
    })

    if err != nil {
        if errors.Is(err, context.Canceled) {
            redis.RemovePlayerList(context.Background(), h.redis, "matchmaking", username)
        }
        if errors.Is(c.Request().Context().Err(), context.DeadlineExceeded) {
            return c.JSON(http.StatusRequestTimeout, map[string]interface{}{
                "message": trans.Get("No body have found. please try again later."),
                "gameId":  nil,
            })
        }
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "message": trans.Get("Something went wrong. please try again later."),
            "gameId":  nil,
        })
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": trans.Get("Game has been made."),
        "gameId":  gameId,
    })
}

func (h *Handler) GetGameData(c echo.Context) error {
    username := c.Get("username").(string)

    gameId := c.Param("gameId")

    if sqlite.HasGameFinished(h.sqlite, gameId) {
        return c.JSON(http.StatusForbidden, map[string]interface{}{
            "message": trans.Get("The game has already finished."),
        })
    }

    if !sqlite.DoesPlayerBelongsToThisGame(h.sqlite, username, gameId) {
        return c.JSON(http.StatusForbidden, map[string]interface{}{
            "message": trans.Get("It's not your game."),
        })
    }

    gameInformation := redis.GetGameInformation(c.Request().Context(), h.redis, gameId)

    players := strings.Split(gameInformation["players"].(string), ",")

    uIndex := my_slice.GetIndex(username, players)

    response := map[string]interface{}{
        "beginnerDirection":    players_service.GetDirection(my_slice.GetIndex(players[0], players), uIndex),
        "players":              players_service.GetPlayersWithDirections(players, uIndex),
        "points":               points_service.GetPoints(gameInformation["points"].(string), uIndex),
        "centerCards":          cards_service.GetCenterCards(gameInformation["center_cards"].(string), uIndex),
        "turn":                 players_service.GetTurn(gameInformation["turn"].(string), uIndex),
        "king":                 cards_service.GetKing(gameInformation["king"].(string), uIndex),
        "kingCards":            cards_service.GetKingCards(gameInformation["king_cards"].(string)),
        "timeRemained":         players_service.GetTimeRemained(gameInformation["last_move_timestamp"].(string)),
        "hasKingCardsFinished": gameInformation["has_king_cards_finished"].(string),
        "trump":                gameInformation["trump"],
    }

    if response["hasKingCardsFinished"] == "true" {
        response["playerCards"] = cards_service.GetPlayerCards(gameInformation["cards"].(string), uIndex)
    } else if response["king"] == "down" {
        response["trumpCards"] = cards_service.GetPlayerCards(gameInformation["cards"].(string), uIndex)[0]
    }

    return c.JSON(http.StatusOK, response)
}

func (h *Handler) ChooseTrump(c echo.Context) error {
    username := c.Get("username").(string)
    gameId := c.Param("gameId")

    var requestBody struct {
        Trump string `json:"trump"`
    }

    if err := c.Bind(&requestBody); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": trans.Get("Invalid JSON.")})
    }

    if !my_slice.Has([]string{"H", "D", "C", "S"}, requestBody.Trump) {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": trans.Get("Invalid trump.")})
    }

    if !sqlite.DoesPlayerBelongsToThisGame(h.sqlite, username, gameId) {
        return c.JSON(http.StatusForbidden, map[string]interface{}{
            "message": trans.Get("It's not your game."),
        })
    }

    gameInformation := redis.GetGameInformation(c.Request().Context(), h.redis, gameId)

    players := strings.Split(gameInformation["players"].(string), ",")

    uIndex := my_slice.GetIndex(username, players)

    if gameInformation["king"].(string) != strconv.Itoa(uIndex) {
        return c.JSON(http.StatusForbidden, map[string]interface{}{
            "message": trans.Get("You're not king in this round."),
        })
    }

    if gameInformation["has_king_cards_finished"].(string) == "true" {
        return c.JSON(http.StatusForbidden, map[string]interface{}{
            "message": trans.Get("You're not allowed to choose a trump at the moment."),
        })
    }

    lastMoveTimestamp := strconv.FormatInt(time.Now().Unix(), 10)

    err := redis.SetTrump(c.Request().Context(), h.redis, gameId, requestBody.Trump, strconv.Itoa(uIndex), lastMoveTimestamp)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "message": trans.Get("Something went wrong, Please try again later."),
        })
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "trump":        requestBody.Trump,
        "cards":        hokm.GetPlayerCards(gameInformation["cards"].(string), uIndex)[1:],
        "timeRemained": hokm.GetTimeRemained(gameInformation["last_move_timestamp"].(string)),
    })
}

func (h *Handler) GetYourCards(c echo.Context) error {
    username := c.Get("username").(string)
    gameId := c.Param("gameId")

    if !sqlite.DoesPlayerBelongsToThisGame(h.sqlite, username, gameId) {
        return c.JSON(http.StatusForbidden, map[string]interface{}{
            "message": trans.Get("It's not your game."),
        })
    }

    gameInformation := redis.GetGameInformation(c.Request().Context(), h.redis, gameId)

    trump := gameInformation["trump"].(string)

    if gameInformation["trump"].(string) == "" {
        err := redis.Subscribe(c.Request().Context(), h.redis, "choosing_trump", func(msg rueidis.PubSubMessage) {
            messages := strings.Split(msg.Message, ",")
            messageId := my_slice.HasLike(messages, func(s string) bool {
                return strings.Contains(s, gameId+"|")
            })
            if messageId != -1 {
                data := strings.Split(messages[messageId], "|")
                trump = data[1]
                redis.Unsubscribe(c.Request().Context(), h.redis, "choosing_trump")
            }
        })

        if err != nil {
            if errors.Is(c.Request().Context().Err(), context.DeadlineExceeded) {
                return c.JSON(http.StatusRequestTimeout, map[string]interface{}{
                    "message": trans.Get("Something went wrong, Please try again later."),
                })
            }
            return c.JSON(http.StatusInternalServerError, map[string]interface{}{
                "message": trans.Get("Something went wrong, Please try again later."),
            })
        }
    }

    players := strings.Split(gameInformation["players"].(string), ",")

    uIndex := my_slice.GetIndex(username, players)

    return c.JSON(http.StatusOK, map[string]interface{}{
        "cards": hokm.GetPlayerCards(gameInformation["cards"].(string), uIndex),
        "trump": trump,
    })
}

func (h *Handler) PlaceCard(c echo.Context) error {
    username := c.Get("username").(string)
    gameId := c.Param("gameId")

    var requestBody struct {
        Card string `json:"card"`
    }

    if err := c.Bind(&requestBody); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": trans.Get("Invalid JSON.")})
    }

    if !my_slice.Has(hokm.Cards, requestBody.Card) {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": trans.Get("Invalid Card.")})
    }

    if !sqlite.DoesPlayerBelongsToThisGame(h.sqlite, username, gameId) {
        return c.JSON(http.StatusForbidden, map[string]interface{}{
            "message": trans.Get("It's not your game."),
        })
    }

    gameInformation := redis.GetGameInformation(c.Request().Context(), h.redis, gameId)

    players := strings.Split(gameInformation["players"].(string), ",")

    uIndex := my_slice.GetIndex(username, players)

    leadSuit := gameInformation["lead_suit"].(string)

    if gameInformation["turn"].(string) != strconv.Itoa(uIndex) {
        return c.JSON(http.StatusForbidden, map[string]interface{}{
            "message": trans.Get("It's not your turn."),
        })
    }

    isSelectedCardForUser := false

    doesPlayerHaveLeadSuitCard := false

    for _, cards := range hokm.GetPlayerCards(gameInformation["cards"].(string), uIndex) {
        for _, card := range cards {
            if card == requestBody.Card {
                isSelectedCardForUser = true
            }
            if hokm.GetCardSuit(card) == leadSuit {
                doesPlayerHaveLeadSuitCard = true
            }
        }
    }

    if leadSuit != "" && doesPlayerHaveLeadSuitCard && hokm.GetCardSuit(requestBody.Card) != leadSuit {
        return c.JSON(http.StatusForbidden, map[string]interface{}{
            "message": trans.Get("You're not allowed to select this card."),
        })
    }

    if !isSelectedCardForUser {
        return c.JSON(http.StatusForbidden, map[string]interface{}{
            "message": trans.Get("You're not allowed to select this card."),
        })
    }

    if gameInformation["trump"].(string) == "" {
        return c.JSON(http.StatusForbidden, map[string]interface{}{
            "message": trans.Get("It's not your turn."),
        })
    }

    if gameInformation["lead_suit"].(string) == "" {
        leadSuit = hokm.GetCardSuit(requestBody.Card)
    }

    centerCards := hokm.UpdateCenterCards(gameInformation["center_cards"].(string), requestBody.Card, uIndex)

    cardsWinner := hokm.FindCardsWinner(
        centerCards, gameInformation["trump"].(string), gameInformation["lead_suit"].(string),
    )

    gameWinner := ""
    roundWinner := ""
    wasKingChanged := false
    isItNewRound := "false"
    king := gameInformation["king"].(string)
    points := gameInformation["points"].(string)

    turn := hokm.GetNewTurn(gameInformation["turn"].(string))

    cards := hokm.UpdateUserCards(gameInformation["cards"].(string), requestBody.Card, uIndex)

    if cardsWinner != "" {
        points, roundWinner, gameWinner = hokm.UpdatePoints(points, cardsWinner)
        if roundWinner == "" {
            turn = cardsWinner
        } else {
            turn = ""
            cards = hokm.DistributeCards()
            isItNewRound = "true"
        }
        king = hokm.GiveKing(roundWinner, king)
        wasKingChanged = king == gameInformation["king"].(string)

        leadSuit = ""
        centerCards = ",,,"
    }

    lastMoveTimestamp := strconv.FormatInt(time.Now().Unix(), 10)

    trump := gameInformation["trump"].(string)

    if roundWinner != "" && wasKingChanged {
        trump = ""
    }

    params := redis.PlaceCardParams{
        GameId:            gameId,
        Card:              requestBody.Card,
        CenterCards:       centerCards,
        LeadSuit:          leadSuit,
        CardsWinner:       cardsWinner,
        Points:            points,
        Turn:              turn,
        King:              king,
        WasKingChanged:    my_bool.ToString(wasKingChanged),
        LastMoveTimestamp: lastMoveTimestamp,
        Trump:             trump,
        IsItNewRound:      isItNewRound,
        Cards:             cards,
        PlayerIndex:       uIndex,
    }

    // Call PlaceCard with the params struct
    err := redis.PlaceCard(
        c.Request().Context(),
        h.redis,
        params,
    )

    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "message": trans.Get("Something went wrong, Please try again later."),
        })
    }

    response := map[string]interface{}{
        "players":           hokm.GetPlayersWithDirections(players, uIndex),
        "points":            hokm.GetPoints(points, uIndex),
        "centerCards":       hokm.GetCenterCards(centerCards, uIndex),
        "turn":              hokm.GetTurn(turn, uIndex),
        "king":              hokm.GetKing(king, uIndex),
        "timeRemained":      hokm.GetTimeRemained(lastMoveTimestamp),
        "playerCards":       hokm.GetPlayerCards(gameInformation["cards"].(string), uIndex),
        "wasKingChanged":    wasKingChanged,
        "trump":             gameInformation["trump"],
        "whoHasWonTheCards": "",
        "whoHasWonTheRound": "",
        "whoHasWonTheGame":  "",
    }

    if cardsWinner != "" {
        cardsWinner, _ := strconv.Atoi(cardsWinner)
        response["whoHasWonTheCards"] = hokm.GetDirection(cardsWinner, uIndex)
    }

    if roundWinner != "" {
        roundWinner, _ := strconv.Atoi(roundWinner)
        response["whoHasWonTheRound"] = hokm.GetDirection(roundWinner, uIndex)
    }

    if gameWinner != "" {
        gameWinner, _ := strconv.Atoi(gameWinner)
        response["whoHasWonTheGame"] = hokm.GetDirection(gameWinner, uIndex)
    }

    return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetUpdate(c echo.Context) error {
    username := c.Get("username").(string)

    var player int
    var card string
    gameId := c.Param("gameId")

    if !sqlite.DoesPlayerBelongsToThisGame(h.sqlite, username, gameId) {
        return c.JSON(http.StatusForbidden, map[string]interface{}{
            "message": trans.Get("It's not your game."),
        })
    }

    gameInformation := redis.GetGameInformation(c.Request().Context(), h.redis, gameId)

    players := strings.Split(gameInformation["players"].(string), ",")

    uIndex := my_slice.GetIndex(username, players)

    err := redis.Subscribe(c.Request().Context(), h.redis, "placing_card", func(msg rueidis.PubSubMessage) {
        messages := strings.Split(msg.Message, "|")
        if messages[0] == gameId {
            gameInformation = redis.GetGameInformation(c.Request().Context(), h.redis, messages[0])
            player, _ = strconv.Atoi(messages[1])
            card = messages[2]
            redis.Unsubscribe(c.Request().Context(), h.redis, "placing_card")
        }
    })
    if err != nil {
        if errors.Is(c.Request().Context().Err(), context.DeadlineExceeded) {
            return c.JSON(http.StatusRequestTimeout, map[string]interface{}{
                "message": trans.Get("Something went wrong, Please try again later."),
            })
        }
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "message": trans.Get("Something went wrong, Please try again later."),
        })
    }

    response := map[string]interface{}{
        "lastMove": map[string]string{
            "from": hokm.GetDirection(player, uIndex),
            "card": card,
        },
        "points":            hokm.GetPoints(gameInformation["points"].(string), uIndex),
        "centerCards":       hokm.GetCenterCards(gameInformation["center_cards"].(string), uIndex),
        "turn":              hokm.GetTurn(gameInformation["turn"].(string), uIndex),
        "king":              hokm.GetKing(gameInformation["king"].(string), uIndex),
        "timeRemained":      hokm.GetTimeRemained(gameInformation["last_move_timestamp"].(string)),
        "wasKingChanged":    gameInformation["was_king_changed"].(string),
        "trump":             gameInformation["trump"],
        "whoHasWonTheCards": "",
        "whoHasWonTheRound": "",
        "whoHasWonTheGame":  "",
    }

    cardsWinner, _ := gameInformation["who_has_won_the_cards"].(string)
    roundWinner, _ := gameInformation["who_has_won_the_round"].(string)
    gameWinner, _ := gameInformation["who_has_won_the_game"].(string)

    if cardsWinner != "" {
        cardsWinner, _ := strconv.Atoi(cardsWinner)
        response["whoHasWonTheCards"] = hokm.GetDirection(cardsWinner, uIndex)
    }

    if roundWinner != "" {
        roundWinner, _ := strconv.Atoi(roundWinner)
        response["whoHasWonTheRound"] = hokm.GetDirection(roundWinner, uIndex)
    }

    if gameWinner != "" {
        gameWinner, _ := strconv.Atoi(gameWinner)
        response["whoHasWonTheGame"] = hokm.GetDirection(gameWinner, uIndex)
    }

    return c.JSON(http.StatusOK, response)
}
