package main

import (
    "fmt"
    "github.com/hiimkina/riotapi-scrapper/matches"
    "github.com/hiimkina/riotapi-scrapper/summoners"
)

func main () {

    var summonersGathered []summoners.SummonerToAnalyse
    summonersGathered = summoners.GatherSummoners("RANKED_SOLO_5x5","DIAMOND","I", 1)

    for i := 0; i < 1; i++ {
        summonerToScrap := new(summoners.Summoner)
        summonerToScrap = summoners.GetSummonerFromId(summonersGathered[i].SummonerId)

        history := new(matches.PlayerHistory)
        history = matches.FillPlayerHistory(summonerToScrap.AccountId)
    }

}
