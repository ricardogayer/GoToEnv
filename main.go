package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load() // este método ler o arquivo .env
	// err = godotenv.Load("prod.env")
	if err != nil {
		log.Println("Erro lendo as variáveis de ambiente")
	}

	url := os.Getenv("DATABASE_URL")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWD")
	linhaInicial := os.Getenv("LINHA_INICIAL")
	linhaFinal := os.Getenv("LINHA_FINAL")

	fmt.Println("URL: ", url)
	fmt.Println("User: ", user)
	fmt.Println("Password: ", password)

	fmt.Println("Linha Inicial: ", linhaInicial)
	fmt.Println("Linha Final: ", linhaFinal)

}
