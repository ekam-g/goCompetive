package main

/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.

Example 1:

Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

Example 2:

Input: height = [1,1]
Output: 1

Constraints:

	n == height.length
	2 <= n <= 105
	0 <= height[i] <= 104

Accepted
2.2M
Submissions
4M
Acceptance Rate
54.0%
*/
func maxArea(height []int) int {
	highest := 0
	for x1, y1 := range height {
		for x2, y2 := range height {
			length := x1 - x2
			if length <= 0 {
				continue
			}
			height := smallest(y1, y2)
			area := height * length
			if highest < area {
				highest = area
			}
		}
	}
	return highest
}

func smallest(x, y int) int {
	if x > y {
		return y
	}
	return x
}
