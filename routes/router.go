package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/oat431/go-blog-std/controllers"
)

func StartServer() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// health check
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome Eartling"))
	})
	r.Get("/ping", controllers.HealthCheck)
	r.Get("/ping/{name}", controllers.HealthCheckWithParam)
	r.Post("/ping", controllers.HealthCheckWithBody)

	http.ListenAndServe(":8080", r)
}
