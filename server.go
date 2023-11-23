package main

import (
	"log"

	"github.com/oat431/go-blog-std/routes"
)

func main() {
	log.Println("Server started on: http://localhost:8080")
	routes.StartServer()
}
