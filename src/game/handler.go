package game

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func HandleQueueGame(writer http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	uid := request.URL.Query().Get("uid")
	if uid == "" {
		http.Error(writer, "no uid", http.StatusBadRequest)
		return
	}

	result, err := QueueGame(ctx, uid)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(writer, "%+v", result)
	return
}
