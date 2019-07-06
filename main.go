package main

import (
  "context"
  "log"
  "net/http"
  "os"
  "os/signal"
  "syscall"
  "github.com/karl-johan-grahn/microtest/handlers"
  "github.com/karl-johan-grahn/microtest/version"
)

func main() {
  log.Printf(
    "Starting the service...\ncommit: %s, build time: %s, release: %s",
    version.Commit, version.BuildTime, version.Release,
  )
  port := os.Getenv("PORT")
  if port == "" {
    log.Fatal("PORT is not set via env variable, quitting.")
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
      log.Printf("%v", err)
    }
  }()

  log.Print("The service is getting ready to listen and serve on port " + port)
  select {
  case killSignal := <-interrupt:
    switch killSignal {
    case os.Interrupt:
      log.Print("Got SIGINT...")
    case syscall.SIGTERM:
      log.Print("Got SIGTERM...")
    }
  case <-shutdown:
    log.Printf("Got an error...")
  }

  log.Print("The service is shutting down...")
  srv.Shutdown(context.Background())
  log.Print("Done")
}
