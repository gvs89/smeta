package main

import (
	"log"
	"net/http"

	"github.com/yourusername/yourprojectname/internal/api"
)

func main() {
	router := api.Router()

	log.Println("Запуск сервера...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
