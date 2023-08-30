package route

import (
	"d_gita_be/controller"
	"d_gita_be/middleware"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func InitRoute() {
	mux := mux.NewRouter()
	mux.HandleFunc("/register", controller.Register).Methods("POST")
	mux.HandleFunc("/login", controller.Login).Methods("POST")
	http.ListenAndServe(":8080", middleware.Logger(os.Stderr, mux))
}
