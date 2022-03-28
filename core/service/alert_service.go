package service

import (
	"alert/core/dao"
	"alert/core/dao/alert_dao"
	"log"
	"time"
)

const ALERTTABLENAME = "alerts"

type AlertServiceImpl struct{}

func (i AlertServiceImpl) Add(taskName, ruleName, ruleCode, expression, value string, roomId uint, time time.Time) bool {
	var count int64
	dao.DB.Debug().Table(ALERTTABLENAME).Where("rule_code = ? AND room_id = ?", ruleCode, roomId).Count(&count)
	if count != 0 {
		resp := dao.DB.Debug().Table(ALERTTABLENAME).Where("rule_code = ? AND room_id = ?", ruleCode, roomId).Update("time", time)
		if resp.Error != nil {
			log.Println(resp.Error)
			return false
		}
		return true
	}
	alert := alert_dao.Alert{
		RuleCode:   ruleCode,
		TaskName:   taskName,
		RuleName:   ruleName,
		Expression: expression,
		Value:      value,
		RoomID:     roomId,
		Time:       time,
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
