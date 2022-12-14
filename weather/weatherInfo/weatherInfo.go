package weather

import (
	"io/ioutil"
	"net/http"

	"github.com/antonholmquist/jason"
)

func WeatherInfo() (string, string, float64) {
	baseUrl := "https://api.openweathermap.org/data/2.5/weather?"
	lat := "lat="
	lon := "&lon="
	lang := "&lang=ja"
	apiKey := ""
	units := "&units=metric" // 摂氏指定

	latitude, longitude := GpsInfo()
	lat += latitude
	lon += longitude
	url := baseUrl + lat + lon + lang + apiKey + units

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	str := string(byteArray)

	v, err := jason.NewObjectFromBytes([]byte(str))
	if err != nil {
		panic(err)
	}

	name, _ := v.GetString("name")
	weatherArray, _ := v.GetObjectArray("weather")
	weather, _ := weatherArray[0].GetString("description")
	tempObject, _ := v.GetObject("main")
	temp, _ := tempObject.GetFloat64("temp")

	return name, weather, temp
}
