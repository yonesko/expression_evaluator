package expressionevaluator

import (
	"github.com/golang-collections/collections/stack"
)

func eval(str string) int32 {
	ops := stack.New()
	vals := stack.New()

	for _, c := range str {
		switch {
		case isOperation(c):
			ops.Push(c)
		case c == ')':
			doAndStack(vals, ops)
		case c >= '1' && c <= '9':
			vals.Push(c - '0')
		}
	}

	for ops.Len() != 0 {
		doAndStack(vals, ops)
	}

	return vals.Pop().(int32)
}

func isOperation(c int32) bool {
	return c == '+' || c == '*' || c == '-'
}

func doAndStack(vals, ops *stack.Stack) {
	vals.Push(do(ops.Pop().(int32), vals.Pop().(int32), vals.Pop().(int32)))
}

func do(action int32, a, b int32) int32 {
	switch action {
	case '+':
		return a + b
	case '*':
		return a * b
	case '-':
		return a - b
	}
	panic("unknown action: ")
}
