package main

import (
	"log"
	"net/http"
	"os"
)

type webHookReqBody struct {
	Message message `json:"message"`
}

type message struct {
	Text string `json:"text"`
	Chat chat   `json:"chat"`
}

type chat struct {
	ID int64 `json:"id"`
}

type reply struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

var token = os.Getenv("token")
var url = "https://api.telegram.org/bot" + token + "/"

func main() {
	port := os.Getenv("PORT")
	err := http.ListenAndServe(":"+port, http.HandlerFunc(handler))
	log.Println(err)
}
