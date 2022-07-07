package main

// Jordan Satterfield
// Description: Create a map and sort the map based on the key length in ascending order
import (
	"fmt"
	"sort"
)

func main() {
	m := make(map[string]int32)
	m["orange"] = 5
	m["apple"] = 7
	m["strawberry"] = 9
	m["mango"] = 3

	keys := make([]string, 0)
	keyLength := make([]int, 0)
	for key,_ := range m {
		keys = append(keys, key)
		keyLength = append(keyLength, len(string(key)))

	}

	sort.SliceStable(keys, func(i, j int) bool {
		return keyLength[i] < keyLength[j]
	})

	for _, j := range keys{
		fmt.Println(j, m[j])
	}
}