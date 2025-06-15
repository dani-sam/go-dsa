package main

import (
	"fmt"
)

func main() {
	fmt.Println("Two Sum")
	nums := []int{2, 5, 3, 1}
	target := 4
	rs := twoSum(nums, target)
	fmt.Println(rs)
}

// brute force
func twoSum(nums []int, target int) []int {

	//taking a res array for output
	res := []int{}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i]+nums[j] == target && i != j {
				res = []int{i, j}
				return res
			}
		}
	}
	return res
}

// Hash table
func twoSum(nums []int, target int) []int {
	prev_map := make(map[int]int)
	for index, val := range nums {
		if j, ok := prev_map[target-val]; ok {
			return []int{j, index}
		}
		prev_map[val] = index
	}
	return nil
}
