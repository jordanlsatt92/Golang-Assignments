package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "firstWord of the string"

	r, _ := regexp.Compile(`[^.][^\s]+`)

	results := r.FindAllString(str, -1)

	for i := range results{
		fmt.Println(results[i])
	} 
}