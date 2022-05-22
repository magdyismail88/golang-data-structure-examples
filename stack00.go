package stack

type Stack[T any] struct {
    storage []T
}

func New[T any]() *Stack[T] {
    return &Stack[T]{}
}

func (self *Stack[T]) Push(value T) {
    self.storage = append(self.storage, value)
}

func (self *Stack[T]) Top() T {
    return self.storage[len(self.storage)-1]
}

func (self *Stack[T]) Pop() {
    self.storage = self.storage[:len(self.storage)-1]
}

func (self *Stack[T]) Values() []T {
    return self.storage
}
