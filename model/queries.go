package model

import (
	"github.com/madeinlobby/R320_backend/model/database"
	"time"
)

func GetMeme(id int64) *database.Meme {
	result := &database.Meme{}
	database.DB.Limit(10).Find(result, id)
	return result
}

func GetTopMeme(time time.Time) (*[]database.Meme, error) {
	result := &[]database.Meme{}
	tx := database.DB.Limit(10).Where("upload_time > ?", time).Order("memes.like desc,upload_time").Find(result)
	return result, tx.Error
}

func GetUser(username string) (*database.User, error) {
	result := &database.User{}
	tx := database.DB.Where("username = ?", username).First(result)
	return result, tx.Error
}

func GetTags(memeId int64) (*[]database.Tag, error) {
	result := &[]database.Tag{}
	tx := database.DB.Where("meme_id = ?", memeId).Order("name").Find(result)
	return result, tx.Error
}
