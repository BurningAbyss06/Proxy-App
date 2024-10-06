package routes

import "net/http"

func HolaMundo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola mundo, te quiero angyee"))
}
