package controllers

import (
	"net/http"
	"strconv"

	"github.com/alexisvt/go-simple-http-server/services"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		return
	}
	user, err := services.GetUser(userId)

}
