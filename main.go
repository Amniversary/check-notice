package main

import (
	"log"
	"github.com/Amniversary/check-notice/config"
	"github.com/Amniversary/check-notice/config/mysql"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	c := config.NewConfig()
	mysql.NewMysql(c)
	ticker := time.NewTicker(time.Second * 5)
	for  {
		select {
		case <- ticker.C:
			go mysql.UpTimeOutNoticeToken()
		}
	}
}
