package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"../db"
)

//AddOpenBet is used for check server state
var AddOpenBet = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("matchID")
	team1 := request.URL.Query().Get("team1")
	team2 := request.URL.Query().Get("team2")
	group, _ := strconv.Atoi(request.URL.Query().Get("group"))
	//Insert into database
	db.InsertOpenBet(id, team1, team2, group)
	fmt.Fprintf(response, "ok")
})
