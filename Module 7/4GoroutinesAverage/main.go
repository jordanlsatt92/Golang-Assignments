//Jordan Satterfield
/*
1)Create a slice of 20 of type int and take 20 number create 4 goroutines to take average of every 5 elements
ex[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20] ,so 1 goroutine should give average of first 5 elements
and next for another 5 so on
*/

package main

import (
	"fmt"
	"sync"
)

func returnAvg(array[] int, c chan  float64){
	sum:=0
	for i:= 0; i < len(array); i++{
		sum += array[i]
	}
	avg:= float64(sum / len(array)) 
	c <- avg
}

var wg sync.WaitGroup

func main(){
	slice:= []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	ch:= make(chan float64)
	
	start, stop:=0,len(slice)/4
	for i:=0; i < 4; i++{
		wg.Add(1)
		go returnAvg(slice[start:stop], ch)
		start, stop = start + len(slice)/4, stop + len(slice)/4
	}
	for i := 0; i < 4; i++{
		fmt.Println("Average:",<-ch)
	}
}