package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rhobro/csclub/handler/algo"
)

func init() {
	// init all endpoints
	algo.Init(root.PathPrefix("/algo").Subrouter())
	// status check endpoint
	root.HandleFunc("/status", func(w http.ResponseWriter, _ *http.Request) {}).
		Methods(http.MethodGet)
}

// root router
var root = mux.NewRouter()

// getter
func Root() *mux.Router {
	return root
}
