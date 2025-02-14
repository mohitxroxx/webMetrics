package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Metrics struct {
	RequestCount int            `json:"request_count"`
	Endpoints    map[string]int `json:"endpoints"`
	mu           sync.Mutex
}

func NewMetrics() *Metrics {
	return &Metrics{
		Endpoints: make(map[string]int),
	}
}

func (m *Metrics) increment(path string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.RequestCount++
	m.Endpoints[path]++
}

func main() {
	metrics := NewMetrics()
	setupRoutes(metrics)
}

func setupRoutes(metrics *Metrics) {

	withMetrics := func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			metrics.increment(r.URL.Path)
			next(w, r)
		}
	}

	http.HandleFunc("/", withMetrics(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Millisecond * 100)
		w.Write([]byte("Hello, World!"))
	}))

	http.HandleFunc("/metrics", withMetrics(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(metrics)
	}))

	fmt.Println("Server is up and running at http://localhost:5000")

	log.Fatal(http.ListenAndServe("localhost:5000",nil))
}
