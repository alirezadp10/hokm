package service

import (
    "context"
    "github.com/alirezadp10/hokm/pkg/repository"
    "strconv"
    "time"
)

type GameService struct {
    GameRepo repository.GameRepository
}

func NewGameService(repo *repository.GameRepository) *GameService {
    return &GameService{
        GameRepo: *repo,
    }
}

func (s *GameService) Matchmaking(ctx context.Context, userId, gameID string, distributedCards []string, king, kingCards string) {
    time.Sleep(1 * time.Second)
    lastMoveTimestamp := strconv.FormatInt(time.Now().Unix(), 10)
    s.GameRepo.Matchmaking(ctx, s.GameRepo.Redis, distributedCards, userId, gameID, lastMoveTimestamp, king, kingCards)
}

func (s *GameService) RemovePlayerFromWaitingList(username string) {
    s.GameRepo.RemovePlayerFromWaitingList(context.Background(), "matchmaking", username)
}