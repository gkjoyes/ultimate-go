// Sample program to show how to implement a custom error type based on the json package in the standard library.
package main

import (
	"fmt"
	"reflect"
)

// An UnmarshalTypeError describes a JSON value that was not appropriate for a value of a specific Go type.
type UnmarshalTypeError struct {
	Value string
	Type  reflect.Type
}

// Error implements the error interface.
func (e *UnmarshalTypeError) Error() string {
	return fmt.Sprintf("json: cannot unmarshal %s into Go value of type %s", e.Value, e.Type.String())
}

// An InvalidUnmarshalError describes an invalid argument passed to Unmarshal.
type InvalidUnmarshalError struct {
	Type reflect.Type
}

// Error implements the error interface.
func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "json: Unmarshal(nil)"
	}

	if e.Type.Kind() != reflect.Ptr {
		return "json: Unmarshal(non-pointer " + e.Type.String() + ")"
	}
	return "json: Unmarshal(nil " + e.Type.String() + ")"
}

// user is a type for use in the Unmarshal call.
type user struct {
	Name int
}

func main() {
	var u user
	if err := Unmarshal([]byte(`{"name":"bill"}`), u); err != nil {
		switch e := err.(type) {
		case *UnmarshalTypeError:
			fmt.Printf("UnmarshalTypeError: Value[%s] Type[%v]\n", e.Value, e.Type)
		case *InvalidUnmarshalError:
			fmt.Printf("InvalidUnmarshalError: Type[%v]\n", e.Type)
		default:
			fmt.Println(err)
		}
		return
	}
	fmt.Println("Name:", u.Name)
}

// Unmarshal simulates an unmarshal call that always fails.
func Unmarshal(data []byte, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}
	return &UnmarshalTypeError{"string", reflect.TypeOf(v)}
}
