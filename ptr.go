package ptr

// To returns a pointer to the given value.
func To[T any](value T) *T {
	return &value
}

// ValueFrom returns the value from a given pointer. If ref is nil, a zero
// value of type T will be returned.
func ValueFrom[T any](ref *T) (value T) {
	if ref != nil {
		value = *ref
	}
	return
}
