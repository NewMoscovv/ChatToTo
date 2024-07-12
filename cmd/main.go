package main

import (
	"ChatToTo/configs/config"
	"ChatToTo/pkg/repository"
	"ChatToTo/pkg/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {
	// создаем конфиг бота, в нем хранятся токены, сообщения, команды и т.д.
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	// создаем объект БотАпи и передаем ему токен
	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		log.Fatal(err)
	}
	// получение логов
	bot.Debug = false
	// в нашу структуру вставляем объект БотАпи
	tgBot := telegram.NewBot(bot, repository.CreateDB(), *cfg)
	// запускаем бота
	if err := tgBot.Start(); err != nil {
		log.Fatal(err)
	}
}
