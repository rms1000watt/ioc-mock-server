package serve_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/rms1000watt/ioc-mock-server/serve"

	log "github.com/sirupsen/logrus"
)

const (
	host = "127.0.0.1"
	port = 7100

	key = "name"
	val = "ryan"
	ok  = "OK"
)

func TestMain(m *testing.M) {
	cfg := serve.Config{
		Host:   host,
		Port:   port,
		DBType: "mock",
	}

	log.SetLevel(log.DebugLevel)
	go serve.Serve(cfg)
	time.Sleep(200 * time.Millisecond)

	code := m.Run()

	time.Sleep(200 * time.Millisecond)
	os.Exit(code)
}

func TestSet(t *testing.T) {
	addr := fmt.Sprintf("%s:%d", host, port)

	req := serve.RequestSet{
		Key: key,
		Val: val,
	}
	jsonBytes, _ := json.Marshal(req)
	buf := bytes.NewBuffer(jsonBytes)

	url := "http://" + addr + "/set"
	res, err := http.Post(url, "application/json", buf)
	if err != nil {
		t.Fatal("Failed posting to server: ", err)
	}

	bodyBytes, _ := ioutil.ReadAll(res.Body)
	if strings.TrimSpace(string(bodyBytes)) != ok {
		t.Fatal("Failed setting value: ", string(bodyBytes))
	}
}

func TestGet(t *testing.T) {
	addr := fmt.Sprintf("%s:%d", host, port)

	reqSet := serve.RequestSet{
		Key: key,
		Val: val,
	}
	jsonBytes, _ := json.Marshal(reqSet)
	buf := bytes.NewBuffer(jsonBytes)

	url := "http://" + addr + "/set"
	res, err := http.Post(url, "application/json", buf)
	if err != nil {
		t.Fatal("Failed posting to server: ", err)
	}

	bodyBytes, _ := ioutil.ReadAll(res.Body)
	if strings.TrimSpace(string(bodyBytes)) != ok {
		t.Fatal("Failed setting value: ", string(bodyBytes))
	}

	reqGet := serve.RequestGet{
		Key: key,
	}
	jsonBytes, _ = json.Marshal(reqGet)
	buf = bytes.NewBuffer(jsonBytes)

	url = "http://" + addr + "/get"
	res, err = http.Post(url, "application/json", buf)
	if err != nil {
		t.Fatal("Failed posting to server: ", err)
	}

	bodyBytes, _ = ioutil.ReadAll(res.Body)
	if strings.TrimSpace(string(bodyBytes)) != val {
		t.Fatal("Failed setting value: ", string(bodyBytes))
	}
}
