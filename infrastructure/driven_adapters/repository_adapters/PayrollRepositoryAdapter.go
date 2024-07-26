package repository_adapters

import "time"

type PayrollModel struct {
	ID        string `gorm:"primaryKey"`
	Document  string
	From      time.Time
	To        time.Time
	WorkHours float32
	Value     float32
	IssuedAt  time.Time
	IssuedBy  string
}

func (*PayrollModel) TableName() string {
	return "payrolls"
}
