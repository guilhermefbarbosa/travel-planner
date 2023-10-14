package structs

import (
	"github.com/google/uuid"
	"time"
)

type JourneyType string

const (
	Flight JourneyType = "flight"
	Train  JourneyType = "train"
	Bus    JourneyType = "bus"
	Car    JourneyType = "car"
	Ride   JourneyType = "ride"
	Bike   JourneyType = "bike"
	Hike   JourneyType = "hike"
)

type Journey struct {
	Metadata
	TravelID      uuid.UUID   `json:"travel_id" gorm:"type:uuid REFERENCES travels(id)"`
	Type          JourneyType `json:"type" gorm:"type:text"`
	Distance      int         `json:"distance" gorm:"type:integer"`
	DepartureTime time.Time   `json:"departure_time" gorm:"type:timestamp"`
	ArrivalTime   time.Time   `json:"arrival_time" gorm:"type:timestamp"`
	From          string      `json:"from" gorm:"type:text"`
	To            string      `json:"to" gorm:"type:text"`
	Price         Money       `json:"price" gorm:"type:text"`
	PaymentID     uuid.UUID   `json:"payment_id" gorm:"type:uuid REFERENCES payments(id)"`
	IsPaid        bool        `json:"is_paid" gorm:"type:boolean"`
	Description   string      `json:"description" gorm:"type:text"`
	Details       string      `json:"details" gorm:"type:text"`
}

func (j Journey) TableName() string {
	return "journeys"
}
