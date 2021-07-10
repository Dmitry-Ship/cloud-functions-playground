// Package p contains an HTTP Cloud Function.
package functions

import (
	"encoding/json"
	"net/http"
)

func GetSentimentAnalysis(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	type ReqBody struct {
		Text string `json:"text"`
	}

	var request ReqBody

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	score := GetSentimentScore(request.Text)
	emoji := GetEmoji(score)

	type ResBody struct {
		Emoji       string `json:"emoji"`
		Score       uint8  `json:"score"`
		Description string `json:"description"`
	}

	resBody := ResBody{Score: score, Emoji: emoji, Description: GetSentimentDesc(score)}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resBody)
}
