package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Message struct {
	Hostname string
	Time     time.Time
}

func getHostname(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	fmt.Println("Handling!!!")
	m := Message{hostname, time.Now()}
	json.NewEncoder(w).Encode(m)
}

func handleRequests() {
	http.HandleFunc("/", getHostname)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
