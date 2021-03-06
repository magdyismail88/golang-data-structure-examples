package set

type Set[T comparable] map[T]int

func (self *Set[T]) Add(value T) {
	if *self == nil {
		*self = make(map[T]int)
	}
	(*self)[value]++
}

func (self *Set[T]) AddMany(values ...T) {
	if *self == nil {
		*self = make(map[T]int)
	}
	for _, v := range values {
		self.Add(v)
	}
}

func (self *Set[T]) Values() []T {
	var values []T
	for key, _ := range *self {
		values = append(values, key)
	}
	return values
}

func (self *Set[T]) Intersections(other *Set[T]) []T {
	var set Set[T]
	if *self == *other {
		return self.Valus()
	}
	for key, _ := range *self {
		set.Add(key)
	}
	for key, _ := range *other {
		set.Add(key)
	}
	return set.Values()
}
