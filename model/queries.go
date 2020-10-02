package model

import (
	"github.com/madeinlobby/R320_backend/model/database"
	"time"
)

//func GetMeme(id int64) *database.Meme {
//	result := &database.Meme{}
//	database.DB.Limit(10).Find(result, id)
//	return result
//}

func GetTopMeme(time *time.Time, limit int) (*[]database.Meme, error) {
	result := &[]database.Meme{}
	tx := database.DB.Limit(limit).Where("upload_time > ?", time).Order("memes.like desc,upload_time").Find(result)
	return result, tx.Error
}

func GetEverTopMeme(limit int) (*[]database.Meme, error) {
	result := &[]database.Meme{}
	tx := database.DB.Limit(limit).Order("memes.like desc,upload_time").Find(result)
	return result, tx.Error
}

func GetRandomMeme(limit int) (*[]database.Meme, error) {
	result := &[]database.Meme{}
	tx := database.DB.Limit(limit).Order("RANDOM()").Find(result)
	return result, tx.Error
}

func GetLastMeme(limit int) (*[]database.Meme, error) {
	result := &[]database.Meme{}
	tx := database.DB.Limit(limit).Order("upload_time desc").Find(result)
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

func GetComment(memeId int64) (*[]database.Comment, error) {
	result := &[]database.Comment{}
	tx := database.DB.Where("meme_id = ? and up_comment_id = ?", memeId, -1).Order("publish_time").Find(result)
	return result, tx.Error
}

func GetReplies(CommentId int64) (*[]database.Comment, error) {
	result := &[]database.Comment{}
	tx := database.DB.Where("up_comment_id = ?", CommentId).Order("publish_time").Find(result)
	return result, tx.Error
}
