package main

import (
	"ChatToTo/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	// создаем объект БотАпи и присваиывем ему свой токен
	bot, err := tgbotapi.NewBotAPI()
	if err != nil {
		log.Fatal(err)
	}
	// получение логов
	bot.Debug = false
	// в нашу структуру вставляем объект БотАпи
	tgBot := telegram.NewBot(bot)
	// запускаем бота
	go func() {
		if err := tgBot.Start(); err != nil {
			log.Fatal(err)
		}
	}()
}
