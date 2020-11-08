package api

import (
	"encoding/json"
	"net/http"

	"github.com/randykinne/configservice/domain"
	"github.com/randykinne/configservice/store"

	"github.com/gorilla/mux"
)

// Response data
type Response struct {
	StatusCode int
	Data       map[string]interface{}
}

// CatchAllHandler for all requests if they don't match an endpoint
func CatchAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

// ConfigHandler for configuration resources
func ConfigHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		data := make(map[string]interface{})
		data["message"] = "success"
		data["type"] = "configuration"

		values, err := store.GetAll()

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data["body"] = values

		response := Response{http.StatusOK, data}

		js, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(response.StatusCode)
		w.Write(js)

	case "POST":
		var c *domain.Config
		err := json.NewDecoder(r.Body).Decode(&c)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		store.Put(c)

		data := make(map[string]interface{})
		data["message"] = "created"
		data["location"] = "lol"
		response := Response{http.StatusCreated, data}

		js, err := json.Marshal(response.Data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(response.StatusCode)
		w.Write(js)
	}
}

// SpecificConfigHandler for specific configuration resources
func SpecificConfigHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	switch r.Method {
	case "GET":
		response, err := store.Get(vars["id"])
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
