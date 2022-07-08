package main

// Jordan Satterfield

/*
Description:
Implement a function for each of the following bullets. Each function will take as an input an array of 100 integers. Do not use the sort package or any built-in function (like min, max, etc.).

    Compute the max of an array of int
    Compute the index of the max of an array of int
    Compute the min of an array of int
    Compute the index of the min of an array of int
    Sort an array of int in descending order and return the new sorted array in a separate array.
    Sort an array of int in ascending order and return the new sorted array in a separate array.
    Compute the mean of an array
    Compute the median of an array
    Identify all positive numbers in the array and return the numbers in a slice
    Identify all negative numbers in the array and return the numbers in a slice
    Compute the longest sequence of sorted numbers (in descending or ascending order) in the array and return in a new array
        Example: input: [1 45 67 87 6 57 0]
        Output: [1 45 67 87]
    Remove duplicates from an array of ints and return the unique elements in a slice
*/
import (
	"fmt"
	"m6DataCollActivity2Module/dependencies"
	"sort"
)

func main(){


	array:=dependencies.GenerateArray()
	fmt.Println("Original Array")
	fmt.Println(array)
	var max, maxInd, min, minInd int = dependencies.ComputeMax(array),dependencies.ComputeMaxIndex(array),dependencies.ComputeMin(array),dependencies.ComputeMinIndex(array)
	descendingSorted, ascendingSorted:= dependencies.SortDescending(array), dependencies.SortAscending(array)
	mean, median:= dependencies.Mean(array), dependencies.Median(array)
	uniqueElements:= dependencies.RemoveDuplicates(array)
	fmt.Println("\nThe max =",max)
	fmt.Println("\nThe index of",max,"=", maxInd)
	fmt.Println("\nThe min =",min)
	fmt.Println("\nThe index of",min,"=", minInd)
	fmt.Println("\nSorted in descending order:")
	fmt.Println(descendingSorted)
	fmt.Println("\nSorted in ascending order:")
	fmt.Println(ascendingSorted)
	fmt.Println("\nThe mean of the array =", mean)
	fmt.Println("\nThe median of the array =", median)
	fmt.Println("\nAll positive numbers:")
	fmt.Println(dependencies.IdPositives(array))
	fmt.Println("\nAll negative numbers:")
	fmt.Println(dependencies.IdNegatives(array))
	fmt.Println("\nThe longest sequence in the array:",dependencies.LongestSequence(array))
	fmt.Println("\nUnique elements in sorted order:")
	sort.Slice(uniqueElements, func (i,j int) bool {
		return uniqueElements[i] < uniqueElements[j]
	})
	fmt.Println(uniqueElements)
}