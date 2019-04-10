package infra

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/yamad07/gRPC-sample/entities/models"
)

var DB *gorm.DB

func InitDB() {
	user := getParamFromEnv("TODO_DB_USER", "root")
	host := getParamFromEnv("TODO_DB_HOST", "127.0.0.1")
	password := getParamFromEnv("TODO_DB_PASSWORD", "")
	name := getParamFromEnv("TODO_DB_NAME", "todo_dev")
	protocol := getParamFromEnv("TODO_DB_PROTOCOL", "tcp")
	port := getParamFromEnv("TODO_DB_PORT", "3306")
	args := getParamFromEnv("TODO_DB_ARGS", "parseTime=true&loc=Asia%2FTokyo&charset=utf8mb4")
	connectionString := fmt.Sprintf("%s:%s@%s([%s]:%s)/%s?%s",
		user, password, protocol, host, port, name, args)

	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Todo{})
	DB = db.Debug()
}

func getParamFromEnv(env, defaultValue string) string {
	param := os.Getenv(env)

	if param == "" {
		param = defaultValue
	}

	return param
}
