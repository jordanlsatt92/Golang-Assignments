package main

// Jordan Satterfield
// Description: Take 20 ints from the console between 1-100, then print the sumary like 1-10: Count-5,
// 11-20: count:7
import (
	"fmt"
	"sort"
)

func main(){

	m:=make(map[int] int)
	m[10] = 0
	m[20] = 0
	m[30] = 0
	m[40] = 0
	m[50] = 0
	m[60] = 0
	m[70] = 0
	m[80] = 0
	m[90] = 0
	m[100] = 0

	keys:=make([]int,0)
	for key,_ := range m{
		keys = append(keys, int(key))
	}
	
	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	fmt.Println(keys)
	var input int
	for i:=0; i < 20; i++{
		fmt.Print("Please enter a number between 1 and 100 (this is going to happen 20 times):")
		fmt.Scan(&input)
		for i:=0; i < len(keys); i++{
			if input <= keys[i]{
				m[keys[i]]++
				break
			}
		}
	}

	for _,val:= range keys{
		fmt.Printf("Numbers between %v - %v: Count=%v\n",val-9,val,m[val])
	}
}