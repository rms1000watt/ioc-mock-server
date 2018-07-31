package serve

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// RequestSet is the input for get request
type RequestSet struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

func handlerSet(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	req := RequestSet{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errStr := "Failed parsing json body: " + err.Error()
		log.Error(errStr)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}

	log.Debug("/set key: ", req.Key)
	log.Debug("/set val: ", req.Val)

	if err := db.Set(req.Key, req.Val); err != nil {
		errStr := "Failed setting key val in DB: " + err.Error()
		log.Error(errStr)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}

	log.Debug("/set OK")
	w.Write([]byte("OK"))
}
