package model

import "time"

//OpenBets struct
type OpenBets struct {
	ID      int `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	MatchID string
	Team1   string
	Team2   string
	IsOpen  bool
	Group 	int
}

//User struct
type User struct {
	ID          int `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	VK          string
	Points      int
	Place       int
	NextBetID   int
	LastCommand string
	SpamCount   int
}

//UserBets struct
type UserBets struct {
	ID        int `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	VK        string
	MatchID   string
	Winner    int
	Goal1     int
	Goal2     int
	Points    int
	Calculate bool
	Added     bool
}

//Info show common info
type Info struct {
	Users    int
	MinBetID int
	MaxBetID int
}

//News struct for news
type News struct {
	ID int `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Header string
	URL string
	Date time.Time
}

//NewsSub struct for people
type NewsSub struct {
	ID int `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	VK string
}