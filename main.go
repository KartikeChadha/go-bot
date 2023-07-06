package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/KartikeChadha/go-bot/helpers"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {

	godotenv.Load(".env")
	os.Getenv("SLACK_BOT_TOKEN")
	os.Getenv("SLACK_APP_TOKEN")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command(helpers.AgeCommand, &slacker.CommandDefinition{
		Description: "birth year calculcator",
		Handler:     helpers.AgeHandler,
	})

	bot.Command(helpers.FileCommand, &slacker.CommandDefinition{
		Description: "file uploader",
		Handler:     helpers.FileHandler,
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen((ctx))
	if err != nil {
		log.Fatal(err)
	}

}
