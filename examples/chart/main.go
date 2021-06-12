package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	aio "github.com/jakefau/goAdafruit"
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

func main() {
	client := connect()
	//options := aio.DataFilter{StartTime: "6/5/2021"}
	//be sure to change the feed to one in your account
	feed, _, err := client.Feed.Get("weather.humidity")
	if err != nil {
		log.Fatal(err)
	}
	client.SetFeed(feed)
	data, _, err := client.Data.GetChartData(nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range data {
		fmt.Printf("%v = %v\n", d.X, d.Y)
	}

}
