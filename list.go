package list

import "errors"

//Elem : Variable type of interface so it can take any datatype
type Elem interface{}

//List is going to be the data structure struct
//I know it's crazy
type List struct {
	elems    []Elem
	elements int //Represents the number of Elem in elems
}

//NewList : returns a new List type
func NewList() List {
	list := List{}
	return list
}

//Add : Adds an element at the end of a List
func (i *List) Add(x Elem) {
	i.elems = append(i.elems, x)
	i.elements = i.elements + 1
}

//AddRange : Adds elements of a list or array, it wil recive an array and it will add it at the end of the List
func (i *List) AddRange(x []Elem) {
	for _, y := range x {
		i.Add(y)
	}
}

//BinarySearch : Looks for an Elem and returns the index of that element in the List
func (i *List) BinarySearch(x Elem) int {
	for index, elem := range i.elems {
		if x == elem {
			return index
		}
	}

	return -1
}

//Clear : Removes all the elements from a List
func (i *List) Clear() {
	i.elems = i.elems[:0]
	i.elements = 0
}

//Contains : Checks whether the speciied element exists or not in a List.
func (i *List) Contains(x Elem) bool {
	exists := i.BinarySearch(x)
	if exists == -1 {
		return false
	}
	return true
}

//Insert : Inserts an Elem in the index specified
func (i *List) Insert(index int, x Elem) error {
	if i.elements < index || index < 0 {
		return errors.New("Index must be within the bound of the List (Parameter 'index')")
	}
	rest := i.elems[index:]
	i.elems[index] = x
	i.elems = i.elems[:index+1]
	i.elems = append(i.elems, rest...)
	i.elements = i.elements + 1
	return nil
}

//InsertRange : Inserts a []Elem from the index specified to the len of []Elem inserted
func (i *List) InsertRange(index int, x []Elem) error {
	if i.elements < index || index < 0 {
		return errors.New("Index must be within the bound of the List (Parameter 'index')")
	}
	rest := i.elems[index:]
	i.elems = i.elems[:index]
	i.elems = append(i.elems, x...)
	i.elems = append(i.elems, rest...)
	i.elements = i.elements + len(x)
	return nil
}

//Remove : Removes the first occurence of the specified element
func (i *List) Remove(x Elem) error {

	if i.Contains(x) {
		index := i.BinarySearch(x)
		i.elems = append(i.elems[:index], i.elems[index+1:]...)
		i.elements = i.elements - 1
		return nil
	}

	return errors.New("Element doesn't exist in the list")
}

//RemoveAt : Removes the element at the specified index
func (i *List) RemoveAt(index int) error {

	if i.elements == 0 {
		return errors.New("List is empty")
	}

	if index >= i.elements || index < 0 {
		return errors.New("Index out of range")
	}

	return i.Remove(i.elems[index])
}
