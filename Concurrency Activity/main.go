package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex


// Runs two go routines, one is sending numbers 1-100, and the other is sending 100-1. The numbers then get printed.
func main(){

	ch1:= make(chan string)
	wg.Add(2)

	go func(){
		wg.Wait()
		close(ch1)
	}()

	go func(){
		for i:=0; i <= 100; i++{
			str:= "One-100 " + strconv.Itoa(i)
			mutex.Lock()
				ch1 <- str
			mutex.Unlock()
			time.Sleep(time.Millisecond * 7)
		}
		wg.Done()

	}()

	go func(){
		for i := 100; i >= 0; i--{
			str:= "100-One " + strconv.Itoa(i)
			mutex.Lock()
				ch1 <- str
			mutex.Unlock()
			time.Sleep(time.Millisecond * 7)
		}
		wg.Done()
	}()
	
	for n:= range ch1{
		fmt.Println(n)
	}

}

