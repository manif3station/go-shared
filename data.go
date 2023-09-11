package shared_lib

func ArrayItem[T int | string](array []T, pos int) T {
	var value T
	if len(array)-1 < pos {
		return value
	} else {
		return array[pos]
	}
}
