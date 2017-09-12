package db

import (
	"log"

	"../model"
)

//InsertNewUser add new user to db
func InsertNewUser(id string) error {
	log.Println("dobavlyaem")
	db := connect()
	defer db.Close()

	user := &model.User{
		VK:          id,
		Points:      0,
		Place:       0,
		NextBetID:   0,
		LastCommand: "",
		SpamCount:   0}

	err := db.Create(&user)

	return err.Error
}

//ExistUserID check exist this ID
func ExistUserID(id string) bool {
	db := connect()
	defer db.Close()
	var user model.User
	db.Where("vk = ?", id).First(&user)
	if len(user.VK) > 1 {
		return true
	}
	return false
}

//GetUser load user profile
func GetUser(id string) model.User {
	db := connect()
	defer db.Close()
	var user model.User
	db.Where("vk = ?", id).First(&user)
	return user
}

//UpdateNextBet replace token
func UpdateNextBet(id string, nextID int) error {
	db := connect()
	defer db.Close()
	var user model.User
	db.Where("vk = ?", id).First(&user)
	newID := user.NextBetID + 1
	if nextID != 0 {
		newID = nextID
	}
	err := db.Model(&user).Where("vk = ?", id).Update("NextBetID", newID)
	return err.Error
}

//GetAllUsers replace token
func GetAllUsers() []model.User {
	db := connect()
	defer db.Close()
	var users []model.User
	db.Find(&users)
	return users
}

//UpdateUserPoints add points
func UpdateUserPoints(id string, points int) error {
	db := connect()
	defer db.Close()
	var user model.User
	db.Where("vk = ?", id).First(&user)
	newpoints := user.Points + points
	err := db.Model(&user).Where("vk = ?", id).Update("Points", newpoints)
	return err.Error
}

//UpdateUserPlace set place
func UpdateUserPlace(id string, place int) error {
	db := connect()
	defer db.Close()
	var user model.User
	db.Where("vk = ?", id).First(&user)
		err := db.Model(&user).Where("vk = ?", id).Update("Place", place)
	return err.Error
}

//UpdateLastCommand set place
func UpdateLastCommand(id string, cmd string) error {
	db := connect()
	defer db.Close()
	var user model.User
	db.Where("vk = ?", id).First(&user)
		err := db.Model(&user).Where("vk = ?", id).Update("LastCommand", cmd)
	return err.Error
}


//INFO

//CreateInfo add new user to db
func CreateInfo(users int, min int, max int) error {
	log.Println("create info")
	db := connect()
	defer db.Close()

	info := &model.Info{
		Users:    users,
		MinBetID: min,
		MaxBetID: max}

	err := db.Create(&info)

	return err.Error
}

//GetInfo load info
func GetInfo() model.Info {
	db := connect()
	defer db.Close()
	var info model.Info
	db.First(&info)
	return info
}

//UpdateInfoUser replace token
func UpdateInfoUser() error {
	db := connect()
	defer db.Close()
	var info model.Info
	db.First(&info)
	newUsers := info.Users + 1
	err := db.Model(&info).Update("Users", newUsers)
	return err.Error
}

//UpdateInfoMin replace token
func UpdateInfoMin(min int) error {
	db := connect()
	defer db.Close()
	var info model.Info
	db.First(&info)
	err := db.Model(&info).Update("MinBetID", min)
	return err.Error
}

//UpdateInfoMax replace token
func UpdateInfoMax(max int) error {
	db := connect()
	defer db.Close()
	var info model.Info
	db.First(&info)
	err := db.Model(&info).Update("MaxBetID", max)
	return err.Error
}
