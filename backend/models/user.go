package models

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Username      string   `json:"username" gorm:"unique;not null"`
	Email         string   `json:"email" gorm:"unique;not null"`
	PasswordHash  string   `json:"-" gorm:"not null"`
	FirstName     string   `json:"first_name" gorm:"not null"`
	LastName      string   `json:"last_name" gorm:"not null"`
	ProfilePicUrl string   `json:"profile_pic"`
	Bio           string   `json:"bio"`
	ProfileLikes  int32    `json:"profile_likes"`
	Folders       []string `json:"folders" gorm:"type:text[]"`
	Favorites     []string `json:"favorites" gorm:"type:text[]"`
	IsAdmin       bool     `json:"is_admin" gorm:"default:false"`
}
