package main

import (
	"dsa/ds"
	"fmt"
)

func MinutesToSeconds(x int) int {
	return x / 60
}

func bitwise(x, y int) {
	fmt.Println("Bitwise AND", x&y)
	fmt.Println("Bitwise OR", x|y)
	fmt.Println("Bitwise XOR", x^y)
	fmt.Println("Right Shift by 2", x>>2)
	fmt.Println("Left shift by 1", x<<1)
}

func exits(x, y int) bool {
	if x >= 0 && y >= 0 {
		return true
	}
	return false
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	testval := image[sr][sc]
	image[sr][sc] = color
	points := [][]int{{sr - 1, sc}, {sr - 1, sc}, {sr, sc - 1}, {sr, sc + 1}}
	for _, p := range points {
		if exits(p[0], p[1]) {
			if testval == image[p[0]][p[1]] {
				image[p[0]][p[1]] = color
				floodFill(image, p[0], p[1], color)
			}
		}
	}
	return image
}

func tryStack() {
	s := ds.NewDeque[int]()
	fmt.Println(s)
	s.Push(10)
	fmt.Println(s)
	s.Push(11)
	fmt.Println(s)
	s.Push(12)
	fmt.Println(s)
	s.Push(13)
	fmt.Println(s)
	for !s.IsEmpty() {
		fmt.Println(s.Pop())
	}
}

type edge struct {
	to   int
	cost int
}

type Node struct {
	vertex int
	cost   int
}

// 724
func pivotIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	} else if len(nums) == 0 {
		return -1
	}
	start := -1
	end := len(nums) - 1
	leftSum := 0
	rightSum := nums[end]
	for start < end {
		if leftSum == rightSum && (end-start) == 2 {
			return (end + start) / 2
		} else if leftSum > rightSum {
			end--
			rightSum += nums[end]
		} else {
			start++
			leftSum += nums[start]
		}
	}
	return -1
}

// problem 1
func twoSum(nums []int, target int) []int {
	var x, y int
	for i, v := range nums {
		for j, u := range nums[i+1:] {
			if v+u == target {
				x, y = i, j+i+1
			}
		}
	}
	return []int{x, y}
}
func maxArea(height []int) int {
	l := len(height)
	var max int
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			area := min(height[i], height[j]) * (j - i)
			if area > max {
				max = area
			}
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	max := len(nums1) + len(nums2)
	i, j := 0, 0
	var arr []int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			arr = append(arr, nums1[i])
			i++
		} else {
			arr = append(arr, nums2[j])
			j++
		}
	}
	if i <= len(nums1)-1 {
		arr = append(arr, nums1[i:]...)
	}
	if j <= len(nums2)-1 {
		arr = append(arr, nums2[j:]...)
	}
	if max%2 == 0 {
		return (float64(arr[max/2] + arr[max/2-1])) / 2
	} else {
		return float64(arr[max/2])
	}
}

func findMaxAverage(nums []int, k int) float64 {
	if len(nums) < 1 {
		return float64(nums[0])
	}
	start := 1
	sum := 0
	for _, v := range nums[0:k] {
		sum += v
	}
	prev_avg := float64(sum) / float64(k)
	for start+k-1 < len(nums) {
		curr_avg := float64(prev_avg*float64(k)+float64(nums[start+k-1])-float64(nums[start-1])) / float64(k)
		if curr_avg > prev_avg {
			prev_avg = curr_avg
		}
		start++
	}
	return prev_avg
}

func totalFruit(fruits []int) int {
	var max int
	i := 0
	for i < len(fruits) {
		t1 := fruits[i]
		t2 := -1
		j := i + 1
		var t2_index int
		for j < len(fruits) {
			if fruits[j] == t1 {
				j++
				continue
			} else if fruits[j] != t1 && t2 == -1 {
				t2 = fruits[j]
				t2_index = j
				j++
				continue
			} else if fruits[j] == t1 || fruits[j] == t2 {
				j++
				continue
			} else {
				break
			}
		}
		if j-i > max {
			max = j - i
		}
		i = t2_index
		if j == len(fruits) {
			return max
		}

	}
	return max
}

func countSubarrays(nums []int, minK int, maxK int) int64 {
	var res int64 = 0
	var minFound, maxFound bool = false, false
	var start, minStart, maxStart int = 0, 0, 0

	for i := 0; i < len(nums); i++ {
		num := nums[i]

		if num < minK || num > maxK {
			minFound = false
			maxFound = false
			start = i + 1
		}

		if num == minK {
			minFound = true
			minStart = i
		}

		if num == maxK {
			maxFound = true
			maxStart = i
		}

		if minFound && maxFound {
			res += int64(min(minStart, maxStart) - start + 1)
		}
	}
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	curr1 := list1
	curr2 := list2
	var outHead *ListNode
	var outCurr *ListNode

	if curr1 == nil && curr2 == nil {
		return nil
	} else if curr1 == nil {
		return list2
	} else if curr2 == nil {
		return list1
	}

	// check until one of the current pointes becomes nil
	for curr1 != nil && curr2 != nil {
		var temp *ListNode
		if curr1.Val > curr2.Val {
			temp = curr1
			curr1 = curr1.Next
		} else {
			temp = curr2
			curr2 = curr2.Next
		}
		if outHead == nil {
			outHead = temp
			outCurr = temp
		}
		temp.Next = nil
		outCurr.Next = temp
		outCurr = temp
	}
	// check if one or both pointers became nil and handle appropriately
	if curr1 == nil && curr2 == nil {
		return outHead
	} else if curr1 == nil {
		outCurr.Next = curr2
	} else if curr2 == nil {
		outCurr.Next = curr1
	}
	return outHead
}

// func threeSum(nums []int) [][]int {
//     sort.Slice(nums, func(i, j int) bool {
//         return nums[i] < nums[j]
//     })
//     var result [][]int
//     for i,_ := range nums {
//         if i > 0 && nums[i]==nums[i-1] {
//             continue
//         }
//         j := i+1
//         k := len(nums)-1
//         for j < k {
//             if nums[i] + nums[j] + nums[k] == 0 {
//                 result = append(result, []int{nums[i], nums[j], nums[k]})
//                 for j < k && nums[j] == nums[j+1] {
//                     j++
//                 }
//                 for j < k && nums[k] == nums[k-1] {
//                     k--
//                 }
//                 j++
//                 k--
//             } else if nums[i] + nums[j] + nums[k] < 0 {
//                 for j < k && nums[j] == nums[j+1] {
//                     j++
//                 }
//                 j++
//             } else {
//                 for j < k && nums[k] == nums[k-1] {
//                     k--
//                 }
//                 k--
//         }
//     }
//     return result

// }
