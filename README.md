# GoList

GoList is a dynamic array data structure based on C#'s dynamic array (List).

## Features
| Feature | Status | Actions |
| ----------------- | -------------- | ----------------- |
| **Add** | ✅ | Adds an element at the end of a List |
| **AddRange** | ✅ | Adds elements of the specified collection at the end of a List. |
| **BinarySearch** | ✅ | Search the element and returns an index of the element |
| **Clear**  | ✅ | Removes all the elements from a List |
| **Contains** | ✅ | Checks whether the speciied element exists or not in a List |
| **Insert** | ✅ | Inserts an element at the specified index in a List |
| **InsertRange** | ✅ | Inserts elements of another collection at the specified index |
| **Remove** | ✅ | Removes the first occurence of the specified element | 
| **RemoveAt** | ✅ | Removes the element at the specified index |
| **RemoveRange** | ✅ | Removes a range of elements from the list |
| **Count** | ✅ | Gets the number of elements contained in the list |
| **IndexOf** | ✅ | Searches for the specified index and returns the value of it |
| **AddInFront** | ✅ | Add an element in front of the list |

# Download

go get github.com/axelrvv/GoList

# Usage 

```Go
package main

import (
	"fmt"

	l "github.com/axelrvv/GoList"
)

func main() {

	list := l.NewList()

	list.Add(100)
	list.Add(200)
	list.Add(300)
	list.Add(400)
	list.Add(500)
	list.Add(600)

	toAdd := []l.Elem{700, 800, 900, 1000}
	toAdd1 := []l.Elem{101, 102, 103}

	list.AddRange(toAdd)
	fmt.Println(list.Count())
	//Returns 10

	fmt.Println(list.BinarySearch(500))
	//Returns 4

	list.Insert(4, 550)
	fmt.Println(list.BinarySearch(550))
	//Returns 4

	list.InsertRange(1, toAdd1)
	fmt.Println(list.Count())
	//Returns 14

	list.Remove(100)
	fmt.Println(list.Count())
	//Returns 13

	fmt.Println(list.Contains(100))
	//Returns false

	fmt.Println(list.IndexOf(1))
	//Returns 102 <nil>

	list.RemoveAt(1)
	fmt.Println(list.IndexOf(1))
	//Returns 103 <nil>
}

/*
Output

10
4
4
14
13
false
102 <nil>
103 <nil>
*/
```

*To make the list receptive for any datatype we deal with interfaces, it's important to know how to do interface assertion when you're storing a struct in the list*

### Example

```Go
package main

import (
	"fmt"

	l "github.com/axelrvv/GoList"
)

type person struct {
	name   string
	age    int
	weight float32
}

func main() {

	list := l.NewList()

	p1 := person{"Axel Vasquez", 22, 180.88}
	p2 := person{"John Doe", 23, 150.50}
	p3 := person{"Doe John", 22, 140.78}

	list.Add(p1)
	list.Add(p2)
	list.Add(p3)

	person1, err := list.IndexOf(0)

	if err != nil {
		panic(err)
	}

	name := person1.(person).name
	age := person1.(person).age
	weight := person1.(person).weight

	fmt.Printf("Name: %v; Age: %v; Weight: %v", name, age, weight)
}

/*
Output

Name: Axel Vasquez; Age: 22; Weight: 180.88
*/
```