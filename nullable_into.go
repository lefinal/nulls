package nulls

import (
	"database/sql/driver"
	"encoding/json"
)

// NullableIntoValue are the requirements for values used in NullableInto as they
// need to implement at least driver.Valuer, and a ScanInto method.
type NullableIntoValue[T any] interface {
	ScanInto(src any, dst *T) error
	driver.Valuer
}

// NullableInto holds a nullable value that satisfies NullableIntoValue. This can
// be used instead of Nullable if the value should not be treated as pointer. It
// then provides the NullableIntoValue.ScanInto method that scans into a passed
// reference.
type NullableInto[T NullableIntoValue[T]] struct {
	// V is the actual value when Valid.
	V T
	// Valid describes whether the Nullable does not hold a NULL value.
	Valid bool
}

// NewNullableInto creates a new valid NullableInto with the given value.
func NewNullableInto[T NullableIntoValue[T]](v T) NullableInto[T] {
	return NullableInto[T]{
		V:     v,
		Valid: true,
	}
}

// MarshalJSON as value. If not vot valid, a NULL-value is returned.
func (n NullableInto[T]) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(n.V)
}

// UnmarshalJSON as value ro sets Valid o false if null.
func (n *NullableInto[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.Valid = false
		return nil
	}
	n.Valid = true
	return json.Unmarshal(data, &n.V)
}

// Scan to value or not valid if nil.
func (n *NullableInto[T]) Scan(src any) error {
	if src == nil {
		n.Valid = false
		return nil
	}
	n.Valid = true
	return n.V.ScanInto(src, &n.V)
}

// Value returns the value for satisfying the driver.Valuer interface.
func (n NullableInto[T]) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.V.Value()
}
