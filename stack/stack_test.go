package stack

import (
	"testing"

	"github.com/babakabdolahi/algo.git/linkedlist"
)

func TestRandom(t *testing.T) {
	// s := []rune{'b', 'c'}

	// for _, v := range s {
	// 	fmt.Printf("%c\n", v)
	// }

	// x := [5]int{0, 2, 1, 5}
	// for _, v := range x {
	// 	fmt.Println(v)
	// }
}

func TestStackInit(t *testing.T) {
	st := Init(10)

	if st == nil {
		t.Fatal("init failed")
	}
}

func TestStackInsert(t *testing.T) {
	st := Init(10)

	err := Insert(st, 3)
	if err != nil {
		t.Fatal(err)
	}

	if st.Data[st.Top] != 3 {
		t.Fatalf("top stack value: Expected: %v, Got: %v", 3, st.Data[st.Top])
	}
}

func TestStackPop(t *testing.T) {
	st := Init(10)

	err := Insert(st, 3)
	if err != nil {
		t.Fatal(err)
	}

	err = Insert(st, 4)
	if err != nil {
		t.Fatal(err)
	}

	err = Insert(st, 5)
	if err != nil {
		t.Fatal(err)
	}

	v, err := Pop(st)
	if err != nil {
		t.Fatal(err)
	}

	if v != 5 {
		t.Errorf("pop operation failed. Expected: %v, Got: %v", 5, v)
	}

	if st.Top != 1 {
		t.Errorf("pop operation failed. Expected Top value: %v, Got: %v", 1, st.Top)
	}

	err = Insert(st, 6)
	if err != nil {
		t.Fatal(err)
	}

	if st.Top != 2 {
		t.Errorf("pop operation failed. Expected Top value: %v, Got: %v", 2, st.Top)
	}

	value, err := Peek(st)
	if err != nil {
		t.Error(err)
	}

	if value != 6 {
		t.Errorf("peek operation failed. Expected: %v, Got: %v", 6, value)
	}
}

func TestStackPeep(t *testing.T) {
	st := Init(10)

	err := Insert(st, 3)
	if err != nil {
		t.Fatal(err)
	}

	err = Insert(st, 4)
	if err != nil {
		t.Fatal(err)
	}

	v, err := Peek(st)
	if err != nil {
		t.Fatal(err)
	}

	if v != 4 {
		t.Errorf("incorrect top element. Expected: %v, Got: %v", 4, v)
	}
}

func TestListStackInsert(t *testing.T) {
	var top *linkedlist.Node

	top = ListStackInsert(top, 5)

	if top.Item != 5 {
		t.Errorf("incorrect top element. Expected: %v, Got: %v", 5, top.Item)
	}
}

func TestListStackPop(t *testing.T) {
	var top *linkedlist.Node

	top = ListStackInsert(top, 5)
	top = ListStackInsert(top, 6)

	top, err := ListStackPop(top)
	if err != nil {
		t.Error(err)
	}

	stackValue, err := ListStackPeek(top)
	if err != nil {
		t.Error(err)
	}

	if stackValue != 5 {
		t.Errorf("incorrect top element. Expected: %v, Got: %v", 5, stackValue)
	}
}

func TestParenthesesValidator(t *testing.T) {
	s := "(())"
	isValid, err := ParenthesesChecker(s)
	if err != nil {
		t.Error(err)
	}

	if !isValid {
		t.Errorf("wrong. Expected: %v, Got: %v", true, isValid)
	}

	s = "(()"

	isValid, err = ParenthesesChecker(s)
	if err != nil {
		t.Error(err)
	}

	if isValid {
		t.Errorf("wrong. Expected: %v, Got: %v", false, isValid)
	}
}

func TestInfixToPostFix(t *testing.T) {
	s := "( a + b )"

	resp, err := InfixToPostFix(s)
	if err != nil {
		t.Error(err)
	}

	expected := "ab+"
	if resp != expected {
		t.Errorf("incorrect. Expected: %v, Got: %v", expected, resp)
	}

	s = "(A+B)*(C+D)"

	resp, err = InfixToPostFix(s)
	if err != nil {
		t.Error(err)
	}

	expected = "AB+CD+*"
	if resp != expected {
		t.Errorf("incorrect. Expected: %v, Got: %v", expected, resp)
	}

	s = "(A+B)*C"

	resp, err = InfixToPostFix(s)
	if err != nil {
		t.Error(err)
	}

	expected = "AB+C*"
	if resp != expected {
		t.Errorf("incorrect. Expected: %v, Got: %v", expected, resp)
	}

	s = "A+B*C"

	resp, err = InfixToPostFix(s)
	if err != nil {
		t.Error(err)
	}

	expected = "ABC*+"
	if resp != expected {
		t.Errorf("incorrect. Expected: %v, Got: %v", expected, resp)
	}
}

func TestPostFixEvaluation(t *testing.T) {
	s := "456*+"

	resp, err := PostFixEvaluation(s)
	if err != nil {
		t.Error(err)
	}

	expected := 34
	if resp != expected {
		t.Errorf("incorrect. Expected: %v, Got: %v", expected, resp)
	}

	s = "78+32+/"

	resp, err = PostFixEvaluation(s)
	if err != nil {
		t.Error(err)
	}

	expected = 3
	if resp != expected {
		t.Errorf("incorrect. Expected: %v, Got: %v", expected, resp)
	}
}
