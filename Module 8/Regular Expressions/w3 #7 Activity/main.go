package main

import (
	"fmt"
	"regexp"
)

func main(){
	str := "a_b_c_d ab_cd abcd"

	r,_ := regexp.Compile(`([a-z]_)+[a-z]`)

	results:= r.FindAllString(str, -1)

	for i := range results{
		fmt.Println(results[i])
	}
}