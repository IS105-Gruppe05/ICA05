package main

import (
	"net/http"
	"./"
)

func main() {
	http.HandleFunc("/", apuxiweather.Foo)
	http.ListenAndServe(":8000", nil)
}
