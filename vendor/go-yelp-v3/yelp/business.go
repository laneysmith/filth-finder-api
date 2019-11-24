package yelp

// A SearchResult is returned from the Search API. It includes
// the region, the total number of results, and a list of matching businesses.
// The business objects returned by this query are shallow - they will not include
// deep results such as reviews.
type SearchResult struct {
	Region     Region     `json:"region"`     // Suggested bounds in a map to display results in
	Total      int        `json:"total"`      // Total number of business results
	Businesses []Business `json:"businesses"` // The list of business entries (see Business)
}

// Region provides the location of a business obtained from search.
type Region struct {
	Center Coordinates `json:"center"` // Center position of map bounds
}

// Deal defines a set of special offerings from the business.
type Deal struct {
	ID                     string       // Deal identifier
	Title                  string       // Deal title
	URL                    string       // Deal URL
	ImageURL               string       `json:"image_URL"`               // Deal image URL
	CurrencyCode           string       `json:"currency_code"`           // ISO_4217 Currency Code
	TimeStart              float32      `json:"time_start"`              // Deal start time (Unix timestamp)
	TimeEnd                float32      `json:"time_end"`                // Deal end time (optional: this field is present only if the Deal ends)
	IsPopular              bool         `json:"is_popular"`              // Whether the Deal is popular (optional: this field is present only if true)
	WhatYouGet             string       `json:"what_you_get"`            // Additional details for the Deal, separated by newlines
	ImportantRestrictions  string       `json:"Important_restrictions"`  // Important restrictions for the Deal, separated by newlines
	AdditionalRestrictions string       `json:"Additional_restrictions"` // Deal additional restrictions
	Options                []DealOption //Deal options

}

// DealOption provides options are optionally included on a deal.
type DealOption struct {
	Title                  string  // Deal option title
	PurchaseURL            string  `json:"Purchase_URL"` // Deal option URL for purchase
	Price                  float32 // Deal option price (in cents)
	FormattedPrice         string  `json:"Formatted_price"`          // Deal option price (formatted, e.g. "$6")
	OriginalPrice          float32 `json:"Original_price"`           // Deal option original price (in cents)
	FormattedOriginalPrice string  `json:"Formatted_original_price"` // Deal option original price (formatted, e.g. "$12")
	IsQuantityLimited      bool    `json:"Is_quantity_limited"`      // Whether the deal option is limited or unlimited
	RemainingCount         float32 `json:"Remaining_count"`          // The remaining deal options available for purchase (optional: this field is only present if the deal is limited)
}

// GiftCertificate defines optional data available on Businesses.
type GiftCertificate struct {
	ID             string                   `json:"id"`              // Gift certificate identifier
	URL            string                   `json:"url"`             // Gift certificate landing page URL
	ImageURL       string                   `json:"image_url"`       //	Gift certificate image URL
	CurrencyCode   string                   `json:"currency_code"`   // ISO_4217 Currency Code
	UnusedBalances string                   `json:"unused_balances"` // Whether unused balances are returned as cash or store credit
	Options        []GiftCertificateOptions `json:"options"`         //	Gift certificate options
}

// GiftCertificateOptions can define a set of pricing options for a gift certificate.
type GiftCertificateOptions struct {
	Price          float32 `json:"price"`           //	Gift certificate option price (in cents)
	FormattedPrice string  `json:"formatted_price"` //	Gift certificate option price (formatted, e.g. "$50")
}

// Review data contains a list of user reviews for a given Business (when queried using the Business API).
type Review struct {
	ID                  string  `json:"id"`                     // Review identifier
	Rating              float32 `json:"rating"`                 // Rating from 1-5
	RatingImageURL      string  `json:"rating_image_url"`       // URL to star rating image for this business (size = 84x17)
	RatingImageSmallURL string  `json:"rating_image_small_url"` // URL to small version of rating image for this business (size = 50x10)
	RatingImageLargeURL string  `json:"Rating_image_large_url"` // URL to large version of rating image for this business (size = 166x30)
	Excerpt             string  `json:"excerpt"`                // Review excerpt
	TimeCreated         float32 `json:"Time_created"`           // Time created (Unix timestamp)
	User                User    `json:"user"`                   // User who wrote the review
}

// User data is linked off of reviews.
type User struct {
	ID       string `json:"id"`        // User identifier
	ImageURL string `json:"image_url"` // User profile image URL
	Name     string `json:"name"`      // User name
}

// Coordinate data is used with location information.
type Coordinate struct {
	Latitude  float32 `json:"latitude"`  // Latitude of current location
	Longitude float32 `json:"longitude"` // Longitude of current location
}

// Coordinates data is used with location information.
type Coordinates struct {
	Latitude  float32 `json:"latitude"`  // Latitude of current location
	Longitude float32 `json:"longitude"` // Longitude of current location
}

// Location information defines the location of a given business.
type Location struct {
	Address1       string   `json:"address1"`
	Address2       string   `json:"address2"`
	Address3       string   `json:"address3"`
	City           string   `json:"city"`
	Country        string   `json:"county"`   // ISO 3166-1 country code for this business
	State          string   `json:"state"`    // ISO 3166-2 state code for this business
	ZipCode        string   `json:"zip_code"` // Postal code for this business
	CrossStreets   string   `json:"cross_streets"`
	DisplayAddress []string `json:"display_address"` // Display address for the business.
	Neighborhoods  []string `json:"neighborhoods"`   // List that provides neighborhood(s) information for business
}

// Category details
type Category struct {
	Alias string
	Title string
}

// Business information is returned in full from the business API, or shallow from the search API.
type Business struct {
	Alias        string      `json:"alias"`      // Alias of this business
	Categories   []Category  `json:"categories"` // Provides a list of category name, alias pairs that this business is associated with. The alias is provided so you can search with the category_filter.
	Coordinates  Coordinates `json:"coordinates"`
	DisplayPhone string      `json:"display_phone"` // Phone number for this business formatted for display
	Distance     float32     `json:"distance"`      // Distance that business is from search location in meters, if a latitude/longitude is specified.
	ID           string      `json:"id"`            // Yelp ID for this business
	ImageURL     string      `json:"image_URL"`     // URL of photo for this business
	IsClosed     bool        `json:"is_closed"`     // Whether business has been (permanently) closed
	Location     Location    `json:"location"`      // Location data for this business
	Name         string      `json:"name"`          // Name of this business
	Phone        string      `json:"phone"`         // Phone number for this business with international dialing code (e.g. +442079460000)
	Price        string      `json:"price"`         // Price level of the business. Value is one of $, $$, $$$ and $$$$.
	Rating       float32     `json:"rating"`        // Rating for this business (value ranges from 1, 1.5, ... 4.5, 5)
	ReviewCount  int         `json:"review_count"`  // Number of reviews for this business
	URL          string      `json:"url"`           // URL for business page on Yelp
	Transactions []string    `json:"transactions"`  // List of Yelp transactions that the business is registered for. Current supported values are pickup, delivery and restaurant_reservation.
}

// Open times by day
type Open struct {
	Day         int    `json:"day"`   // From 0 to 6, day of the week from Monday to Sunday. You may get the same day of the week more than once if the business has more than one opening time slots.
	Start       string `json:"start"` // Opening time in 24-hour clock notation
	End         string `json:"end"`   // Closing time in 24-hour clock notation
	IsOvernight string `json:"is_overnight"`
}

// Hours for business
type Hours struct {
	HoursType string `json:"hours_type"` // The type of the opening hours information. Right now, this is always REGULAR.
	IsOpenNow string `json:"Is_open_now"`
	Open      []Open
}

// BusinessDetail is the business object returned from detailed call
type BusinessDetail struct {
	Business
	Hours     Hours `json:"hours "`
	IsClaimed bool  `json:"Is_claimed"` // Whether business has been claimed by a business owner
}
