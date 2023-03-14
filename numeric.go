package books

import "golang.org/x/exp/constraints"

// Min compares two variables and returns the smallest of them.
// When passing strings, this returns the lexically first one.
func Min[T constraints.Ordered](t1, t2 T) T {
	if t1 > t2 {
		return t2
	}

	return t1
}

// Max compares two variables and returns the largest of them.
// When passing strings, this returns the lexically last one.
func Max[T constraints.Ordered](t1, t2 T) T {
	if t1 < t2 {
		return t2
	}

	return t1
}
