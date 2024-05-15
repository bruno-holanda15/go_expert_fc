package main

import (
	"fmt"
	"math"
)

var squareRoots map[int]struct{}

func main() {
	fmt.Println(NextPerfectSquarePermutation(100, 2))
	fmt.Println(NextPerfectSquarePermutation(100, 3))
	fmt.Println(NextPerfectSquarePermutation(100, 4))
}

func NextPerfectSquarePermutation(lowerLimit, numberOfSquares int) int {
  squareRoots = make(map[int]struct{})
	nextSquareNumber := int(math.Sqrt(float64(lowerLimit))) + 1

	for {
		gratest, squares := permutation(nextSquareNumber)
		if squares == numberOfSquares {
			return gratest
		}

		nextSquareNumber += 1
	}
}

func permutation(nextSquareRoot int) (gtsNumber int, numOfSquares int) {
	if _, ok := squareRoots[nextSquareRoot]; ok {
		return 0, 0
	}

	squareRoots[nextSquareRoot] = struct{}{}
	realNumber := nextSquareRoot * nextSquareRoot

	if nextSquareRoot%10 == 0 || hasZero(realNumber) {
		return 0, 0
	}

	group := heapAlgorithm(realNumber)

	auxFilter := make(map[int]struct{})

	for _, numSLice := range group {
		num := sliceToInt(numSLice)

		if _, ok := auxFilter[num]; ok {
			continue
		}

		auxFilter[num] = struct{}{}
		squareRoot := math.Sqrt(float64(num))
		parsedSquareRoot := int(squareRoot)

		if squareRoot == float64(parsedSquareRoot) {
			squareRoots[parsedSquareRoot] = struct{}{}
			numOfSquares += 1

			// Get the greatest number
			if gtsNumber < num {
				gtsNumber = num
			}
		}
	}

	return
}

func hasZero(number int) bool {
	numbers := intToSlice(number)

	for _, value := range numbers {
		if value == 0 {
			return true
		}
	}

	return false
}

func heapAlgorithm(number int) [][]int {
	var helper func([]int, int)
	res := [][]int{}
	numGroup := intToSlice(number)

	helper = func(numbers []int, n int) {
		if n == 1 {
			aux := make([]int, len(numbers))
			copy(aux, numbers)
			res = append(res, aux)
		} else {
			for i := 0; i < n; i++ {
				helper(numbers, n-1)
				if n%2 == 1 {
					numbers[i], numbers[n-1] = numbers[n-1], numbers[i]
				} else {
					numbers[0], numbers[n-1] = numbers[n-1], numbers[0]
				}
			}
		}
	}

	helper(numGroup, len(numGroup))
	return res
}

func sliceToInt(group []int) int {
	result := 0
	powAux := 0
	for i := len(group) - 1; i >= 0; i-- {
		result += group[i] * int(math.Pow10(powAux))

		powAux += 1
	}

	return result
}

func intToSlice(k int) []int {
	var helper func(num int, sequence []int) []int
	arr := []int{}

	helper = func(num int, sequence []int) []int {
		if num > 0 {
			sequence = append([]int{num % 10}, sequence...)
			return helper(num/10, sequence)
		}

		return sequence
	}

	return helper(k, arr)
}