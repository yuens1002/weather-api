package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

// Denvor
// 40.0149856,-105.2705456

// Weather is the returned payload
type Weather struct {
	Currently `json:"currently"`
}

//Currently is the current forecast from darksky
type Currently struct {
	Temperature       float64 `json:"temperature"`
	Humidity          float64 `json:"humidity"`
	PrecipProbability float64 `json:"precipProbability"`
	Summary           string  `json:"summary"`
}

// GetWeather fetches weather from DarkSky
func GetWeather(geoData LatLng) Currently {

	lat := fmt.Sprintf("%f", geoData.Lat)
	lng := fmt.Sprintf("%f", geoData.Lng)
	// fmt.Println(lat, lng)
	latlng := lat + "," + lng
	base := viper.GetString("darksky-api.url")
	key := viper.GetString("darksky-api.key")
	url := fmt.Sprintf("%s/%s/%s", base, key, latlng)

	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	w := &Weather{}
	// var obj map[string]interface{}
	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// json.Unmarshal([]byte(body), &obj)
	// f := colorjson.NewFormatter()
	// f.Indent = 2
	// s, _ := f.Marshal(obj)
	// fmt.Println("--------res.body starts-----------")
	// fmt.Println(string(s))
	// fmt.Println("--------res.body ends-----------")
	if err := json.NewDecoder(res.Body).Decode(w); err != nil {
		log.Fatalln(err)
	}
	log.Println(w.Currently)
	return w.Currently
}
