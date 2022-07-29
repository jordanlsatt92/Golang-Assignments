package main

import (
	"fmt"
	"strconv"
	"time"
)

func main(){
	currentTime := time.Now().Unix()
	i,_ := strconv.ParseInt(strconv.Itoa(int(currentTime)), 10 , 64)
	t := time.Unix(i,0)
	fmt.Println(t)
}