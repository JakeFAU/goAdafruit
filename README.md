# goAdafruit
Go SDK for Adafruit IO

[![GoDoc](https://pkg.go.dev/github.com/jakefau/goAdafruit?status.svg)](https://pkg.go.dev/github.com/jakefau/goAdafruit)

A go client library for talking to your io.adafruit.com account. This SDK is for V2 of the io.adafruit.com API.  It works basically as the V1 version, just expanded to take advantage of all the features of V2

Requires go version 1.6 or better. Running tests requires the github.com/stretchr/testify library, which can be installed with:

    $ go get github.com/stretchr/testify

## Usage
    import aio "github.com/jakefau/goAdafruit"

Authentication for Adafruit IO is managed by providing your Adafruit IO token in the head of all web requests via the *X-AIO-Key header*. This is handled for you by the client library, which expects you API Token when it is initialized.

To connect to the API you need a user name and a secret key

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

The entire http API is provided by this SDK, to get a sense of what the API can do please see https://io.adafruit.com/api/docs/

### Feeds
Feeds are what contain the individual datapoints.  Before you can get the data points in JSON or in Charting format you have to set the current feed in the client.


    func GetDataForFeed() {
        client := connect()
        feed, _, err := client.Feed.Get("weather.humidity")
        if err != nil {
            log.Fatal(err)
        }
        client.SetFeed(feed)
        // choose the key from one of your feeds
        data, _, err := client.Data.All(nil)

Please let me know what you think and be sure to let me know if there are any bugs you find or enhancements you would like to see.  Feel free to fork the project and add to it, I will monitor pull requests.
