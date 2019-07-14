package handlers

import (
  "encoding/json"
  "net/http"
  "github.com/rs/zerolog/log"
//  "math"
//  "math/rand"
//  "time"
)

// hello returns a simple HTTP handler function which writes a response.
func hello(buildTime, commit, release string) http.HandlerFunc {
  return func(w http.ResponseWriter, _ *http.Request) {
    info := struct {
      BuildTime string `json:"buildTime"`
      Commit    string `json:"commit"`
      Release   string `json:"release"`
    }{
      buildTime, commit, release,
    }

    body, err := json.Marshal(info)
    if err != nil {
      log.Error().Msg("Could not encode info data")
      http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
      return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(body)
  }
}

/* func IsPrime(value int) bool {
    for i := 2; i <= int(math.Floor(float64(value) / 2)); i++ {
        if value%i == 0 {
            return false
        }
    }
    return value > 1
}

func Jitter(d time.Duration) time.Duration {
  jitter := time.Duration(rand.Int63n(int64(d)))
  d = d + jitter/2
  return time.Duration(d)
} */
