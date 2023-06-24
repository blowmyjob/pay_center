package model

type Activity struct {
	Id int `gorm:"column:id"`
	A  int `gorm:"column:a"`
	B  int `gorm:"column:b"`
}
