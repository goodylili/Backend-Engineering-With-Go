package mirrorFinder

import (
	"encoding/json"
	"fmt"
	"github.com/theghostmac/Backend-Engineering-With-Go/RestAPIs/DebianMirrorsFinder/mirrors"
	"log"
	"net/http"
	"time"
)

type response struct {
	FastestURL string        `json:"fastest_url"` // url for fastest mirror site
	Latency    time.Duration `json:"latency"`     // time taken to download the README from Debian repository
}

func main() {
	http.HandleFunc("/fastest-mirror", func(writer http.ResponseWriter, request *http.Request) {
		response := findFastest(mirrors.MirrorList)
		repsJSON, _ := json.Marshal(response)

		writer.Header().Set("Content-Type", "application/json")
		writer.Write(repsJSON)
	})
	port := ":8080"
	server := &http.Server{
		Addr:           port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Starting server on port %sn", port)
	log.Fatal(server.ListenAndServe())
}

func findFastest(urls [...]string) response {
	urlChannel := make(chan string)
	latencyChannel := make(chan time.Duration)

	for _, url := range urls {
		mirrorURL := url
		go func() {
			start := time.Now()
			_, err := http.Get(mirrorURL + "/README")
			latency := time.Now().Sub(start) / time.Millisecond
			if err == nil {
				urlChannel <- mirrorURL
				latencyChannel <- latency
			}
		}()
	}
	return response{<-urlChannel, <-latencyChannel}
}