package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	port := 8080

	http.HandleFunc("/decode", unMarshalJSON)
	log.Printf("serving in port: 8080")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

type JSONresponse struct {
	Message string `json:"message"`
}

type Requests struct {
	Method string
	Name   string
}

func unMarshalJSON(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	var request Requests
	err = json.Unmarshal(body, &request)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	response := JSONresponse{Message: "Hello " + request.Name}
	encoder := json.NewEncoder(w)

	encoder.Encode(response)
}
