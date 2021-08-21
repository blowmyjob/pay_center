package service

import (
	"pay_center/config"
	"pay_center/model"
	"pay_center/utils"
)

var IdWorker = new(utils.IdWorker)

type PayService struct {
}

func (payService *PayService) insertPayOrder(u model.PayRecord) (reqNo int64, err error) {
	id, _ := IdWorker.NextId()
	u.Id = id
	u.SeqId = id
	err = config.GVA_DB.Table("payRecord").Create(&u).Error
	return id, err
}

func (payService *PayService) updateOrderStatus(id int64, status string) bool {
	err := config.GVA_DB.Table("payRecord").Where("id=?", id).Update("status", status).Error
	if err != nil {
		return false
	}
	return true
}
