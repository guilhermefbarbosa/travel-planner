package structs

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Currency string

const (
	USD Currency = "USD"
	EUR Currency = "EUR"
	GBP Currency = "GBP"
	BRL Currency = "BRL"
	JPY Currency = "JPY"
	AUD Currency = "AUD"
	CAD Currency = "CAD"
	HDK Currency = "HDK"
	RMB Currency = "RMB"
	CNY Currency = "CNY"
)

type Money struct {
	Amount   int    `json:"amount" gorm:"type:integer"`
	Currency string `json:"currency" gorm:"type:text"`
}

func (m *Money) GetAmount() int {
	return m.Amount
}

func (m *Money) String() string {
	return fmt.Sprintf("%s %f", m.Currency, float64(m.Amount/100))
}

func (m *Money) Scan(value interface{}) error {
	strValue, ok := value.(string)
	if !ok {
		return fmt.Errorf("unhandled type received: %v", value)
	}
	_, err := fmt.Sscanf(strValue, "%s %f", &m.Currency, &m.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (m *Money) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(Format)+2)
	b = append(b, '"')
	b = append(b, []byte(m.String())...)
	b = append(b, '"')
	return b, nil
}

func (m *Money) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	parts := strings.Split(string(data), " ")
	if len(parts) != 2 {
		return errors.New("invalid money format")
	}
	m.Currency = parts[0]
	floatAmount, err := strconv.ParseFloat(string(parts[1]), 64)
	if err != nil {
		return err
	}
	m.Amount = int(floatAmount * 100)
	return nil
}

func (m *Money) Value() (interface{}, error) {
	return m.String(), nil
}

func (m *Money) Add(money Money) {
	m.Amount += money.Amount
}

func (m *Money) Sub(money Money) {
	m.Amount -= money.Amount
}

func (m *Money) Mul(multiplier int) {
	m.Amount *= multiplier
}

func (m *Money) Div(divisor int) {
	m.Amount /= divisor
}
