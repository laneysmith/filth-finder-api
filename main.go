package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-yelp-v3/yelp"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/guregu/null"
)

func custom(w http.ResponseWriter, r *http.Request) {
	options, err := getCredentials(w)
	if err != nil {
		fmt.Println(err)
	}

	client := yelp.New(options, nil)

	term := r.URL.Query().Get("term")
	longitude := r.URL.Query().Get("longitude")
	latitude := r.URL.Query().Get("latitude")

	// build SearchOptions from incoming query params
	latitudeFloat, err := strconv.ParseFloat(latitude, 64)
	longitudeFloat, err := strconv.ParseFloat(longitude, 64)
	generalOptions := yelp.GeneralOptions{
		Term: term,
	}
	coordinateOptions := yelp.CoordinateOptions{
		Latitude:  null.FloatFrom(latitudeFloat),
		Longitude: null.FloatFrom(longitudeFloat),
	}
	locationOptions := yelp.LocationOptions{
		CoordinateOptions: &coordinateOptions,
	}
	searchOptions := yelp.SearchOptions{
		LocationOptions: &locationOptions,
		GeneralOptions:  &generalOptions,
	}
	results, err := client.DoSearch(searchOptions)
	if err != nil {
		fmt.Println("error:", err)
	}

	bytes, err := json.Marshal(results)
	if err != nil {
		fmt.Println("error:", err)
	}

	io.WriteString(w, string(bytes))
}

func main() {
	http.HandleFunc("/custom", custom)
	// TODO: add simple search option
	// http.HandleFunc("/simple", simple)
	http.ListenAndServe(":8000", nil)
}

func getCredentials(w http.ResponseWriter) (options *yelp.AuthOptions, err error) {
	c := &yelp.AuthOptions{
		YelpAPIKey: os.Getenv("YELP_API_KEY"),
	}

	if c.YelpAPIKey == "" {
		return nil, errors.New("Missing env variable for YELP_API_KEY")
	}

	return c, nil
}
