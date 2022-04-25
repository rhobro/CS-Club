package algo

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router) {
	r.HandleFunc("/gcd", giveGCD).
		Methods(http.MethodGet)
	r.HandleFunc("/gcd", receiveGCD).
		Methods(http.MethodPost)
}
