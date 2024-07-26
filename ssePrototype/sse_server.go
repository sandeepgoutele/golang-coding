package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/events", eventHandler)
	http.ListenAndServe(":8080", nil)
}

func eventHandler(respWriter http.ResponseWriter, req *http.Request) {
	// Set CORS headers to allow all origins. You may want to restrict this to specific origins in a production environment.
	respWriter.Header().Set("Access-Control-Allow-Origin", "*")
	respWriter.Header().Set("Access-Control-Expose-Headers", "Content-Type")

	respWriter.Header().Set("Content-Type", "text/event-stream")
	respWriter.Header().Set("Cache-Control", "no-cache")
	respWriter.Header().Set("Connection", "keep-alive")

	for idx := 1; idx < 10; idx++ {
		fmt.Fprintf(respWriter, "data: %s\n\n", fmt.Sprintf("Event %d", idx))
		time.Sleep(1 * time.Second)
		respWriter.(http.Flusher).Flush()
	}

	closeNotify := respWriter.(http.CloseNotifier).CloseNotify()
	<-closeNotify
}
