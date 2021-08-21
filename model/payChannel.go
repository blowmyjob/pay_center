package model

type payChannel struct {
	Id         int64  `gorm:"id"`
	PayName    string `gorm:"pay_name"`
	PayCode    string `gorm:"pay_code"`
	AppId      string `gorm:"app_id"`
	AppSecret  string `gorm:"app_secret"`
	MerchantId string `gorm:"merchant_id"`
	Status     string `gorm:"status"`
}
