package nulls

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// String holds a nullable string.
type String struct {
	// String is the actual value when Valid.
	String string `exhaustruct:"optional"`
	// Valid when no NULL-value is represented.
	Valid bool
}

// NewString returns a valid String with the given value.
func NewString(s string) String {
	return String{
		String: s,
		Valid:  true,
	}
}

// MarshalJSON marshals the string. If not valid, a NULL-value is returned.
func (s String) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(s.String)
}

// UnmarshalJSON as string or sets Valid to false if null.
func (s *String) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		s.Valid = false
		return nil
	}
	s.Valid = true
	return json.Unmarshal(data, &s.String)
}

// Scan to string value or not valid if nil.
func (s *String) Scan(src any) error {
	var sqlString sql.NullString
	err := sqlString.Scan(src)
	if err != nil {
		return err
	}
	s.Valid = sqlString.Valid
	s.String = sqlString.String
	return nil
}

// Value returns the value for satisfying the driver.Valuer interface.
func (s String) Value() (driver.Value, error) {
	return sql.NullString{
		String: s.String,
		Valid:  s.Valid,
	}.Value()
}
