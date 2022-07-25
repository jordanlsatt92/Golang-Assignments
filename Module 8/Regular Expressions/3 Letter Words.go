package main

import (
	"fmt"
	"regexp"
)

func main(){

	str := "did"// dad gig gag tic tac toc"

	r,_:= regexp.Compile(`(^.)./1`)

	results:= r.FindAllString(str,-1)

	for i:= range results{
		fmt.Println(results[i])
	}
}