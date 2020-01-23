package main

import (
    "fmt"
    "github.com/hiimkina/riotapi-scrapper/summoners"
)

func main () {
    var summonerToScrap = new(summoners.Summoner)
    summonerToScrap = summoners.GetSummonerInfo("GV kina")

    fmt.Printf("%+v\n", summonerToScrap)
    fmt.Printf("%+v\n", summonerToScrap.Ranks)
}
