package int_info

import (
	"fmt"
	"math/rand"
	"time"
)

func main1() {
	cards := []int{2, 4, 51, 52, 3}
	fmt.Println(cards)
	shuffle(cards)
	fmt.Println(cards)

}

func shuffle(nums []int) {

	rand.Seed(time.Now().UnixNano())
	for i := len(nums) - 1; i > 0; i-- {

		j := rand.Intn(i + 1)

		nums[i], nums[j] = nums[j], nums[i]
	}
}
