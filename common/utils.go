package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func SendJSONresponse(response interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func MakeRequest(url string, respInterface interface{}, channel chan interface{}) {
	resp, fetchErr := http.Get(url)

	if fetchErr != nil {
		log.Fatal("failed to fetch")
	}

	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		log.Fatal("failed to read response")
	}

	jsonResult := respInterface

	jsonErr := json.Unmarshal(body, &jsonResult)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	channel <- jsonResult

}
