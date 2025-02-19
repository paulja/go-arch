package main

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/paulja/go-arch/web/config"
	"github.com/paulja/go-arch/web/internal/middleware"
	"github.com/paulja/go-arch/web/internal/search"
)

func main() {
	vlog := slog.Default()
	slog.SetLogLoggerLevel(slog.LevelDebug)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/users/{exp}", handleUsers)

	vlog.Info("server", "endpoint", config.GetServePort())
	log.Fatal(http.ListenAndServe(
		config.GetServePort(),
		middleware.Logger(mux),
	))
}

func handleRoot(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func handleUsers(w http.ResponseWriter, req *http.Request) {
	exp := req.PathValue("exp")

	c := search.NewSearchClient()
	err := c.Connect()
	defer c.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u, err := c.FindUsers(exp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(u) == 0 {
		http.Error(w, fmt.Sprintf("no users found using exp: %s", exp), http.StatusNotFound)
		return
	}

	buf, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(buf)
}
