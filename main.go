/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
    "github.com/alirezadp10/hokm/cmd"
    "github.com/joho/godotenv"
)

func main() {
    _ = godotenv.Load()
    cmd.Execute()
}
