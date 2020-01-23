package summoners

import (
    "encoding/json"
    "github.com/hiimkina/riotapi-scrapper/utils"
    "log"
)

type RankedInfo struct {
    QueueType string
    Wins int
    Losses int
    Rank string
    Tier string
    LeaguePoints int
}

type Summoner struct {
    Id string           `json:"id"`
    AccountId string    `json:"accountId"`
    Puuid string        `json:"puuid"`
    Name string         `json:"name"`
    ProfileIconId int   `json:"profileIconId"`
    RevisionDate int    `json:"revisionDate"`
    SummonerLevel int   `json:"summonerLevel"`
    Ranks []RankedInfo
}

func GetSummonerFromName ( summonerName string ) *Summoner {
    apiEndpoint :=  "https://euw1.api.riotgames.com/lol/summoner/v4/summoners/by-name/" + summonerName

    stringOutput := utils.DoHttpRequest(apiEndpoint)
    summoner := new(Summoner)
    err := json.Unmarshal([]byte(stringOutput), &summoner)
    if err != nil {
        log.Fatal("Error converting JSON to struct in GetMatch")
    }
    summoner = getSummonerRankedInfo(summoner)

    return summoner
}

func GetSummonerFromId ( summonerId string ) *Summoner {
    apiEndpoint :=  "https://euw1.api.riotgames.com/lol/summoner/v4/summoners/" + summonerId

    stringOutput := utils.DoHttpRequest(apiEndpoint)
    summoner := new(Summoner)
    err := json.Unmarshal([]byte(stringOutput), &summoner)
    if err != nil {
        log.Fatal("Error converting JSON to struct in GetMatch")
    }
    summoner = getSummonerRankedInfo(summoner)

    return summoner
}

func getSummonerRankedInfo ( summoner *Summoner ) *Summoner {
    apiEndpoint := "https://euw1.api.riotgames.com/lol/league/v4/entries/by-summoner/" + summoner.Id

    stringOutput := utils.DoHttpRequest(apiEndpoint)
    var rankedQueues []RankedInfo
    err := json.Unmarshal([]byte(stringOutput), &rankedQueues)
    if err != nil {
        log.Fatal("Error converting JSON to struct in GetMatch")
    }
    summoner.Ranks = rankedQueues

    return summoner
}
