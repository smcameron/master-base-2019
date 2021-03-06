package baseapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/HackRVA/master-base-2019/filelogging"
	gm "github.com/HackRVA/master-base-2019/game"
)

var logger = log.Ger.With().Str("pkg", "baseapi").Logger()

// NewGame - function to schedule newgame
func NewGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var e gm.Game
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &e)
	ScheduleGame(e)

	j, _ := json.Marshal(e)
	w.Write(j)
}

// NextGame -- returns the game that is sheduled next
func NextGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	next := GetNext()
	if next.AbsStart == 0 {
		j, _ := json.Marshal("There are no games scheduled")
		w.Write(j)
	} else {
		j, _ := json.Marshal(next)
		w.Write(j)
	}
}

// AllGames - returns all scheduled games
func AllGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(GetGames())
	w.Write(j)
}
