package handlers

import (
  "log"
  "sync/atomic"
  "time"
  "github.com/gorilla/mux"
  "github.com/prometheus/client_golang/prometheus/promhttp"
)

// Router register necessary routes and returns an instance of a router.
func Router(buildTime, commit, release string) *mux.Router {
  // Should use channels instead of counters for managing state?
  isReady := &atomic.Value{}
  isReady.Store(false)
  go func() {
    log.Printf("Readyz endpoint is negative by default...")
    // Simulate warmup by waiting a little
    time.Sleep(10 * time.Second)
    isReady.Store(true)
    log.Printf("Readyz endpoint is positive.")
  }()
  r := mux.NewRouter()
  r.HandleFunc("/hello", hello(buildTime, commit, release)).Methods("GET")
  r.HandleFunc("/healthz", healthz)
  r.HandleFunc("/readyz", readyz(isReady))
  r.Handle("/metrics", promhttp.Handler())
  return r
}
