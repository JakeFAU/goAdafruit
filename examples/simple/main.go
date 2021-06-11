package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

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

func ListAllGroups() {
	client := connect()
	groups, _, err := client.Group.All()
	if err != nil {
		log.Fatal(err)
	}
	for _, feed := range groups {
		feedJSON, err := json.MarshalIndent(feed, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(string(feedJSON))
	}
}

func ListAllFeeds() {
	client := connect()
	feeds, _, err := client.Feed.All()
	if err != nil {
		log.Fatal(err)
	}
	for _, feed := range feeds {
		feedJSON, err := json.MarshalIndent(feed, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(string(feedJSON))
	}
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
	ListAllGroups()
	time.Sleep(5 * time.Second)
	ListAllFeeds()
	time.Sleep(5 * time.Second)
	GetDataForFeed()
}
