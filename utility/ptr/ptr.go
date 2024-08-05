package ptr

func ToPointer[T any](data T) *T {
	return &data
}
