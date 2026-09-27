package queue

import "testing"

func TestQueue(t *testing.T) {
	q := Init(5)

	err := Insert(q, 20)
	if err != nil {
		t.Error(err)
	}

	expected := 20

	if Peek(q) != expected {
		t.Errorf("incorrect. Expected: %v, Got: %v", expected, Peek(q))
	}

	err = Insert(q, 30)
	if err != nil {
		t.Error(err)
	}

	deleted, err := DeleteElement(q)
	if err != nil {
		t.Error(err)
	}

	expected = 20

	if deleted != expected {
		t.Errorf("incorrect. Expected: %v, Got: %v", expected, deleted)
	}

	err = Insert(q, 40)
	if err != nil {
		t.Error(err)
	}

	expected = 30

	if Peek(q) != expected {
		t.Errorf("incorrect. Expected: %v, Got: %v", expected, Peek(q))
	}
}
