package route

import (
	"d_gita_be/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoute() {
	mux := mux.NewRouter()
	mux.HandleFunc("users/", controller.Register).Methods("POST")

	http.ListenAndServe(":8080", mux)
}
