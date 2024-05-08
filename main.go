package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot_id = "6927276355:AAFyFficeYeNWEbKaghIu53A1rTIYxbf3iU"
var chat_id int64 = 7188356683

func main() {
    http.HandleFunc("/bot", hanlde_bot)
    http.HandleFunc("/sendbot", handle_sendbot)

    http.ListenAndServe(":8080", nil)
}

func hanlde_bot(w http.ResponseWriter, r *http.Request) {
    
    bot, _ := tgbotapi.NewBotAPI(bot_id)
    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates := bot.GetUpdatesChan(u)

    for update := range updates {
        if update.Message != nil {
            fmt.Printf("chat id: %s", update.Message.Chat.ID)
        }
    }
}

func handle_sendbot(w http.ResponseWriter, r *http.Request) {
    bot, _ := tgbotapi.NewBotAPI(bot_id)

    data, _ := ioutil.ReadAll(r.Body)

    requestAsString := string(data)

    msg := tgbotapi.NewMessage(chat_id, requestAsString)

    bot.Send(msg)
}