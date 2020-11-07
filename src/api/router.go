package api

import (
    "net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "hello world"}`))
}

func route(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    switch r.Method {
    case "GET":    
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"message": "hello world"}`))
    case "POST":
        w.WriteHeader(http.StatusCreated)
        w.Write([]byte(`{"message": "created"}`))
    case "PUT":
        w.WriteHeader(http.StatusAccepted)
        w.Write([]byte(`{"message": "accepted"}`))
    case "DELETE":
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"message": "OK"}`))
    default:
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"message": "Not Found"}`))
    }
}