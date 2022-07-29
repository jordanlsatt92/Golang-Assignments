package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "askdjhfsb a@#FSDFxdfb asdlfiahj weoijgsldikijb"

	r, _ := regexp.Compile(`a[^\s]*b`)

	results := r.FindAllString(str, -1)

	for i := range results{
		fmt.Println(results[i])
	} 
}