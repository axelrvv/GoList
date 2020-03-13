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
	list := NewList()
	el := []Elem{
		1, 2, 3, 4, 5,
	}
	list.AddRange(el)

	var expected int

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
	t.Run("When index out of range", func(t *testing.T) {
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
	t.Run("When index out of range", func(t *testing.T) {
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
