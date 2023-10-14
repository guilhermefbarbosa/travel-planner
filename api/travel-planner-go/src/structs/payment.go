package structs

import (
	"github.com/google/uuid"
	"time"
)

type PaymentMethod string

const (
	CreditCard  PaymentMethod = "credit_card"
	DebitCard   PaymentMethod = "debit_card"
	PayPal      PaymentMethod = "paypal"
	PaymentSlip PaymentMethod = "slip"
	Cash        PaymentMethod = "cash"
	Pix         PaymentMethod = "pix"
)

type Payment struct {
	Metadata
	TravelID    uuid.UUID     `json:"travel_id" gorm:"type:uuid REFERENCES travels(id)"`
	Method      PaymentMethod `json:"method" gorm:"type:text"`
	Description string        `json:"description" gorm:"type:text"`
	Amount      Money         `json:"amount" gorm:"type:text"`
	PaidAt      time.Time     `json:"paid_at" gorm:"type:timestamp"`
}

func (p Payment) TableName() string {
	return "payments"
}
