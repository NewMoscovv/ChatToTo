package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

// NewBot - функция, добавляющая в структуру Bot апи телеграма
func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
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
		if update.Message == nil {
			continue
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			b.bot.Send(msg)
		}
	}
}
