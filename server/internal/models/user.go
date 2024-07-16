package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	ID         uint   `gorm:"primary_key"`
	LastName   string `gorm:"type:varchar(255)"`
	FirstName  string `gorm:"type:varchar(255)"`
	MiddleName string `gorm:"type:varchar(255)"`
	Phone      string `gorm:"type:varchar(255);unique_index"`
	Password   string `gorm:"type:varchar(255)"`
	Position   string `gorm:"type:varchar(255)"`
	Role       string `gorm:"type:enum('Admin', 'User')"`
	LastLogin  time.Time
}

func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
	return db
}
