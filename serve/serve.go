package serve

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

var db DB

// Serve starts the server
func Serve(cfg Config) {
	var err error

	db, err = newDB(cfg)
	if err != nil {
		log.Error("Failed getting new DB: ", err)
		return
	}

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	srv := http.Server{
		Addr:              addr,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       30 * time.Second,
		Handler:           ServerHandler(),
	}

	log.Info("Starting HTTP server at: ", addr)
	log.Fatal(srv.ListenAndServe())
}

// ServerHandler maps all the paths to handlers via mux
func ServerHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/set", handlerSet)
	mux.HandleFunc("/get", handlerGet)
	mux.HandleFunc("/", handlerRoot)

	return mux
}
