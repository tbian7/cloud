// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type firstPerson struct {
	name string
	age  int
}

type secondPerson struct {
	name string
	age  int
}

func main() {
	f := firstPerson{
		name: "Bob",
		age:  50,
	}
	var g struct {
		name string
		age  int
	}
	g = f
	fmt.Println(f == g)

	var s secondPerson = secondPerson(f)
	// fmt.Println(f == s) //panic
	fmt.Println(g == s) //true

	more()
}

type myInterface interface {
	Print()
}

type myType struct{}

func (myType) Print() {}

func more() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	a := func() {}
	describe(a)

	var b myInterface
	describe(b)
	if b == nil {
		fmt.Println("b == nil")
	}

	b = myType{}
	describe(b)
	if b == nil {
		fmt.Println("b with type == nil")
	}
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
