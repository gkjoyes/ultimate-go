package main

import (
	"log"
	"net/http"

	"github.com/gkjoyes/ultimate-go/lesson_12/example4/handlers"
)

func main() {
	handlers.Routes()
	log.Println("Listening on: http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}
