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
