package dependencies

import (
	"fmt"
	"strings"
)

func TakeInput(n int) []string {
	inputArray := make([]string, 0)
	var input string
	for i := 0; i < n; i++ {
		fmt.Print("Please enter a string:")
		fmt.Scan(&input)
		inputArray = append(inputArray, strings.ToLower(input))
	}
	return inputArray
}

func Average(array []string) float64 {
	var sum,count int
	for i := 0; i < len(array) ; i++ {
		sum += len(array[i])
		count++
	}
	return float64(sum) / float64(count)
}

func IdLongWords(array []string, avg float64) []string {
	longWords := make([]string, 1)
	for i := 0; i < len(array); i++ {
		if float64(len(array[i])) > avg {
			longWords = append(longWords, array[i])
		}
	}
	return longWords
}

func IdShortWords(array []string, avg float64) []string {
	shortWords := make([]string, 1)
	for i := 0; i < len(array); i++ {
		if float64(len(array[i])) < avg {
			shortWords = append(shortWords, array[i])
		}
	}
	return shortWords
}