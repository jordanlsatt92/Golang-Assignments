package main

import (
	"fmt"
)

// routine that loops throught 100 numbers and sends each number to the channel
func increment() chan int{
	out:=make(chan int)

	go func(){
		for i:=0; i <= 100; i++{
			out <-i
			// fmt.Println(i)
		}
		close(out)
	}()

	return out
}

// Receives numbers from a channel and the finds if it is prime then sends it
func primePrinter(c chan int) chan int{
	out:=make(chan int)
	go func(){
		for n:=range c{
			if n < 2{
				continue
			}
			isPrime := true
			for i:=2; i < n; i++{
				if n % i == 0{
					isPrime = false
					break
				}
			}
			if isPrime{
				out <- n
			}
		}
		close(out)
	}()
	return out
}

func main(){
	n1:=increment()
	p1:=primePrinter(n1)
	for n:= range p1{
		fmt.Println(n,"is prime.")
	}
}