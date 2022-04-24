package nulls

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// Float32 holds a nullable float32.
type Float32 struct {
	// Float32 is the actual value when Valid.
	Float32 float32
	// Valid when no NULL-value is represented.
	Valid bool
}

// NewFloat32 returns a valid Float32 with the given value.
func NewFloat32(f float32) Float32 {
	return Float32{
		Float32: f,
		Valid:   true,
	}
}

// MarshalJSON marshals the float32. If not valid, a NULL-value is returned.
func (f Float32) MarshalJSON() ([]byte, error) {
	if !f.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(f.Float32)
}

// UnmarshalJSON as float32 or sets Valid to false if null.
func (f *Float32) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		f.Valid = false
		return nil
	}
	f.Valid = true
	return json.Unmarshal(data, &f.Float32)
}

// Scan to float32 value or not valid if nil.
func (f *Float32) Scan(src any) error {
	var sqlFloat sql.NullFloat64
	err := sqlFloat.Scan(src)
	if err != nil {
		return err
	}
	f.Valid = sqlFloat.Valid
	f.Float32 = float32(sqlFloat.Float64)
	return nil
}

// Value returns the value for satisfying the driver.Valuer interface.
func (f *Float32) Value() (driver.Value, error) {
	return sql.NullFloat64{
		Float64: float64(f.Float32),
		Valid:   f.Valid,
	}.Value()
}
