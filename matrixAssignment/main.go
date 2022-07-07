//Jordan Satterfield
package main

import (
	"fmt"
	"matrixModule/dependencies"
)

func main(){

	var matrix [][] int

	//Append user input slice to matrix
	for i:=0; i < 3; i++{
		matrix = append(matrix, dependencies.ReturnSlice())
	}

	// Print the matrix ints by row
	for i:=0; i < len(matrix);i++{
		for j:=0; j <= i; j++{
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}

	// Switch top and bottom row of matrix
	fmt.Println("\nSwitch top and bottom row:")
	dependencies.SwitchRows(matrix)
	for i:=0; i < len(matrix);i++{
		for j:=0; j < len(matrix); j++{
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
	dependencies.SwitchRows(matrix) // switch back to original
	
	// Switch top and bottom row of matrix
	fmt.Println("\nSwitch left-most and right-most columns:")
	dependencies.SwitchColumns(matrix)
	for i:=0; i < len(matrix);i++{
		for j:=0; j < len(matrix); j++{
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}