// Jordan Satterfield
/*
Create Two go routine where routine 1 generate random number and append them in  slice where another
goroutine sort the slice at the same time . print the slice after every append and sorted array at the same time  side by side
*/

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

var wg sync.WaitGroup

/*
@description: Creates an empty slice and appends random numbers to slice. Sends slice to channel after every append
@param: ch (chan<- []int) that channel to which the function sends slices
*/
func generateNum(ch chan<- []int) {
	array:= make([]int,0)
	var arrayCopy []int
	for i:=0; i < 20; i++ {
		rand.Seed(time.Now().UnixNano())
		randNum:= rand.Intn(50)
		array = append(array, randNum)
		arrayCopy = make([]int, len(array))
		copy(arrayCopy, array)
		ch <- arrayCopy
	}
	close(ch)
	wg.Done()
}

/*
@description: Function that receives an array from the channel, prints the array, sorts the array, and prints the sorted array
@param: ch (<-chan []int) the channel from which the function receives arrays
*/
func sortNum(ch <-chan []int){
	for array := range ch{
		fmt.Println("Unsorted:", array)
		sort.Slice(array, func(i, j int) bool{
			return array[i] < array[j]
		})
		fmt.Println("Sorted:",array)
	}
	wg.Done()
}


func main(){
	ch:= make(chan []int)

	wg.Add(2)
	go generateNum(ch)
	go sortNum(ch)
	wg.Wait()

}