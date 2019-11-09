// Package yelp provides a lightweight wrapper around the Yelp REST API.  It supports authentication with
// OAuth 1.0, the Search API, and the Business API.
package yelp

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

const (
	rootURI      = "https://api.yelp.com/"
	businessArea = "/v3/businesses"
	searchArea   = "/v3/businesses/search"
)

var (
	errUnspecifiedLocation = errors.New("location must be specified")
	errBusinessNotFound    = errors.New("business not found")
)

// AuthOptions provide keys required for using the Yelp API.  Find more
// information here:  https://www.yelp.com/developers/documentation/v3/authentication
type AuthOptions struct {
	YelpAPIKey string // API key from the yelp fusion API access site
}

// Client manages all searches.  All searches are performed from an instance of a client.
// It is the top level object used to perform a search or business query.  Client objects
// should be created through the createClient API.
type Client struct {
	Options *AuthOptions
	Client  *http.Client
}

// DoSimpleSearch performs a simple search with a term and location.
func (client *Client) DoSimpleSearch(term, location string) (result SearchResult, err error) {

	// verify the term and location are not empty
	if location == "" {
		return SearchResult{}, errUnspecifiedLocation
	}

	// set up the query options
	params := map[string]string{
		"term":     term,
		"location": location,
	}

	// perform the search request
	_, err = client.makeRequest(searchArea, "", params, &result)
	if err != nil {
		return SearchResult{}, err
	}
	return result, nil
}

// DoSearch performs a complex search with full search options.
func (client *Client) DoSearch(options SearchOptions) (result SearchResult, err error) {

	// get the options from the search provider
	params, err := options.getParameters()
	if err != nil {
		return SearchResult{}, err
	}

	// perform the search request
	_, err = client.makeRequest(searchArea, "", params, &result)
	if err != nil {
		return SearchResult{}, err
	}

	return result, nil
}

// GetBusiness obtains a single business by name.
func (client *Client) GetBusiness(name string) (result Business, err error) {
	statusCode, err := client.makeRequest(businessArea, name, nil, &result)
	if err != nil {
		// At some point the Yelp API stopped reporting 404s for missing business names, and
		// started reporting 400s :(
		if statusCode == 400 {
			return Business{}, errBusinessNotFound
		}
		return Business{}, err
	}
	return result, nil
}

// makeRequest is an internal/private API used to make underlying requests to the Yelp API.
func (client *Client) makeRequest(area string, id string, params map[string]string, v interface{}) (statusCode int, err error) {

	// get the base url
	queryURI, err := url.Parse(rootURI)
	if err != nil {
		return 0, err
	}

	// add the type of request we're making (search|business)
	queryURI.Path = area

	// if an id is provided, add it too
	if id != "" {
		queryURI.Path += "/" + id
	}

	// add query params
	if len(params) > 0 {
		queryString := queryURI.Query()
		for key, val := range params {
			queryString.Set(key, val)
		}
		queryURI.RawQuery = queryString.Encode()
	}

	// create a new request using http
	request, err := http.NewRequest("GET", queryURI.String(), nil)

	// add Authorization header to request
	bearer := "Bearer " + client.Options.YelpAPIKey
	request.Header.Add("Authorization", bearer)

	// execute the request using http Client
	response, err := client.Client.Do(request)

	if err != nil {
		if response != nil {
			return response.StatusCode, err
		}
		return 500, err
	}

	// close the request when done
	defer response.Body.Close()

	// ensure the request returned a 200
	if response.StatusCode != 200 {
		return response.StatusCode, errors.New(response.Status)
	}

	err = json.NewDecoder(response.Body).Decode(v)
	return response.StatusCode, err
}

// New will create a new yelp search client.  All search operations should go through this API.
func New(options *AuthOptions, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		Options: options,
		Client:  httpClient,
	}
}
