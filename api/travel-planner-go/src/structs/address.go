package structs

type Address struct {
	Metadata
	Street       string `json:"street" gorm:"type:text"`
	StreetNumber string `json:"street_number" gorm:"type:text"`
	Neighborhood string `json:"neighborhood" gorm:"type:text"`
	Complement   string `json:"complement" gorm:"type:text"`
	City         string `json:"city" gorm:"type:text"`
	State        string `json:"state" gorm:"type:text"`
	Country      string `json:"country_name" gorm:"type:text"`
	ZipCode      string `json:"zip_code" gorm:"type:text"`
}

func (a Address) TableName() string {
	return "addresses"
}
