package nulls

import (
	"encoding/json"
)

// jsonNull is the marshalled nil-value.
var jsonNull []byte

func init() {
	var err error
	jsonNull, err = json.Marshal(nil)
	if err != nil {
		panic(err)
	}
}

// marshalMust marshals the given value and panics if marshalling fails.
func marshalMust(v any) []byte {
	raw, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return raw
}
