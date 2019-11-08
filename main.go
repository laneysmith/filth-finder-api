package main

import (
	"errors"
	"fmt"
	"go-yelp-v3/yelp"
	"io"
	"net/http"
	"os"
)

func res(w http.ResponseWriter, r *http.Request) {
	options, err := getCredentials(w)
	if err != nil {
		fmt.Println(err)
	}

	client := yelp.New(options, nil)

	term := r.URL.Query().Get("term")
	longitude := r.URL.Query().Get("longitude")
	latitude := r.URL.Query().Get("latitude")

	// TODO: add coords
	coordinateOptions := yelp.CoordinateOptions{
		// Latitude:  null.FloatFrom(latitude),
		// Longitude: null.FloatFrom(longitude),
	}
	locationOptions := yelp.LocationOptions{
		CoordinateOptions: &coordinateOptions,
	}
	searchOptions := yelp.SearchOptions{
		LocationOptions: &locationOptions,
	}
	results, err := client.DoSearch(searchOptions)
	if err != nil {
		fmt.Println(err)
	}

	io.WriteString(w, fmt.Sprintf("Found a total of %v results for \"%v\" at longitude \"%v\" and latitude \"%v\"", results.Total, term, longitude, latitude))
	for i := 0; i < len(results.Businesses); i++ {
		io.WriteString(w, fmt.Sprintf("<div>%v, %v</div>", results.Businesses[i].Name, results.Businesses[i].Rating))
	}
}

func main() {
	http.HandleFunc("/", res)
	http.ListenAndServe(":8000", nil)
}

func getCredentials(w http.ResponseWriter) (options *yelp.AuthOptions, err error) {
	c := &yelp.AuthOptions{
		YelpClientID: os.Getenv("YELP_CLIENT_ID"),
		YelpAPIKey:   os.Getenv("YELP_API_KEY"),
	}

	if c.YelpClientID == "" || c.YelpAPIKey == "" {
		return nil, errors.New("Missing env variables for YELP_CLIENT_ID and/or YELP_API_KEY")
	}

	return c, nil
}
