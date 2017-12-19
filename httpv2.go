package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080

	http.HandleFunc("/jsonapi", indexFunction)
	log.Printf("running in port 8080")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}

type JSONTest struct {
	// Build Constructor for JSONTest
	Message string `json:"pesan"`
	Name    string `json:"nama"`
	Address string `json:"alamat"`
	Gender  string `json:"jenis_kelamin"`
}

func indexFunction(w http.ResponseWriter, r *http.Request) {
	// insert JSONTest to message
	message := JSONTest{
		Message: "hello guys!",
		Name:    "Rangga Adhitya",
		Address: "Jl. Kuala",
		Gender:  "Pria",
	}

	// insert data and error if has
	data, err := json.Marshal(message)
	if err != nil {
		// return panic error
		panic("Opps.. error")
	}

	// return response data in JSON
	fmt.Fprintf(w, string(data))
}
