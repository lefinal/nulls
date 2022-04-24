package nulls

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// NullableValue are the requirements for values used in Nullable as they need
// to implement at least sql.Scanner and driver.Valuer.
type NullableValue interface {
	sql.Scanner
	driver.Valuer
}

// Nullable holds a nullable value.
type Nullable[T NullableValue] struct {
	// V is the actual value when Valid.
	V T
	// Valid describes whether the Nullable does not hold a NULL value.
	Valid bool
}

// NewNullable creates a new valid Nullable with the given value.
func NewNullable[T NullableValue](v T) Nullable[T] {
	return Nullable[T]{
		V:     v,
		Valid: true,
	}
}

// MarshalJSON as value. If not vot valid, a NULL-value is returned.
func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(n.V)
}

// UnmarshalJSON as value ro sets Valid o false if null.
func (n *Nullable[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.Valid = false
		return nil
	}
	n.Valid = true
	return json.Unmarshal(data, &n.V)
}

// Scan to value or not valid if nil.
func (n *Nullable[T]) Scan(src any) error {
	if src == nil {
		n.Valid = false
		return nil
	}
	n.Valid = true
	return n.V.Scan(src)
}

// Value returns the value for satisfying the driver.Valuer interface.
func (n Nullable[T]) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.V.Value()
}
