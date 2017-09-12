package db

import (
	"log"

	gormigrate "gopkg.in/gormigrate.v1"
	"time"
	"../model"
	"github.com/jinzhu/gorm"
)

//Migrate func for migrate
func Migrate() error {
	db := connect()
	defer db.Close()

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "201707241800",
			Migrate: func(tx *gorm.DB) error {
				type OpenBets struct {
					ID      int `sql:"AUTO_INCREMENT" gorm:"primary_key"`
					MatchID string
					Team1   string
					Team2   string
					IsOpen  bool
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
				return tx.AutoMigrate(&model.OpenBets{}, &model.User{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("open_bets", "users").Error
			},
		},
		{
			ID: "201707290005",
			Migrate: func(tx *gorm.DB) error {
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
				return tx.AutoMigrate(&model.UserBets{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("users_bets").Error
			},
		},
		{
			ID: "201707290009",
			Migrate: func(tx *gorm.DB) error {
				type Info struct {
					Users    int
					MinBetID int
					MaxBetID int
				}
				return tx.AutoMigrate(&model.Info{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("info").Error
			},
		},
		{
			ID: "201707290012",
			Migrate: func(tx *gorm.DB) error {
				type OpenBets struct {
					ID      int `sql:"AUTO_INCREMENT" gorm:"primary_key"`
					MatchID string
					Team1   string
					Team2   string
					IsOpen  bool
					Group	int
				}
				return tx.AutoMigrate(&model.OpenBets{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropColumn("group").Error
			},
		},
		{
			ID: "201708311828",
			Migrate: func(tx *gorm.DB) error {
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
				return tx.AutoMigrate(&model.News{},&model.NewsSub{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropColumn("group").Error
			},
		},
	})
	err := m.Migrate()
	if err == nil {
		log.Printf("Migration did run successfully")
	} else {
		log.Printf("Could not migrate: %v", err)
	}

	return nil
}
