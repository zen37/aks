package main

import (
	"fmt"

	"time"

	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/go-redis/redis"

	"github.com/paulbellamy/ratecounter"
)

// album represents data about a record album.
type intro struct {
	Prefix    string `json:"prefix"`
	Timestamp string `json:"timestamp"`
}

// intros slice to seed record intro data.
var intros = []intro{
	{Prefix: "Hello Word", Timestamp: "Dec 30, 2021 07:30:58 AM"},
}

// This getIntros function creates JSON from the slice of album structs, writing the JSON into the response.

// getIntros responds with the list of all albums as JSON.
func getIntros(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, intros)
}

// Record requests-per-minute
var counter_perMin = ratecounter.NewRateCounter(60 * time.Second)

// postIntros adds an intro from JSON received in the request body.
func postIntros(c *gin.Context) {
	var newIntro intro

	// Web app called and Increment our rates counter
	counter_perMin.Incr(1)

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newIntro); err != nil {
		return
	}

	// Add the new intro to the slice.
	// intros = append(intros, newIntro)
	c.IndentedJSON(http.StatusCreated, newIntro)

	// Using redis
	// Client for storage1
	storage01_client := redis.NewClient(&redis.Options{
		// Addr: "localhost:6379",
		Addr:     "gogin-redis-service:6379",
		Password: "",
		DB:       0,
	})

	// Client for storage2
	storage02_client := redis.NewClient(&redis.Options{
		// Addr: "localhost:6379",
		Addr:     "gogin-redis-service:6379",
		Password: "",
		DB:       1,
	})

	// Determine the db client to use based on call rates per minute

	// get the current requests-per-minute
	call_rate := counter_perMin.Rate()
	fmt.Println("Request per min:", call_rate)

	// Using 200 calls/min as yardstick

	var db_client *redis.Client
	if call_rate < 10 {
		db_client = storage01_client
	} else if call_rate > 10 {
		db_client = storage02_client
	} else {
		db_client = storage01_client
	}

	// Test the connection
	pong, err := db_client.Ping().Result()
	fmt.Println(pong, err)

	// Check the value
	fmt.Println(newIntro)

	// Set value
	// we can call set with a `Key` and a `Value`.
	// I am using the timestamp as key here, you can use whatever key u want

	// Using Timestamp as keys in redis db
	// Use time.Now and one of time.Unix or time.UnixNano to get a timestamp.
	now := time.Now()          // current local time
	secTimestamp := now.Unix() // number of seconds since January 1, 1970 UTC
	// nsecTimestamp := now.UnixNano() // number of nanoseconds since January 1, 1970 UTC
	// fmt.Println("Timestamp (Secs):", secTimestamp)
	// fmt.Println("Timestamp (NanoSecs):", nsecTimestamp)

	// err = db_client.Set(secTimestamp, newIntro.Prefix+" "+newIntro.Timestamp, 0).Err()
	err = db_client.Set(strconv.FormatInt(secTimestamp, 10), newIntro.Prefix+" "+newIntro.Timestamp, 0).Err()
	// if there has been an error setting the value
	// handle the error
	if err != nil {
		fmt.Println(err)
	}

}

// This sets up an association in which getAlbums handles requests to the /albums endpoint path.
func main() {
	router := gin.Default()
	router.GET("/intros", getIntros)
	router.POST("/intros", postIntros)

	router.Run("0.0.0.0:8080")

}
