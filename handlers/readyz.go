package handlers

import (
  "net/http"
  "sync/atomic"
)

// readiness endpoint - is the application ready to serve traffic?
func readyz(isReady *atomic.Value) http.HandlerFunc {
  return func(w http.ResponseWriter, _ *http.Request) {
    if isReady == nil || !isReady.Load().(bool) {
      http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
      return
    }
    w.WriteHeader(http.StatusOK)
  }
}
