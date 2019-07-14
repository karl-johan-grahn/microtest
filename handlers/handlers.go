package handlers

import (
  "net/http"
  "sync/atomic"
  "time"
  "github.com/rs/zerolog/log"
  "github.com/gorilla/mux"
  "github.com/prometheus/client_golang/prometheus/promhttp"
)

func openapi(w http.ResponseWriter, r *http.Request) {
  // Note that path to file is defined in dockerfile
  http.ServeFile(w, r, "./openapi.yaml")
}

// Router register necessary routes and returns an instance of a router.
func Router(buildTime, commit, release string) *mux.Router {
  // Should use channels instead of counters for managing state?
  isReady := &atomic.Value{}
  isReady.Store(false)
  go func() {
    log.Info().Msg("Readyz endpoint is negative by default...")
    // Simulate warmup by waiting a little
    time.Sleep(10 * time.Second)
    isReady.Store(true)
    log.Info().Msg("Readyz endpoint is positive.")
  }()
  r := mux.NewRouter()
  api := r.PathPrefix("/api").Subrouter()
  api.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
  })
  api1 := api.PathPrefix("/v1").Subrouter()
  api1.HandleFunc("/hello", hello(buildTime, commit, release)).Methods("GET")
  api1.HandleFunc("/openapi", openapi).Methods("GET")
  api1.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusForbidden)
  })
  r.HandleFunc("/live", healthz).Methods("GET")
  r.HandleFunc("/ready", readyz(isReady)).Methods("GET")
  r.Handle("/metrics", promhttp.Handler()).Methods("GET")
  return r
}
