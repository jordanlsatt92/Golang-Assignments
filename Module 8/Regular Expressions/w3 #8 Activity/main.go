package main

import (
	"fmt"
	"regexp"
)

func main(){
	str := "You will See asdfThis"

	r,_ := regexp.Compile(`[A-Z][a-z]+`)

	results := r.FindAllString(str, -1)

	for i := range results{
		fmt.Println(results[i])
	}
}