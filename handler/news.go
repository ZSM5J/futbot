package handler

import (
	"fmt"
	"net/http"
	"../db"
	"strings"
	"log"
)

//AddNews is used for check server state
var AddNews = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	head := request.URL.Query().Get("head")
	url := request.URL.Query().Get("url")
	id := ""
	subs := db.GetNewsSubscribers()
	for i, sub := range subs {
		id = id + sub.VK
		if i != len(subs) -1 {
			id = id + ","
		}
	}
	log.Println(id)
	SendNews(id, head, url)
	//Insert into database
	db.InsertNews(head, url)
	fmt.Fprintf(response, "news is inserted")
})

//SubscribeNews subscribe on news
func SubscribeNews(id, arg string) {
	switch arg {
		case "+":
			if	db.ExistNewsSub(id) {
				SendMessage(id, "Вы уже подписаны на новости!")
			} else {
				db.InsertSubNews(id)
				SendMessage(id, "Договорились🤝. Я буду присылать тебе новости в тот же миг, когда сам узнаю B-)")
			}
		case "-":
			if	!db.ExistNewsSub(id) {
				SendMessage(id, "Ты не подписан на новости.")
			} else {
				db.DeleteSubNews(id)
				SendMessage(id, "Я больше не буду присылать тебе свежие новости из мира футбола.")
			}
		
		default:
		SendMessage(id, "Неверно использована команда. Наберите:%0A новости %2b чтобы подписаться%0A новости - чтобы отписаться ")
	}

}


//SendNews answer to user
func SendNews(id, head, url string) {
	
	message := head + "%0A " + url
	log.Println(message)
	message = strings.Replace(message, " ", "%20", -1)
	_, err := http.Get("https://api.vk.com/method/messages.send?access_token=449c941ace6703fd55cfaee6b5863da12ecd29ad56168d842b73809a747bd2987f49a0ad4f3e64ab81178" + "&user_ids=" + id + "&message=" + message)
	if err != nil {
		log.Println(err)
	}
}