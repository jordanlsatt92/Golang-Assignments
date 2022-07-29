package main

/*
Write a program that finds all strings that include the letter i followed by one or two instances of the letter n.
*/
import (
	"fmt"
	"regexp"
)

func findAllStrings(str string) []string{
	r,_ := regexp.Compile(`[^ ]*in{1,2}[^ |n]+`)
	retString:= r.FindAllString(str,-1)
	return retString
}

func main(){ 
	str1:=" thing pinnng strinnng snip innner innning pies nope not here "
	results := findAllStrings(str1)
	for i := range results {		
		fmt.Println(results[i])
	}
}