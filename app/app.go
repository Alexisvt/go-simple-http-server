package app

import (
	"net/http"

	"github.com/alexisvt/go-simple-http-server/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":3333", nil); err != nil {
		panic(err)
	}
}
