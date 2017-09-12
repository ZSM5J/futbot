package handler

import (
	"math/rand"
	"time"
)

//RandomGoodBetMessage is a good function
func RandomGoodBetMessage(message string) string {
	rand.Seed(time.Now().UnixNano())
	str := []string{"Близко к правде!", message + "! Уже слышу как Соловьёв кричит этот счёт после финального свистка.",
		message + "!:ok:", "Были бы у меня ноги, поставил бы также.", "Принято.", "Мой любимый счет!",
		"Твоя жизнь - твои прогнозы.", "Как-то слабо верится, но принято!", "Популярный счет.", "Записал!", "В лучших традициях!"}
	str2 := str[rand.Intn(len(str))]
	return str2
}

//RandomVideoURL is a good function
func RandomVideoURL() string {
	rand.Seed(time.Now().UnixNano())
	str := []string{"video-152793181_456239017",
		"video-152793181_456239018", "video-152793181_456239019",
		"video-152793181_456239020", "video-152793181_456239022",
		"video-152793181_456239023", "video-152793181_456239024",
		"video-152793181_456239025", "video-152793181_456239026",
		"video-152793181_456239027", "video-152793181_456239028",}
	str2 := str[rand.Intn(len(str))]
	return str2
}

//RandomGifURL is a good function
func RandomGifURL() string {
	rand.Seed(time.Now().UnixNano())
	str := []string{"doc-152793181_449961124",
		"doc-152793181_449961216", "doc-152793181_449961227",
		"doc-152793181_449961233", "doc-152793181_449961242",
		"doc-152793181_449961256", "doc-152793181_449961261",
		"doc-152793181_449961283", "doc-152793181_449961331",
		"doc-152793181_449961299", "doc-152793181_449961321",
		"doc-152793181_449961345", "doc-152793181_449961360",
		"doc-152793181_449961383", "doc-152793181_449961390",
		"doc-152793181_449961395", "doc-152793181_449961403",
		"doc-152793181_449961412", "doc-152793181_449961423",
		"doc-152793181_449961451", "doc-152793181_449961488",
		"doc-152793181_449961504", "doc-152793181_449961529",
		"doc-152793181_449961547", "doc-152793181_449961556",}
	str2 := str[rand.Intn(len(str))]
	return str2
}

//RandomHello is a good function
func RandomHello() string {
	rand.Seed(time.Now().UnixNano())
	str := []string{"Привет:-)",
		"Бонжур!", "О!:-D Приветик",
		"дороу", "хай", "Дорова, братан!", "Йеп! Как сам?", "Хеллоу, нига!&#128526;"}
	str2 := str[rand.Intn(len(str))]
	return str2
}

//RandomHowAreYou is a good function
func RandomHowAreYou() string {
	rand.Seed(time.Now().UnixNano())
	str := []string{"После того, как ты написал - хорошо:-)",
		"Работаю:-(", "Думаю подкатить к кофеварке из бухгалтерии. А ты чем занят?",
		"устал:((", "купил очки B-) ", "нормально...", "всё ок, а ты?", "Как пальцем пиханная колбаса!"}
	str2 := str[rand.Intn(len(str))]
	return str2
}

//RandomFuckingMessage is a good function
func RandomFuckingMessage(message string) string {
	rand.Seed(time.Now().UnixNano())
	str := []string{"Давай нормально!", message + " у тебя в штанах", "не балуйся.", "Введи корректную команду", "Ай, я конечно отвечу, но лучше бы ты писал по делу", "нет такой команды", "первый раз слышу"}
	str2 := str[rand.Intn(len(str))]
	return str2
}
