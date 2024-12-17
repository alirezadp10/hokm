package cmd

import (
    "context"
    "fmt"
    "github.com/alirezadp10/hokm/internal/database/redis"
    "github.com/alirezadp10/hokm/internal/database/sqlite"
    "github.com/alirezadp10/hokm/internal/handler"
    "github.com/alirezadp10/hokm/internal/middleware"
    "github.com/alirezadp10/hokm/internal/telegram"
    "github.com/labstack/echo/v4"
    "github.com/spf13/cobra"
    "time"
)

var serveCmd = &cobra.Command{
    Use:   "serve",
    Short: "run http server",
    Run:   serve,
}

func init() {
    rootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, args []string) {
    sqliteClient := sqlite.GetNewConnection()
    redisClient := redis.GetNewConnection()
    ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
    defer cancel()

    go telegram.Start(sqliteClient)

    h := handler.NewHandler(sqliteClient, redisClient, ctx)

    e := echo.New()
    e.Static("/assets", "assets")
    e.GET("/", h.GetSplashPage)
    e.GET("/menu", h.GetMenuPage)
    e.GET("/game", h.GetGamePage)
    e.GET("/game/start", h.GetGameId, middleware.AuthMiddleware)
    e.GET("/game/:gameId", h.GetGameData, middleware.AuthMiddleware)
    e.POST("/game/choose-trump", h.ChooseTrump, middleware.AuthMiddleware)
    e.GET("/game/:gameId/cards", h.GetYourCards, middleware.AuthMiddleware)
    e.POST("/game/:gameId/place", h.PlaceCard, middleware.AuthMiddleware)
    e.GET("/game/:gameId/refresh", h.GetUpdate, middleware.AuthMiddleware)

    fmt.Println("Server is running at 9090")
    e.Logger.Fatal(e.Start("0.0.0.0:9090"))
}
