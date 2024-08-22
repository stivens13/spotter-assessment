package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/stivens13/spotter-assessment/helper/constants"
)

/*
Use castom Date type to ensure that date is always formatted in the same way
*/
type Date time.Time

func NewDate(date string) Date {
	t, err := time.Parse(constants.VIDEO_DATE_FORMAT, date)
	if err != nil {
		log.Errorf("failed to parse time: %v", err)
	}
	return Date(t)
}

func (d Date) Equal(u Date) bool {
	return time.Time(d).Equal(time.Time(u))
}

func (d Date) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	if t.IsZero() {
		return []byte("null"), nil
	}
	formatted := t.Format(constants.VIDEO_DATE_FORMAT)

	jsonStr := "\"" + formatted + "\""
	return []byte(jsonStr), nil
}

/*
Custom Marshalling and Unmarshalling (serialization and deserialization)
*/
func (d *Date) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		*d = Date(time.Time{})
		return nil
	}
	if len(b) < 2 || b[0] != '"' || b[len(b)-1] != '"' {
		return errors.New("not a json string")
	}

	b = b[1 : len(b)-1]

	t, err := time.Parse(constants.VIDEO_DATE_FORMAT, string(b))
	if err != nil {
		return fmt.Errorf("failed to parse time: %w", err)
	}

	*d = Date(t)

	return nil
}

/*
Custom Value and Scan to help db driver handle custom time struct automatically
*/
func (d Date) Value() (driver.Value, error) {
	t := time.Time(d)
	if t.IsZero() {
		return nil, nil
	}
	return t.Format(constants.VIDEO_DATE_FORMAT), nil
}

func (d *Date) Scan(value interface{}) error {
	if value == nil {
		*d = Date(time.Time{})
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*d = Date(v)
	case []byte:
		t, err := time.Parse(constants.VIDEO_DATE_FORMAT, string(v))
		if err != nil {
			return err
		}
		*d = Date(t)
	case string:
		t, err := time.Parse(constants.VIDEO_DATE_FORMAT, v)
		if err != nil {
			return err
		}
		*d = Date(t)
	default:
		return fmt.Errorf("cannot scan type %T into Date", value)
	}
	return nil
}
