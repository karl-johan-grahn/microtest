package handlers

import "net/http"

// liveness endpoint - is the application running?
// Scale the service based on number of requests via HPA instead?
func healthz(w http.ResponseWriter, _ *http.Request) {
  w.WriteHeader(http.StatusOK)
}
