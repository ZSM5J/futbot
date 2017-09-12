package db

import (
	"log"
	"time"
	"../model"
)


//InsertNews add news
func InsertNews(header string, url string) error {
	log.Println("dobavlyaem news")
	db := connect()
	defer db.Close()

	news := &model.News{
		Header: header,
		URL:   url,
		Date:   time.Now()}

	err := db.Create(&news)

	return err.Error
}

//InsertSubNews add subscriber
func InsertSubNews(id string) error {
	log.Println("dobavlyaem")
	db := connect()
	defer db.Close()

	sub := &model.NewsSub{
		VK:          id}

	err := db.Create(&sub )

	return err.Error
}

//DeleteSubNews add subscriber
func DeleteSubNews(id string) {
	log.Println("dobavlyaem")
	db := connect()
	defer db.Close()
	db.Where("vk = ?", id).Delete(&model.NewsSub{})
}

//GetNewsSubscribers replace token
func GetNewsSubscribers() []model.NewsSub {
	db := connect()
	defer db.Close()
	var subs []model.NewsSub
	db.Find(&subs)
	return subs
}

//ExistNewsSub check exist this ID
func ExistNewsSub(id string) bool {
	db := connect()
	defer db.Close()
	var user model.NewsSub
	db.Where("vk = ?", id).First(&user)
	if len(user.VK) > 1 {
		return true
	}
	return false
}
