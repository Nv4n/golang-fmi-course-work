package main

import (
	"fmt"
	"sort"
)

func Part(n int) string {
	// your code
	prod := Prod(Enum(n))
	rangeVal := prod[len(prod)-1] - prod[0]
	sum := 0
	for _, val := range prod {
		sum += val
	}
	length := len(prod)
	avg := float64(sum) / float64(length)

	var median float64
	if length == 1 {
		median = float64(prod[0])
	} else if length%2 != 0 {
		median = float64(prod[length/2+1])
	} else {
		median = float64(prod[length/2]+prod[length/2+1]) / 2.0
	}

	return fmt.Sprintf(
		"Range: %v Average: %.02f Median: %.02f",
		rangeVal, avg, median)
}

func Enum(n int) [][]int {
	var res [][]int
	for i := n; i > 0; i-- {
		subEnums := Enum(n - i)
		for i, sub := range subEnums {
			subEnums[i] = append(sub, i)
		}
		if subEnums == nil {
			subEnums = [][]int{{i}}
		}
		res = append(res, subEnums...)
	}
	return res
}

func Prod(enum [][]int) (res []int) {
	prodsMap := make(map[int]bool)
	for _, subEnum := range enum {
		subProd := 1
		for _, val := range subEnum {
			subProd *= val
		}
		prodsMap[subProd] = true
	}

	for k, _ := range prodsMap {
		res = append(res, k)
	}
	sort.Ints(res)
	return res
}
func main() {
	fmt.Println(Part(2))
}
