package utils

// Union unions two slice
func Union[T comparable](a []T, b []T) []T {
	m := make(map[T]struct{})
	c := make([]T, 0)
	for _, ele := range a {
		m[ele] = struct{}{}
	}
	for _, ele := range b {
		if _, exist := m[ele]; exist {
			c = append(c, ele)
		}
	}
	return c
}

func Contains[T comparable](target []T, needle T) bool {
	for _, ele := range target {
		if ele == needle {
			return true
		}
	}
	return false
}
