package jsonurlv2

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	//"io/ioutil"
	"log"
	//"net/http"
	"strings"
)

func JsonUrlV2() {

	url := "http://api.apixu.com/v1/current.json?key=3cc8e1ca15b24b72a60121151172803&q=Paris"
	fmt.Printf("HTML code of %s \n", url)
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// show the HTML code as a string %s
	//fmt.Printf("%s\n\n", html)

	var jsonStream = string(html[:])

	type Location struct {
		Name string
		//Region []string
		Lat float64
		Lon float64
	}
	type Measurements struct {
		Temp_c   float64
		Humidity float64
		Wind_kph float64
	}
	type Weather struct {
		Location Location
		Current  Measurements
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var w Weather
		if err := dec.Decode(&w); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Location: %s\n", w.Location.Name)
		fmt.Printf("Coordinates are: longitude %.2f and latitude %.2f\n", w.Location.Lon, w.Location.Lat)
		fmt.Printf("Temperature Celcius: %.2f\n", w.Current.Temp_c)
		fmt.Printf("Humidity: %.f\n", w.Current.Humidity)
		fmt.Printf("Wind kph: %.2f\n\n", w.Current.Wind_kph)
	}
}
