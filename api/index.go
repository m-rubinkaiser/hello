package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type Greet struct {
	Greeting string
}

func Hello(w http.ResponseWriter, r *http.Request) {
	greet := Greet{Greeting: "hello"}
	json.NewEncoder(w).Encode(greet)
}

func Handler() {
	http.HandleFunc("/", Hello)
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", nil))
}
