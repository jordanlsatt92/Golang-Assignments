package main

/*
Write a program that finds all strings that include the letter i followed by one or more instances of the letter n.
*/
import (
	"fmt"
	"regexp"
)

func findAllStrings(str string) []string{
	r,_ := regexp.Compile(`[^\s]*in+[^\s]*`)
	retString:= r.FindAllString(str,-1)
	return retString
}

func main(){ 
	str1:=" thing ping string snip inner inning pies nope not here "
	results := findAllStrings(str1)
	for i,_ := range results{
		fmt.Println(results[i])
	}
}