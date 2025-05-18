package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type SnippetModel struct {
	gorm.Model
	Title       string         `json:"title" gorm:"not null"`
	Code        string         `json:"code" gorm:"type:text"`
	Language    string         `json:"language" gorm:"not null"`
	Description string         `json:"description"`
	Folders     pq.StringArray `json:"folders" gorm:"type:text[]"`
	Tags        pq.StringArray `gorm:"type:text[]" json:"tags"`
	Username    string         `json:"username" gorm:"not null"`
	Likes       int32          `json:"likes"`
	Favorite    bool           `json:"favorite"`
	UserID      uint           `json:"user_id" gorm:"not null"`
}
