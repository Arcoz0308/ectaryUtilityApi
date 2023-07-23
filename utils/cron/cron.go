package cron

import (
	"ectary/handlers/mysql"
	"github.com/jasonlvhit/gocron"
)

func LoadCron() {
	err := gocron.Every(30).Minute().Do(mysql.LoadRanks)
	if err != nil {
		panic(err)
	}
}
