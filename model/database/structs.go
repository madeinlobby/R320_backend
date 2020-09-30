package database

import (
	"time"
)

type User struct {
	Username     string `gorm:"primaryKey"`
	Password     string
	Email        string
	Avatar       string
	IsRegistered bool
	LastLogin    time.Time
	SignUpTime   time.Time
}
type Meme struct {
	ID               int64 `gorm:"primaryKey" gorm:"autoincrement"`
	UploaderUsername string
	ImageAddress     string
	Title            string
	Content          string
	Like             int
	UploadTime       time.Time
}

type Tag struct {
	MemeID int64
	Name   string
}

type Comment struct {
	ID          int64 `gorm:"primaryKey" gorm:"autoincrement"`
	Username    string
	Text        string
	Like        int
	MemeId      int64
	UpCommentID int64
}
