package set

type Set[T comparable] struct {
	storage map[T]int
}

func New[T comparable]() *Set[T] {
	return &Set[T]{storage: make(map[T]int),}
}

func (self *Set[T]) Add(value T) {
	self.storage[value]++
}

func (self *Set[T]) AddMany(values ...T) {
	for _, v := range values {
		self.Add(v)
	}
}

func (self *Set[T]) Values() []T {
	var values []T
	for key, _ := range self.storage {
		values = append(values, key)
	}
	return values
}

func (self *Set[T]) Intersections(other *Set[T]) []T {
	set := New[T]()
	for _, v := range self.Values() {
		set.Add(v)
	}
	for _, v := range other.Values() {
		set.Add(v)
	}
	return set.Values()
}
