package dependencies

import "fmt"

// Create 1 dimensional slice of size 3
// Return slice of user input

// Creates a slice containing the integers entered by the user
// Returns the slice
func ReturnSlice() []int{
	row:=make([]int,3)
	for j:=0; j < 3;j++{
		var input int
		fmt.Println("Please enter an integer:")
		fmt.Scan(&input)
		row[j] = input
	}
	return row
}

// Switches the top and bottom rows of the matrix
func SwitchRows(mat [][]int){
	temp:=mat[0]
	mat[0] = mat[len(mat)-1]
	mat[len(mat)-1] = temp
}

// Switches the left-most and right-most columns of the matrix
func SwitchColumns(mat [][]int){
	var temp int
	for i:=0; i < len(mat); i++{
		temp = mat[i][0]
		mat[i][0] = mat[i][len(mat)-1]
		mat[i][len(mat)-1] = temp
	}
}
