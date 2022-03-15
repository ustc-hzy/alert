package service

import (
	"alert/core/dao"
	"alert/core/dao/alert_dao"
	"log"
	"time"
)

const ALERTTABLENAME = "alerts"

type AlertServiceImpl struct{}

func (i AlertServiceImpl) Add(ruleName string, roomId uint, time time.Time) bool {
	alert := alert_dao.Alert{
		RuleName: ruleName,
		RoomID:   roomId,
		Time:     time,
	}
	resp := dao.DB.Debug().Table(ALERTTABLENAME).Create(&alert)
	if resp.Error != nil {
		log.Println(resp.Error)
		return false
	}
	return true
}

func (i AlertServiceImpl) Query() []alert_dao.Alert {
	var alertList []alert_dao.Alert
	res := dao.DB.Debug().Table(ALERTTABLENAME).Find(&alertList)
	if res.Error != nil {
		log.Println(res.Error)
		return nil
	}
	return alertList
}
