package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/aws/aws-xray-sdk-go/xray"
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

var (
	enableXrayTracing bool
	payloads          []Payload
)

func init() {
	if enable, err := strconv.ParseBool(os.Getenv("ENABLE_XRAY_TRACING")); err == nil {
		enableXrayTracing = enable
	}

	if enableXrayTracing {
		xray.Configure(xray.Config{
			LogLevel: "warn",
		})
	}
}

type giphyHandler struct{}

func (h *giphyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var out GiphyStruct

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	req, err := http.NewRequest(http.MethodGet, "https://api.giphy.com/v1/gifs/random?api_key=0UTRbFtkMxAplrohufYco5IY74U8hOes&tag=fail&rating=g", nil)
	if err != nil {
		return
	}

	var client *http.Client
	if enableXrayTracing {
		client = xray.Client(&http.Client{})
	} else {
		client = &http.Client{}
	}

	for i := 0; i < 2; i++ {
		resp, err := client.Do(req.WithContext(r.Context()))
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

type indexHandler struct{}

func (h *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func main() {

	var giphy http.Handler
	var healthcheck http.Handler
	if enableXrayTracing {
		xraySegmentNamer := xray.NewFixedSegmentNamer("api")
		giphy = xray.Handler(xraySegmentNamer, &giphyHandler{})
		healthcheck = xray.Handler(xraySegmentNamer, &indexHandler{})
	} else {
		giphy = &giphyHandler{}
		healthcheck = &indexHandler{}
	}

	http.Handle("/", healthcheck)
	http.Handle("/giphy", giphy)
	log.Fatal(http.ListenAndServe(":80", nil))
}
