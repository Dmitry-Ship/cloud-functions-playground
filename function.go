// Package p contains an HTTP Cloud Function.
package p

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

func GetMessage(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Message string `json:"message"`
	}

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "no message provided!")
		return
	}

	if d.Message == "" {
		fmt.Fprint(w, "empty message provided!")
		return
	}
	fmt.Fprint(w, html.EscapeString(d.Message))
}
