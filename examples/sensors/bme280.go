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
	tempFed := aio.Feed{Name: "Temperature"}
	humidFed := aio.Feed{Name: "Humidity"}
	pressFed := aio.Feed{Name: "Pressure"}
	//create the feeds
	temp, _, err := client.Feed.Create(&tempFed)
	if err != nil {
		log.Println(err)
	}
	humid, _, err := client.Feed.Create(&humidFed)
	if err != nil {
		log.Println(err)
	}
	press, _, err := client.Feed.Create(&pressFed)
	if err != nil {
		log.Println(err)
	}
	//update the data
	client.SetFeed(temp)
	_, _, err = client.Data.Create(&aio.Data{ID: "0", Value: "76"})
	if err != nil {
		log.Println(err)
	}
	client.SetFeed(humid)
	_, _, err = client.Data.Create(&aio.Data{ID: "1", Value: "53"})
	if err != nil {
		log.Println(err)
	}
	client.SetFeed(press)
	_, _, err = client.Data.Create(&aio.Data{ID: "2", Value: "1125"})
	if err != nil {
		log.Println(err)
	}

}
