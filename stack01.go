package stack

type Stack[T any] []T

func (self *Stack[T]) Push(value T) {
    *self = append(*self, value)
}

func (self *Stack[T]) Pop() {
    *self = (*self)[:len(*self)-1]
}

func (self *Stack[T]) Top() T {
    return (*self)[len(*self)-1]
}

func (self *Stack[T]) Values() []T {
    return *self
}
