package routing

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/madeinlobby/R320_backend/configuration"
	"github.com/madeinlobby/R320_backend/model/database"
	"github.com/madeinlobby/R320_backend/view"
	"net/http"
	"os"
	"strconv"
	"time"
)

var Router *mux.Router

func LunchServer() error {
	var err error
	if err = database.LunchDB(); err != nil {
		return err
	}
	err = configureRouter()
	return err
}

func configureRouter() error {
	Router = mux.NewRouter()
	var err error
	if err = staticFile(Router); err != nil {
		return err
	}
	if err = memeRouting(Router); err != nil {
		return err
	}
	// routing
	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, Router),
		Addr:         configuration.Address + ":" + strconv.Itoa(configuration.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	err = srv.ListenAndServe()
	return err
}

func staticFile(router *mux.Router) error {
	router.PathPrefix("/files/").Handler(http.StripPrefix("/files/",
		http.FileServer(http.Dir(configuration.FilesAddress))))
	return nil
}

func memeRouting(router *mux.Router) error {
	router.Methods("GET").Path("/meme/top/day").HandlerFunc(view.TopDayMeme)
	router.Methods("GET").Path("/meme/top/week").HandlerFunc(view.TopWeekMeme)
	router.Methods("GET").Path("/meme/top/ever").HandlerFunc(view.TopEverMeme)
	return nil
}
