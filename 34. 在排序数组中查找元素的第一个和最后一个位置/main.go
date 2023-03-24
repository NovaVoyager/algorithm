package main

import "fmt"

//https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
func main() {
	//fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	//fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	//fmt.Println(searchRange([]int{}, 0))
	//fmt.Println(searchRange([]int{1, 4}, 4))

	//fmt.Println(searchRange2([]int{5, 7, 7, 8, 8, 10}, 8))
	//fmt.Println(searchRange2([]int{5, 7, 7, 8, 8, 10}, 6))
	//fmt.Println(searchRange2([]int{}, 0))
	//fmt.Println(searchRange2([]int{1, 4}, 4))

	fmt.Println(searchRange3([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange3([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange3([]int{}, 0))
	fmt.Println(searchRange3([]int{1, 4}, 4))
}

//二分法
func searchRange3(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	return []int{searchFirst(nums, target), searchLast(nums, target)}
}

func searchFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if left > len(nums)-1 || nums[left] != target {
		return -1
	}

	return left
}

func searchLast(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if right < 0 || nums[right] != target {
		return -1
	}

	return right
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	start, end := -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == target && start == -1 { //第一个位置
			start = i
		}
		if start != -1 {
			if nums[i] > target { //查找到最后了
				end = i - 1
				break
			} else if nums[i] == target {
				end = i
			}
		}
	}

	if len(nums) == 1 && start == 0 {
		end = start
	}

	return []int{start, end}
}

//searchRange2 双指针
func searchRange2(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	startIndex, endIndex := 0, len(nums)-1
	startFlag, endFlag := false, false
	for range nums {
		if startFlag && endFlag {
			break
		}

		if nums[startIndex] != target && !startFlag {
			startIndex++
		} else {
			startFlag = true
		}

		if nums[endIndex] != target && !endFlag {
			endIndex--
		} else {
			endFlag = true
		}
	}

	if !startFlag && !endFlag {
		return []int{-1, -1}
	}

	return []int{startIndex, endIndex}
}
