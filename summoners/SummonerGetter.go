package summoners

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

type RankedInfo struct {
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

func GetSummonerInfo ( summonerName string ) *Summoner {
    apiEndpoint :=  "https://euw1.api.riotgames.com/lol/summoner/v4/summoners/by-name/" + summonerName

    stringOutput := doHttpRequest(apiEndpoint)
    summoner := new(Summoner)
    json.Unmarshal([]byte(stringOutput), &summoner)
    summoner = getSummonerRankedInfo(summoner)

    return summoner
}



func getSummonerRankedInfo ( summoner *Summoner ) *Summoner {
    apiEndpoint := "https://euw1.api.riotgames.com/lol/league/v4/entries/by-summoner/" + summoner.Id

    stringOutput := doHttpRequest(apiEndpoint)
    var rankedQueues []RankedInfo
    json.Unmarshal([]byte(stringOutput), &rankedQueues)
    summoner.Ranks = rankedQueues

    return summoner
}

func doHttpRequest ( apiEndpoint string ) string {
    apiToken := "RGAPI-80136d22-3d28-43ff-a80c-fee11ce34569"
    req, err := http.NewRequest("GET", apiEndpoint, nil)
    if err != nil {
        fmt.Printf("Error generating HTTP request")
    }
    req.Header.Set("X-Riot-Token", apiToken)

    response, err := http.DefaultClient.Do(req)
    defer response.Body.Close()
    if err != nil {
        fmt.Printf("Error doing the HTTP request")
    }

    if response.StatusCode != http.StatusOK {
        log.Fatal("Error in your request: %s", response.StatusCode)
    }

    bodyBytes, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    return string(bodyBytes)
}
