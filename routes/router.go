package routes

import (
	"log"
	"net/http"

	"github.com/oat431/go-blog-std/controllers"
)

func StartServer() {
	http.HandleFunc("/", controllers.HealthCheck)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
