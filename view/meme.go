package view

import (
	"encoding/json"
	"github.com/madeinlobby/R320_backend/model"
	"log"
	"net/http"
	"time"
)

func TopDayMeme(responseWriter http.ResponseWriter, request *http.Request) {
	memes, err := model.GetTopMeme(time.Now().Add(-24 * time.Hour))
	if err != nil {
		writeError(err, responseWriter)
		return
	}
	var result []Meme
	for _, element := range *memes {
		meme := Meme{
			Title:    element.Title,
			Id:       element.ID,
			Tag:      nil,
			Username: element.UploaderUsername,
			Picture:  element.ImageAddress,
			Like:     element.Like,
		}
		user, err := model.GetUser(element.UploaderUsername)
		if err != nil {
			writeError(err, responseWriter)
			return
		}
		tags, err := model.GetTags(element.ID)
		if err != nil {
			writeError(err, responseWriter)
			return
		}
		meme.UsernameAvatarUrl = user.Avatar
		var stringTags []string
		for _, tag := range *tags {
			stringTags = append(stringTags, tag.Name)
		}
		meme.Tag = &stringTags
		result = append(result[:], meme)
	}
	sendResponse(responseWriter, request, result)
}

func writeError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	_, err = w.Write([]byte(err.Error()))
	if err != nil {
		log.Printf(err.Error())
	}
}

func sendResponse(writer http.ResponseWriter, request *http.Request, resp interface{}) {
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err = writer.Write(jsonResp); err != nil {
		log.Printf("could not write response: %s", request.RequestURI)
	}
}
