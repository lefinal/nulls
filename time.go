package nulls

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"time"
)

// Time holds a nullable time.Time.
type Time struct {
	// Time is the actual value when Valid.
	Time time.Time `exhaustruct:"optional"`
	// Valid when no NULL-value is represented.
	Valid bool
}

// NewTime returns a valid Time with the given value.
func NewTime(t time.Time) Time {
	return Time{
		Time:  t,
		Valid: true,
	}
}

// MarshalJSON marshals the time.Time. If not valid, a NULL-value is returned.
func (t Time) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(t.Time)
}

// UnmarshalJSON as time.Time or sets Valid to false if null.
func (t *Time) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		t.Valid = false
		return nil
	}
	t.Valid = true
	return json.Unmarshal(data, &t.Time)
}

// Scan to time.Time value or not valid if nil.
func (t *Time) Scan(src any) error {
	var sqlTime sql.NullTime
	err := sqlTime.Scan(src)
	if err != nil {
		return err
	}
	t.Valid = sqlTime.Valid
	t.Time = sqlTime.Time
	return nil
}

// Value returns the value for satisfying the driver.Valuer interface.
func (t Time) Value() (driver.Value, error) {
	return sql.NullTime{
		Time:  t.Time,
		Valid: t.Valid,
	}.Value()
}

// UTC returns the UTC time.
func (t Time) UTC() Time {
	return Time{
		Time:  t.Time.UTC(),
		Valid: t.Valid,
	}
}
