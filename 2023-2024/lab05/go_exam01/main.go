package main

import "fmt"

func findWinner(n, m int) (p int) {
	people := make(map[int]bool)
	for i := 0; i < n; i++ {
		people[i] = true
	}
	i := 0
	for free, count := len(people), 0; free > 1; i = i + 1 {
		if i == len(people) {
			i = 0
		}
		if !people[i] {
			continue
		}

		count++
		if count == m {
			people[i] = false
			free = free - 1
			count = 0
		}
	}
	if i == len(people) {
		return 0
	}
	return i
}

func main() {
	var n, m int
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &m)
	p := findWinner(n, m)
	fmt.Printf("N: %d; M: %d; P: %d\n", n, m, p)

	//p = findWinner(8, 3)
	//fmt.Printf("N: %d; M: %d; P: %d\n", 8, 3, p)
	//p = findWinner(11, 5)
	//fmt.Printf("N: %d; M: %d; P: %d\n", 11, 5, p)

}
