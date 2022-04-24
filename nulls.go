// Package nulls provides nullable types like the ones from sql package. However,
// (un)marshalling support is available as well.
package nulls

// isNull checks if the given byte slice represents NULL value.
func isNull(b []byte) bool {
	return b == nil || string(b) == "null"
}
