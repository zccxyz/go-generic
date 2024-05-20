package slice

// Contains 判断元素是否在slice中，简化版
func Contains[T comparable](slice []T, val T) bool {
	return ContainsFunc(slice, func(t T) bool {
		return t == val
	})
}

// ContainsFunc 判断元素是否在slice中，自定义比较函数
func ContainsFunc[T any](slice []T, fn func(T) bool) bool {
	for _, v := range slice {
		if fn(v) {
			return true
		}
	}

	return false
}

// ContainsAny s1中是否包含s2中的一个元素，简化版
func ContainsAny[T comparable](s1, s2 []T) bool {
	m := toMap(s1)

	for _, v2 := range s2 {
		if _, ok := m[v2]; ok {
			return true
		}
	}

	return false
}

// ContainsAnyFunc s1中是否包含s2中的一个元素
func ContainsAnyFunc[T any](s1, s2 []T, fn func(v1, v2 T) bool) bool {
	for _, v2 := range s2 {
		for _, v1 := range s1 {
			if fn(v1, v2) {
				return true
			}
		}
	}
	return false
}

func ContainsAll[T comparable](s1, s2 []T) bool {
	m := toMap(s1)

	for _, v2 := range s2 {
		if _, ok := m[v2]; !ok {
			return false
		}
	}

	return true
}

func ContainsAllFunc[T comparable](s1, s2 []T, fn func(v1, v2 T) bool) bool {
	for _, v2 := range s2 {
		for _, v1 := range s1 {
			if !fn(v1, v2) {
				return false
			}
		}
	}

	return true
}

func toMap[T comparable](s []T) map[T]struct{} {
	m := make(map[T]struct{}, len(s))
	for _, v := range s {
		m[v] = struct{}{}
	}
	return m
}
