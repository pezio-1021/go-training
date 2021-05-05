package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type UserService struct {
	db *gorm.DB
}

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null; unique_index"`
}

func NewUserService(connectioninfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectioninfo)
	if err != nil {
		return nil, err
	}

	db.LogMode(true)
	return &UserService{
		db: db,
	}, nil
}
