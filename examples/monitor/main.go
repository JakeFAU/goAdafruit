package main

import (
	"net/url"
	"os"
	"strconv"
	"time"

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
	// get the three feeds
	client := connect()
	//update the feeds in a loop
	for {
		// Get the data
		readings := bme.BME280ReadValues()
		// update the feeds
		tData := aio.Data{FeedKey: "temperature", Value: strconv.FormatFloat(float64(readings.Temperature), 'e', 2, 32)}
		client.Data.Create(&tData)
		hData := aio.Data{FeedKey: "humidity", Value: strconv.FormatFloat(float64(readings.Humidity), 'e', 2, 32)}
		client.Data.Create(&hData)
		pData := aio.Data{FeedKey: "pressure", Value: strconv.FormatFloat(float64(readings.Pressure), 'e', 2, 32)}
		client.Data.Create(&pData)
		//sleep so we don't overwhelm adafruit io
		time.Sleep(time.Second * 20)
	}

}
