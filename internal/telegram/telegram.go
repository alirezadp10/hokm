package telegram

import (
    "fmt"
    "github.com/alirezadp10/hokm/internal/database"
    "github.com/joho/godotenv"
    "gopkg.in/telebot.v4"
    "gorm.io/gorm"
    "log"
    "os"
    "time"
)

func Start(db *gorm.DB) {
    _ = godotenv.Load()
    pref := telebot.Settings{
        Token:  os.Getenv("TELEGRAM_BOT_TOKEN"),
        Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
    }
    bot, err := telebot.NewBot(pref)

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
    _, err := database.SavePlayer(db, c.Message().Sender)
    if err != nil {
        log.Fatalf("couldn't save: %v", err)
    }
    return c.Send("Let's Play", &telebot.ReplyMarkup{
        InlineKeyboard: [][]telebot.InlineButton{
            {
                {
                    Text:   "شروع بازی",
                    WebApp: &telebot.WebApp{URL: os.Getenv("APP_URL")},
                },
            },
        },
    })
    //client, _ := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"redis:6379"}})
    //defer client.Close()
    //subClient := client.B().Subscribe().Channel("").Build() // Subscription command
    //
    //message := c.Message()
    //user := message.Sender
    //
    //foo, _ := json.MarshalIndent(user, "", " ")
    // Print message details
    //fmt.Println(string(foo))
    //return c.Send("سلام خوش اومدی، صبر کن بقیه هم بیان خبرت میکنم")
}
