package nulls

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
)

// JSONRawMessage holds a json.RawMessage. Keep in mind, that the JSON NULL
// value will be represented with Valid being false.
type JSONRawMessage struct {
	// RawMessage is the actual json.RawMessage when Valid.
	RawMessage json.RawMessage
	// Valid when no NULL-value is represented.
	Valid bool
}

// NewJSONRawMessage returns a valid JSONRawMessage with the given value.
func NewJSONRawMessage(raw json.RawMessage) JSONRawMessage {
	return JSONRawMessage{
		RawMessage: raw,
		Valid:      true,
	}
}

// MarshalJSON marshals the RawMessage. If not valid, a NULL-value is returned.
func (rm JSONRawMessage) MarshalJSON() ([]byte, error) {
	if !rm.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(rm.RawMessage)
}

// UnmarshalJSON as json.RawMessage or sets Valid to false if empty.
func (rm *JSONRawMessage) UnmarshalJSON(data []byte) error {
	// Do NOT use regular NULL-check here.
	if isNull(data) {
		rm.Valid = false
		rm.RawMessage = nil
		return nil
	}
	rm.Valid = true
	return json.Unmarshal(data, &rm.RawMessage)
}

// Scan to json.RawMessage value or not valid if nil.
func (rm *JSONRawMessage) Scan(src any) error {
	if src == nil {
		rm.Valid = false
		rm.RawMessage = nil
		return nil
	}
	rm.Valid = true
	// Copy bytes.
	srcBytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("cannot convert to byte slice: %s", reflect.TypeOf(src).String())
	}
	b := make([]byte, len(srcBytes))
	copy(b, srcBytes)
	rm.RawMessage = b
	return nil
}

// Value returns the value for satisfying the driver.Valuer interface.
func (rm JSONRawMessage) Value() (driver.Value, error) {
	if !rm.Valid {
		return nil, nil
	}
	return []byte(rm.RawMessage), nil
}
