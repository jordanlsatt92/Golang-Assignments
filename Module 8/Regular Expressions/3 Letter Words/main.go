package main

// Find words 3 letter words that start and end with same letter
import (
	"fmt"

	"github.com/dlclark/regexp2"
)

// FindAllString like function
func regexp2FindAllString(re *regexp2.Regexp, s string) []string {
	var matches []string
	m, _ := re.FindStringMatch(s)
	for m != nil {
		matches = append(matches, m.String())
		m, _ = re.FindNextMatch(m)
	}
	return matches
}


func main() {
	str := "did dad pig pip pop you peep"

	r:= regexp2.MustCompile(`([^\s])[^\s]\1`,0)

	results := regexp2FindAllString(r, str)

	for i := range results{
		fmt.Println(results[i])
	} 
}