// Jordan Satterfield
// Assignment: Create a map where values must be a structure type

package main

import "fmt"

type Employee struct{
	employeeID int
	name string
}

func main(){

	e1:=Employee{1, "Sarah"}
	e2:=Employee{2, "James"}
	e3:=Employee{3, "John"}
	e4:=Employee{4, "Nancy"}

	m:=make(map[int]Employee)
	m[e1.employeeID] = e1
	m[e2.employeeID] = e2
	m[e3.employeeID] = e3
	m[e4.employeeID] = e4

	fmt.Println(m)
}