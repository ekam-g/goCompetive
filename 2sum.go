package main

func twoSum(nums []int, target int) []int {
	for ans1, val1 := range nums {
		for ans2, val2 := range nums {
			if ans2 == ans1 {
				continue
			} else if (val1 + val2) == target {
				return []int{ans1, ans2}
			}
		}
	}
	return []int{}
}
