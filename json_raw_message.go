package nulls

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
)

// JSONRawMessage holds a json.RawMessage. However, this differs from the other
// types: Actually, "null" is a valid json.RawMessage (NOT nil!) and for example
// PostgreSQL distinguishes between regular NULL-values and "null"-values, being
// regular JSONB. Marshalling-logic however, only knows "null". Therefore, in
// UnmarshalJSON we do NOT set Valid to false, when unmarshalling "null" (as we
// do for the other types), but only, if the passed data is nil.
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
//
// Warning: This differs from unmarshalling-logic of other types. For more
// information see documentation of JSONRawMessage.
func (rm *JSONRawMessage) UnmarshalJSON(data []byte) error {
	// Do NOT use regular NULL-check here.
	if data == nil {
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
