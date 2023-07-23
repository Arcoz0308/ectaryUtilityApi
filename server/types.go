package server

import (
	"ectary/handlers/mysql"
	"strconv"
	"time"
)

type PlayerInfos struct {
	Username        string     `json:"username"`
	Coins           int        `json:"coins"`
	Gems            int        `json:"gems"`
	XUID            string     `json:"xuid"`
	Friends         *string    `json:"friends"`
	Cosmetics       *string    `json:"cosmetics"`
	Rank            string     `json:"rank"`
	RankId          int        `json:"rank_id"`
	GameVersion     *string    `json:"game_version"`
	Language        *string    `json:"language"`
	DeviceOS        string     `json:"device_os"`
	FirstConnection *time.Time `json:"first_connection"`
	LastConnection  *time.Time `json:"last_connection"`
}

type FullInfos struct {
	Username        string        `json:"username"`
	Coins           int           `json:"coins"`
	Gems            int           `json:"gems"`
	XUID            string        `json:"xuid"`
	Friends         *string       `json:"friends"`
	Cosmetics       *string       `json:"cosmetics"`
	Rank            string        `json:"rank"`
	RankId          int           `json:"rank_id"`
	GameVersion     *string       `json:"game_version"`
	Language        *string       `json:"language"`
	DeviceOS        string        `json:"device_os"`
	FirstConnection *time.Time    `json:"first_connection"`
	LastConnection  *time.Time    `json:"last_connection"`
	BedWars         BedWarsInfos  `json:"bedwars_stats"`
	Practice        PracticeInfos `json:"practice_stats"`
}

type BedWarsInfos struct {
	Wins         int `json:"wins"`
	Losses       int `json:"losses"`
	WinStreak    int `json:"win_streak"`
	BrokenBeds   int `json:"broken_beds"`
	Kills        int `json:"kills"`
	FinalKills   int `json:"final_kills"`
	RankedPoints int `json:"ranked_points"`
}
type PracticeInfos struct {
	Wins   int `json:"wins"`
	Losses int `json:"losses"`
	Elo    int `json:"elo"`
	Kills  int `json:"kills"`
	Deaths int `json:"deaths"`
}

func BedWarsInfosFromMysql(d mysql.BedwarsStats) BedWarsInfos {
	return BedWarsInfos{
		Wins:         d.BedwarsWinCount,
		Losses:       d.BedwarsLossesNumber,
		WinStreak:    d.BedwarsWinStreak,
		BrokenBeds:   d.BedwarsBedBrokenNumber,
		Kills:        d.BedwarsKills,
		FinalKills:   d.BedwarsFinalKills,
		RankedPoints: d.BedwarsRankedPoints,
	}
}
func PracticeInfosFromMysql(d mysql.PracticeStats) PracticeInfos {
	return PracticeInfos{
		Wins:   d.PracticeWinCount,
		Losses: d.PracticeLossesNumber,
		Elo:    d.PracticeElo,
		Kills:  d.PracticeKills,
		Deaths: d.PracticeDeaths,
	}
}
func PlayerInfosFromMysql(d mysql.PlayerData) PlayerInfos {
	p := PlayerInfos{
		Username: d.GameUsername,
		Coins:    d.GameCoins,
		Gems:     d.GameGems,
		XUID:     d.GameXUID,
	}
	p.Friends = nil
	p.Cosmetics = nil
	if d.GameRank.Valid {
		p.Rank = mysql.Ranks[d.GameRank.Int32].RankName
		p.RankId = int(d.GameRank.Int32)
	} else {
		p.Rank = mysql.DefaultRank.RankName
		p.RankId = mysql.DefaultRank.Id
	}
	if d.GameVersion.Valid {
		p.GameVersion = &d.GameVersion.String
	} else {
		p.GameVersion = nil
	}
	if d.GameLanguage.Valid {
		p.Language = &d.GameLanguage.String
	} else {
		p.Language = nil
	}
	if d.GameDeviceOS.Valid {
		p.DeviceOS = getOs(d.GameDeviceOS.String)
	} else {
		p.DeviceOS = OS[0]
	}
	if d.GameDateFirstConnection.Valid {
		p.FirstConnection = &d.GameDateFirstConnection.Time
	} else {
		p.FirstConnection = nil
	}
	if d.GameDataLastConnection.Valid {
		p.LastConnection = &d.GameDataLastConnection.Time
	} else {
		p.LastConnection = nil
	}
	return p
}
func PlayerFullInfosFromMysql(d mysql.PlayerDataFull) FullInfos {
	p := FullInfos{
		Username: d.GameUsername,
		Coins:    d.GameCoins,
		Gems:     d.GameGems,
		XUID:     d.GameXUID,
		BedWars: BedWarsInfos{
			Wins:         d.BedwarsWinCount,
			Losses:       d.BedwarsLossesNumber,
			WinStreak:    d.BedwarsWinStreak,
			BrokenBeds:   d.BedwarsBedBrokenNumber,
			Kills:        d.BedwarsKills,
			FinalKills:   d.BedwarsFinalKills,
			RankedPoints: d.BedwarsRankedPoints,
		},
		Practice: PracticeInfos{
			Wins:   d.PracticeWinCount,
			Losses: d.PracticeLossesNumber,
			Elo:    d.PracticeElo,
			Kills:  d.PracticeKills,
			Deaths: d.PracticeDeaths,
		},
	}
	p.Friends = nil
	p.Cosmetics = nil
	if d.GameRank.Valid {
		p.Rank = mysql.Ranks[d.GameRank.Int32].RankName
		p.RankId = int(d.GameRank.Int32)
	} else {
		p.Rank = mysql.DefaultRank.RankName
		p.RankId = mysql.DefaultRank.Id
	}
	if d.GameVersion.Valid {
		p.GameVersion = &d.GameVersion.String
	} else {
		p.GameVersion = nil
	}
	if d.GameLanguage.Valid {
		p.Language = &d.GameLanguage.String
	} else {
		p.Language = nil
	}
	if d.GameDeviceOS.Valid {
		p.DeviceOS = getOs(d.GameDeviceOS.String)
	} else {
		p.DeviceOS = OS[0]
	}
	if d.GameDateFirstConnection.Valid {
		p.FirstConnection = &d.GameDateFirstConnection.Time
	} else {
		p.FirstConnection = nil
	}
	if d.GameDataLastConnection.Valid {
		p.LastConnection = &d.GameDataLastConnection.Time
	} else {
		p.LastConnection = nil
	}
	return p
}

var OS = []string{
	"Unknown",
	"Android",
	"IOS",
	"OSX",
	"Amazon",
	"Gear_VR",
	"Hololens",
	"Windows 10",
	"WIN32",
	"Dedicated",
	"TVOS",
	"Playstation",
	"Nintendo",
	"Xbox",
	"Windows Phone",
}

func getOs(osInt string) string {
	i, err := strconv.Atoi(osInt)
	if err != nil {
		return OS[0]
	}
	if i < 1 || i > 14 {
		return OS[0]
	}
	return OS[i]
}
