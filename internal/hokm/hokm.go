package hokm

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/alirezadp10/hokm/internal/database/redis"
    "github.com/redis/rueidis"
    "math/rand"
    "strconv"
    "strings"
    "time"
)

// Global cards list representing all cards in a deck
var cards = []string{
    "01C", "02C", "03C", "04C", "05C", "06C", "07C", "08C", "09C", "10C", "JC", "QC", "KC", // Clubs
    "01D", "02D", "03D", "04D", "05D", "06D", "07D", "08D", "09D", "10D", "JD", "QD", "KD", // Diamonds
    "01H", "02H", "03H", "04H", "05H", "06H", "07H", "08H", "09H", "10H", "JH", "QH", "KH", // Hearts
    "01S", "02S", "03S", "04S", "05S", "06S", "07S", "08S", "09S", "10S", "JS", "QS", "KS", // Spades
}

// chooseFirstKing selects the first player with a king card and returns the cards sequence and the player's index
func chooseFirstKing() (string, string) {
    localCards := make([]string, len(cards)) // Create a copy of the cards to shuffle
    copy(localCards, cards)

    rand.Shuffle(len(localCards), func(i, j int) { localCards[i], localCards[j] = localCards[j], localCards[i] }) // Shuffle cards

    var cardsList []string // To track the sequence of dealt cards
    i := 0

    // Deal cards until a king is found
    for {
        card := localCards[0]
        localCards = localCards[1:] // Remove dealt card from the deck
        cardsList = append(cardsList, card)

        if card[:2] == "01" { // Check if the card is "01" (e.g., Ace of any suit)
            return strings.Join(cardsList, ","), strconv.Itoa(i % 4) // Return the sequence and player index
        }

        i++
    }
}

// GetKingCards splits a comma-separated king cards string into a slice
func GetKingCards(kingCardsString string) []string {
    return strings.Split(kingCardsString, ",")
}

// GetKing returns the direction of the king player relative to the current user
func GetKing(kingIndex string, uIndex int) string {
    kingI, _ := strconv.Atoi(kingIndex) // Convert king index to integer
    return GetDirection(kingI, uIndex)
}

// GetTurn returns the direction of the player whose turn it is
func GetTurn(turnIndex string, uIndex int) string {
    if turnIndex == "" { // Handle empty turn index
        return ""
    }
    turnI, _ := strconv.Atoi(turnIndex) // Convert turn index to integer
    return GetDirection(turnI, uIndex)
}

// GetTimeRemained calculates the time remaining for the player's turn
func GetTimeRemained(lastMoveTimestampStr string) time.Duration {
    lastMoveTimestampInt, err := strconv.ParseInt(lastMoveTimestampStr, 10, 64) // Convert string to integer
    if err != nil {
        fmt.Println("Error parsing timestamp:", err)
        return 15 * time.Second // Default time if parsing fails
    }

    lastMoveTimestamp := time.Unix(lastMoveTimestampInt, 0) // Convert to time.Time
    timeElapsed := time.Since(lastMoveTimestamp)            // Time elapsed since the move
    timeRemaining := 15*time.Second - timeElapsed           // Calculate remaining time

    if timeRemaining < 0 { // Ensure non-negative duration
        timeRemaining = 0
    }

    return timeRemaining.Round(time.Second) // Round to the nearest second
}

// GetPlayersWithDirections maps usernames to relative directions
func GetPlayersWithDirections(players []string, uIndex int) map[string]interface{} {
    return map[string]interface{}{
        "left":  map[string]string{"username": players[(4+(uIndex-1))%4]},
        "down":  map[string]string{"username": players[(4+(uIndex+0))%4]},
        "right": map[string]string{"username": players[(4+(uIndex+1))%4]},
        "up":    map[string]string{"username": players[(4+(uIndex+2))%4]},
    }
}

// GetPoints processes the points string and organizes them by player direction
func GetPoints(pointsString string, uIndex int) map[string]map[string]string {
    points := make(map[string]string)
    err := json.Unmarshal([]byte(pointsString), &points) // Parse JSON string
    if err != nil {
        fmt.Println("Error unmarshalling:", err)
    }

    var downTotalPoints, rightTotalPoints, downRoundPoints, rightRoundPoints string
    if uIndex%2 == 0 { // Determine point order based on user index
        downTotalPoints = strings.Split(points["total"], ",")[0]
        downRoundPoints = strings.Split(points["round"], ",")[0]
        rightTotalPoints = strings.Split(points["total"], ",")[1]
        rightRoundPoints = strings.Split(points["round"], ",")[1]
    } else {
        downTotalPoints = strings.Split(points["total"], ",")[1]
        downRoundPoints = strings.Split(points["round"], ",")[1]
        rightTotalPoints = strings.Split(points["total"], ",")[0]
        rightRoundPoints = strings.Split(points["round"], ",")[0]
    }

    return map[string]map[string]string{
        "total":        {"down": downTotalPoints, "right": rightTotalPoints},
        "currentRound": {"down": downRoundPoints, "right": rightRoundPoints},
    }
}

// GetCenterCards processes the center cards string and maps them to directions
func GetCenterCards(centerCards string, uIndex int) map[string]interface{} {
    result := make(map[string]interface{})
    for key, val := range strings.Split(centerCards, ",") {
        result[GetDirection(key, uIndex)] = val
    }
    return result
}

// GetDirection determines the relative direction of a player
func GetDirection(pIndex, uIndex int) string {
    directions := []string{"down", "left", "up", "right"}
    diff := (4 + (uIndex - pIndex)) % 4 // Calculate relative direction
    return directions[diff]
}

// Matchmaking assigns players to a game and initializes game data in Redis
func Matchmaking(ctx context.Context, client rueidis.Client, userId, gameId string) {
    time.Sleep(1 * time.Second)                                   // Simulate delay for matchmaking
    distributedCards := distributeCards()                         // Distribute cards among players
    lastMoveTimestamp := strconv.FormatInt(time.Now().Unix(), 10) // Record timestamp
    kingCards, king := chooseFirstKing()                          // Determine king cards and player
    redis.Matchmaking(ctx, client, distributedCards, userId, gameId, lastMoveTimestamp, king, kingCards)
}

// distributeCards shuffles and deals cards to 4 players
func distributeCards() []string {
    localCards := make([]string, len(cards))
    copy(localCards, cards)
    rand.Shuffle(len(localCards), func(i, j int) { localCards[i], localCards[j] = localCards[j], localCards[i] })

    hands := make([][]string, 4) // Initialize hands for 4 players
    for i, card := range cards {
        player := i % 4 // Determine player index
        hands[player] = append(hands[player], card)
    }

    result := make([]string, 4)
    for i := 0; i < 4; i++ {
        result[i] = `["` + strings.Join(hands[i], `","`) + `"]` // Format hands as JSON strings
    }

    return result
}

// GetPlayerCards parses and chunks the cards for a specific player
func GetPlayerCards(gameCardsString string, uIndex int) [][]string {
    gameCards := make(map[int][]string)
    err := json.Unmarshal([]byte(gameCardsString), &gameCards) // Parse JSON string
    if err != nil {
        fmt.Println("Error unmarshalling:", err)
    }

    return chunkCards(gameCards, uIndex)
}

// chunkCards divides a player's cards into groups of 5
func chunkCards(gameCards map[int][]string, uIndex int) [][]string {
    var result [][]string
    var chunk []string
    for i, card := range gameCards[uIndex] {
        chunk = append(chunk, card)
        if (i+1)%5 == 0 { // Create a new chunk every 5 cards
            result = append(result, chunk)
            chunk = []string{}
        }
    }
    result = append(result, chunk) // Add remaining cards
    return result
}
