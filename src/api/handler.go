package api

import (
    "net/http"

    "github.com/randykinne/configservice/domain"

    "github.com/gorilla/mux"
)

// Handler for all requests if they don't match an endpoint
func CatchAllHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "hello world"}`))
}

// Handler for something
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    switch r.Method {
    case "GET":    
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"message": "hello world"}`))
    case "POST":
        w.WriteHeader(http.StatusCreated)
        w.Write([]byte(`{"message": "created"}`))
    case "PUT":
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(`{"message": "accepted"}`))
    case "DELETE":
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"message": "OK"}`))
    default:
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"message": "Not Found"}`))
    }
}

// Handler for configuration resources
func ConfigHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    switch r.Method {
    case "GET":
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"message": "hello world"}`))
    }
}

// Handler for specific configuration resources
func SpecificConfigHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    w.Header().Set("Content-Type", "application/json")
    switch r.Method {
    case "GET":
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"message": "hello world"}`))
    }
}