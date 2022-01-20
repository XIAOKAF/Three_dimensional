package model

import "time"

type Transfer struct {
	PaymentId  int
	Payer      string
	Payee      string
	Amount     float32
	Postscript string
	Time       time.Time
}
