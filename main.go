package main

import (
  "context"
  "net/http"
  "os"
  "os/signal"
  "syscall"
  "time"
  "github.com/rs/zerolog"
  "github.com/rs/zerolog/log"
  "github.com/karl-johan-grahn/microtest/handlers"
  "github.com/karl-johan-grahn/microtest/version"
)

func main() {
  zerolog.TimeFieldFormat = time.RFC3339
  log.Info().Str("commit", version.Commit).Str("build time", version.BuildTime).Str("release", version.Release).Msg("Starting the service...")
  port := os.Getenv("PORT")
  if port == "" {
    log.Fatal().Msg("PORT is not set via env variable, quitting.")
  }
  router := handlers.Router(version.BuildTime, version.Commit, version.Release)
  interrupt := make(chan os.Signal, 1)
  signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

  srv := &http.Server{
    Addr:    ":" + port,
    Handler: router,
  }

  // this channel is for graceful shutdown:
  // if we receive an error, we can send it here to notify the server to be stopped
  shutdown := make(chan struct{}, 1)
  go func() {
    err := srv.ListenAndServe()
    if err != nil {
      shutdown <- struct{}{}
      log.Error().Msg(err.Error())
    }
  }()

  log.Info().Str("port", port).Msg("The service is getting ready to listen and serve")
  select {
  case killSignal := <-interrupt:
    switch killSignal {
    case os.Interrupt:
      log.Info().Msg("Got SIGINT...")
    case syscall.SIGTERM:
      log.Info().Msg("Got SIGTERM...")
    }
  case <-shutdown:
    log.Error().Msg("Got an error...")
  }

  log.Info().Msg("The service is shutting down...")
  srv.Shutdown(context.Background())
  log.Info().Msg("Done")
}
