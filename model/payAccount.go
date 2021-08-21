package model

type PayAccount struct {
	Id         int64   `gorm:"id"`
	AppId      string  `gorm:"app_id"`
	Amount     float64 `gorm:"amount"`
	Status     string  `gorm:"status"`
	CreateTime int64   `gorm:"create_time"`
	UpdateTime int64   `gorm:"update_time"`
}
