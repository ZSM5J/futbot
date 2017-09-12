package db

import (
	"log"

	"../model"
)

//InsertOpenBet add new user to db
func InsertOpenBet(matchID string, team1 string, team2 string, group int) error {
	log.Println("dobavlyaem")
	db := connect()
	defer db.Close()

	bet := &model.OpenBets{
		MatchID: matchID,
		Team1:   team1,
		Team2:   team2,
		IsOpen:  true,
		Group: group}

	err := db.Create(&bet)

	return err.Error
}

//GetOpenBetByID load user profile
func GetOpenBetByID(id string) model.OpenBets {
	db := connect()
	defer db.Close()
	var bet model.OpenBets
	db.Where("match_id = ?", id).First(&bet)
	return bet
}

//InsertBet add new bet
func InsertBet(vk string, matchID string, winner int, goal1 int, goal2 int) error {
	log.Println("delaem stavochky")
	db := connect()
	defer db.Close()

	bet := &model.UserBets{
		VK:        vk,
		MatchID:   matchID,
		Winner:    winner,
		Goal1:     goal1,
		Goal2:     goal2,
		Points:    0,
		Calculate: false,
		Added:     false}

	err := db.Create(&bet)

	return err.Error
}

//UpdateBetPoints replace token
func UpdateBetPoints(points int, id int) {
	log.Println("obnovlyaem")
	db := connect()
	defer db.Close()
	var bet model.UserBets
	db.Where("id = ?", id).First(&bet)
	db.Model(&bet).Update("Points", points)
	db.Model(&bet).Update("Calculate", true)
}

//UpdateBetAdd replace token
func UpdateBetAdd(id int) {
	log.Println("obnovlyaem add")
	db := connect()
	defer db.Close()
	var bet model.UserBets
	db.Where("id = ?", id).First(&bet)
	db.Model(&bet).Update("Added", true)
}

//UpdateToken replace token
func UpdateToken(strID string, strToken string) error {
	db := connect()
	defer db.Close()
	var user model.User
	db.Where("vk_id = ?", strID).First(&user)
	//проверить, мб можно удалить какую-то из строчек^
	err := db.Model(&user).Where("vk_id = ?", strID).Update("Token", strToken)
	return err.Error
}

//GetBetsByID replace token
func GetBetsByID(matchID string) []model.UserBets {
	db := connect()
	defer db.Close()
	var bets []model.UserBets
	db.Where("match_id = ? and calculate = ?", matchID, false).Find(&bets)
	return bets
}

//GetBetsByVK replace token
func GetBetsByVK(vk string) []model.UserBets {
	db := connect()
	defer db.Close()
	var bets []model.UserBets
	db.Where("vk = ? and calculate = ? and added = ?", vk, true, false).Find(&bets)
	return bets
}
