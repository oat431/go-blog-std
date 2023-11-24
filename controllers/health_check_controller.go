package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	req "github.com/oat431/go-blog-std/payloads/request"
	res "github.com/oat431/go-blog-std/payloads/response"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, res.HealthCheckDto{
		Message: "pong",
	})
}

func HealthCheckWithParam(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	render.JSON(w, r, res.HealthCheckDto{
		Message: "pong " + name,
	})
}

func HealthCheckWithBody(w http.ResponseWriter, r *http.Request) {
	var ping req.HealthCheckRequest
	render.DecodeJSON(r.Body, &ping)
	name := ping.Ping
	render.JSON(w, r, res.HealthCheckDto{
		Message: "pong " + name,
	})
}
