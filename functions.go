// Package p contains an HTTP Cloud Function.
package p

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/cdipaolo/sentiment"
)

func getEmoji(sentiment uint8) string {
	type emojiList = []string
	type emojiSentimentMap = map[uint8]emojiList

	PosEmoList := emojiList{"🎁", "😙", "💞", "💃", "🎊", "🏆", "☺", "🐾", "😋", "😛", "🌸", "🐱", "😃", "🍜", "💪"}
	NegEmoList := emojiList{"👿", "😕", "😐", "😒", "😿", "😦", "😾", "😠", "👺", "😡", "😨", "💩", "😭", "😓", "👹"}

	EmojiSentimentMap := emojiSentimentMap{
		0: NegEmoList,
		1: PosEmoList,
	}
	emojis := EmojiSentimentMap[sentiment]

	randomIndex := rand.Intn(len(emojis))
	randomEmoji := emojis[randomIndex]

	return randomEmoji
}

func getSentiment(text string) *sentiment.Analysis {
	model, err := sentiment.Restore()
	if err != nil {

		panic(fmt.Sprintf("Could not restore model!\n\t%v\n", err))
	}

	analysis := model.SentimentAnalysis(text, sentiment.English)
	return analysis
}

func GetSentimentAnalysis(w http.ResponseWriter, r *http.Request) {
	type ReqBody struct {
		Text string `json:"text"`
	}

	var request ReqBody

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	analysis := getSentiment(request.Text)
	score := analysis.Score
	emoji := getEmoji(analysis.Score)

	type ResBody struct {
		Emoji string `json:"emoji"`
		Score uint8  `json:"score"`
	}

	resBody := ResBody{Score: score, Emoji: emoji}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resBody)
}
