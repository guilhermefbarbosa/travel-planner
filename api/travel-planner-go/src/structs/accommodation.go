package structs

import (
	"github.com/google/uuid"
)

type AccommodationType string

const (
	Apartment    AccommodationType = "apartment"
	CouchSurfing AccommodationType = "couch_surfing"
	GuestHouse   AccommodationType = "guest_house"
	HolidayHome  AccommodationType = "holiday_home"
	Homestay     AccommodationType = "homestay"
	Hotel        AccommodationType = "hotel"
	Hostel       AccommodationType = "hostel"
	Resort       AccommodationType = "resort"
)

type Accommodation struct {
	Metadata
	Type        AccommodationType `json:"type" gorm:"type:text"`
	StartsAt    Date              `json:"starts_at" gorm:"type:date"`
	EndsAt      Date              `json:"ends_at" gorm:"type:date"`
	Price       Money             `json:"price" gorm:"type:text"`
	IsPaid      bool              `json:"is_paid" gorm:"type:boolean"`
	Name        string            `json:"name" gorm:"type:text"`
	Description string            `json:"description" gorm:"type:text"`
	PaymentID   uuid.UUID         `json:"payment_id" gorm:"type:uuid REFERENCES payments(id)"`
	TravelID    uuid.UUID         `json:"travel_id" gorm:"type:uuid REFERENCES travels(id)"`
	AddressID   uuid.UUID         `json:"address_id" gorm:"type:uuid REFERENCES addresses(id)"`
}

func (a Accommodation) TableName() string {
	return "accommodations"
}
