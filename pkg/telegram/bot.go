package telegram

import (
	"ChatToTo/configs/config"
	"ChatToTo/pkg/keyboards"
	"ChatToTo/pkg/repository"
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	db  *sql.DB
	cfg config.Config
}

// NewBot - функция, добавляющая в структуру Bot апи телеграма
func NewBot(bot *tgbotapi.BotAPI, db *sql.DB, cfg config.Config) *Bot {
	return &Bot{bot: bot, db: db, cfg: cfg}
}

// Start - функция запуска бота
func (b *Bot) Start() error {
	// лог запуска в терминал
	log.Println("Telegram bot started")
	// все обновления инициализируем в переменную через функцию
	updates := b.initUpdatesChannel()
	// обрабатываем апдейты
	b.handleUpdates(updates)
	return nil
}

// initUpdatesChannel - функция инициализации обновлений
func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	// Формируем конфиг обновления, с помощью которого бот будет получать обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	// канал, передающий обновления
	//(внутри го-рутина, которая позволяет постоянно получать сообщения)
	return b.bot.GetUpdatesChan(u)
}

// handleUpdates = обработчик событий, там все понятно
func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		// проверяем на наличие сообщения
		if update.Message == nil {
			continue
		}
		// проверяем человека на наличие в базе данных
		if CheckUser(update, b.db) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "ты зареган уже броски")
			msg.ReplyToMessageID = update.Message.MessageID
			b.bot.Send(msg)
		} else {
			// если текст, отправленный юзером, который не удовлетворяет бд, содержит в себе 'Пройти регистрацию', то начинается процесс регистрации
			if update.Message.Text == "Пройти регистрацию" {

			} else {
				// в противном случае мы приветствуем пользователя
				b.greetUnregedUsers(update)
			}

		}
	}
}

// greetUnregedUsers - фукнция приветсвия пользователя, который либо не содержится в бд, либо забанен
func (b *Bot) greetUnregedUsers(update tgbotapi.Update) {
	// объявляем переменную типа User и собираем в него информацию из БД
	currentUser := Gather(update, b.db)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = keyboards.RegistrationButton

	// исходя из данных формируем сообщение
	switch currentUser.Status {
	case "BANNED":
		msg.Text = fmt.Sprintf(b.cfg.Messages.Greetings.ToBannedUsers, update.Message.From.UserName)
		b.bot.Send(msg)
	default:
		if currentUser.Status == "" {
			repository.InputID(update, b.db)
		}
		msg.Text = fmt.Sprintf(b.cfg.Messages.Greetings.ToUnregisteredUsers, update.Message.From.UserName)
		b.bot.Send(msg)
	}
}
