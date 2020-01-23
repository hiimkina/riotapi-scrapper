package matches

import (
    "encoding/json"
    "github.com/hiimkina/riotapi-scrapper/utils"
    "log"
)

type ChampionBan struct {
    PickTurn int
    ChampionId int
}

type MatchDetails struct {
    GameId int
    ParticipantIdentities []ParticipantDetails
    GameVersion string
    Teams []TeamDetails
    Participants []PlayerStats
    GameDuration int
    GameCreation int
}

type ParticipantDetails struct {
    Player PlayerDetails
    ParticipantId int
}

type PlayerAdvancedStats struct {
    Assists int
    ChampLevel int
    CombatPlayerScore int
    DamageDoneToTurrets int
    DamageDealtToObjectives int
    DamageSelfMitigated int
    Deaths int
    DoubleKills int
    FirstBloodAssist bool
    FirstBloodKill bool
    FirstTowerAssist bool
    FirstTowerKill bool
    GoldEarned bool
    GoldSpent bool
    InhibitorKills int
    Item0 int
    Item1 int
    Item2 int
    Item3 int
    Item4 int
    Item5 int
    Item6 int
    KillingSprees int
    Kills int
    LargestCriticalStrike int
    LargestKillingSpree int
    LargestMultiKill int
    LongestTimeSpentLiving int
    MagicDamageDealt int
    MagicDamageDealtToChampions int
    MagicDamageTaken int
    NeutralMinionsKilled int
    NeutralMinionsKilledTeamJungle int
    ObjectivePlayerScore int
    ParticipantId int
    PentaKills int
    Perk0 int
    Perk0Var1 int
    Perk0Var2 int
    Perk0Var3 int
    Perk1 int
    Perk1Var1 int
    Perk1Var2 int
    Perk1Var3 int
    Perk2 int
    Perk2Var1 int
    Perk2Var2 int
    Perk2Var3 int
    Perk3 int
    Perk3Var1 int
    Perk3Var2 int
    Perk3Var3 int
    Perk4 int
    Perk4Var1 int
    Perk4Var2 int
    Perk4Var3 int
    Perk5 int
    Perk5Var1 int
    Perk5Var2 int
    Perk5Var3 int
    PerkPrimaryStyle int
    PerkSubStyle int
    PhysicalDamageDealt int
    PhysicalDamageDealtToChampions int
    PhysicalDamageTaken int
    QuadraKills int
    SightWardsBoughtInGame int
    StatPerk0 int
    StatPerk1 int
    StatPerk2 int
    TimeCCingOthers int
    TotalDamageDealt int
    TotalDamageDealtToChampions int
    TotalDamageTaken int
    TotalHeal int
    TotalMinionKilled int
    TotalPlayerScore int
    TotalScoreRank int
    TotalTimeCrowdControlDealt int
    TotalUnitsHealed int
    TripleKills int
    TrueDamageDealt int
    TrueDamageDealtToChampions int
    TrueDamageTaken int
    TurretKills int
    UnrealKills int
    VisionScore int
    VisionWardsBoughtInGame int
    WardsPlaced int
    WardsKilled int
    Win bool
}

type PlayerDetails struct {
    CurrentPlatformId string
    SummonerName string
    CurrentAccountId string
    SummonerId string
}

type PlayerStats struct {
    Spell1Id int
    ParticipantId int
    Spell2Id int
    TeamId int
    Stats PlayerAdvancedStats
    ChampionId int
}

type TeamDetails struct {
    FirstDragon bool
    Bans []ChampionBan
    FirstInhibitor bool
    Win bool
    FirstRiftHerald bool
    FirstBaron bool
    BaronKills int
    RiftHeraldKills int
    FirstBlood bool
    TeamId int
    FirstTower bool
    InhibitorKills int
    TowerKills int
    DragonKills int
}

func GetMatch ( matchId int ) *PlayerHistory {
    apiEndpoint :=  "https://euw1.api.riotgames.com/lol/match/v4/matches/" + string(matchId)

    stringOutput := utils.DoHttpRequest(apiEndpoint)
    history := new(PlayerHistory)
    err := json.Unmarshal([]byte(stringOutput), &history)
    if err != nil {
        log.Fatal("Error converting JSON to struct in GetMatch")
    }

    return history
}