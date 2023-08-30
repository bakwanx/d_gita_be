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
	var imgServer = http.FileServer(http.Dir("./public/"))
	mux.PathPrefix("/public/").Handler(http.StripPrefix("/public/", imgServer))

	mux.HandleFunc("/register", controller.Register).Methods("POST")
	mux.HandleFunc("/login", controller.Login).Methods("POST")
	mux.HandleFunc("/receipt", controller.CreateReceipt).Methods("POST")

	http.ListenAndServe(":8080", middleware.Logger(os.Stderr, mux))
}
