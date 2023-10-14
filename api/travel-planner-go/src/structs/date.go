package structs

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
	"time"
)

const Format = "2006-01-02"

type Date struct {
	time time.Time
}

func (d *Date) Scan(value interface{}) error {
	var err error
	timeValue, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("unhandled type received: %v", reflect.TypeOf(value).Kind())
	}
	d.time = timeValue
	return err
}

func (d Date) Value() (driver.Value, error) {
	return d.time.Format(Format), nil
}

func (d Date) String() string {
	return d.time.Format(Format)
}

func (d Date) MarshalJSON() ([]byte, error) {
	if y := d.time.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(Format)+2)
	b = append(b, '"')
	b = d.time.AppendFormat(b, Format)
	b = append(b, '"')
	return b, nil
}

func (d *Date) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	d.time, err = time.Parse(`"`+Format+`"`, string(data))
	if err != nil {
		return d.time.UnmarshalJSON(data)
	}
	return nil
}

func (d Date) MarshalText() ([]byte, error) {
	if y := d.time.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("MarshalText: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(Format))
	return d.time.AppendFormat(b, Format), nil
}

func (d *Date) UnmarshalText(data []byte) error {
	var err error
	d.time, err = time.Parse(Format, string(data))
	return err
}

func New(year int, month time.Month, day int) Date {
	return Date{
		time: time.Date(year, month, day, 0, 0, 0, 0, time.UTC),
	}
}

func Parse(value string) (Date, error) {
	parsed, err := time.Parse(Format, value)
	if err != nil {
		return Date{}, err
	}
	return Date{
		time: parsed,
	}, nil
}

func (d Date) Year() int {
	return d.time.Year()
}

func (d Date) Month() int {
	return int(d.time.Month())
}

func (d Date) Day() int {
	return d.time.Day()
}

func (d Date) YearDay() int {
	return d.time.YearDay()
}

func (d Date) IsZero() bool {
	return d.time.IsZero()
}

func (d Date) Time() time.Time {
	return d.time
}

func (d Date) Add(year int, month int, day int) Date {
	d.time = d.time.AddDate(year, month, day)
	return d
}

func (d Date) After(u time.Time) bool {
	return d.time.After(u)
}

func FromTime(baseTime time.Time) Date {
	return Date{
		time: time.Date(baseTime.Year(), baseTime.Month(), baseTime.Day(), 0, 0, 0, 0, time.UTC),
	}
}

func Today() Date {
	now := time.Now()
	return Date{
		time: time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC),
	}
}
