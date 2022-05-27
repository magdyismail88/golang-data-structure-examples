package main

import (
	"fmt"
	"strconv"
)

type Stack[T any] []T

func (stack *Stack[T]) Push(value T) {
	*stack = append(*stack, value)
}

func (stack *Stack[T]) Pop() {
	*stack = remove(*stack, len(*stack)-1)
}

func (stack *Stack[T]) Top() T {
	return (*stack)[len(*stack)-1]
}

func (stack *Stack[T]) IsEmpty() bool {
	return len(*stack) == 0
}

func remove[T any](haystack []T, index int) []T {
	keep := haystack[:]
	haystack = haystack[:0]
	haystack = keep[:index]
	haystack = append(haystack, keep[index+1:]...)
	return haystack
}

func convertFromInfixToPostfix(expr string) string {
	stack := Stack[string]{}
	var stmt string

	for _, val := range expr {

		if string(val) == " " {
			continue
		}

		switch string(val) {
		case "(":
			stack.Push(string(val))
			break
		case ")":
			for stack.Top() != "(" {
				stmt += string(stack.Top())
				stack.Pop()
			}

			stack.Pop()
			break
		case "+", "-":

			if !stack.IsEmpty() {
				for stack.Top() == "*" || stack.Top() == "/" {
					stmt += stack.Top()
					stack.Pop()
				}
			}

			stack.Push(string(val))

			break
		case "*", "/":

			stack.Push(string(val))

			break
		default:
			stmt += string(val)
			break
		}
	}

	if !stack.IsEmpty() {
		stmt += stack.Top()
	}

	return stmt
}

func handlePostfix(expr string) float64 {

	stack := Stack[string]{}

	for _, val := range expr {

		switch string(val) {
		case "*", "/", "+", "-":
			snumber, _ := strconv.ParseFloat(stack.Top(), 64)
			stack.Pop()
			fnumber, _ := strconv.ParseFloat(stack.Top(), 64)
			stack.Pop()
			result := calc(fnumber, snumber, string(val))
			stack.Push(fmt.Sprintf("%f", result))
			break
		default:
			stack.Push(string(val))
			break
		}

	}

	result, _ := strconv.ParseFloat(stack.Top(), 64)
	return result
}

func calc(fnumber, snumber float64, oper string) float64 {
	switch oper {
	case "*":
		return fnumber * snumber
	case "/":
		return fnumber / snumber
	case "+":
		return fnumber + snumber
	case "-":
		return fnumber - snumber
	default:
		return 0

	}
}

func pair(open, close rune) bool {

	switch {
	case open == '{' && close == '}':
		return true
	case open == '(' && close == ')':
		return true
	case open == '[' && close == ']':
		return true
	default:
		return false
	}

}

// example for expr: {([])}
func isBalanced(expr string) bool {
	openBraces := new(Stack[rune])
	chars := []rune(expr)

	for _, char := range chars {

		if char == '{' || char == '(' || char == '[' {

			openBraces.Push(char)

		} else if char == '}' || char == ')' || char == ']' {

			if openBraces.IsEmpty() {
				return false
			} else if pair(openBraces.Top(), char) == false {
				return false
			}

			openBraces.Pop()
		}

	}

	return openBraces.IsEmpty()
}

func main() {
	// expression := "((3+2)*5-2)"
	expression := "((3+2)*5-2+2(2-1))"
	// expression := "3+2"

	// result = 32+5*2-
	postfix := convertFromInfixToPostfix(expression)
	fmt.Println(postfix)

	result := handlePostfix(postfix)

	fmt.Println(result)

	braces := "{([])}"

	ok := isBalanced(braces)

	fmt.Println(ok)

}
