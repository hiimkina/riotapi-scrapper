package summoners

import (
    "encoding/json"
    "github.com/hiimkina/riotapi-scrapper/utils"
)

type SummonerToAnalyse struct {
    LeagueId string
    SummonerId string
}

// https://euw1.api.riotgames.com/lol/league/v4/entries/RANKED_SOLO_5x5/DIAMOND/I?page=1
func GatherSummoners ( queueType string, tier string, rank string ) []SummonerToAnalyse{
    apiEndpoint :=  "https://euw1.api.riotgames.com/lol/league/v4/entries/" + queueType + "/" + tier + "/" + rank

    stringOutput := utils.DoHttpRequest(apiEndpoint)
    var summonersToAnalyse []SummonerToAnalyse
    json.Unmarshal([]byte(stringOutput), &summonersToAnalyse)

    return summonersToAnalyse
}
