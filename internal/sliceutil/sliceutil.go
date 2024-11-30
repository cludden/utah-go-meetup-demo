package sliceutil

func Map[A, B any](s []A, fn func(A) B) []B {
	result := make([]B, len(s))
	for i, a := range s {
		result[i] = fn(a)
	}
	return result
}
