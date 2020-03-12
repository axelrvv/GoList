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
	tests := []struct {
		elems []Elem
		len   int
	}{
		{x, 2},
		{y, 5},
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
