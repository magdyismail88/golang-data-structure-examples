package bst

import (
       "fmt"
       "golang.org/x/exp/constraints"
)

type node[T constraints.Ordered] struct {
	left *node[T]
	right *node[T]
	value T
}

type bst[T constraints.Ordered] struct {
	root *node[T]
}

func New[T constraints.Ordered]() *bst[T] {
	return &bst[T]{}
}

func (self *bst[T]) insertRecursion(obj *node[T], value T) *node[T] {
	if self.root == nil {
	   	newNode := &node[T]{value: value,}
		self.root = newNode
		return self.root
	}
	if obj == nil {
		return &node[T]{value: value}
	}
	if value <= self.root.value {
		obj.left = self.insertRecursion(obj.left, value)
	}
	if value > self.root.value {
		obj.right = self.insertRecursion(obj.right, value)
	}
	return obj
}

func (self *bst[T]) Insert(value T) {
	self.insertRecursion(self.root, value)
}

func (self *bst[T]) findRecursion(obj *node[T], value T) *node[T] {
	if obj == nil {
		return nil
	}
	if value == obj.value {
		return obj
	}
	if value < obj.value {
		return self.findRecursion(obj.left, value)
	}
	return self.findRecursion(obj.right, value)
}

func (self *bst[T]) Find(value T) error {
	node := self.findRecursion(self.root, value)
	if node == nil {
		return fmt.Errorf("Value: %v not founded in tree", value)
	}
	return nil
}

func (self *bst[T]) deleteRecursion(obj *node[T], value T) *node[T] {
	if obj == nil {
		return nil
	}
	if value < obj.value {
		obj.left = self.deleteRecursion(obj.left, value)
	} else if value > obj.value {
		obj.right = self.deleteRecursion(obj.right, value)
	} else {
		if obj.left == nil {
			return obj.right
		} else if obj.right == nil {
			return obj.left
		} else {
			maxValue := self.GetMax()
			obj.value = maxValue
			self.deleteRecursion(obj, maxValue)
		}
	}
	return obj
}

func (self *bst[T]) Delete(value T) {
	self.root = self.deleteRecursion(self.root, value)
}

func (self *bst[T]) GetMin() T {
	node := self.root
	for node.left != nil {
		node = node.left
	}
	return node.value
}

func (self *bst[T]) GetMax() T {
	node := self.root
	for node.right != nil {
		node = node.right
	}
	return node.value
}


