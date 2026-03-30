package baselearn

import (
	"fmt"
	"testing"
)

/* 
Example test for interface implementation 
go test -v -run TestInterfaceExample
*/
func TestInterfaceExample(t *testing.T) {
	fmt.Printf("Hello, World!\n")
	interfaceExample()
}

/** 
Example test for empty interface 
go test -v -run TestPrintValue
*/
func TestPrintValue(t *testing.T) {
	printValue(42)
	printValue("Hello, Go!")
	printValue(3.14)
	printValue([]int{1, 2, 3})
	printValue(map[string]int{"a": 1, "b": 2})
}