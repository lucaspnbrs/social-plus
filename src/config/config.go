package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//String database connection
	StringConnection = " "

	//Port API is running
	Port = 0
)

//Initialize Enviroment Variables 
func LoadSys() {
	var erro error

	if erro := godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Port, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Port = 9000
	}

	StringConnection = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
    os.Getenv("DB_HOST_LOCAL"),
    os.Getenv("DB_PORT_LOCAL"),
    os.Getenv("DB_USER_LOCAL"),
    os.Getenv("DB_PASSWORD_LOCAL"),
    os.Getenv("DB_NAME"),
)
}