package model

type payRecord struct {
	Id         int64   `gorm:"id"`
	SeqId      string  `gorm:"seq_id"`
	AppId      string  `gorm:"app_id"`
	ProductId  int64   `gorm:"product_id"`
	UserId     int64   `gorm:"user_id"`
	PayAmount  float64 `gorm:"pay_amount"`
	PayStatus  string  `gorm:"pay_status"`
	PayMethod  string  `gorm:"pay_method"`
	CreateTime int64   `gorm:"create_time"`
	UpdateTime int64   `gorm:"update_time"`
}
