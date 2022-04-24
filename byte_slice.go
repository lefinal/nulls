package nulls

import (
	"database/sql"
	"database/sql/driver"
	"encoding/base64"
	"encoding/json"
)

// ByteSlice holds a nullable byte slice.
type ByteSlice struct {
	// ByteSlice is the actual byte slice when Valid.
	ByteSlice []byte
	// Valid when no NULL-value is represented.
	Valid bool
}

// NewByteSlice returns a valid ByteSlice with the given value.
func NewByteSlice(b []byte) ByteSlice {
	return ByteSlice{
		ByteSlice: b,
		Valid:     true,
	}
}

// MarshalJSON marshals the ByteSlice. If not valid, a NULL-value is returned.
func (b ByteSlice) MarshalJSON() ([]byte, error) {
	if !b.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(b.ByteSlice)
}

// UnmarshalJSON as byte slice or sets Valid to false if null.
func (b *ByteSlice) UnmarshalJSON(data []byte) error {
	if isNull(data) {
		b.Valid = false
		return nil
	}
	b.Valid = true
	return json.Unmarshal(data, &b.ByteSlice)
}

// Scan to byte slice value or not valid if nil.
func (b *ByteSlice) Scan(src any) error {
	var sqlString sql.NullString
	err := sqlString.Scan(src)
	if err != nil {
		return err
	}
	b.Valid = sqlString.Valid
	b.ByteSlice, err = base64.StdEncoding.DecodeString(sqlString.String)
	if err != nil {
		return err
	}
	return nil
}

// Value returns the value for satisfying the driver.Valuer interface.
func (b ByteSlice) Value() (driver.Value, error) {
	if !b.Valid {
		return nil, nil
	}
	return base64.StdEncoding.EncodeToString(b.ByteSlice), nil
}
