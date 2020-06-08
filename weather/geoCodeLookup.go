package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

/*********** sample data struture from res
{
	"results": [
		{
			...
			formatted_address: "Seattle, WA, USA",
			geometry: {
				location: {
					lat: 0.3234,
					lng: 0.2342,
				}
			}
		}
	]
}
*/

// Geo returns geo coodinates
type Geo struct {
	Results []GeometryAddress `json:"results"`
}

// GeometryAddress returns a set of lat lng & city/state
type GeometryAddress struct {
	Geometry Location `json:"geometry"`
	Address  string   `json:"formatted_address"`
}

// Location embeds lat, lng properties
type Location struct {
	Location LatLng `json:"location"`
}

// LatLng lat, lng properties
type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// GetLatLngAddress fetchs geo-codes (address) from the google maps geocode endpt
func GetLatLngAddress(cityState string) GeometryAddress {
	base := viper.GetString("g-maps-api.url")
	key := viper.GetString("g-maps-api.key")
	options := "components=locality"
	country := "US"
	url := fmt.Sprintf("%s/json?%s:%s|country:%s&key=%s", base, options, cityState, country, key)

	return makeGeoRequest(url)

}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func makeGeoRequest(url string) GeometryAddress {
	res, err := http.Get(url)
	checkErr(err)

	defer res.Body.Close()

	w := &Geo{}
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

	err = json.NewDecoder(res.Body).Decode(w)
	checkErr(err)

	return w.Results[0]
}
