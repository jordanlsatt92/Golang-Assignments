package main

import (
	"fmt"
	"regexp"
)

func findAllStrings(str string) []string{
	r,_ := regexp.Compile(`[^ ]*in{3}[^ n]+`)
	retString:= r.FindAllString(str,-1)
	return retString
}

func main(){ 
	str1:=" thing pinnng strinnng snip innner innnning pies nope not here "
	results := findAllStrings(str1)
	for i,_ := range results{
		fmt.Println(results[i])
	}
}