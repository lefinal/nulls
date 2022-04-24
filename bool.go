package nulls

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// Bool holds a nullable boolean value.
type Bool struct {
	// Bool is the actual boolean value when Valid.
	Bool bool
	// Valid when no NULL-value is represented.
	Valid bool
}

// NewBool returns a valid Bool with the given value.
func NewBool(b bool) Bool {
	return Bool{
		Bool:  b,
		Valid: true,
	}
}

// MarshalJSON marshals the Bool. If not valid, a NULL-value is returned.
func (b Bool) MarshalJSON() ([]byte, error) {
	if !b.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(b.Bool)
}

// UnmarshalJSON as boolean or sets Valid to false if null.
func (b *Bool) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		b.Valid = false
		return nil
	}
	b.Valid = true
	return json.Unmarshal(data, &b.Bool)
}

// Scan to boolean value or not valid if nil.
func (b *Bool) Scan(src any) error {
	var sqlBool sql.NullBool
	err := sqlBool.Scan(src)
	if err != nil {
		return err
	}
	b.Valid = sqlBool.Valid
	b.Bool = sqlBool.Bool
	return nil
}

// Value returns the value for satisfying the driver.Valuer interface.
func (b Bool) Value() (driver.Value, error) {
	return sql.NullBool{
		Bool:  b.Bool,
		Valid: b.Valid,
	}.Value()
}
