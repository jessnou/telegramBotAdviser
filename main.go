package main

import (
	"adviser_bot/clients/telegram"
	"flag"
	"log"
)

func main() {

	tgClient := telegram.New(mustToken())

	// fetcher := fetcher.New()

	// processor := processor.New()

	// consumer.Start(fetcher, processor)
}

const (
	tgBotHost = "api.telegram.org"
)

func mustToken() string {
	token := flag.String("token-bot-token",
		"",
		"token for acces to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
}
