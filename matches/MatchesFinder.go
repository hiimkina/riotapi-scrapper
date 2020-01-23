package matches

import (
    "encoding/json"
    "github.com/hiimkina/riotapi-scrapper/utils"
    "log"
)

type Match struct {
    Lane string
    GameId int
    Champion int
    Timestamp int
    Queue int
    Role string
}

type PlayerHistory struct {
    Matches []Match
    endIndex int
    startIndex int
}

func FillPlayerHistory ( accountId string ) *PlayerHistory {
    apiEndpoint :=  "https://euw1.api.riotgames.com/lol/match/v4/matchlists/by-account/" + accountId

    stringOutput := utils.DoHttpRequest(apiEndpoint)
    history := new(PlayerHistory)
    err := json.Unmarshal([]byte(stringOutput), &history)
    if err != nil {
        log.Fatal("Error converting JSON to struct in GetMatch")
    }

    return history
}