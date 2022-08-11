package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func Test(t *testing.T) {
	testURL = "http://" + getEnvValue("HOST") + ":" + getEnvValue("PORT")
	testClient = http.Client{}
}

func TestServerPing(t *testing.T) {
	res, err := http.Get(testURL + "/ping")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("status not OK")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Error(err)
		}
	}(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	b := string(body)
	if !strings.Contains(b, "Todo") {
		t.Fatal()
	}
}

func TestLoadPing(t *testing.T) {
	rate := vegeta.Rate{Freq: 1000, Per: time.Second}
	duration := 5 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    testURL + "/ping",
	})
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()
	log.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}

func TestServerCreate(t *testing.T) {
	body := []byte(`[{"title":  "Buy cinnabon"}]`)
	req, err := http.NewRequest(http.MethodPost, testURL+"/todo", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	res, err := testClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("status not OK")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Error(err)
		}
	}(res.Body)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLoadCreate(t *testing.T) {
	rate := vegeta.Rate{Freq: 1000, Per: time.Second}
	duration := 5 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    testURL + "/todo?title=Buy%20coffee",
	})
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()
	log.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}

func TestServerGetOne(t *testing.T) {
	res, err := http.Get(testURL + "/todo?title=Buy%20coffee")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("status not OK")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Error(err)
		}
	}(res.Body)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLoadGetOne(t *testing.T) {
	rate := vegeta.Rate{Freq: 1000, Per: time.Second}
	duration := 5 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    testURL + "/todo?id=1",
	})
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()
	log.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}

func TestServerGetAll(t *testing.T) {
	res, err := http.Get(testURL + "/todos")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("status not OK")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Error(err)
		}
	}(res.Body)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLoadGetAll(t *testing.T) {
	rate := vegeta.Rate{Freq: 1000, Per: time.Second}
	duration := 5 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    testURL + "/todos",
	})
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()
	log.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}

func TestServerUpdate(t *testing.T) {
	// TODO: Corrections are needed. Incorrect request
	body := []byte(`[{"id":  "1"}, {"comleted": "true"}]`)
	req, err := http.NewRequest(http.MethodPatch, testURL+"/todo?", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	res, err := testClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(res)
	if res.StatusCode != http.StatusOK {
		t.Errorf("status not OK")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Error(err)
		}
	}(res.Body)
	if err != nil {
		t.Fatal(err)
	}
}

func TestServerDeleteOne(t *testing.T) {
	body := []byte(`[{"id":  "1"}]`)
	req, err := http.NewRequest(http.MethodDelete, testURL+"/todo", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	res, err := testClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusNoContent {
		t.Errorf("status not OK")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Error(err)
		}
	}(res.Body)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLoadDeleteOne(t *testing.T) {
	rate := vegeta.Rate{Freq: 1000, Per: time.Second}
	duration := 5 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "DELETE",
		URL:    testURL + "/todo?id=1",
	})
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()
	log.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}

func TestServerDeleteAll(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, testURL+"/todos", nil)
	if err != nil {
		t.Fatal(err)
	}
	res, err := testClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusNoContent {
		t.Errorf("status not OK")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Error(err)
		}
	}(res.Body)
	if err != nil {
		t.Fatal(err)
	}
}
