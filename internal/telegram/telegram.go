package telegram

import (
    "fmt"
    "github.com/alirezadp10/hokm/internal/database/sqlite"
    "gopkg.in/telebot.v4"
    "gorm.io/gorm"
    "log"
    "os"
    "time"
)

func Start(db *gorm.DB) {
    bot, err := telebot.NewBot(telebot.Settings{
        Token:  os.Getenv("TELEGRAM_BOT_TOKEN"),
        Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
    })

    if err != nil {
        log.Fatal(err)
    }

    bot.Handle("/start", func(context telebot.Context) error {
        return startHandler(context, db)
    })

    fmt.Println("Bot is running...")
    bot.Start()
}

func startHandler(c telebot.Context, db *gorm.DB) error {
    player, err := sqlite.SavePlayer(db, c.Sender(), c.Chat().ID)
    if err != nil {
        log.Fatalf("couldn't save: %v", err)
    }

    return c.Send("Let's Play", &telebot.ReplyMarkup{
        InlineKeyboard: [][]telebot.InlineButton{{{
            Text:   "شروع بازی",
            WebApp: &telebot.WebApp{URL: os.Getenv("APP_URL") + "?username=" + player.Username},
        }}},
    })
}
