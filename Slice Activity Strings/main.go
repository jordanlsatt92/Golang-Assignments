// Jordan Satterfield

package main

import (
	"fmt"
	"sliceActivityStringsModule/dependencies"
)

func main(){
	inputArray:= dependencies.TakeInput(10)
	fmt.Printf("\nThe average length of words entered : %.2f\n",dependencies.Average(inputArray))
	fmt.Println("Original slice:", inputArray)
	fmt.Println("Words longer than average:", dependencies.IdLongWords(inputArray, dependencies.Average(inputArray)))
	fmt.Println("Words shorter than average:", dependencies.IdShortWords(inputArray, dependencies.Average(inputArray)))
}