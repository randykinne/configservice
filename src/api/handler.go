package api

import (
	"encoding/json"
	"net/http"

	"github.com/randykinne/configservice/data"
	"github.com/randykinne/configservice/domain"

	"github.com/gorilla/mux"
)

// Response data
type Response struct {
	Message string `json:"message"`
}

// CatchAllHandler for all requests if they don't match an endpoint
func CatchAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "hello world"}`))
}

// ConfigHandler for configuration resources
func ConfigHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "hello world"}`))
	case "POST":
		var c *domain.Config
		err := json.NewDecoder(r.Body).Decode(&c)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data.Put(c)

		response := Response{"created"}
		js, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

// SpecificConfigHandler for specific configuration resources
func SpecificConfigHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	switch r.Method {
	case "GET":
		response, err := data.Get(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		js, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}
