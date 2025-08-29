package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type CustomRFC3339Time struct {
	time.Time
}

func (customTime CustomRFC3339Time) Std() time.Time {
	return customTime.Time
}

// Force +00:00 instead of Z
func (t CustomRFC3339Time) String() string {
	return t.Format("2006-01-02T15:04:05-07:00")
}

func (t CustomRFC3339Time) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("\"%s\"", t.Format("2006-01-02T15:04:05-07:00"))
	return []byte(s), nil
}

// Make GORM recognize this type for auto migration
func (t CustomRFC3339Time) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *CustomRFC3339Time) Scan(value interface{}) error {
	if value == nil {
		*t = CustomRFC3339Time{}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*t = CustomRFC3339Time{v}
		return nil
	default:
		return fmt.Errorf("cannot convert %T to CustomRFC3339Time", value)
	}
}
