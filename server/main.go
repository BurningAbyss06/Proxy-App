package main

import (
	"net/http"

	"github.com/Angxandralol/Proxy-App/server/database"
	"github.com/Angxandralol/Proxy-App/server/routes"
	"github.com/gorilla/mux"
)

func main() {
	database.DBConnection()
	r := mux.NewRouter()

	r.HandleFunc("/", routes.HolaMundo)

	http.ListenAndServe(":8080", r)
}
