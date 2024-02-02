package nulls

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// JSONNullable holds a nullable value. Keep in mind, that T must be
// (un)marshallable. However, it cannot be used as sql.Scanner or driver.Valuer.
type JSONNullable[T any] struct {
	// V is the actual value when Valid.
	V T `exhaustruct:"optional"`
	// Valid describes whether the JSONNullable does not hold a NULL value.
	Valid bool
}

// NewJSONNullable creates a new valid JSONNullable with the given value.
func NewJSONNullable[T any](v T) JSONNullable[T] {
	return JSONNullable[T]{
		V:     v,
		Valid: true,
	}
}

// MarshalJSON as value. If not vot valid, a NULL-value is returned.
func (n JSONNullable[T]) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(n.V)
}

// UnmarshalJSON as value ro sets Valid o false if null.
func (n *JSONNullable[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.Valid = false
		return nil
	}
	n.Valid = true
	return json.Unmarshal(data, &n.V)
}

// Scan to value or not valid if nil.
func (n *JSONNullable[T]) Scan(src any) error {
	return errors.New("unsupported operation")
}

// Value returns the value for satisfying the driver.Valuer interface.
func (n JSONNullable[T]) Value() (driver.Value, error) {
	return nil, errors.New("unsupported operation")
}
