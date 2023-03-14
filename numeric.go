package books

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](t1, t2 T) T {
	if t1 > t2 {
		return t2
	}

	return t1
}

func Max[T constraints.Ordered](t1, t2 T) T {
	if t1 < t2 {
		return t2
	}

	return t1
}
