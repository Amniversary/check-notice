package mysql

import (
	"time"
	"log"
)

func UpTimeOutNoticeToken() bool {
	err := db.Table("NoticeToken").Where("expire < ?", time.Now().Unix()).Update("status", 0).Error
	if err != nil {
		log.Printf("update expire time out token err: [%v]", err)
		return false
	}
	return true
}