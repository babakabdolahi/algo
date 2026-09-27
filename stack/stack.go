package stack

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/babakabdolahi/algo.git/linkedlist"
)

const MaxStackSize = 100

type Stack struct {
	Max  int
	Top  int
	Data []int
}

func Init(max int) *Stack {
	st := Stack{
		Max: max,
		Top: -1,
	}

	st.Data = make([]int, max)

	return &st
}

func Insert(stack *Stack, value int) error {
	if stack == nil {
		return fmt.Errorf("stack is not initialized")
	}

	if stack.Top == stack.Max-1 {
		return fmt.Errorf("stack overflow")
	}

	stack.Top++

	stack.Data[stack.Top] = value

	return nil
}

func Pop(stack *Stack) (int, error) {
	if stack.Top == -1 {
		return 0, fmt.Errorf("stack underflow")
	}

	v := stack.Data[stack.Top]

	stack.Top--

	return v, nil
}

// Peek returns the value of the stack located at Top
func Peek(stack *Stack) (int, error) {
	if stack == nil {
		return 0, fmt.Errorf("nil stack")
	}

	return stack.Data[stack.Top], nil
}

func ListStackInsert(top *linkedlist.Node, value int) *linkedlist.Node {
	if top == nil {
		return linkedlist.Init(value)
	}

	newNode := linkedlist.Init(value)
	newNode.Next = top

	return newNode
}

func ListStackPop(top *linkedlist.Node) (*linkedlist.Node, error) {
	if top == nil {
		return nil, errors.New("stack is empty")
	}

	top = top.Next

	return top, nil
}

func ListStackPeek(top *linkedlist.Node) (int, error) {
	if top == nil {
		return 0, errors.New("stack is empty")
	}

	return top.Item, nil
}

func ParenthesesChecker(s string) (bool, error) {
	var top *linkedlist.Node
	var err error

	for _, v := range s {
		if v == '(' {
			top = ListStackInsert(top, int(v))

			continue
		}

		if v == ')' {
			top, err = ListStackPop(top)
			if err != nil {
				return false, err
			}
		}
	}

	if top != nil {
		return false, err
	}

	return true, err
}

// This function assumes that the input is valid
func InfixToPostFix(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}

	stk := Init(MaxStackSize)

	var result string

	for _, v := range s {
		if v == '(' || isOperand(v) {
			err := Insert(stk, int(v))
			if err != nil {
				return "", err
			}
		} else if isCharacter(v) {
			result = result + string(v)
		} else if v == ')' {
			for {
				top, err := Pop(stk)
				if err != nil {
					return "", err
				}

				if top == int('(') {
					break
				}

				result = result + string(rune(top))
			}
		}
	}

	if stk.Top != -1 {
		for {
			top, err := Pop(stk)
			if err != nil {
				return "", err
			}

			result = result + string(rune(top))

			if stk.Top == -1 {
				break
			}
		}
	}

	return result, nil
}

func PostFixEvaluation(s string) (int, error) {
	if len(s) == 0 {
		return 0, nil
	}

	stk := Init(MaxStackSize)

	for _, v := range s {
		if isInteger(v) {
			value, err := strconv.Atoi(string(v))
			if err != nil {
				return 0, err
			}

			err = Insert(stk, value)
			if err != nil {
				return 0, err
			}
		} else if isOperand(v) {
			v1, err := Pop(stk)
			if err != nil {
				return 0, err
			}

			v2, err := Pop(stk)
			if err != nil {
				return 0, err
			}

			var res int

			if v == '+' {
				res = v2 + v1
			} else if v == '-' {
				res = int(v2) - int(v1)
			} else if v == '/' {
				res = int(v2) / int(v1)
			} else if v == '*' {
				res = int(v2) * int(v1)
			} else {
				return 0, errors.New("invalid operand")
			}

			err = Insert(stk, res)
			if err != nil {
				return 0, err
			}
		}
	}

	resp, err := Pop(stk)
	if err != nil {
		return 0, err
	}

	return resp, nil
}

func isOperand(c rune) bool {
	for _, v := range "+*/-" {
		if c == v {
			return true
		}
	}

	return false
}

func isCharacter(c rune) bool {
	for _, v := range "aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ" {
		if c == v {
			return true
		}
	}

	return false
}

func isInteger(c rune) bool {
	for _, v := range "0123456789" {
		if c == v {
			return true
		}
	}

	return false
}
