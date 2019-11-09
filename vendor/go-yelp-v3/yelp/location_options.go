package yelp

import (
	"errors"
	"fmt"
)

// LocationOptions enable specifing a Location by Neighborhood, Address, or City.
// The location format is defined: ?location=location
type LocationOptions struct {
	Location          string             // Specifies the combination of "address, neighborhood, city, state or zip, optional country" to be used when searching for businesses. (required)
	CoordinateOptions *CoordinateOptions // An optional latitude, longitude parameter can also be specified as a hint to the geocoder to disambiguate the location text. The format for this is defined as:   ?cll=latitude,longitude
}

// getParameters will reflect over the values of the given
// struct, and provide a type appropriate set of querystring parameters
// that match the defined values.
func (o *LocationOptions) getParameters() (params map[string]string, err error) {
	params = make(map[string]string)

	locationProvided := o.Location != ""
	coordinatesProvided := o.CoordinateOptions != nil &&
		o.CoordinateOptions.Latitude.Valid &&
		o.CoordinateOptions.Longitude.Valid

	if !locationProvided && !coordinatesProvided {
		return params, errors.New("To perform a location based search, you must provide either the coordinates or the location of the area to search.  For coordinate based searches, use the CoordinateOption class.  For location based searches, use the Location class")
	}

	// if location is specified add it to the parameters hash
	if locationProvided {
		params["location"] = o.Location
	}

	// if coordinates are specified add those to the parameters hash
	if coordinatesProvided {
		params["latitude"] = fmt.Sprintf("%v", o.CoordinateOptions.Latitude.Float64)
		params["longitude"] = fmt.Sprintf("%v", o.CoordinateOptions.Longitude.Float64)
	}

	return params, nil
}
