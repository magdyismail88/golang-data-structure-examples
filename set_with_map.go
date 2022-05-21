package set

import (
	"fmt"
)

type Set[T comparable] map[T]int

func (self *Set[T]) Add(value T) {
	if *self == nil {
		*self = make(map[T]int)
	}
	(*self)[value]++
}

func (self *Set[T]) Stream(values []T) {
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

func HasDuplications[T comparable](values []T) bool {
	var set Set[T]
	for _, v := range values {
		set.Add(v)
	}
	return len(set.Values) != len(values)
}

func RemoveDuplications[T any](values []T) []T {
	var set Set[T]
	set.Stream(values)
	return set.Values()
}

func main() {
	
	var set Set[int]
	var other Set[int]

	set.Add(1)
	set.Add(2)
	set.Add(1)
	set.Add(1)
	set.Add(1)
	set.Add(3)

	other.Add(3)
	other.Add(4)
	other.Add(5)

	interVals := set.Intersections(&other)

	for _, v := range interVals {
		fmt.Println(v)
	}

	values := []int{1, 1, 2, 3, 4}

	if HasDuplications[int](values) {
		// set := new(Set[int])
		// values = set.Stream(values)
		values = RemoveDuplications(values)
	}

	fmt.Println(values)

	// for _, v := range set.Values() {
	// 	fmt.Printf("%d\n", v)
	// }


}



