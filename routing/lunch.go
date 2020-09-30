package routing

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/madeinlobby/R320_backend/model/database"
	"github.com/madeinlobby/R320_backend/view"
	"net/http"
	"os"
	"strconv"
	"time"
)

var Router *mux.Router

func LunchServer() error {
	err := database.LunchDB()
	if err != nil {
		return err
	}
	err = configureRouter()
	return err
}

func configureRouter() error {
	Router = mux.NewRouter()
	err := staticFile(Router)
	if err != nil {
		return err
	}
	err = memeRouting(Router)
	if err != nil {
		return err
	}
	// routing
	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, Router),
		Addr:         Address + ":" + strconv.Itoa(Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	err = srv.ListenAndServe()
	return err
}

func staticFile(router *mux.Router) error {
	router.PathPrefix("/files/").Handler(http.StripPrefix("/files/",
		http.FileServer(http.Dir(FilesAddress))))
	return nil
}

func memeRouting(router *mux.Router) error {
	router.Methods("GET").Path("/meme/top/day").HandlerFunc(view.TopDayMeme)
	return nil
}
