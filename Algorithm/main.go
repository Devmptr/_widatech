package main

import "fmt"

func main() {
	res := generate(3, 8)
	fmt.Printf("%+v", res)
}

func generate(length int, target int) [][]int {
	goodCombination := [][]int{}
	arrInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	temp := generatePairs(arrInt, length, 0, []int{}, [][]int{})

	for _, combination := range temp {
		if sum(combination) == target {
			goodCombination = append(goodCombination, combination)
		}
	}

	return goodCombination
}

func sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func generatePairs(arrData []int, length int, index int, tempPair []int, result [][]int) [][]int {
	if index == len(arrData) {
		if len(tempPair) == length {
			result = append(result, append([]int{}, tempPair...))
		}
		return result
	}

	tempPair = append(tempPair, arrData[index])
	result = generatePairs(arrData, length, index+1, tempPair, result)

	tempPair = tempPair[:len(tempPair)-1]
	result = generatePairs(arrData, length, index+1, tempPair, result)

	return result
}
