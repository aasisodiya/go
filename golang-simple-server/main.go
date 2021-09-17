package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Sample struct {
	Property1 string `json:"property1"`
}

func GetSample(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Get Sample API Called")
	w.Header().Set("Content-Type", "application/json")
	data := Sample{
		Property1: "property one",
	}
	jsonResponse, err := json.Marshal(data)
	if err != nil {
		// handle error
	}
	w.Write(jsonResponse)
	// fmt.Fprintf(w, "{\"property1\": \"property one\"}")
}

func main() {

	http.HandleFunc("/get-sample", GetSample)

	http.ListenAndServe(":8090", nil)
}
