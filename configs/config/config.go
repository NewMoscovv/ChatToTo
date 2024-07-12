package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	Token string

	Status   Status
	Messages Messages
}

type Messages struct {
	Greetings Greetings `mapstructure:"greeting"`
}

type Greetings struct {
	ToUnregisteredUsers string `mapstructure:"to_unregistered_users"`
	ToBannedUsers       string `mapstructure:"to_banned_users"`
}

type Status struct {
	Unregistered string `mapstructure:"unregistered"`
}

// InitConfig - функция, создающая конфигурацию бота. Она парсит токены, сообщения из окружения
func InitConfig() (*Config, error) {
	// указываем папку в которой хранится файл для парсинга
	viper.AddConfigPath("configs")
	// указываем файл для парсинга
	viper.SetConfigName("main")
	// считывание файла
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	var config Config
	// парсим файл конфига
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	if err := viper.UnmarshalKey("messages.greetings", &config.Messages.Greetings); err != nil {
		return nil, err
	}
	if err := viper.UnmarshalKey("status", &config.Status); err != nil {
		return nil, err
	}
	// парсим .env
	if err := parseEnv(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

// parseEnv - функция, которая передает токен в конфиг
func parseEnv(cfg *Config) error {
	// загружаем файл .env
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	// передаем токен в конфиг
	if value, exists := os.LookupEnv("TOKEN"); exists {
		cfg.Token = value
	}
	return nil

}
