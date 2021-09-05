package service

import (
	"pay_center/config"
	"pay_center/model"
)

type PayAccountService struct {
}

func (payAccountService *PayAccountService) insertAccount(account model.PayAccount) (err error) {
	err = config.GVA_DB.Table("pay_account").Create(&account).Error
	return err
}

func (payAccountService *PayAccountService) selectAccount(id int64) (account model.PayAccount) {
	config.GVA_DB.Table("pay_account").Where("id=?", id).First(&account)
	return account
}

func (payAccountService *PayAccountService) updateStatus(id int64, status string) (err error) {
	err = config.GVA_DB.Table("pay_account").Where("id=?", id).Update("status=?", status).Error
	return err
}
