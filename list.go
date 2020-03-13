package list

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

//Add : Adds an element into the list
func (i *List) Add(x Elem) {
	i.elems = append(i.elems, x)
	i.elements = i.elements + 1
}

//AddRange : Adds elements of a list or array, it wil recive an array and it will add it at the end of the list
func (i *List) AddRange(x []Elem) {
	i.elems = append(i.elems, x...)
	i.elements = i.elements + len(x)
}

//BinarySearch : Looks for an Elem and returns the index of that element
func (i *List) BinarySearch(x Elem) int {
	for index, elem := range i.elems {
		if x == elem {
			return index
		}
	}

	return -1
}

//Clear : Removes all the elements from the list
func (i *List) Clear() {
	i.elems = i.elems[:0]
	i.elements = 0
}

//Contains checks whether the speciied element exists or not in a List.
func (i *List) Contains(x Elem) bool {
	exists := i.BinarySearch(x)
	if exists == -1 {
		return false
	}
	return true
}
