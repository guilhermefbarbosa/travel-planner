package main

import "github.com/guilhermefbarbosa/travel-planner/travel-planner-go/src/http"

func main() {

	http.NewServer(http.Config{
		Port:             5001,
		AppName:          "Travel Planner",
		TimeoutInSeconds: 10,
	}).Start()

}
