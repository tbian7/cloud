package main

import "fmt"

type Inner struct {
	A int
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)
}

func (i Inner) Double() string {
	result := i.A * 2
	return i.IntPrinter(result)
}

type Outer struct {
	Inner
	A int
	S string
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}

func (o Outer) Double() string {
	result := o.A * 2
	return o.IntPrinter(result)
}

func main() {
	o := Outer{
		Inner: Inner{
			A: 10,
		},
		S: "Hello",
		A: 100,
	}
	fmt.Println(o.Inner.Double())
}
