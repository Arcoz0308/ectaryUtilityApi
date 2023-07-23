package mysql

import (
	"database/sql"
	"ectary/handlers/config"
	error2 "ectary/utils/error"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init() {
	c := config.C.Database
	conf := mysql.Config{
		User:                 c.User,
		Passwd:               c.Password,
		Net:                  c.NetType,
		Addr:                 c.Host,
		DBName:               c.DbName,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	d, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		panic(err)
	}
	err = d.Ping()
	if err != nil {
		panic(err)
	}
	DB = d
	loadPrepares()
	LoadRanks()
}

func PlayerSkinByUsername(username string) ([]uint8, error) {
	r := prepares["select-skin-by-username"].QueryRow(username)
	s := sql.NullString{}
	err := r.Scan(&s)
	if err != nil {
		return nil, err
	}
	if s.Valid && len(s.String) > 0 {
		return decompressSkin([]byte(s.String))
	}
	return nil, &error2.ApiError{
		Code:           error2.CodeErrNoSkin,
		AdditionalData: nil,
		Err:            error2.ErrNoSkin,
	}
}

func PlayerSkinByXUID(xuid string) ([]uint8, error) {
	r := prepares["select-skin-by-xuid"].QueryRow(xuid)
	s := sql.NullString{}
	err := r.Scan(&s)
	if err != nil {
		return nil, err
	}
	if s.Valid && len(s.String) > 0 {
		return decompressSkin([]byte(s.String))
	}
	return nil, &error2.ApiError{
		Code:           error2.CodeErrNoSkin,
		AdditionalData: nil,
		Err:            error2.ErrNoSkin,
	}
}

func PlayerCapeByUsername(username string) ([]uint8, error) {
	r := prepares["select-cape-by-username"].QueryRow(username)
	s := sql.NullString{}
	err := r.Scan(&s)
	if err != nil {
		return nil, err
	}
	if s.Valid && len(s.String) > 0 {
		return decompressCape([]byte(s.String))
	}
	return nil, &error2.ApiError{
		Code:           error2.CodeErrNoCape,
		AdditionalData: nil,
		Err:            error2.ErrNoCape,
	}
}

func PlayerCapeByXUID(xuid string) ([]uint8, error) {
	r := prepares["select-cape-by-xuid"].QueryRow(xuid)
	s := sql.NullString{}
	err := r.Scan(&s)
	if err != nil {
		return nil, err
	}
	if s.Valid && len(s.String) > 0 {
		return decompressCape([]byte(s.String))
	}
	return nil, &error2.ApiError{
		Code:           error2.CodeErrNoCape,
		AdditionalData: nil,
		Err:            error2.ErrNoCape,
	}
}

func ranks() (ranks []*RankData) {
	r, err := prepares["select-ranks"].Query()
	if err != nil {
		panic(err)
	}
	defer r.Close()
	for r.Next() {
		var rd RankData
		err = r.Scan(&rd.Id, &rd.RankName, &rd.RankDefault)
		if err != nil {
			panic(err)
		}
		ranks = append(ranks, &rd)
	}
	return
}

func PlayerFullInfosByUsername(username string) (PlayerDataFull, error) {
	r := prepares["select-player-info-full-by-username"].QueryRow(username)
	var d PlayerDataFull
	err := r.Scan(
		&d.GameUsername,
		&d.GameCoins,
		&d.GameGems,
		&d.GameXUID,
		&d.GameFriends,
		&d.GameCosmetics,
		&d.GameRank,
		&d.GameVersion,
		&d.GameLanguage,
		&d.GameDeviceOS,
		&d.GameDateFirstConnection,
		&d.GameDataLastConnection,
		&d.BedwarsWinCount,
		&d.BedwarsLossesNumber,
		&d.BedwarsWinStreak,
		&d.BedwarsBedBrokenNumber,
		&d.BedwarsFinalKills,
		&d.BedwarsKills,
		&d.BedwarsRankedPoints,
		&d.PracticeWinCount,
		&d.PracticeLossesNumber,
		&d.PracticeElo,
		&d.PracticeKills,
		&d.PracticeDeaths,
	)
	if err != nil {
		return PlayerDataFull{}, err
	}
	return d, nil
}

func PlayerFullInfosByXUID(xuid string) (PlayerDataFull, error) {
	r := prepares["select-player-info-full-by-xuid"].QueryRow(xuid)
	var d PlayerDataFull
	err := r.Scan(
		&d.GameUsername,
		&d.GameCoins,
		&d.GameGems,
		&d.GameXUID,
		&d.GameFriends,
		&d.GameCosmetics,
		&d.GameRank,
		&d.GameVersion,
		&d.GameLanguage,
		&d.GameDeviceOS,
		&d.GameDateFirstConnection,
		&d.GameDataLastConnection,
		&d.BedwarsWinCount,
		&d.BedwarsLossesNumber,
		&d.BedwarsWinStreak,
		&d.BedwarsBedBrokenNumber,
		&d.BedwarsFinalKills,
		&d.BedwarsKills,
		&d.BedwarsRankedPoints,
		&d.PracticeWinCount,
		&d.PracticeLossesNumber,
		&d.PracticeElo,
		&d.PracticeKills,
		&d.PracticeDeaths,
	)
	if err != nil {
		return PlayerDataFull{}, err
	}
	return d, nil
}

func PlayerInfosByUsername(username string) (PlayerData, error) {
	r := prepares["select-player-info-by-username"].QueryRow(username)
	var d PlayerData
	err := r.Scan(
		&d.GameUsername,
		&d.GameCoins,
		&d.GameGems,
		&d.GameXUID,
		&d.GameFriends,
		&d.GameCosmetics,
		&d.GameRank,
		&d.GameVersion,
		&d.GameLanguage,
		&d.GameDeviceOS,
		&d.GameDateFirstConnection,
		&d.GameDataLastConnection,
	)
	if err != nil {
		return PlayerData{}, err
	}
	return d, nil
}

func PlayerInfosByXUID(xuid string) (PlayerData, error) {
	r := prepares["select-player-info-by-xuid"].QueryRow(xuid)
	var d PlayerData
	err := r.Scan(
		&d.GameUsername,
		&d.GameCoins,
		&d.GameGems,
		&d.GameXUID,
		&d.GameFriends,
		&d.GameCosmetics,
		&d.GameRank,
		&d.GameVersion,
		&d.GameLanguage,
		&d.GameDeviceOS,
		&d.GameDateFirstConnection,
		&d.GameDataLastConnection,
	)
	if err != nil {
		return PlayerData{}, err
	}
	return d, nil
}

func PlayerBedwarsInfosByUsername(username string) (BedwarsStats, error) {
	r := prepares["select-player-info-bedwars-by-username"].QueryRow(username)
	var d BedwarsStats
	err := r.Scan(
		&d.GameXUID,
		&d.BedwarsWinCount,
		&d.BedwarsLossesNumber,
		&d.BedwarsWinStreak,
		&d.BedwarsBedBrokenNumber,
		&d.BedwarsFinalKills,
		&d.BedwarsKills,
		&d.BedwarsRankedPoints,
	)
	if err != nil {
		return BedwarsStats{}, err
	}
	return d, nil
}

func PlayerBedwarsInfosByXUID(xuid string) (BedwarsStats, error) {
	r := prepares["select-player-info-bedwars-by-xuid"].QueryRow(xuid)
	var d BedwarsStats
	err := r.Scan(
		&d.GameXUID,
		&d.BedwarsWinCount,
		&d.BedwarsLossesNumber,
		&d.BedwarsWinStreak,
		&d.BedwarsBedBrokenNumber,
		&d.BedwarsFinalKills,
		&d.BedwarsKills,
		&d.BedwarsRankedPoints,
	)
	if err != nil {
		return BedwarsStats{}, err
	}
	return d, nil
}

func PlayerPracticeInfosByUsername(username string) (PracticeStats, error) {
	r := prepares["select-player-info-practice-by-username"].QueryRow(username)
	var d PracticeStats
	err := r.Scan(
		&d.GameXUID,
		&d.PracticeWinCount,
		&d.PracticeLossesNumber,
		&d.PracticeElo,
		&d.PracticeKills,
		&d.PracticeDeaths,
	)
	if err != nil {
		return PracticeStats{}, err
	}
	return d, nil
}

func PlayerPracticeInfosByXUID(xuid string) (PracticeStats, error) {
	r := prepares["select-player-info-practice-by-xuid"].QueryRow(xuid)
	var d PracticeStats
	err := r.Scan(
		&d.GameXUID,
		&d.PracticeWinCount,
		&d.PracticeLossesNumber,
		&d.PracticeElo,
		&d.PracticeKills,
		&d.PracticeDeaths,
	)
	if err != nil {
		return PracticeStats{}, err
	}
	return d, nil
}
