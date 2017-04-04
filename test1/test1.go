package apuxiweather

import(
	"html/template"
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	// "os"
	// "runtime"
	"path"
)

type Location struct {
	Name   string `json:"Name"`
	Region string `json:"Region"`
	Country string `json:"Country"`
	Lat float64 `json:"Lat"`
	Lon float64 `json:"Lon"`
}
type Current struct {
	Temp_c   float64 `json:"Temp_c"`
	Humidity float64 `json:"Humidity"`
	Wind_kph float64 `json:"Wind_kph"`
}
type Weather struct {
	Loc Location
	Cur  Current
}

type API struct {

}

func (r API) GetWeather() (*Weather, error) {

	body, err := MakeRequest("http://api.apixu.com/v1/current.json?key=3cc8e1ca15b24b72a60121151172803&q=London")
	if err != nil {
		return nil, err
	}
	s, err := ParseWeather(body)
	return  s, err
}


func MakeRequest(url string) ([]byte, error) {

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return []byte(body), err
}


func ParseWeather(body []byte) (*Weather, error) {
	var w = new(Weather)
	err := json.Unmarshal(body, &w)
	if err != nil {
		fmt.Println("error:", err)
	}
	return w, err
}

func Foo(w http.ResponseWriter, r *http.Request) {

fp := path.Join("templates", "index.html")
tmpl, err := template.ParseFiles(fp)
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}

if err := tmpl.Execute(w, Weather{}); err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
}
}