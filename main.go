package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-yelp-v3/yelp"
	"net/http"
	"os"
	"strconv"

	"github.com/guregu/null"
)

func main() {
	port := getPort()

	http.HandleFunc("/custom", custom)
	http.ListenAndServe(":"+port, nil)
}

func custom(w http.ResponseWriter, r *http.Request) {
	options, err := getAPICredentials(w)
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

	// perform search
	results, err := client.DoSearch(searchOptions)
	if err != nil {
		fmt.Println("error:", err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

func getAPICredentials(w http.ResponseWriter) (options *yelp.AuthOptions, err error) {
	o := &yelp.AuthOptions{
		YelpAPIKey: os.Getenv("YELP_API_KEY"),
	}

	if o.YelpAPIKey == "" {
		return nil, errors.New("Missing environment variable for YELP_API_KEY")
	}

	return o, nil
}

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		fmt.Println("No PORT environment variable detected, defaulting to " + port)
	}

	return port
}
