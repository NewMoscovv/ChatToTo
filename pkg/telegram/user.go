package telegram

import (
	"ChatToTo/pkg/repository"
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type User struct {
	ID     string
	Name   string
	Age    int
	Sex    string
	Status string
}

// Info - интерфейс, необходимыЙ для того, чтобы в дальнейшем в случае чего нам легче было менять тип бд
type Info interface {
	Gather(currentUser *repository.CurrentUser) (User, error)
}

// Gather - функция, которая принимает структуру *CurrentUser из database.go и затем создает объект типа User
// и возвращает его
func Gather(update tgbotapi.Update, db *sql.DB) *User {
	currentUser, err := repository.GatherInfo(update, db)
	if err != nil {
		log.Fatal(err)
	}
	return &User{ID: currentUser.ID, Name: currentUser.Name, Age: currentUser.Age, Sex: currentUser.Sex, Status: currentUser.Status}
}

func CheckUser(update tgbotapi.Update, db *sql.DB) bool {
	currentUser := Gather(update, db)
	if currentUser.Status == "" || currentUser.Status == "UNREGISTERED" {
		return false
	}
	return true
}
