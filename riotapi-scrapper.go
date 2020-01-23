package main

import (
    "fmt"
    "github.com/hiimkina/riotapi-scrapper/summoners"
)

func main () {

    var summonersGathered []summoners.SummonerToAnalyse
    summonersGathered = summoners.GatherSummoners("RANKED_SOLO_5x5","DIAMOND","I")

    fmt.Printf("%+v\n\n", summonersGathered[0])

    for i := 0; i < 10; i++ {
        summonerToScrap := new(summoners.Summoner)
        fmt.Printf("\nScrapping: %s\n", summonersGathered[i].SummonerId)
        summonerToScrap = summoners.GetSummonerFromId(summonersGathered[i].SummonerId)
        fmt.Printf("%+v\n\n", summonerToScrap)
    }

    //summonerToScrap := new(summoners.Summoner)
    //summonerToScrap = summoners.GetSummonerInfo("GV kina")

    //fmt.Printf("%+v\n", summonerToScrap)
    //fmt.Printf("%+v\n", summonerToScrap.Ranks)
}
