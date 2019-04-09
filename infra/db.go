package infra

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Init() {
	user := getParamFromEnv("EMODE_DB_USER", "root")
	host := getParamFromEnv("EMODE_DB_HOST", "127.0.0.1")
	password := getParamFromEnv("EMODE_DB_PASSWORD", "")
	name := getParamFromEnv("EMODE_DB_NAME", "e_mode_dev")
	protocol := getParamFromEnv("EMODE_DB_PROTOCOL", "tcp")
	port := getParamFromEnv("EMODE_DB_PORT", "3306")
	args := getParamFromEnv("EMODE_DB_ARGS", "parseTime=true&loc=Asia%2FTokyo&charset=utf8mb4")
	connectionString := fmt.Sprintf("%s:%s@%s([%s]:%s)/%s?%s",
		user, password, protocol, host, port, name, args)

	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	DB = db.Debug()
}

func getParamFromEnv(env, defaultValue string) string {
	param := os.Getenv(env)

	if param == "" {
		param = defaultValue
	}

	fmt.Println(param)

	return param
}
