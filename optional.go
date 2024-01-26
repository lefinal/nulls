package nulls

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// Optional holds a nullable value. This can be used instead of Nullable or
// NullableInto if database support is not required.
type Optional[T any] struct {
	// V is the actual value when Valid.
	V T
	// Valid describes whether the Nullable does not hold a NULL value.
	Valid bool
}

// NewOptional creates a new valid Optional with the given value.
func NewOptional[T any](v T) Optional[T] {
	return Optional[T]{
		V:     v,
		Valid: true,
	}
}

// MarshalJSON as value. If not vot valid, a NULL-value is returned.
func (n Optional[T]) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(n.V)
}

// UnmarshalJSON as value or sets Valid or false if null.
func (n *Optional[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.Valid = false
		return nil
	}
	n.Valid = true
	return json.Unmarshal(data, &n.V)
}

// Scan returns an error as this is currently not supported on Optional. Use
// Nullable or NullableInto instead.
func (n *Optional[T]) Scan(_ any) error {
	return fmt.Errorf("cannot scan optional")
}

// Value returns an error as this is currently not supported on Optional. Use
// Nullable or NullableInto instead.
func (n Optional[T]) Value() (driver.Value, error) {
	return nil, fmt.Errorf("cannot value optional")
}
