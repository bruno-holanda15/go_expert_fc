package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

