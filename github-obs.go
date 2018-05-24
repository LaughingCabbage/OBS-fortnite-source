package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	tracker "github.com/LaughingCabbage/fortnite-tracker/v1"
	"github.com/LaughingCabbage/tracker-bot/key"
)

func main() {
	fmt.Println("start")

	router := mux.NewRouter()

	router.HandleFunc("/obs/fortnite.html", handleFortniteData)
	router.PathPrefix("/obs/assets/").Handler(http.StripPrefix("/obs/assets/", http.FileServer(http.Dir("assets"))))
	log.Println(templates.DefinedTemplates())

	log.Fatal(http.ListenAndServe(":8000", router))
}

type wins struct {
	Value int
}

func handleFortniteData(w http.ResponseWriter, r *http.Request) {
	Key := key.LoadKey(".key")
	profile, err := tracker.GetProfile("pc", "laughingcabbage", Key.Value)
	if err != nil {
		panic(err)
	}

	winCount, err := tracker.GetWins(profile)
	if err != nil {
		panic(err)
	}

	renderTemplate(w, "fortnite", wins{Value: winCount})
}
