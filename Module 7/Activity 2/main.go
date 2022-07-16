// Jordan Satterfield
/*
Take a Paragraphs of text Max word  up to 200 words in a Slice or String and then Pass each Sentence(up to full stop)
to another go routine/routines   1)print it in reverse order(1 goroutine) 2 )count its words(2 goroutine)
*/
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func reverse(s string, ch chan string){
	reversed:=" "
	for i:=len(s)-1; i >= 0; i--{
		reversed += string(s[i])
	}
	ch <-reversed
	wg.Done()
}

func wordCount(s string, ch chan int){
	count:=1
	i:=0
	if string(s[0]) == " "{
		i = 1
	}
	for ; i < len(s); i++{
		if string(s[i]) == " "{
			count++
		}
	}
	ch <- count
	wg.Done()
}

func main(){

	var str string =  "One two three. Four five six seven! eight \"nine?\""

	start,stop:= 0,0

	ch1Reverse := make(chan string,100)
	ch2Count := make(chan int,100)

	for i:=0; i < len(str);i++{
		if string(str[i]) == "." || string(str[i]) == "?" || string(str[i]) == "!"{
			if i != len(str)-1 && string(str[i+1]) == "\""{
				stop = i+2
			}else{
				stop = i+1
			}
			wg.Add(2)
			if start == 0{
				go reverse(str[:stop], ch1Reverse)
				go wordCount(str[:stop],ch2Count)
				continue
			}
			go reverse(str[start:stop], ch1Reverse)
			go wordCount(str[start:stop],ch2Count)
		}
		start = stop
	}
	wg.Wait()
	close(ch1Reverse)
	close(ch2Count)

	for val := range ch1Reverse{
		fmt.Println(val)
	}
	for count:= range ch2Count{
		fmt.Println(count)
	}

}