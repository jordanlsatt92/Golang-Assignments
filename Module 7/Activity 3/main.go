//Jordan Satterfield
/*
Create a goroutine/channels to send and receive structure data type free to design for any purpose
*/
package main

import (
	"SendReceiveStructsModule/dependencies"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

//changes the id number of the employee
func changeIdNumber(employees <-chan dependencies.Employee, changes chan<- dependencies.Employee){
	defer wg.Done()
	for emp := range employees{
		fmt.Printf("Changing %v's ID number from %v to %v\n",emp.Name, emp.IdNum, emp.IdNum+1)
		emp.IdNum ++
		changes <- emp
	}
}

func main(){

	structArray:= [] dependencies.Employee{dependencies.Employee{"Jordan", 1}, dependencies.Employee{"Hank", 2}, dependencies.Employee{"William", 3}}
	fmt.Println("Original structs:",structArray)

	ch1 := make(chan dependencies.Employee, len(structArray))
	ch2 := make(chan dependencies.Employee, len(structArray))

	wg.Add(1)
	go changeIdNumber(ch1, ch2)
	fmt.Println("Called go routine")
	for i:=0; i < len(structArray);i++{
		ch1 <- structArray[i]
	}

	close(ch1)
	wg.Wait()
	close(ch2)
	

	for val := range ch2{
		fmt.Println(val)
	}

}