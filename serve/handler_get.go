package serve

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// RequestGet is the input for get request
type RequestGet struct {
	Key string `json:"key"`
}

func handlerGet(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	req := RequestGet{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errStr := "Failed parsing json body: " + err.Error()
		log.Error(errStr)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}

	log.Debug("/get key: ", req.Key)
	val, err := db.Get(req.Key)
	if err != nil {
		errStr := "Failed getting key from DB: " + err.Error()
		log.Error(errStr)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}

	log.Debug("/get val: ", val)
	w.Write([]byte(val))
}

func bodyGetKey(body io.ReadCloser) (key string, err error) {
	bodyBytes := []byte{}
	buf := bytes.NewBuffer(bodyBytes)
	if _, err = io.Copy(buf, body); err != nil {
		log.Debug("/get body copy failed: ", err)
		return
	}

	key = string(bodyBytes)
	return
}
