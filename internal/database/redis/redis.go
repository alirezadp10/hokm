package redis

import (
    _ "embed"
    "github.com/redis/rueidis"
    "log"
)

//go:embed setupAGame.lua
var setupAGameScript string

func GetNewConnection() rueidis.Client {
    client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
    if err != nil {
        log.Fatal("couldn't connect to redis")
    }
    return client
}