package main

import "fmt"

func main() {
	angka := []int{1, 5, 7, 9, 11}
	var target int = 17
	tampung := findComb(angka, target)
	fmt.Println(tampung)
	// fmt.Println(angka[1:])
	// fmt.Println(len(angka))
}

func findComb(nums []int, total int) [][]int {
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}
	t := total - nums[0]

	if t == 0 {
		res = append(res, []int{nums[0]})
	} else if t > 0 {
		res = findComb(nums[1:], t)
		for i, v := range res {
			res[i] = append([]int{nums[0]}, v...)
		}
	}

	for len(nums) > 1 && nums[0] == nums[1] {
		nums = nums[1:]
	}

	res = append(res, findComb(nums[1:], total)...)
	return res
}
