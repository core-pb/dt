package ut

func Array2Array[A1, A2 any](arr []A1, getKey func(A1) A2) []A2 {
	val := make([]A2, 0, len(arr))
	for i := range arr {
		val = append(val, getKey(arr[i]))
	}
	return val
}

func Array2Map[K comparable, V any](arr []V, getKey func(V) K) map[K]V {
	m := make(map[K]V)
	for i := range arr {
		m[getKey(arr[i])] = arr[i]
	}

	return m
}

func MapKey2Array[K comparable, V any](m map[K]V) []K {
	arr := make([]K, 0, len(m))
	for k := range m {
		arr = append(arr, k)
	}
	return arr
}
