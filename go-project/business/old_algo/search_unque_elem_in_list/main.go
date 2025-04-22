package main

import "fmt"

func solution(cards []int) int {
	// Edit your code here
	result := 0
	answerMap := make(map[int]int, 0)
	//answerList := make([]int,0)
	for _, card := range cards {
		if _, ok := answerMap[card]; !ok {
			answerMap[card] = 1
		} else {
			answerMap[card]++
		}
	}
	for k, v := range answerMap {
		if v == 1 {
			result = k
		}
	}
	return result

}

func main() {
	// Add your test cases here

	fmt.Println(solution([]int{1, 1, 2, 2, 3, 3, 4, 5, 5}) == 4)
	fmt.Println(solution([]int{0, 1, 0, 1, 2}) == 2)
}
