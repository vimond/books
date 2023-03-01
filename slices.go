package books

// RemoveDupes removes duplicates from a slice
func RemoveDupes[T comparable, S ~[]T](sl S) S {
	m := make(map[T]struct{})
	idx := 0

	// If an item has not been seen before, adding it in-place and incrementing the index
	for _, s := range sl {
		if _, ok := m[s]; !ok {
			m[s] = struct{}{}
			sl[idx] = s
			idx++
		}
	}

	return sl[:idx]
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
