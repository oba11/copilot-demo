package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Payload struct for response passed to frontend
type Payload struct {
	Title interface{} `json:"title"`
	URL   interface{} `json:"url"`
}

// GiphyStruct lazy struct for giphy API
type GiphyStruct struct {
	Data map[string]interface{} `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}

var payloads []Payload

func getGiphy(w http.ResponseWriter, r *http.Request) {
	var out GiphyStruct

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	for i := 0; i < 2; i++ {
		resp, err := http.Get("https://api.giphy.com/v1/gifs/random?api_key=0UTRbFtkMxAplrohufYco5IY74U8hOes&tag=fail&rating=g")
		if err != nil {
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
		json.Unmarshal([]byte(body), &out)
		payloads = append(payloads, Payload{Title: out.Data["title"], URL: out.Data["image_original_url"]})
	}

	json.NewEncoder(w).Encode(payloads)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/giphy", getGiphy).Methods("GET")
	r.HandleFunc("/", indexHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":80", r))
}
