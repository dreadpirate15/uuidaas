package main

import "fmt"
import "github.com/satori/go.uuid"

func main() {
	fmt.Printf("Got UUID: %s\n", uuid.NewV4())
}
