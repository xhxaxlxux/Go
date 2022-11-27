package weather

import (
	"io/ioutil"
	"net/http"

	"github.com/antonholmquist/jason"
)

func GpsInfo() (string, string) {
	url := "https://get.geojs.io/v1/ip/geo.json"

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	str := string(byteArray)

	v, err := jason.NewObjectFromBytes([]byte(str))
	if err != nil {
		panic(err)
	}

	latitude, _ := v.GetString("latitude")
	longitude, _ := v.GetString("longitude")

	return latitude, longitude
}
