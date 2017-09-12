package handler

import (
	"log"
	"strconv"
	"regexp"
	"../db"
)

//MakeBets help you make bets
func MakeBets(id string) {
	if db.ExistUserID(id) {
		user := db.GetUser(id)
		info := db.GetInfo()
		log.Println(info)
		if user.NextBetID == 0 && info.MinBetID == 1 {
			bet := db.GetOpenBetByID("1")
			match := bet.Team1 + " - " + bet.Team2
			str := "Вот и пришло время тебе сделать первые прогнозы. Давай так, я пишу матч, ты пишешь счет в формате '2-0' или '3:1'. Первый матч: %0A" + match
			SendMessage(id, str)
			db.UpdateNextBet(id, 0)
			db.UpdateLastCommand(id, "bets")
		}
		if (user.NextBetID >= info.MinBetID) && user.NextBetID <= info.MaxBetID {
			bet := db.GetOpenBetByID(strconv.Itoa(user.NextBetID))
			match := bet.Team1 + " - " + bet.Team2
			str := "Жду прогноз по матчу: %0A" + match
			db.UpdateLastCommand(id, "bets")
			SendMessage(id, str)
		}
		if user.NextBetID > info.MaxBetID {
			SendMessage(id, "Ты сделал все прогнозы. Я сообщу тебе, когда появятся новые")
		}
		if user.NextBetID < info.MinBetID && info.MinBetID != 1 {
			db.UpdateNextBet(id, info.MinBetID)
			bet := db.GetOpenBetByID(strconv.Itoa(info.MinBetID))
			match := bet.Team1 + " - " + bet.Team2
			str := "Ты немножко отстал, но ничего! Жду прогноз: %0A" + match
			SendMessage(id, str)
			db.UpdateLastCommand(id, "bets")
		}

	} else {
		SendMessage(id, "Чтобы делать ставки вам необходимо зарегистрироваться в лиге. Напишите: рега")
	}
}

//BetMatch help you make bets
func BetMatch(id string, arg string) {
	if db.ExistUserID(id) {
		user := db.GetUser(id)
		info := db.GetInfo()
		log.Println(info)
		//don't see what match need to bet
		if user.LastCommand != "bets" && user.NextBetID <= info.MaxBetID  {
			str := "Напиши ставки, чтобы знать на что ставишь"
			SendMessage(id, str)
		}
		//do all bets;
		if user.NextBetID > info.MaxBetID {
			SendMessage(id, "Ты сделал все прогнозы. Я сообщу тебе, когда появятся новые")
		}

		//all good, make bet.
		if user.NextBetID >= info.MinBetID && user.NextBetID <= info.MaxBetID && user.LastCommand == "bets"  {
			if CheckArg(arg) {
				goals1, _ := strconv.Atoi(arg[:1])
				goals2, _ := strconv.Atoi(arg[2:])
				winner := WhoWin(goals1, goals2)
				log.Println(goals1, ":", goals2, "winner: ", winner)
				betID := strconv.Itoa(user.NextBetID)
				db.InsertBet(user.VK, betID, winner, goals1, goals2)
				db.UpdateNextBet(id, 0)
				if user.NextBetID+1 < info.MaxBetID {
					str := RandomGoodBetMessage(arg)
					bet := db.GetOpenBetByID(strconv.Itoa(user.NextBetID + 1))
					match := bet.Team1 + " - " + bet.Team2
					answer := str + " Следующий матч:%0A" + match
					SendMessage(id, answer)
				}
				if user.NextBetID+1 == info.MaxBetID {
					str := RandomGoodBetMessage(arg)
					bet := db.GetOpenBetByID(strconv.Itoa(user.NextBetID + 1))
					match := bet.Team1 + " - " + bet.Team2
					answer := str + " И напоследок:%0A" + match
					SendMessage(id, answer)
				}
				if user.NextBetID+1 > info.MaxBetID {
					SendMessage(id, "Принято! Ты сделал все прогнозы. Я сообщу тебе, когда появятся новые")
					db.UpdateLastCommand(id, "none")
				}
			} else {
				var validID = regexp.MustCompile(`[0-9]{1}[:-]{1}[0-9]{1}`)
				if validID.MatchString(arg) {
					SendMessage(id, "Это не баскетбол. Ставьте количество голов меньше 10")
				} else {
					SendMessage(id, "Неверный формат счёта.")
				}
				
			}
		}

		if user.NextBetID < info.MinBetID && user.NextBetID > 0   && user.LastCommand == "bets" {
			db.UpdateNextBet(id, info.MinBetID)
			str := "Напиши ставки, чтобы знать на что ставишь. ;-)"
			SendMessage(id, str)
		}

	} else {
		SendMessage(id, "Чтобы делать ставки вам необходимо зарегистрироваться в лиге. Напишите: рега")
	}
}
