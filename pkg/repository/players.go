package repository

import (
    "github.com/alirezadp10/hokm/pkg/api/request"
    "github.com/alirezadp10/hokm/pkg/model"
    "github.com/redis/rueidis"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
    "log"
    "time"
)

type PlayersRepositoryContract interface {
    CheckPlayerExistence(username string) bool
    SavePlayer(user request.User, chatId int64) (*model.Player, error)
    AddPlayerToGame(username, gameID string) (*model.Game, error)
}

var _ PlayersRepositoryContract = &PlayersRepository{}

type PlayersRepository struct {
    sqlite gorm.DB
    redis  rueidis.Client
}

func NewPlayersRepository(sqliteClient *gorm.DB, redisClient *rueidis.Client) *PlayersRepository {
    return &PlayersRepository{
        sqlite: *sqliteClient,
        redis:  *redisClient,
    }
}

func (r *PlayersRepository) CheckPlayerExistence(username string) bool {
    var count int64

    err := r.sqlite.Table("players").Where("username = ?", username).Count(&count).Error

    if err != nil {
        log.Fatal(err)
        return false
    }

    return count > 0
}

func (r *PlayersRepository) SavePlayer(user request.User, chatId int64) (*model.Player, error) {
    newPlayer := model.Player{
        Id:        user.ID,
        FirstName: user.FirstName,
        LastName:  user.LastName,
        Username:  user.Username,
        ChatId:    chatId,
        UpdatedAt: time.Now(),
        JoinedAt:  time.Now(),
    }

    err := r.sqlite.Clauses(clause.OnConflict{
        Columns:   []clause.Column{{Name: "id"}},
        DoNothing: true,
    }).Create(&newPlayer).Error

    return &newPlayer, err
}

func (r *PlayersRepository) AddPlayerToGame(username, gameID string) (*model.Game, error) {
    var player model.Player
    r.sqlite.First(&player, "username = ?", username)

    newGame := model.Game{GameId: gameID, PlayerId: player.Id, CreatedAt: time.Now()}

    err := r.sqlite.Create(&newGame).Error

    return &newGame, err
}
