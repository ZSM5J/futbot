package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"../db"
	"../model"
)

//CulculateBets is func
var CulculateBets = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

	matchID := request.URL.Query().Get("id")
	winner, _ := strconv.Atoi(request.URL.Query().Get("win"))
	goals1, _ := strconv.Atoi(request.URL.Query().Get("goals1"))
	goals2, _ := strconv.Atoi(request.URL.Query().Get("goals2"))
	bets := db.GetBetsByID(matchID)

	for i := 0; i < len(bets); i++ {
		points := PointsFromBet(winner, goals1, goals2, bets[i].Winner, bets[i].Goal1, bets[i].Goal2)
		db.UpdateBetPoints(points, bets[i].ID)
	}
	fmt.Fprintf(response, "рассчитано")

})

//AddPoints is func
var AddPoints = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var i, k int
	var bets []model.UserBets
	id := request.URL.Query().Get("id")
	if id == "16154083" {
		users := db.GetAllUsers()
		for i = 0; i < len(users); i++ {
			bets = db.GetBetsByVK(users[i].VK)
			for k = 0; k < len(bets); k++ {
				db.UpdateUserPoints(users[i].VK, bets[k].Points)
				db.UpdateBetAdd(bets[k].ID)
			}
		}
		log.Println(users)
	}

	fmt.Fprintf(response, "добавлено")

})

//SortPlaces is func
var SortPlaces = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var i, j, k int
	var temp model.User
	id := request.URL.Query().Get("id")
	if id == "16154083" {
		users := db.GetAllUsers()
		
		for j = 0; j < len(users); j++{
			for k = 0; k < len(users) - 1; k++ {
				if (users[k].Points < users[k + 1].Points) {
					temp = users[k];
					users[k] = users[k + 1];
					users[k + 1] = temp;
				}
			}
		}

		
		for i = 0; i < len(users); i++ {
			db.UpdateUserPlace(users[i].VK, i+1)
		}
	}

	fmt.Fprintf(response, "места распределены")

})

//PointsFromBet is func
func PointsFromBet(res int, g1 int, g2 int, betres int, betg1 int, betg2 int) int {
	if g1 == betg1 && g2 == betg2 {
		log.Println("3")
		return 3
	}
	if res == betres && g1-g2 == betg1-betg2 && res != 0 {
		log.Println("2")
		return 2
	}
	if res == betres {
		log.Println("1")
		return 1
	}
	return 0
}
