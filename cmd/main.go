package main

import (
	"WithGres/internal/server"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Server is running")
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
