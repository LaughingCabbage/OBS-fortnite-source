package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	tracker "github.com/LaughingCabbage/fortnite-tracker/v1"
	"github.com/LaughingCabbage/tracker-bot/key"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("start")

	router := mux.NewRouter()

	router.HandleFunc("/obs/laughingcabbage", handleFortniteData)
	router.PathPrefix("/obs/assets/").Handler(http.StripPrefix("/obs/assets/", http.FileServer(http.Dir("assets"))))
	log.Fatal(http.ListenAndServe(":8000", router))
}

// Data holds response data to serve as json
type Data struct {
	Wins  int     `json:"wins"`
	KDR   float64 `json:"kdr"`
	Kills int     `json:"kills"`
}

func handleFortniteData(w http.ResponseWriter, r *http.Request) {
	Key := key.LoadKey(".key")
	profile, err := tracker.GetProfile("pc", "laughingcabbage", Key.Value)
	if err != nil {
		panic(err)
	}
	data := Data{}

	kills, err := tracker.GetKills(profile)
	handleError(err)
	data.Kills = kills

	wins, err := tracker.GetWins(profile)
	handleError(err)
	data.Wins = wins

	kdr, err := tracker.GetCurrentKDR(profile)
	handleError(err)
	data.KDR = kdr

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
