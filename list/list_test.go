package list

import "testing"

// insert an intermediate element
func TestListInsert1(t *testing.T) {
	list := List{
		Elements:       []int{1, 4, 6},
		LastElPosition: 2,
	}

	err := list.Insert(10, 1)
	if err != nil {
		t.Fatal(err)
	}

	if len(list.Elements) != 4 {
		t.Errorf("expected list length: %v\n got list length: %v", 4, len(list.Elements))
	}

	if list.Elements[0] != 1 || list.Elements[1] != 10 || list.Elements[2] != 4 || list.Elements[3] != 6 {
		t.Fatal("incorrect insertion")
	}

	if list.LastElPosition != 3 {
		t.Errorf("expected last element position: %v\n got last element position: %v", 3, list.LastElPosition)
	}
}
