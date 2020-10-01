package view

import (
	"encoding/json"
	"github.com/madeinlobby/R320_backend/configuration"
	"github.com/madeinlobby/R320_backend/model"
	"github.com/madeinlobby/R320_backend/model/database"
	"log"
	"net/http"
	"strconv"
	"time"
)

func TopDayMeme(writer http.ResponseWriter, request *http.Request) {
	query := func() (*[]database.Meme, error) {
		pageSize, pageNumber := getPageInfo(request)
		t := time.Now().Add(-24 * time.Hour)
		return model.GetTopMeme(&t, pageNumber*pageSize)
	}
	TopMeme(query, writer, request)
}

func TopWeekMeme(writer http.ResponseWriter, request *http.Request) {
	query := func() (*[]database.Meme, error) {
		pageSize, pageNumber := getPageInfo(request)
		t := time.Now().Add(-7 * 24 * time.Hour)
		return model.GetTopMeme(&t, pageNumber*pageSize)
	}
	TopMeme(query, writer, request)
}

func TopEverMeme(writer http.ResponseWriter, request *http.Request) {
	query := func() (*[]database.Meme, error) {
		pageSize, pageNumber := getPageInfo(request)
		return model.GetEverTopMeme(pageNumber * pageSize)
	}
	TopMeme(query, writer, request)
}

func TopMeme(query func() (*[]database.Meme, error), writer http.ResponseWriter, request *http.Request) {
	pageSize, pageNumber := getPageInfo(request)
	memes, err := query()
	if err != nil {
		writeError(err, writer)
		return
	}
	serveMemes(memes, pageSize, pageNumber, writer, request)
}

func serveMemes(memes *[]database.Meme, pageSize, pageNumber int, w http.ResponseWriter, request *http.Request) {
	var result []Meme
	for index, element := range *memes {
		if pageSize*(pageNumber-1) <= index && index < pageSize*pageNumber {
			meme, err := processMeme(&element)
			if err != nil {
				writeError(err, w)
				return
			}
			result = append(result[:], *meme)
		}
	}
	sendResponse(w, request, result)
}

func getPageInfo(request *http.Request) (int, int) {
	pn := request.URL.Query().Get("page_number")
	ps := request.URL.Query().Get("page_size")
	pageNumber, err := strconv.Atoi(pn)
	if err != nil {
		pageNumber = configuration.DefaultPageNumber
	}
	pageSize, err := strconv.Atoi(ps)
	if err != nil {
		pageSize = configuration.DefaultPageSize
	}
	return pageSize, pageNumber
}

func processMeme(element *database.Meme) (*Meme, error) {
	meme := Meme{
		Title:    element.Title,
		Id:       element.ID,
		Username: element.UploaderUsername,
		Picture:  element.ImageAddress,
		Like:     element.Like,
	}
	user, err := model.GetUser(element.UploaderUsername)
	if err != nil {
		return nil, err
	}
	tags, err := model.GetTags(element.ID)
	if err != nil {
		return nil, err
	}
	meme.UsernameAvatarUrl = user.Avatar
	var stringTags []string
	stringTags = append(stringTags)
	for _, tag := range *tags {
		stringTags = append(stringTags, tag.Name)
	}
	meme.Tag = stringTags
	return &meme, nil
}

func writeError(err error, writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusInternalServerError)
	_, err = writer.Write([]byte(err.Error()))
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
