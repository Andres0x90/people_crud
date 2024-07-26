package domain

import "time"

type Payroll struct {
	ID        string
	From      time.Time
	To        time.Time
	WorkHours float32
	Value     float32
	IssuedAt  time.Time
	IssuedBy  string
}
