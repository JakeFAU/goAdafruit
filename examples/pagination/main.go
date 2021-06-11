package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"

	aio "github.com/jakefau/goAdafruit"
)

var (
	baseURL  string
	key      string
	username string
)

func connect() aio.Client {
	// basic stuff
	username := "JakeFau"
	baseURL := "https://io.adafruit.com/"
	// Hide your key
	key := os.Getenv("ADAFRUIT_IO_KEY")
	// get a client
	client := aio.NewClient(key, username)
	client.BaseURL, _ = url.Parse(baseURL)
	return *client
}

func GetDataForFeed() {
	client := connect()
	feed, _, err := client.Feed.Get("weather.humidity")
	if err != nil {
		log.Fatal(err)
	}
	client.SetFeed(feed)
	// choose the key from one of your feeds
	data, _, err := client.Data.All(nil)
	dataJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(dataJSON))
}

func main() {
	GetDataForFeed()
}
