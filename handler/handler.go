package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	_"../config"
	"../db"
	"github.com/Jeffail/gabs"
)

func responseMessage(data interface{}, response http.ResponseWriter, request *http.Request) {
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(js)
}

//ServerStatus is struct
type ServerStatus struct {
	Alive bool
	Time  time.Time
}

//StatusHandler is used for check server state
var StatusHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var status ServerStatus
	status.Alive = true
	status.Time = time.Now()
	responseMessage(status, response, request)
})

//MinBetChanger is used for check server state
var MinBetChanger = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	value, _ := strconv.Atoi(request.URL.Query().Get("value"))
	db.UpdateInfoMin(value)
	fmt.Fprintf(response, "установлено min")
})

//MaxBetChanger is used for check server state
var MaxBetChanger = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	value, _ := strconv.Atoi(request.URL.Query().Get("value"))
	db.UpdateInfoMax(value)
	fmt.Fprintf(response, "установлено max")
})

//ConfirmHandler is used for check server state
var ConfirmHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "515e7ecd")
})

//MessageHandler is used for check server state
var MessageHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	body, _ := ioutil.ReadAll(request.Body)
	json, _ := gabs.ParseJSON(body)
	message := strings.ToLower(ClearString(json.Path("object.body").String()))
	id := ClearID(json.Path("object.user_id").String())
	command, arg := Split(message)
	log.Println(id)
	log.Println(command, arg)
	switch command {
	case "help":
		ShowHelp(id)
	case "хелп":
		ShowHelp(id)
	case "помощь":
		ShowHelp(id)
	case "?":
		ShowHelp(id)

	case "привет":
		SendMessage(id, RandomHello())
	case "ок":
		SendMessage(id, ":like:")
	case "окей":
		SendMessage(id, ":like:")
	case "ok":
		SendMessage(id, ":like:")
	case "круто":
		SendMessage(id, ":like:")
	case "клево":
		SendMessage(id, ":like:")
	case "заебись":
		SendMessage(id, ";-]")
	case "збс":
		SendMessage(id, ";-]")
	case "здорово":
		SendMessage(id, ";-]")
	case "отлично":
		SendMessage(id, ";-]")
	case "хорошо":
		SendMessage(id, ";-)")
	case "пока":
		SendMessage(id, ":v:")
	case "прет":
		SendMessage(id, RandomHello())
	case "хай":
		SendMessage(id, RandomHello())
	case "hello":
		SendMessage(id, RandomHello())
	case "hi":
		SendMessage(id, RandomHello())
	case "добрый день":
		SendMessage(id, RandomHello())
	case "как":
		if arg == "дела" || arg == "сам" || arg == "че" || arg == "ты" || arg == "дела?" || arg == "сам?" || arg == "ты?" {
			SendMessage(id, RandomHowAreYou())
		} else {
			SendMessage(id, RandomFuckingMessage(message))
		}
	case "соси":
		SendMessage(id, "давай как-то сам")
	case "нахуй":
		SendMessage(id, "пожалуй, откажусь")
	case "пидор":
		SendMessage(id, "зачем обзываешься?")
	case "сука":
		SendMessage(id, "грубо.")
	case "блять":
		SendMessage(id, "...")
	case "блядь":
		SendMessage(id, "...")
	case "иди":
		if arg == "нахуй" || arg == "на" || arg == "в" || arg == "впизду" {
		SendMessage(id, "неа")
		} else {
			SendMessage(id, RandomFuckingMessage(message))
		}


	case "рега!":
		Register(id)
	case "рега":
		Register(id)
	case "рег":
		Register(id)
	case "место":
		ShowPlace(id)
	case "рейтинг":
		ShowPlace(id)
	case "рейт":
		ShowPlace(id)

	case "ставки":
		MakeBets(id)
	case "бет":
		BetMatch(id, arg)
	case "нарезка":
		SendVideo(id)
	case "гифка":
		SendGif(id)
	case "гиф":
		SendGif(id)
	case "новости":
		SubscribeNews(id, arg)
	default:
		if CheckArg(command) {
			BetMatch(id, command)
		} else {
			SendMessage(id, RandomFuckingMessage(message) + "%0AЧтобы увидеть список команд, набери: хелп ")
			db.UpdateLastCommand(id, "spam")
		}
	}
	fmt.Fprintf(response, "ok")
})

//SendVideo is for narezka
func SendVideo(id string) {
	url := RandomVideoURL()
	SendContent(id,"Лови", url)
	db.UpdateLastCommand(id, "video")
}

//SendGif is for narezka
func SendGif(id string) {
	url := RandomGifURL()
	SendContent(id,"Держи", url)
	db.UpdateLastCommand(id, "gif")
}

func registerUser(id string) {
	SendMessage(id, "Ты принят в лигу! %0A Правила лиги: https://vk.cc/73P0af %0A Теперь можешь делать прогнозы и участвовать в обсуждениях.⚽ %0A  ❗Для прогнозов набери: ставки")
}

//SendMessage answer to user
func SendMessage(id, message string) {
	log.Println(id)
	message = strings.Replace(message, " ", "%20", -1)
	_, err := http.Get("https://api.vk.com/method/messages.send?access_token=449c941ace6703fd55cfaee6b5863da12ecd29ad56168d842b73809a747bd2987f49a0ad4f3e64ab81178" + "&user_ids=" + id + "&message=" + message)
	if err != nil {
		log.Println(err)
	}
}

//Register add user to db or not
func Register(id string) {
	if db.ExistUserID(id) {
		SendMessage(id, "Вы уже учавствуете в лиге прогнозистов.")
	} else {
		db.UpdateInfoUser()
		db.InsertNewUser(id)
		registerUser(id)
	}
}

//ShowHelp show help
func ShowHelp(id string) {
	help := "📌 Список команд:%0A ▶хелп/?  - список команд%0A ▶нарезка - случайная видео подборка голов %0A ▶гифка - случайная гифка %0A ▶новости %2b(знак плюс) - подписаться на новости%0A ▶новости -(знак минус) - отписаться от новостей %0A ▶рега - вступить в лигу прогнозистов%0A ▶ставки - матчи для прогнозов %0A ▶бет 2-0(счёт) - прогноз на матч %0A ▶место - узнать свой рейтинг"
	SendMessage(id, help)
	db.UpdateLastCommand(id, "help")
}

//ShowPlace show здфсу
func ShowPlace(id string) {
	str := ""
	if db.ExistUserID(id) {
		user := db.GetUser(id)
		info := db.GetInfo()
		usersCount := strconv.Itoa(info.Users)
		place := strconv.Itoa(user.Place)
		points := strconv.Itoa(user.Points)
		if user.Points == 0 && info.MinBetID == 1 {
			str = "У вас 0 очков, как и у всех. Ждем пока сыграют первые матчи ;-)"
		} 
		if user.Points == 0 && info.MinBetID > 1 {
			str = "У вас 0 очков. Вы на последнем месте из " + usersCount
		}
		if user.Points > 0 {
			str = "Набрано поинтов: " + points + ". Вы занимаете " + place + " место из " + usersCount 
		}

	} else {
		str = "Чтобы узнать свою позицию в лиге, надо в ней зарегистрироваться. Напишите: рег"
	}
	
	SendMessage(id, str)
	db.UpdateLastCommand(id, "place")
}
