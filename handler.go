package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/TechProber/pro-bot/method"
	"github.com/TechProber/pro-bot/model"
)

func Handler(res http.ResponseWriter, req *http.Request) {
	body := &model.WebhookReqBody{} // decode the JSON response body
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		log.Println("could not decode request body", err)
		return
	}

	if !strings.Contains(strings.ToLower(body.Message.Text), "morning") {
		return
	}

	if err := method.Hello(body.Message.Chat.ID, body.Message.Text); err != nil {
		log.Println("error in sending reply:", err)
		return
	}

	log.Println("reply sent")
}

func main() {
	http.ListenAndServe(":3000", http.HandlerFunc(Handler))
}
