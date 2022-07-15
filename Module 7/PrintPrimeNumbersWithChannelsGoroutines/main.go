package main

import (
	"fmt"
)

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

func primPrinter(c chan int) chan int{
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
	p1:=primPrinter(n1)
	for n:= range p1{
		fmt.Println(n,"is prime.")
	}
}