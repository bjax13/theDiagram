package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	// API the api server
	API struct {
		Port string
	}
)

// RunServer runs the http server
func (api API) RunServer() error {

	r := mux.NewRouter()
	r.HandleFunc("/api/helloworld", MyHandler).Methods("GET")

	fs := http.FileServer(http.Dir("client/build/"))
	r.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	fmt.Printf("Running server on port %s\n", api.Port)
	http.ListenAndServe(":"+api.Port, r)

	return nil
}

// MyHandler a route handler
func MyHandler(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Word string
	}{
		Word: "yo",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
