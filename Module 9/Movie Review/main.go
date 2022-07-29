package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type WordScore struct{
	Word string
	Score int
}
func main(){
	var bad, good int

	goodWords := make(map[string]*WordScore)
	badWords := make(map[string]string)

	goodFile,_ := os.Open("C:/Users/User/Downloads/GoodWords.txt")

	Scanner := bufio.NewScanner(goodFile)
	Scanner.Split(bufio.ScanWords)

	i:= 1
	var word string
	var score int
	for Scanner.Scan() {
		if i % 2 != 0{
			word = Scanner.Text()
		} else{
			score,_ = strconv.Atoi(Scanner.Text())
			goodWords[word] = &WordScore{word, score}
		}
		i ++
	}


	fmt.Println(goodWords["timeless"])
	badFile,_ := os.Open("C:/Users/User/Downloads/BadWords.txt")
	Scanner = bufio.NewScanner(badFile)
	Scanner.Split(bufio.ScanWords)

	for Scanner.Scan(){
		word = Scanner.Text()
		badWords[word] = word
	}

	review,_ := os.Open("C:/Users/User/Downloads/Review.txt")

	Scanner = bufio.NewScanner(review)
	Scanner.Split(bufio.ScanWords)

	for Scanner.Scan(){
		word = strings.ToLower(Scanner.Text())
		// fmt.Println(word)
		if goodWords[word] != nil {
			fmt.Println(goodWords[word].Word)
			good += goodWords[word].Score
		}else if badWords[word] != ""{
			bad += 1
		}
	}
	fmt.Println("Good:", good, "; Bad:", bad)
	fmt.Println("Score:", int(float64(good) / (float64(good) +float64(bad)) * 100), "%")
}