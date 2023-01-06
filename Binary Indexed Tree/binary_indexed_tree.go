package main

import "fmt"

type NumArray struct {
	arrayLen int
	nums     []int
	BIT      []int
}

// Time: nlog(n) | Space: O(n)
func Constructor(nums []int) NumArray {
	BITArr := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		BITIndex := i + 1
		for BITIndex <= len(nums) {
			BITArr[BITIndex] += nums[i]
			BITIndex += BITIndex & (-BITIndex)
		}
	}

	return NumArray{
		arrayLen: len(nums),
		nums:     nums,
		BIT:      BITArr,
	}
}

// Time: log(n) | Space: O(1)
func (this *NumArray) Update(index int, val int) {
	delta := val - this.nums[index]
	this.nums[index] = val
	BITIndex := index + 1
	for BITIndex <= this.arrayLen {
		this.BIT[BITIndex] += delta
		BITIndex += BITIndex & (-BITIndex)
	}
}

// Time: log(n) | Space: O(1)
func (this *NumArray) SumRange(left int, right int) int {
	return this.getRangeSum(right) - this.getRangeSum(left-1)
}

func (this *NumArray) getRangeSum(index int) int {
	rangeSum := 0
	BITIndex := index + 1
	for BITIndex > 0 {
		rangeSum += this.BIT[BITIndex]
		BITIndex -= BITIndex & (-BITIndex)
	}
	return rangeSum
}

func main() {
	//index 0   1  2  3   4   5   6  7  8  9  10  11  12  13 14 15
	nums := []int{2, -1, 8, 9, 22, -10, 1, 3, 6, 7, 11, -6, 12, -2, 1, 4}
	BIT := Constructor(nums)
	fmt.Println("sumRange(5, 13)", BIT.SumRange(5, 13))
	fmt.Println("Before update")
	fmt.Println("nums", BIT.nums)
	fmt.Println("BIT array", BIT.BIT)
	BIT.Update(5, -16)
	fmt.Println("After update")
	fmt.Println("nums", BIT.nums)
	fmt.Println("BIT array", BIT.BIT)
}

/*
output:
sumRange(5, 13) 22
Before update
nums [2 -1 8 9 22 -10 1 3 6 7 11 -6 12 -2 1 4]
BIT array [0 2 1 8 18 22 12 1 34 6 13 11 18 12 10 1 67]
After update
nums [2 -1 8 9 22 -16 1 3 6 7 11 -6 12 -2 1 4]
BIT array [0 2 1 8 18 22 6 1 28 6 13 11 18 12 10 1 61]
*/
