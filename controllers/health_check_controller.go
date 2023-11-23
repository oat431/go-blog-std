package controllers

import (
	"net/http"

	res "github.com/oat431/go-blog-std/payloads/response"
	utils "github.com/oat431/go-blog-std/utils"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.SendJson(http.StatusBadGateway, w, res.HealthCheckDto{
		Message: "Server is running",
	})
}
