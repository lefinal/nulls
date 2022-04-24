package nulls

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// Float64 holds a nullable float64.
type Float64 struct {
	// Float64 is the actual value when Valid.
	Float64 float64
	// Valid when no NULL-value is represented.
	Valid bool
}

// NewFloat64 returns a valid Float64 with the given value.
func NewFloat64(f float64) Float64 {
	return Float64{
		Float64: f,
		Valid:   true,
	}
}

// MarshalJSON marshals the float64. If not valid, a NULL-value is returned.
func (f Float64) MarshalJSON() ([]byte, error) {
	if !f.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(f.Float64)
}

// UnmarshalJSON as float64 or sets Valid to false if null.
func (f *Float64) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		f.Valid = false
		return nil
	}
	f.Valid = true
	return json.Unmarshal(data, &f.Float64)
}

// Scan to float64 value or not valid if nil.
func (f *Float64) Scan(src any) error {
	var sqlFloat sql.NullFloat64
	err := sqlFloat.Scan(src)
	if err != nil {
		return err
	}
	f.Valid = sqlFloat.Valid
	f.Float64 = sqlFloat.Float64
	return nil
}

// Value returns the value for satisfying the driver.Valuer interface.
func (f *Float64) Value() (driver.Value, error) {
	return sql.NullFloat64{
		Float64: f.Float64,
		Valid:   f.Valid,
	}.Value()
}
