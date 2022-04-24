package nulls

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// Int16 holds a nullable int16.
type Int16 struct {
	// Int16 is the actual value when Valid.
	Int16 int16
	// Valid when no NULL-value is represented.
	Valid bool
}

// NewInt16 returns a valid Int16 with the given value.
func NewInt16(i int16) Int16 {
	return Int16{
		Int16: i,
		Valid: true,
	}
}

// MarshalJSON marshals the int. If not valid, a NULL-value is returned.
func (i Int16) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(i.Int16)
}

// UnmarshalJSON as int or sets Valid to false if null.
func (i *Int16) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		i.Valid = false
		return nil
	}
	i.Valid = true
	return json.Unmarshal(data, &i.Int16)
}

// Scan to int value or not valid if nil.
func (i *Int16) Scan(src any) error {
	var sqlInt16 sql.NullInt16
	err := sqlInt16.Scan(src)
	if err != nil {
		return err
	}
	i.Valid = sqlInt16.Valid
	i.Int16 = sqlInt16.Int16
	return nil
}

// Value returns the value for satisfying the driver.Valuer interface.
func (i Int16) Value() (driver.Value, error) {
	return sql.NullInt16{
		Int16: i.Int16,
		Valid: i.Valid,
	}.Value()
}
