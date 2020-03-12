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
		{list.size, 0},
		{list.capacity, 0},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("Value: %v, Expected: %v", test.value, test.wanted)
		t.Run(testname, func(t *testing.T) {
			if test.value != test.wanted {
				t.Errorf("Expected %v", test.wanted)
			}
		})
	}

}
