package repository

import (
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// CreateDB - функция, создающая базу данных.
func CreateDB() *sql.DB {
	// подключение к БД
	db, _ := sql.Open("sqlite3", "database/main.db")
	// команда, создающая таблицу users
	statement, err := db.Prepare("CREATE TABLE IF " +
		"NOT EXISTS users(ID TEXT, NAME, TEXT, AGE INT, SEX TEXT, STATUS TEXT)")
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
		" INTO users(ID, NAME) VALUES(?, ?)")
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

// CheckDB - функция, проверяющая на наличие в БД юзера с определенным ID
func CheckDB(update tgbotapi.Update, db *sql.DB) bool {
	// получаем строку, где ячейка ID равна введенному значению ID
	statement, err := db.Query("SELECT "+
		"ID FROM users WHERE ID = ?", update.Message.From.UserName)
	// обработчик ошибки
	if err != nil {
		log.Fatal(err)
	}
	// проверка на наличие ячейки с введенным ID
	for statement.Next() {
		var value string
		err := statement.Scan(&value)
		if err != nil {
			log.Fatal(err)
		}
		if value == update.Message.From.UserName {
			return true
		}
	}
	return false
}
