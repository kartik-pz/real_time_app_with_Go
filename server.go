package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

var addr = flag.String("addr", ":8000", "address to listen on")

func main() {
	flag.Parse()

	http.HandleFunc("/", home)
	http.HandleFunc("/events", events)
	fmt.Printf("Server starting on %s\n", *addr)
	http.ListenAndServe(*addr, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func events(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")

	tokens := []string{"this", "is", "a", "live", "event", "test", "for", "stimulating", "llm", "response."}

	for _, token := range tokens {
		content := fmt.Sprintf("data: %s\n\n", string(token))
		w.Write([]byte(content))
		w.(http.Flusher).Flush()

		// intentional wait
		time.Sleep(time.Millisecond * 420)
	}
}
