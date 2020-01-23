package summoners

import (
    "encoding/json"
    "github.com/hiimkina/riotapi-scrapper/utils"
    "log"
    "strconv"
)

type SummonerToAnalyse struct {
    LeagueId string
    SummonerId string
}

// https://euw1.api.riotgames.com/lol/league/v4/entries/RANKED_SOLO_5x5/DIAMOND/I?page=1
func GatherSummoners ( queueType string, tier string, rank string, page int ) []SummonerToAnalyse{
    apiEndpoint :=  "https://euw1.api.riotgames.com/lol/league/v4/entries/" + queueType + "/" + tier + "/" + rank + "?page=" + strconv.Itoa(page)

    stringOutput := utils.DoHttpRequest(apiEndpoint)
    var summonersToAnalyse []SummonerToAnalyse
    err := json.Unmarshal([]byte(stringOutput), &summonersToAnalyse)
    if err != nil {
        log.Fatal("Error converting JSON to struct in GetMatch")
    }

    return summonersToAnalyse
}

