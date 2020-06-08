package handler

import (
	"encoding/json"
	"net/http"

	weather "github.com/yuens1002/server-graphql-go-weather/weather"
)

// CityState is the format of the body request
type CityState struct {
	City  string `json:"city"`
	State string `json:"state"`
}

// Payload to send down the pipe
type Payload struct {
	Weather weather.Currently `json:"weather"`
	Address string            `json:"formatted_address"`
}

// ServeHTTP configures city state inputs
func (cs *CityState) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Parse json request body and use it to set fields on user
	// Note that user is passed as a pointer variable so that it's fields can be modified
	if err := json.NewDecoder(r.Body).Decode(cs); err != nil {
		panic(err)
	}

	lga := weather.GetLatLngAddress(cs.City + "+" + cs.State)
	address := lga.Address
	currently := weather.GetWeather(lga.Geometry.Location)
	payload := Payload{currently, address}

	// Marshal or convert user object back to json and write to response
	cJSON, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	// Set Content-Type header so that clients will know how to read response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//Write json response back to response
	w.Write(cJSON)
}
