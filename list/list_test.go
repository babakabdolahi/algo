package list

import "testing"

// insert an intermediate element
func TestListInsert1(t *testing.T) {
	list := New(10)

	err := list.Insert(1, 0)
	if err != nil {
		t.Fatal(err)
	}

	err = list.Insert(2, 1)
	if err != nil {
		t.Fatal(err)
	}

	err = list.Insert(3, 2)
	if err != nil {
		t.Fatal(err)
	}

	err = list.Insert(4, 1)
	if err != nil {
		t.Fatal(err)
	}

	if len(list.Elements) != 4 {
		t.Errorf("expected list length: %v\n got list length: %v", 4, len(list.Elements))
	}

	if list.Elements[0] != 1 || list.Elements[1] != 4 || list.Elements[2] != 2 || list.Elements[3] != 3 {
		t.Error("incorrect insertion")
	}

	if list.LastElPosition != 3 {
		t.Errorf("expected last element position: %v\n got last element position: %v", 3, list.LastElPosition)
	}
}

// insert into nil list
func TestListInsert2(t *testing.T) {
	list := List{
		Elements: nil,
	}

	err := list.Insert(10, 1)
	if err != nil {
		expectedErr := "nil list"

		if err.Error() != expectedErr {
			t.Errorf("wrong error, expected %v\n got %v", expectedErr, err.Error())
		}
	} else {
		t.Error("the purpose of the test is not met")
	}
}

// list max length limit
func TestListInsert3(t *testing.T) {
	list := New(2)

	err := list.Insert(1, 0)
	if err != nil {
		t.Error(err)
	}

	err = list.Insert(2, 1)
	if err != nil {
		t.Error(err)
	}

	err = list.Insert(3, 2)
	if err != nil {
		expectedErr := "list is full"

		if err.Error() != expectedErr {
			t.Errorf("wrong error, expected %v\n got %v", expectedErr, err.Error())
		}
	} else {
		t.Error("the purpose of the test is not met")
	}
}

// list max length limit
func TestListInsert4(t *testing.T) {
	list := New(2)

	err := list.Insert(1, 0)
	if err != nil {
		t.Error(err)
	}

	err = list.Insert(2, 1)
	if err != nil {
		t.Error(err)
	}

	err = list.Insert(3, 2)
	if err != nil {
		expectedErr := "list is full"

		if err.Error() != expectedErr {
			t.Errorf("wrong error, expected %v\n got %v", expectedErr, err.Error())
		}
	} else {
		t.Error("the purpose of the test is not met")
	}
}

// out of range position
func TestListInsert5(t *testing.T) {
	list := New(3)

	err := list.Insert(1, 0)
	if err != nil {
		t.Error(err)
	}

	err = list.Insert(2, 1)
	if err != nil {
		t.Error(err)
	}

	// instead of position 2, insert into position 3
	err = list.Insert(3, 3)
	if err != nil {
		expectedErr := "position is out of range"

		if err.Error() != expectedErr {
			t.Errorf("wrong error, expected %v\n got %v", expectedErr, err.Error())
		}
	} else {
		t.Error("the purpose of the test is not met")
	}
}

// negative position
func TestListInsert6(t *testing.T) {
	list := New(1)

	err := list.Insert(1, -1)
	if err != nil {
		expectedErr := "negative position"

		if err.Error() != expectedErr {
			t.Errorf("wrong error, expected %v\n got %v", expectedErr, err.Error())
		}
	} else {
		t.Error("the purpose of the test is not met")
	}
}

// locate elements
func TestListLocate1(t *testing.T) {
	list := New(2)

	err := list.Insert(1, 0)
	if err != nil {
		t.Fatal(err)
	}

	err = list.Insert(2, 1)
	if err != nil {
		t.Fatal(err)
	}

	p1, ok := list.Locate(2)
	if !ok || p1 != 1 {
		t.Errorf("expected: %v\n got: %v", 1, p1)
	}

	p2, ok := list.Locate(5)
	if ok {
		t.Errorf("value %v should not exist", p2)
	}
}

// return element
func TestListRetrieve1(t *testing.T) {
	list := New(2)

	err := list.Insert(1, 0)
	if err != nil {
		t.Fatal(err)
	}

	err = list.Insert(2, 1)
	if err != nil {
		t.Fatal(err)
	}

	v, err := list.Retrieve(1)
	if err != nil {
		t.Fatal(err)
	}

	if v != 2 {
		t.Errorf("expected: %v\n got: %v", 2, v)
	}
}

// delete element
func TestListDelete1(t *testing.T) {
	list := New(3)

	err := list.Insert(1, 0)
	if err != nil {
		t.Fatal(err)
	}

	err = list.Insert(2, 1)
	if err != nil {
		t.Fatal(err)
	}

	err = list.Insert(3, 2)
	if err != nil {
		t.Fatal(err)
	}

	err = list.Delete(1)
	if err != nil {
		t.Fatal(err)
	}

	if len(list.Elements) != 2 {
		t.Errorf("expected list length: %v\n got: %v", 2, len(list.Elements))
	}

	v1, err := list.Retrieve(0)
	if err != nil {
		t.Fatal(err)
	}

	v2, err := list.Retrieve(1)
	if err != nil {
		t.Fatal(err)
	}

	if v1 != 1 || v2 != 3 {
		t.Errorf("delete operation done wrong")
	}
}

func TestEndMethod(t *testing.T) {
	list := New(2)

	err := list.Insert(1, 0)
	if err != nil {
		t.Fatal(err)
	}

	err = list.Insert(2, 1)
	if err != nil {
		t.Fatal(err)
	}

	endPosition := list.End()

	if endPosition != 2 {
		t.Errorf("expected end position: %v\n got: %v", 2, endPosition)
	}
}

func TestMakeEmpty(t *testing.T) {
	list := New(2)

	err := list.Insert(1, 0)
	if err != nil {
		t.Fatal(err)
	}

	err = list.Insert(2, 1)
	if err != nil {
		t.Fatal(err)
	}

	list.MakeEmpty()

	if len(list.Elements) != 0 {
		t.Errorf("list is not emptied")
	}
}
