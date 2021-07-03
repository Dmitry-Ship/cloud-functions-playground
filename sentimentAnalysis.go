package functions

import (
	"fmt"
	"math/rand"

	"github.com/cdipaolo/sentiment"
)

func GetEmoji(sentiment uint8) string {
	type emojiList = []string

	PosEmoList := emojiList{"🎁", "😙", "💞", "💃", "🎊", "🏆", "☺", "🐾", "😋", "😛", "🌸", "🐱", "😃", "🍜", "💪"}
	NegEmoList := emojiList{"👿", "😕", "😐", "😒", "😿", "😦", "😾", "😠", "👺", "😡", "😨", "💩", "😭", "😓", "👹"}

	type emojiSentimentMap = map[uint8]emojiList

	EmojiSentimentMap := emojiSentimentMap{
		0: NegEmoList,
		1: PosEmoList,
	}

	emojis := EmojiSentimentMap[sentiment]
	randomIndex := rand.Intn(len(emojis))

	return emojis[randomIndex]
}

func GetSentimentScore(text string) uint8 {
	model, err := sentiment.Restore()
	if err != nil {
		panic(fmt.Sprintf("Could not restore model!\n\t%v\n", err))
	}

	analysis := model.SentimentAnalysis(text, sentiment.English)
	return analysis.Score
}
