package methods

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

func SayHi(chatID int64) error {
	reqBody := &MessageReqBody{
		ChatID: chatID,
		Text:   "Nihao!!",
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	// Send a post request with your token
	res, err := http.Post("https://api.telegram.org/bot5239780065:AAH9p8NF3zxEwjGSrrTaQQLS83PJN_fz8eA/sendMessage", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}