package repository

import (
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type CurrentUser struct {
	ID     string
	Name   string
	Age    int
	Sex    string
	Status string
}

// CreateDB - функция, создающая базу данных.
func CreateDB() *sql.DB {
	// подключение к БД
	db, _ := sql.Open("sqlite3", "database/main.db")
	// команда, создающая таблицу users
	statement, err := db.Prepare("CREATE TABLE IF " +
		"NOT EXISTS users(ID TEXT, NAME TEXT, AGE INT, SEX TEXT, STATUS TEXT)")
	// обработчик ошибки
	if err != nil {
		log.Fatal(err)
	}
	// запускаем команду и создаем таблицу
	_, err = statement.Exec()
	// обработчик ошибки
	if err != nil {
		log.Println("Error creating users table")
		log.Fatal(err)
	}
	return db
}

// InputID - функция, вносящая данные в БД
func InputID(update tgbotapi.Update, db *sql.DB) {
	// командца инъекции данных
	statement, err := db.Prepare("INSERT" +
		" INTO users(ID, NAME, AGE, SEX, STATUS) VALUES(?, ?, 0, '?',  'UNREGISTERED')")
	// обработчик ошибки
	if err != nil {
		log.Fatal(err)
	}
	// запускаем команду и производим инъекцию
	_, err = statement.Exec(update.Message.From.UserName, update.Message.From.FirstName)
	// обработчик ошибки
	if err != nil {
		log.Fatal(err)
	}
}

// GatherInfo - функция, которая собирает данные о пользователе из бд и возвращает их
func GatherInfo(update tgbotapi.Update, db *sql.DB) (CurrentUser, error) {
	var user CurrentUser
	statement, err := db.Query("SELECT"+
		" * FROM users WHERE ID = ?", update.Message.From.UserName)
	if err != nil {
		log.Fatal(err)
	}
	for statement.Next() {
		if err = statement.Scan(&user.ID, &user.Name, &user.Age, &user.Sex, &user.Status); err != nil {
			return user, err
		}
	}
	return user, nil
}
