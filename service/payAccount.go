package service

import (
	"pay_center/config"
	"pay_center/model"
)

type PayAccountService struct {
}

func (payAccountService *PayAccountService) InsertAccount(account model.PayAccount) (err error) {
	err = config.GVA_DB.Table("pay_account").Create(&account).Error
	return err
}

func (payAccountService *PayAccountService) SelectAccount(appId string) (account model.PayAccount, err error) {
	err = config.GVA_DB.Table("pay_account").Where("app_id=?", appId).First(&account).Error
	return account, err
}

func (payAccountService *PayAccountService) UpdateStatus(id int64, status string) (err error) {
	err = config.GVA_DB.Table("pay_account").Where("id=?", id).Update("status=?", status).Error
	return err
}
