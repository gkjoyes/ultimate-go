// Sample program to show how to declare function variables.
package main

import "fmt"

type data struct {
	name string
	age  int
}

// displayName provides a pretty print view of the name.
func (d data) displayName() {
	fmt.Println("My Name Is", d.name)
}

// setAge sets the age and display the value.
func (d *data) setAge(age int) {
	d.age = age
	fmt.Println(d.name, "Is Age", d.age)
}

func main() {
	d := data{
		name: "x1",
	}

	fmt.Println("Proper calls to methods:")
	d.displayName()
	d.setAge(30)

	fmt.Println("What the compiler is doing:")
	data.displayName(d)
	(*data).setAge(&d, 30)

	fmt.Println("Call value receiver methods with variable:")
	valueSematics(d)
	pointerSematics(d)
}

func valueSematics(d data) {
	// Declare a function variable for the method bound to the d variable.
	// The function variable will get its own copy of d becuase the method is using a value receiver.
	f1 := d.displayName

	// Call the method via the variable.
	f1()

	// Change the value of d.
	d.name = "x2"

	// Call the method via the variable. We don't see the change.
	f1()
}

func pointerSematics(d data) {
	// Declare a function variable for the method bound to the d variable.
	// The function variable will get the address of d because the method is using a pointer receiver.
	f2 := d.setAge

	// Call the method via the variable.
	f2(45)

	// Change the value of d.
	d.name = "x2"

	// Call the method via the variable. We see the change.
	f2(45)
}
