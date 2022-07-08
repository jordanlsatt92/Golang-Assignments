package main

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