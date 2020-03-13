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
	}
	for _, test := range tests {
		list.Add(test.added)
		if len(list.elems) != test.expected {
			t.Errorf("Expected %v", test.expected)
		}
	}
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
