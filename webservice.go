package main

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
	"time"
)

type UUIDResponse struct {
	UUID      string `json:"uuid,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index")
}

func generateUUID(w http.ResponseWriter, r *http.Request) {
	var u UUIDResponse
	u.UUID = uuid.NewV4().String()
	u.Timestamp = time.Now().Format(time.RFC822)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

func SetHandlers() {
	http.HandleFunc("/", index)
	http.HandleFunc("/uuid", generateUUID)
}

func main() {
	fmt.Println("Init")
	SetHandlers()
	//fmt.Println("Got UUID: %s", uuid.NewV4())
	http.ListenAndServe(":8080", nil)
}
