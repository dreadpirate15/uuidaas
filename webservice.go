package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index")
}

func generateUUID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, uuid.NewV4().String())
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
