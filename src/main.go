package main

import (
	"fmt"
	"log"
	"net/http"

	"game"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/game/queue_game", game.HandleQueueGame)

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// The "/" pattern matches everything, so we need to check that we're at the root here.
		if request.URL.Path != "/" {
			http.NotFound(writer, request)
			return
		}
		fmt.Fprint(writer, "Welcome to the home page!")
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
