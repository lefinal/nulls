package nulls

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// Int64 holds a nullable int64.
type Int64 struct {
	// Int64 is the actual value when Valid.
	Int64 int64
	// Valid when no NULL-value is represented.
	Valid bool
}

// NewInt64 returns a valid Int64 with the given value.
func NewInt64(i int64) Int64 {
	return Int64{
		Int64: i,
		Valid: true,
	}
}

// MarshalJSON marshals the int. If not valid, a NULL-value is returned.
func (i Int64) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(i.Int64)
}

// UnmarshalJSON as int or sets Valid to false if null.
func (i *Int64) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		i.Valid = false
		return nil
	}
	i.Valid = true
	return json.Unmarshal(data, &i.Int64)
}

// Scan to int value or not valid if nil.
func (i *Int64) Scan(src any) error {
	var sqlInt64 sql.NullInt64
	err := sqlInt64.Scan(src)
	if err != nil {
		return err
	}
	i.Valid = sqlInt64.Valid
	i.Int64 = sqlInt64.Int64
	return nil
}

// Value returns the value for satisfying the driver.Valuer interface.
func (i Int64) Value() (driver.Value, error) {
	return sql.NullInt64{
		Int64: i.Int64,
		Valid: i.Valid,
	}.Value()
}
