package handlers

import (
  "encoding/json"
  "io/ioutil"
  "net/http"
  "net/http/httptest"
  "testing"
  "time"
)

func TestHello(t *testing.T) {
  w := httptest.NewRecorder()
  buildTime := time.Now().Format("20190705_10:28:00")
  commit := "t3sThasH"
  release := "0.0.1"
  h := hello(buildTime, commit, release)
  h(w, nil)

  resp := w.Result()
  if have, want := resp.StatusCode, http.StatusOK; have != want {
    t.Errorf("Status code is wrong. Have: %d, want: %d.", have, want)
  }

  greeting, err := ioutil.ReadAll(resp.Body)
  resp.Body.Close()
  if err != nil {
    t.Fatal(err)
  }
  info := struct {
    BuildTime string `json:"buildTime"`
    Commit    string `json:"commit"`
    Release   string `json:"release"`
  }{}
  err = json.Unmarshal(greeting, &info)
  if err != nil {
    t.Fatal(err)
  }
  if info.Release != release {
    t.Errorf("Release version is wrong. Have: %s, want: %s", info.Release, release)
  }
  if info.BuildTime != buildTime {
    t.Errorf("Build time is wrong. Have: %s, want: %s", info.BuildTime, buildTime)
  }
  if info.Commit != commit {
    t.Errorf("Commit is wrong. Have: %s, want: %s", info.Commit, commit)
  }
}
