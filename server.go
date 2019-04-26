package main

import (
	"fmt"
	"net/http"
	"os"

	api "github.com/HackRVA/master-base-2019/baseapi"
	log "github.com/HackRVA/master-base-2019/filelogging"
	lb "github.com/HackRVA/master-base-2019/leaderboard"

	ss "github.com/HackRVA/master-base-2019/serverstartup"
	"github.com/gorilla/mux"
)

var logger = log.Ger

func main() {
	uri := os.Getenv("LEADERBOARD_API")
	if uri == "" {
		os.Setenv("LEADERBOARD_API", "http://localhost:5000/api/")
	}

	r := mux.NewRouter()
	r.HandleFunc("/api/newgame", api.NewGame).Methods("POST")
	r.HandleFunc("/api/nextgame", api.NextGame).Methods("GET")
	r.HandleFunc("/api/games", api.AllGames).Methods("GET")
	http.Handle("/", r)
	fmt.Println("running web server on port 8000")
	lb.StartLeaderboardLoop()
	ss.StartBadgeWrangler()
	http.ListenAndServe(":8000", nil)
}
