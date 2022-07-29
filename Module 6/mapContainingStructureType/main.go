package main

// Create a map where values must be a structure type
import (
	"fmt"
	"mapContainingStructureTypeModule/dependencies"
)

func main(){

	e1:=dependencies.Employee{1, "Sarah"}
	e2:=dependencies.Employee{2, "James"}
	e3:=dependencies.Employee{3, "John"}
	e4:=dependencies.Employee{4, "Nancy"}

	m:=make(map[int]dependencies.Employee)
	m[e1.EmployeeID] = e1
	m[e2.EmployeeID] = e2
	m[e3.EmployeeID] = e3
	m[e4.EmployeeID] = e4

	fmt.Println(m)
}