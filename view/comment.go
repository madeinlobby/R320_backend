package view

import (
	"github.com/gorilla/mux"
	"github.com/madeinlobby/R320_backend/model"
	"github.com/madeinlobby/R320_backend/model/database"
	"net/http"
	"strconv"
)

func CommentByMemeID(writer http.ResponseWriter, request *http.Request) {
	memeIDString := mux.Vars(request)["meme_id"]
	memeId, err := strconv.ParseInt(memeIDString, 10, 64)
	if err != nil {
		writeError(err, writer)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	comments, err := model.GetComment(memeId)
	if err != nil {
		writeError(err, writer)
		return
	}
	var result []Comment
	for _, element := range *comments {
		a, err := commentToView(&element)
		if err != nil {
			writeError(err, writer)
			return
		}
		result = append(result, *a)
	}
	sendResponse(writer, request, result)
}

func commentToView(comment *database.Comment) (*Comment, error) {
	result := Comment{
		ID:     comment.ID,
		Author: comment.Username,
		Text:   comment.Text,
		Like:   comment.Like,
		Date:   comment.PublishTime,
	}
	user, err := model.GetUser(comment.Username)
	if err != nil {
		return nil, err
	}
	result.Avatar = user.Avatar
	comments, err := model.GetReplies(comment.ID)
	if err != nil {
		return nil, err
	}
	var replies []Comment
	for _, element := range *comments {
		t, err := commentToView(&element)
		if err != nil {
			return nil, err
		}
		replies = append(replies, *t)
	}
	result.Replies = &replies
	return &result, nil
}
