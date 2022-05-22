package main

type set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable]() *set[T] {
	s := &set[T]{}
	s.m = make(map[T]struct{})
	return s
}

func (s *set[T]) Add(value T) {
	s.m[value] = struct{}{}
}

func (s *set[T]) Remove(value T) {
	delete(s.m, value)
}

func (s *set[T]) Contains(value T) bool {
    _, c := s.m[value]
    return c
}
