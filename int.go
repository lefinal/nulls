package nulls

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// Int holds a nullable int.
type Int struct {
	// Int is the actual value when Valid.
	Int int
	// Valid when no NULL-value is represented.
	Valid bool
}

// NewInt returns a valid Int with the given value.
func NewInt(i int) Int {
	return Int{
		Int:   i,
		Valid: true,
	}
}

// MarshalJSON marshals the int. If not valid, a NULL-value is returned.
func (i Int) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(i.Int)
}

// UnmarshalJSON as int or sets Valid to false if null.
func (i *Int) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		i.Valid = false
		return nil
	}
	i.Valid = true
	return json.Unmarshal(data, &i.Int)
}

// Scan to int value or not valid if nil.
func (i *Int) Scan(src any) error {
	var sqlInt sql.NullInt64
	err := sqlInt.Scan(src)
	if err != nil {
		return err
	}
	i.Valid = sqlInt.Valid
	i.Int = int(sqlInt.Int64)
	return nil
}

// Value returns the value for satisfying the driver.Valuer interface.
func (i Int) Value() (driver.Value, error) {
	return sql.NullInt64{
		Int64: int64(i.Int),
		Valid: i.Valid,
	}.Value()
}
