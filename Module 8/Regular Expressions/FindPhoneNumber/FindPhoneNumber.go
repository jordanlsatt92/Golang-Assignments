package main

import (
	"fmt"
	"regexp"
)

func main(){
	str:= "1234 10a203984932 18284003000 1.828.400.3000 1-828-400-3000 1(828)-400-3000 1-(828)-400-3000"

	r,_ := regexp.Compile(`[^\s\D]*[0-9][\.\(-]*\d{3}[-\.)]*\d{3}[\.-]?\d{4}`)

	results:= r.FindAllString(str,-1)
	for i := range results{
		fmt.Println(results[i])
	}
}