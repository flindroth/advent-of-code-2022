package util

type Stack[T any] []T

func (s *Stack[T]) Push(t T) {
	*s = append(*s, t)
}

func (s *Stack[T]) Pop() T {
	r := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return r
}

func (s *Stack[T]) Peek() T {
	return (*s)[len(*s)-1]
}
