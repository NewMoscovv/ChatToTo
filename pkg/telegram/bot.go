package telegram

import (
	"ChatToTo/pkg/repository"
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	db  *sql.DB
}

// NewBot - функция, добавляющая в структуру Bot апи телеграма
func NewBot(bot *tgbotapi.BotAPI, db *sql.DB) *Bot {
	return &Bot{bot: bot, db: db}
}

// Start - функция запуска бота
func (b *Bot) Start() error {
	// лог запуска в терминал
	log.Println("Telegram bot started")
	// все обновления инициализируем в переменную через функцию
	updates, err := b.initUpdatesChannel()
	if err != nil {
		return err
	}
	// обрабатываем апдейты
	b.handleUpdates(updates)
	return nil
}

// initUpdatesChannel - функция инициализации обновлений
func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	// Формируем конфиг обновления, с помощью которого бот будет получать обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	// канал, передающий обновления. (внутри го-рутина, которая позволяет постоянно получать сообщения)
	return b.bot.GetUpdatesChan(u)
}

// handleUpdates = обработчик событий, там все понятно
func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if repository.CheckDB(update, b.db) {
			if update.Message == nil {
				continue
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				b.bot.Send(msg)
			}
		} else {
			repository.InputID(update, b.db)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("<b>%s</b>\nПривествуем, вам необходимо пройти <i>регистрацию</i>\nприступим?", update.Message.From.UserName))
			msg.ParseMode = "HTML"
			b.bot.Send(msg)
		}
	}
}
