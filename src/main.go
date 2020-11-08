package main

import (
	"github.com/randykinne/configservice/api"
	"github.com/randykinne/configservice/data"

	"net/http"
	"os"

	"github.com/gorilla/mux"
	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	"github.com/slok/go-http-metrics/middleware/std"
)

// Our main function
func main() {

	mdlw := middleware.New(middleware.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})

	var log = logrus.New()

	log.Out = os.Stdout

	r := mux.NewRouter()
	r.Use(std.HandlerProvider("", mdlw))

	log.Info("HTTP Server Ready")
	data.Initialize()
	log.Info("Data Store Initialized")

	configHandler := r.PathPrefix("/api/configurations").Subrouter()
	configHandler.HandleFunc("", api.ConfigHandler)
	configHandler.HandleFunc("/{id:[0-9]+}", api.SpecificConfigHandler)
	//r.PathPrefix("/").HandlerFunc(api.CatchAllHandler)

	r.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":8080", r))
}
