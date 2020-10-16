package main

import "fmt"

func main() {
	var n int
	var m int
	var sum int

	fmt.Scan(&n)
	s := make([]int, 0, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		s = append(s, m)
	}

	for _, v := range s {
		sum += v
	}

	fmt.Println(sum)
}
