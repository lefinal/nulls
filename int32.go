package nulls

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// Int32 holds a nullable int32.
type Int32 struct {
	// Int32 is the actual value when Valid.
	Int32 int32 `exhaustruct:"optional"`
	// Valid when no NULL-value is represented.
	Valid bool
}

// NewInt32 returns a valid Int32 with the given value.
func NewInt32(i int32) Int32 {
	return Int32{
		Int32: i,
		Valid: true,
	}
}

// MarshalJSON marshals the int. If not valid, a NULL-value is returned.
func (i Int32) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(i.Int32)
}

// UnmarshalJSON as int or sets Valid to false if null.
func (i *Int32) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		i.Valid = false
		return nil
	}
	i.Valid = true
	return json.Unmarshal(data, &i.Int32)
}

// Scan to int value or not valid if nil.
func (i *Int32) Scan(src any) error {
	var sqlInt32 sql.NullInt32
	err := sqlInt32.Scan(src)
	if err != nil {
		return err
	}
	i.Valid = sqlInt32.Valid
	i.Int32 = sqlInt32.Int32
	return nil
}

// Value returns the value for satisfying the driver.Valuer interface.
func (i Int32) Value() (driver.Value, error) {
	return sql.NullInt32{
		Int32: i.Int32,
		Valid: i.Valid,
	}.Value()
}
