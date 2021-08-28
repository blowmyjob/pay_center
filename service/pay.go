package service

import (
	"pay_center/config"
	"pay_center/model"
	"pay_center/utils"
	"time"
)

var IdWorker = new(utils.IdWorker)

type PayService struct {
}

func (payService *PayService) InsertPayOrder(u model.PayRecord) (reqNo int64, err error) {
	id, _ := IdWorker.NextId()
	u.Id = id
	u.SeqId = id
	u.CreateTime = time.Now()
	u.UpdateTime = time.Now()
	err = config.GVA_DB.Table("pay_record").Create(&u).Error
	return id, err
}

func (payService *PayService) UpdateOrderStatus(id int64, status string) bool {
	total := 0
	config.GVA_DB.Table("pay_record").Select("count(*) as total").Count(&total)
	if total == 0 {
		return false
	}
	tx := config.GVA_DB.Begin()
	err := tx.Table("pay_record").Where("id=?", id).Update("status", status).Error
	if err != nil {
		tx.Callback()
		return false
	}
	tx.Commit()
	return true
}

func (payService *PayService) SelectPayRecord(id int64) model.PayRecord {
	var payRecord model.PayRecord
	config.GVA_DB.Table("pay_record").Where("id", id).Model(&payRecord)
	return payRecord
}
