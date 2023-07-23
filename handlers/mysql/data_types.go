package mysql

import "database/sql"

type RankData struct {
	Id          int
	RankName    string
	RankDefault bool
}
type PlayerData struct {
	GameUsername            string
	GameCoins               int
	GameGems                int
	GameXUID                string
	GameFriends             sql.NullString // soon
	GameCosmetics           sql.NullString // soon
	GameRank                sql.NullInt32
	GameVersion             sql.NullString
	GameLanguage            sql.NullString
	GameDeviceOS            sql.NullString
	GameDateFirstConnection sql.NullTime
	GameDataLastConnection  sql.NullTime
}
type BedwarsStats struct {
	GameXUID               string
	BedwarsWinCount        int
	BedwarsLossesNumber    int
	BedwarsWinStreak       int
	BedwarsBedBrokenNumber int
	BedwarsFinalKills      int
	BedwarsKills           int
	BedwarsRankedPoints    int
}
type PracticeStats struct {
	GameXUID             string
	PracticeWinCount     int
	PracticeLossesNumber int
	PracticeElo          int
	PracticeKills        int
	PracticeDeaths       int
}
type PlayerDataFull struct {
	GameUsername            string
	GameCoins               int
	GameGems                int
	GameXUID                string
	GameFriends             sql.NullString // soon
	GameCosmetics           sql.NullString // soon
	GameRank                sql.NullInt32
	GameVersion             sql.NullString
	GameLanguage            sql.NullString
	GameDeviceOS            sql.NullString
	GameDateFirstConnection sql.NullTime
	GameDataLastConnection  sql.NullTime
	BedwarsWinCount         int
	BedwarsLossesNumber     int
	BedwarsWinStreak        int
	BedwarsBedBrokenNumber  int
	BedwarsFinalKills       int
	BedwarsKills            int
	BedwarsRankedPoints     int
	PracticeWinCount        int
	PracticeLossesNumber    int
	PracticeElo             int
	PracticeKills           int
	PracticeDeaths          int
}
