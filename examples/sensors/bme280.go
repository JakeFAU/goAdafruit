package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/davidgs/bme280_go"
	aio "github.com/jakefau/goAdafruit"
)

// connect to the API
func connect() aio.Client {
	// basic stuff
	username := "JakeFau"
	baseURL := "https://io.adafruit.com/"
	// Hide your key
	key := os.Getenv("ADAFRUIT_IO_KEY")
	// get a client
	client := aio.NewClient(key, username)
	// set the base url, aka the host
	client.BaseURL, _ = url.Parse(baseURL)
	return *client
}

func main() {
	// the bme280 Go driver
	bme := bme280_go.BME280{}
	// The bme280 uses i2c to communicate
	dev := "/dev/i2c-1"
	bme.BME280Init(dev, 0x77)
	// Get the data
	readings := bme.BME280ReadValues()
	fmt.Printf("%v, %v, %v", readings.Temperature, readings.Humidity, readings.Pressure)

	//get the client
	client := connect()
	//build the feeds
	tempFed := aio.FeedCreation{Name: "Temperature"}
	humidFed := aio.FeedCreation{Name: "Humidity"}
	pressFed := aio.FeedCreation{Name: "Pressure"}
	//create the feeds
	_, _, err := client.Feed.Create(&tempFed)
	if err != nil {
		log.Println(err)
	}
	_, _, err = client.Feed.Create(&humidFed)
	if err != nil {
		log.Println(err)
	}
	_, _, err = client.Feed.Create(&pressFed)
	if err != nil {
		log.Println(err)
	}
	tempFeed, _, err := client.Feed.Get("Temperature")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	humidFeed, _, err := client.Feed.Get("Humidity")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pressureFeed, _, err := client.Feed.Get("Pressure")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//update the data
	client.SetFeed(tempFeed)
	_, _, err = client.Data.Create(&aio.Data{ID: "0", Value: "76"})
	if err != nil {
		log.Println(err)
	}
	client.SetFeed(humidFeed)
	_, _, err = client.Data.Create(&aio.Data{ID: "1", Value: "53"})
	if err != nil {
		log.Println(err)
	}
	client.SetFeed(pressureFeed)
	_, _, err = client.Data.Create(&aio.Data{ID: "2", Value: "1125"})
	if err != nil {
		log.Println(err)
	}

}
