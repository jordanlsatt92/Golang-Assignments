// Jordan Satterfield
// Assignment 1
// package main

// import "fmt"

// func sum(numbers []int) int {
// 	var total int
// 	for _, val := range numbers {
// 		total += val
// 	}
// 	return total
// }

// func main() {
// 	slice := []int{1, 2, 3, 4}
// 	fmt.Println(sum(slice))
// 	slice2 := []int{1, 2, -1, -2}
// 	fmt.Println(sum(slice2))
// }

// Assignment 2
// Jordan Satterfield

// package main

// import "fmt"

// func half(n int) (int, bool){
// 	answer:= n / 2
// 	even:= n % 2 == 0
// 	return answer, even
// }

// func main() {
// 	fmt.Println(half(1))
// 	fmt.Println(half(2))
// }

// Assignment 3
// Jordan Satterfield

// package main

// import "fmt"

// func greatest(numbers... int) int{
// 	var largest int = 0
// 	for _,val:= range numbers{
// 		if val > largest{
// 			largest = val
// 		}
// 	}
// 	return largest
// }

// func main() {
// 	fmt.Println(greatest(10,9,21,8,7,12))
// 	fmt.Println(greatest(100,100,100,101,100,100))
// }

// // Assignment 4
// // Jordan Satterfield

// package main

// import "fmt"

// func makeOddGenerator() func() int{
// 	var number int = 1
// 	var temp int
// 	return func() int{
// 		temp = number
// 		number += 2
// 		return temp
// 	}
// }

// func main() {
// 	f:=makeOddGenerator()
// 	fmt.Println(f())
// 	fmt.Println(f())
// 	fmt.Println(f())
// 	fmt.Println(f())
// 	fmt.Println(f())
// }

// Assignment 5
// Jordan Satterfield

package main

import "fmt"

func fibonacii(number uint) int{
	if number == 2{
		return 1
	}else if number == 1{
		return 0
	}else{
		return fibonacii(number-1) + fibonacii(number-2)
	}
}

func main() {
	fmt.Println(fibonacii(1))
	fmt.Println(fibonacii(2))
	fmt.Println(fibonacii(3))
	fmt.Println(fibonacii(4))
	fmt.Println(fibonacii(5))
	fmt.Println(fibonacii(6))
}

