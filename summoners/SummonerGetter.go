package summoners

import (
    "fmt"
    "net/http"
    "encoding/json"
)

type Summoner struct {
    Id string
    AccountId string
    Puuid string
    Name string
    ProfileIconId int
    RevisionDate int
    SummonerLevel int
}

func ConvertSummoner ( toConvert *http.Response, target interface{} ) error {
    return json.NewDecoder(toConvert.Body).Decode(target)
}

func GetSummonerInfo ( summonerName string ) {
    var apiEndpoint string

    apiEndpoint =  "https://euw1.api.riotgames.com/lol/summoner/v4/summoners/by-name/" + summonerName

    req, err := http.NewRequest("GET", apiEndpoint, nil)
    if err != nil {
       	fmt.Printf("Error generating HTTP request")
    }
    req.Header.Set("X-Riot-Token", "RGAPI-80136d22-3d28-43ff-a80c-fee11ce34569")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Printf("Error doing the HTTP request")
    }

    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
        summoner := new(Summoner)
        ConvertSummoner(resp, summoner)

        fmt.Printf("%+v", summoner)
    }
}
