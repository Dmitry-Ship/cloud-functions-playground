// Package p contains an HTTP Cloud Function.
package p

import (
	"net/http"

	"example.com/cloudfunction/common"
)

func GetMessage(w http.ResponseWriter, r *http.Request) {
	result1 := make(chan interface{})
	result2 := make(chan interface{})

	changelogs := Changelogs{}

	go common.MakeRequest("https://manychat.com/changelog/get", &changelogs, result1)
	go common.MakeRequest("https://jsonplaceholder.typicode.com/todos/1", &changelogs, result2)

	select {
	case msg1 := <-result1:
		common.SendJSONresponse(msg1, w)

	case msg2 := <-result2:
		common.SendJSONresponse(msg2, w)
	}
}
