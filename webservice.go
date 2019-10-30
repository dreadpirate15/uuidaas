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
	u, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
    ut := time.Now().Format(time.RFC822)
    uuid := UUIDResponse{u.String(), ut}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(uuid)
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
