package dependencies

import (
	"math/rand"
	"sort"
	"time"
)

func GenerateArray() [100] int{
	var array [100] int
	for i:=0; i < len(array); i++{
		rand.Seed(time.Now().UnixNano()+int64(i)) //added the i because I was getting multiple of the same number in a row
		var randNum = rand.Intn(99+99+1)-99
		if randNum==-100{randNum++}
		array[i] = randNum
	}
	return array
}

func ComputeMax(array [100]int) int{
	var max int = -200
	for i:=0; i < len(array);i++{
		if array[i] > max{
			max = array[i]
		}
	}
	return max
}

func ComputeMaxIndex(array [100]int) int{
	var max int = -200
	var maxInd int
	for pos, val:=range array{
		if val > max{
			max = val
			maxInd = pos
		}
	}
	return maxInd
}

func ComputeMin(array [100]int) int{
	var min int = 200
	for i:=0; i < len(array);i++{
		if array[i] < min{
			min = array[i]
		}
	}
	return min
}

func ComputeMinIndex(array [100]int) int{
	var min int = 200
	var minInd int
	for pos, val:=range array{
		if val < min{
			min = val
			minInd = pos
		}
	}
	return minInd
}

func SortDescending(array [100]int) [100]int{
	var duplicate = array[:]
	var retArray [100] int 
	sort.Slice(duplicate, func(i, j int) bool {
		return duplicate[i] > duplicate[j]
	})
	copy(retArray[:], duplicate)
	return retArray
}

func SortAscending(array [100]int) [100]int{
	var duplicate = array[:]
	var retArray [100] int
	sort.Ints(duplicate)
	copy(retArray[:], duplicate)
	return retArray
}

func Mean(array [100]int) float64{
	var total float64
	for _,val := range array{
		total+=float64(val)
	}
	return total / float64(len(array))
}

func Median(array [100] int) int{
	var duplicate = array
	duplicate = SortAscending(array)
	return duplicate[len(duplicate) / 2]
}

func IdPositives(array [100] int) [] int{
	sorted:= SortAscending(array)
	var retArray []int
	for pos,val:= range sorted{
		if val < 1{
			continue
		}else{
			retArray = sorted[pos:]
			break
		}
	}
	return retArray
}

func IdNegatives(array [100] int) [] int{
	sorted:= SortAscending(array)
	var retArray []int
	for pos,val:= range sorted{
		if val < 0{
			continue
		}else{
			retArray = sorted[:pos]
			break
		}
	}
	return retArray
}

func LongestSequence(array [100] int) [] int{
	var longestSeq,temp = make([]int,0),make([]int,0)
	//var pos, val int
	
	for i:=0; i < len(array); i++{
		if len(temp) == 0{
			temp = append(temp, array[i])
			continue
		}
		if temp[len(temp)-1] < array[i]{
			temp = append(temp, array[i])
			continue
		}else{
			if len(temp) > len(longestSeq){
				longestSeq=make([]int, len(temp))
				copy(longestSeq,temp)
			}
			temp = make([]int,0)
		}
	}
	return longestSeq
}

func RemoveDuplicates(array [100] int) [] int{
	m:=make(map[int]int)
	retArray := make([]int,0)
	for i:=0; i < len(array); i++{
		if (m[array[i]] == 0){
			if (array[i] == 0){
				retArray = append(retArray, array[i])
				m[array[i]] = 1000
				continue
			}
			m[array[i]] = array[i]
			retArray = append(retArray, array[i])
		}
		//retArray = append(retArray,array[i])
	}
	return retArray
}
