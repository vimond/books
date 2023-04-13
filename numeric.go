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

// Clamp returns the value of x clamped to the range [min, max].
// If max is less than min, the function will return max, no checks are performed to ensure that min is less than max.
func Clamp[T constraints.Ordered](x, min, max T) T {
	return Min(Max(x, min), max)
}
