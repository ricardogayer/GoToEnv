package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	err = godotenv.Load("prod.env")
	if err != nil {
		log.Println("Erro lendo as variáveis de ambiente")
	}

	url := os.Getenv("DATABASE_URL")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWD")

	fmt.Println("URL: ",url)
	fmt.Println("User: ",user)
	fmt.Println("Password: ",password)
	
}
