package books

// Ptr takes any type and returns a pointer to it.
func Ptr[T any](t T) *T {
	return &t
}

// Deref returns the underlying value of the pointer. If it is nil, it returns a newly initialized type.
func Deref[T any](p *T) T {
	if p == nil {
		var n T

		return n
	}

	return *p
}
