package main

import (
	"./jsonurl"
	"./jsonurlv2"
	"./jsonurlv3"
	"./jsonurlv4"
	"./jsonurlv5"
)

func main() {
	jsonurl.JsonUrl()
	go jsonurlv2.JsonUrlV2()
	go jsonurlv3.JsonUrlV3()
	go jsonurlv4.JsonUrlV4()
	jsonurlv5.JsonUrlV5()
}
