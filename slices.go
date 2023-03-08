package books

// RemoveDupes removes duplicates from a slice, order is preserved
func RemoveDupes[T comparable, S ~[]T](s S) S {
	m := make(map[T]struct{})
	idx := 0

	// If an item has not been seen before, adding it in-place and incrementing the index
	for i := range s {
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = struct{}{}
			s[idx] = s[i]
			idx++
		}
	}

	return s[:idx]
}

// ReverseSlice reverses the order of the elements in the array passed
func ReverseSlice[T any, S ~[]T](s S) S {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

// Contains checks if the slice contains the given element
func Contains[T comparable, S ~[]T](s S, e T) bool {
	for i := range s {
		if e == s[i] {
			return true
		}
	}

	return false
}

// Filter filters a slice by passed function
// The passed function should return true to keep item, and false to remove it
func FilterSlice[T any, S ~[]T](s S, f func(T) bool) S {
	idx := 0

	for i := range s {
		if f(s[i]) {
			s[idx] = s[i]
			idx++
		}
	}

	return s[:idx]
}

// EqualSlices checks if two slices of comparable types are equal
func EqualSlices[T comparable, S ~[]T](s1, s2 S) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
