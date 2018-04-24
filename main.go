package main

import (
	"log"
	"github.com/Amniversary/check-notice/config"
	"github.com/Amniversary/check-notice/config/mysql"
	"time"
	//"github.com/jinzhu/now"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	c := config.NewConfig()
	mysql.NewMysql(c)
	for {
		mysql.UpTimeOutNoticeToken()
		time.Sleep(time.Minute * 5)
	}
}
