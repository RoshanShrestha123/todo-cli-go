package utils

func RemoveArrayElementOnIndex[T any](index int, arr []T) []T {
	return append(arr[:index], arr[index+1:]...)
}
