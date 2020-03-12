package list

//Elem : Variable type of interface so it can take any datatype
type Elem interface{}

//List is goint to be the data structure struct
//I know it's crazy
type List struct {
	elems    []Elem
	size     int
	capacity int
}

//NewList returns a new List type
func NewList() List {
	list := List{}
	return list
}
