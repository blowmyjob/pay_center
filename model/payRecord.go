package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type PayRecord struct {
	Id         int64           `gorm:"id"`
	SeqId      int64           `gorm:"seq_id"`
	AppId      string          `gorm:"app_id"`
	ProductId  int64           `gorm:"product_id"`
	UserId     int64           `gorm:"user_id"`
	PayAmount  decimal.Decimal `gorm:"pay_amount"`
	PayStatus  string          `gorm:"pay_status"`
	PayMethod  string          `gorm:"pay_method"`
	CreateTime time.Time       `gorm:"create_time"`
	UpdateTime time.Time       `gorm:"update_time"`
}
