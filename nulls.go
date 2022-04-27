// Package nulls provides nullable types like the ones from sql package.
// However, (un)marshalling support is available as well. Keep in mind, that
// NULL-values and "undefined"-values (JS-style) are treated the same.
package nulls

// isNull checks if the given byte slice represents NULL-value or "nothing" (in
// JS: undefined; here: nil).
func isNull(b []byte) bool {
	return b == nil || string(b) == "null"
}
