package list

import (
	"fmt"
	"testing"
)

func TestNewList(t *testing.T) {
	list := NewList()

	tests := []struct {
		value  int
		wanted int
	}{
		{len(list.elems), 0},
		{list.elements, 0},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("Value:%v, Expected:%v", test.value, test.wanted)
		t.Run(testname, func(t *testing.T) {
			if test.value != test.wanted {
				t.Errorf("Expected:%v", test.wanted)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	list := NewList()

	tests := []struct {
		added    interface{}
		expected int
	}{
		{10, 1},
		{"Test", 2},
		{true, 3},
	}
	for _, test := range tests {
		list.Add(test.added)
		if len(list.elems) != test.expected {
			t.Errorf("Expected %v", test.expected)
		}
	}
	t.Run("If size increases", func(t *testing.T) {
		if list.elements != 3 {
			t.Errorf("Expected %v", 3)
		}
	})
}

func TestAddRange(t *testing.T) {
	list := NewList()
	x := []Elem{"uno", "dos"}
	y := []Elem{"uno", "dos", "tres"}
	z := []Elem{}
	tests := []struct {
		elems []Elem
		len   int
	}{
		{x, 2},
		{y, 5},
		{z, 5},
	}

	for index, test := range tests {
		testname := fmt.Sprintf("TestAddRange %v", index+1)
		list.AddRange(test.elems)
		t.Run(testname, func(t *testing.T) {
			if len(list.elems) != test.len {
				t.Errorf("Expected %v", test.len)
			}
		})
	}
	t.Run("If size increases", func(t *testing.T) {
		if len(list.elems) != 5 {
			t.Errorf("Expected %v", 5)
		}
	})
}

func TestBinarySearch(t *testing.T) {
	var expected int
	list := NewList()

	t.Run("If list is empty", func(t *testing.T) {
		expected = list.BinarySearch(1)
		if expected != -1 {
			t.Errorf("Expected %v", -1)
		}
	})

	el := []Elem{
		1, 2, 3, 4, 5,
	}
	list.AddRange(el)

	t.Run("If the values is not in the list", func(t *testing.T) {
		expected = list.BinarySearch(1233)
		if expected != -1 {
			t.Errorf("Expected %v", -1)
		}
	})

	t.Run("If values is in the list", func(t *testing.T) {
		expected = list.BinarySearch(5)
		if expected != 4 {
			t.Errorf("Expected %v", 4)
		}
	})
}

func TestClear(t *testing.T) {
	list := NewList()
	listElements := []Elem{
		"Uno",
		"Dos",
		"Tres",
	}

	list.AddRange(listElements)
	list.Clear()

	t.Run("list.elements should be 0", func(t *testing.T) {
		if list.elements != 0 {
			t.Errorf("Expected %v", 0)
		}
	})

	if len(list.elems) != 0 {
		t.Errorf("Expected %v", 0)
	}
}

func TestContains(t *testing.T) {
	list := NewList()
	list.Add(1)
	list.Add("Uno")

	test := list.Contains(1)
	t.Run("If Elem exists", func(t *testing.T) {
		if test != true {
			t.Errorf("Expected %v", true)
		}
	})
	test = list.Contains(4)
	t.Run("If Elem doesn't exist", func(t *testing.T) {
		if test != false {
			t.Errorf("Expected %v", false)
		}
	})
}

func TestInsert(t *testing.T) {
	list := NewList()
	add := []Elem{
		1, 2, 3, 4, 5,
	}
	list.AddRange(add)

	test := list.Insert(7, 8)
	t.Run("When index out of range (Positive)", func(t *testing.T) {
		if test == nil {
			t.Error("Expected en error")
		}
	})

	test = list.Insert(-7, 8)
	t.Run("When index out of range (Negative)", func(t *testing.T) {
		if test == nil {
			t.Error("Expected en error")
		}
	})

	test = list.Insert(1, 11)
	t.Run("When index in range", func(t *testing.T) {
		if test != nil {
			t.Errorf("Expected %v", nil)
		}
	})

	t.Run("If value was added", func(t *testing.T) {
		if list.elems[1] != 11 {
			t.Errorf("Expected %v", 1)
		}
	})

	t.Run("If size increases", func(t *testing.T) {
		if list.elements != 6 {
			t.Errorf("Expected %v", list.elements)
		}
	})
}

func TestInsertRange(t *testing.T) {
	list := NewList()
	add := []Elem{
		1, 2, 3, 4, 5,
	}
	list.AddRange(add)
	inserted := []Elem{6, 7, 8}

	test := list.InsertRange(7, inserted)
	t.Run("When index out of range (Positive)", func(t *testing.T) {
		if test == nil {
			t.Error("Expected an error")
		}
	})

	test = list.InsertRange(-7, inserted)
	t.Run("When index out of range (Negative)", func(t *testing.T) {
		if test == nil {
			t.Error("Expected an error")
		}
	})

	test = list.InsertRange(1, inserted)
	t.Run("When index in range", func(t *testing.T) {
		if test != nil {
			t.Errorf("Expected %v", nil)
		}
	})

	for j := 1; j <= len(inserted); j++ {
		testname := fmt.Sprintf("If value is %v", list.elems[j])
		t.Run(testname, func(t *testing.T) {
			if inserted[j-1] != list.elems[j] {
				t.Errorf("Expected %v", inserted[j-1])
			}
		})
	}

	t.Run("If size increased", func(t *testing.T) {
		if list.elements != 8 {
			t.Errorf("Expected %v", list.elements)
		}
	})
}

func TestRemove(t *testing.T) {
	var expected error
	list := NewList()

	t.Run("If element not in list", func(t *testing.T) {
		expected = list.Remove("Test")
		if expected == nil {
			t.Error("Expected an error")
		}
	})

	add := []Elem{
		1, 2, 3, 4, 5,
	}
	list.AddRange(add)

	expected = list.Remove(3)
	t.Run("If expected return", func(t *testing.T) {
		if expected != nil {
			t.Errorf("Expected %v", nil)
		}
	})

	t.Run("If Elem was removed", func(t *testing.T) {
		test := list.Contains(3)
		if test == true {
			t.Errorf("Expected %v", false)
		}
	})

	t.Run("If size reduces", func(t *testing.T) {
		if list.elements != 4 {
			t.Errorf("Expected %v", 4)
		}
	})
}

func TestRemoveAt(t *testing.T) {
	var expected error
	list := NewList()

	t.Run("When list is empty", func(t *testing.T) {
		expected = list.RemoveAt(4)
		if expected == nil {
			t.Error("Expected an error")
		}
	})

	add := []Elem{
		1, 2, 3, 4, 5,
	}
	list.AddRange(add)

	t.Run("When index out of range (Positive)", func(t *testing.T) {
		expected = list.RemoveAt(7)
		if expected == nil {
			t.Error("Expected an error")
		}
	})

	t.Run("When index out of range (Negative)", func(t *testing.T) {
		expected = list.RemoveAt(-7)
		if expected == nil {
			t.Error("Expected an error")
		}
	})

	expected = list.RemoveAt(0)
	t.Run("If expected return", func(t *testing.T) {
		if expected != nil {
			t.Errorf("Expected %v", nil)
		}
	})

	t.Run("If Elem was removed", func(t *testing.T) {
		if list.elems[0] != 2 {
			t.Errorf("Expected %v", 2)
		}
	})

	t.Run("If size reduces", func(t *testing.T) {
		if list.elements != 4 {
			t.Errorf("Expected %v", 4)
		}
	})
}

func TestRemoveRange(t *testing.T) {
	var expected error
	list := NewList()

	t.Run("When list is empty", func(t *testing.T) {
		expected = list.RemoveRange(0, 4)
		if expected == nil {
			t.Error("Expected an error")
		}
	})

	add := []Elem{
		1, 2, 3, 4, 5,
	}
	list.AddRange(add)

	t.Run("When start is bigger than the end", func(t *testing.T) {
		expected = list.RemoveRange(8, 1)
		if expected == nil {
			t.Error("Expected en error")
		}
	})

	t.Run("When startIndex is bigger than the endIndex", func(t *testing.T) {
		expected = list.RemoveRange(6, 1)
		if expected == nil {
			t.Error("Expected ")
		}
	})

	t.Run("When start is out of range (Positive)", func(t *testing.T) {
		expected = list.RemoveRange(6, 9)
		if expected == nil {
			t.Error("Expected an error")
		}
	})

	t.Run("When start is out of range (Negative)", func(t *testing.T) {
		expected = list.RemoveRange(-6, 0)
		if expected == nil {
			t.Error("Expected an error")
		}
	})

	t.Run("When end is out of range (Positive)", func(t *testing.T) {
		expected = list.RemoveRange(0, 8)
		if expected == nil {
			t.Error("Expected an error")
		}
	})

	expected = list.RemoveRange(0, 3)
	t.Run("If expected return", func(t *testing.T) {
		if expected != nil {
			t.Errorf("Expected %v", nil)
		}
	})

	t.Run("If Elems removed", func(t *testing.T) {
		if list.elems[0] != 5 {
			t.Errorf("Expected %v", list.elems[0])
		}
	})

	t.Run("If size reduces", func(t *testing.T) {
		if list.elements != 1 {
			t.Errorf("Expected %v", list.elements)
		}
	})
}

func TestCount(t *testing.T) {
	var expected int
	list := NewList()

	t.Run("When list is empty", func(t *testing.T) {
		expected = list.Count()
		if expected != 0 {
			t.Errorf("Expected %v", 0)
		}
	})

	add := []Elem{
		1, 2, 3, 4, 5,
	}
	list.AddRange(add)

	t.Run("When list not empty", func(t *testing.T) {
		expected = list.Count()
		if expected != 5 {
			t.Errorf("Expected %v", 5)
		}
	})

	list.RemoveAt(0)
	t.Run("When one removed", func(t *testing.T) {
		expected = list.Count()
		if expected != 4 {
			t.Errorf("Expected %v", 4)
		}
	})
}

/*
Test output

=== RUN   TestNewList
=== RUN   TestNewList/Value:0,_Expected:0
=== RUN   TestNewList/Value:0,_Expected:0#01
--- PASS: TestNewList (0.00s)
    --- PASS: TestNewList/Value:0,_Expected:0 (0.00s)
    --- PASS: TestNewList/Value:0,_Expected:0#01 (0.00s)
=== RUN   TestAdd
=== RUN   TestAdd/If_size_increases
--- PASS: TestAdd (0.00s)
    --- PASS: TestAdd/If_size_increases (0.00s)
=== RUN   TestAddRange
=== RUN   TestAddRange/TestAddRange_1
=== RUN   TestAddRange/TestAddRange_2
=== RUN   TestAddRange/TestAddRange_3
=== RUN   TestAddRange/If_size_increases
--- PASS: TestAddRange (0.00s)
    --- PASS: TestAddRange/TestAddRange_1 (0.00s)
    --- PASS: TestAddRange/TestAddRange_2 (0.00s)
    --- PASS: TestAddRange/TestAddRange_3 (0.00s)
    --- PASS: TestAddRange/If_size_increases (0.00s)
=== RUN   TestBinarySearch
=== RUN   TestBinarySearch/If_list_is_empty
=== RUN   TestBinarySearch/If_the_values_is_not_in_the_list
=== RUN   TestBinarySearch/If_values_is_in_the_list
--- PASS: TestBinarySearch (0.00s)
    --- PASS: TestBinarySearch/If_list_is_empty (0.00s)
    --- PASS: TestBinarySearch/If_the_values_is_not_in_the_list (0.00s)
    --- PASS: TestBinarySearch/If_values_is_in_the_list (0.00s)
=== RUN   TestClear
=== RUN   TestClear/list.elements_should_be_0
--- PASS: TestClear (0.00s)
    --- PASS: TestClear/list.elements_should_be_0 (0.00s)
=== RUN   TestContains
=== RUN   TestContains/If_Elem_exists
=== RUN   TestContains/If_Elem_doesn't_exist
--- PASS: TestContains (0.00s)
    --- PASS: TestContains/If_Elem_exists (0.00s)
    --- PASS: TestContains/If_Elem_doesn't_exist (0.00s)
=== RUN   TestInsert
=== RUN   TestInsert/When_index_out_of_range_(Positive)
=== RUN   TestInsert/When_index_out_of_range_(Negative)
=== RUN   TestInsert/When_index_in_range
=== RUN   TestInsert/If_value_was_added
=== RUN   TestInsert/If_size_increases
--- PASS: TestInsert (0.00s)
    --- PASS: TestInsert/When_index_out_of_range_(Positive) (0.00s)
    --- PASS: TestInsert/When_index_out_of_range_(Negative) (0.00s)
    --- PASS: TestInsert/When_index_in_range (0.00s)
    --- PASS: TestInsert/If_value_was_added (0.00s)
    --- PASS: TestInsert/If_size_increases (0.00s)
=== RUN   TestInsertRange
=== RUN   TestInsertRange/When_index_out_of_range_(Positive)
=== RUN   TestInsertRange/When_index_out_of_range_(Negative)
=== RUN   TestInsertRange/When_index_in_range
=== RUN   TestInsertRange/If_value_is_6
=== RUN   TestInsertRange/If_value_is_7
=== RUN   TestInsertRange/If_value_is_8
=== RUN   TestInsertRange/If_size_increased
--- PASS: TestInsertRange (0.00s)
    --- PASS: TestInsertRange/When_index_out_of_range_(Positive) (0.00s)
    --- PASS: TestInsertRange/When_index_out_of_range_(Negative) (0.00s)
    --- PASS: TestInsertRange/When_index_in_range (0.00s)
    --- PASS: TestInsertRange/If_value_is_6 (0.00s)
    --- PASS: TestInsertRange/If_value_is_7 (0.00s)
    --- PASS: TestInsertRange/If_value_is_8 (0.00s)
    --- PASS: TestInsertRange/If_size_increased (0.00s)
=== RUN   TestRemove
=== RUN   TestRemove/If_element_not_in_list
=== RUN   TestRemove/If_expected_return
=== RUN   TestRemove/If_Elem_was_removed
=== RUN   TestRemove/If_size_reduces
--- PASS: TestRemove (0.00s)
    --- PASS: TestRemove/If_element_not_in_list (0.00s)
    --- PASS: TestRemove/If_expected_return (0.00s)
    --- PASS: TestRemove/If_Elem_was_removed (0.00s)
    --- PASS: TestRemove/If_size_reduces (0.00s)
=== RUN   TestRemoveAt
=== RUN   TestRemoveAt/When_list_is_empty
=== RUN   TestRemoveAt/When_index_out_of_range_(Positive)
=== RUN   TestRemoveAt/When_index_out_of_range_(Negative)
=== RUN   TestRemoveAt/If_expected_return
=== RUN   TestRemoveAt/If_Elem_was_removed
=== RUN   TestRemoveAt/If_size_reduces
--- PASS: TestRemoveAt (0.00s)
    --- PASS: TestRemoveAt/When_list_is_empty (0.00s)
    --- PASS: TestRemoveAt/When_index_out_of_range_(Positive) (0.00s)
    --- PASS: TestRemoveAt/When_index_out_of_range_(Negative) (0.00s)
    --- PASS: TestRemoveAt/If_expected_return (0.00s)
    --- PASS: TestRemoveAt/If_Elem_was_removed (0.00s)
    --- PASS: TestRemoveAt/If_size_reduces (0.00s)
=== RUN   TestRemoveRange
=== RUN   TestRemoveRange/When_list_is_empty
=== RUN   TestRemoveRange/When_start_is_bigger_than_the_end
=== RUN   TestRemoveRange/When_startIndex_is_bigger_than_the_endIndex
=== RUN   TestRemoveRange/When_start_is_out_of_range_(Positive)
=== RUN   TestRemoveRange/When_start_is_out_of_range_(Negative)
=== RUN   TestRemoveRange/When_end_is_out_of_range_(Positive)
=== RUN   TestRemoveRange/If_expected_return
=== RUN   TestRemoveRange/If_Elems_removed
=== RUN   TestRemoveRange/If_size_reduces
--- PASS: TestRemoveRange (0.00s)
    --- PASS: TestRemoveRange/When_list_is_empty (0.00s)
    --- PASS: TestRemoveRange/When_start_is_bigger_than_the_end (0.00s)
    --- PASS: TestRemoveRange/When_startIndex_is_bigger_than_the_endIndex (0.00s)
    --- PASS: TestRemoveRange/When_start_is_out_of_range_(Positive) (0.00s)
    --- PASS: TestRemoveRange/When_start_is_out_of_range_(Negative) (0.00s)
    --- PASS: TestRemoveRange/When_end_is_out_of_range_(Positive) (0.00s)
    --- PASS: TestRemoveRange/If_expected_return (0.00s)
    --- PASS: TestRemoveRange/If_Elems_removed (0.00s)
    --- PASS: TestRemoveRange/If_size_reduces (0.00s)
=== RUN   TestCount
=== RUN   TestCount/When_list_is_empty
=== RUN   TestCount/When_list_not_empty
=== RUN   TestCount/When_one_removed
--- PASS: TestCount (0.00s)
    --- PASS: TestCount/When_list_is_empty (0.00s)
    --- PASS: TestCount/When_list_not_empty (0.00s)
    --- PASS: TestCount/When_one_removed (0.00s)
PASS
ok      List    0.966s
*/
