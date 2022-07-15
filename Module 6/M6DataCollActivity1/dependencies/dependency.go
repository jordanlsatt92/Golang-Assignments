package dependencies

import (
	"math/rand"
	"time"
)

func GenerateArray(n int) [] int{
	var array = make([]int,0)
	for i:=0; i < n; i++{
		rand.Seed(time.Now().UnixNano()+int64(i)) //added the i because I was getting multiple of the same number in a row
		var randNum = rand.Intn(99+99+1)-99
		if randNum==-100{randNum++}
		array = append(array, randNum)
	}
	return array

}