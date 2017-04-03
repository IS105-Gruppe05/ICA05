package main

import (
	"html/template"
	"net/http"
	"path"
	"encoding/json"
	"log"
)

type GeneratedWeather struct {
	Location struct {
		Name string `json:"name"`
		Region string `json:"region"`
		Country string `json:"country"`
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
		TzID string `json:"tz_id"`
		LocaltimeEpoch int `json:"localtime_epoch"`
		Localtime string `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int `json:"last_updated_epoch"`
		LastUpdated string `json:"last_updated"`
		TempC float64 `json:"temp_c"`
		TempF float64 `json:"temp_f"`
		IsDay int `json:"is_day"`
		Condition struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int `json:"code"`
		} `json:"condition"`
		WindMph float64 `json:"wind_mph"`
		WindKph float64 `json:"wind_kph"`
		WindDegree int `json:"wind_degree"`
		WindDir string `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm float64 `json:"precip_mm"`
		PrecipIn float64 `json:"precip_in"`
		Humidity int `json:"humidity"`
		Cloud int `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		VisKm float64 `json:"vis_km"`
		VisMiles float64 `json:"vis_miles"`
	} `json:"current"`
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8001", nil)

}

func foo(w http.ResponseWriter, r *http.Request) {


	url := "http://api.apixu.com/v1/current.json?key=3cc8e1ca15b24b72a60121151172803&q=London"
	//  fmt.Printf("HTML code of %s ...\n", url)
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var w1 GeneratedWeather
	err = json.NewDecoder(resp.Body).Decode(&w1)
	if err != nil {
		log.Fatal(err)
	}

		fp := path.Join("templates", "index.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w, w1); err != nil { //været og ikke w´en til høyre for den orginale)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}



}