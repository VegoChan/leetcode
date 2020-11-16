//leetcode题目

package main



//sumOfSubarray[i]:以num[i]结尾的最大连续数组和
//sumOfSubarray[i] = max{sumOfSubarray[i-1]+nums[i],nums[i]}
func maximumSubarray(nums []int) int{
	length := len(nums)
	sumOfSubarray := make([]int, length)
	sumOfSubarray[0] = nums[0]

	for i := 1; i < length; i++{
		if sumOfSubarray[i-1] > 0{
			sumOfSubarray[i] = sumOfSubarray[i-1]+nums[i]
		}else{
			sumOfSubarray[i] = nums[i]
		}
	}

	maxSumOfSubarray := sumOfSubarray[0]
	for i := 1; i < len(sumOfSubarray); i++{
		if maxSumOfSubarray < sumOfSubarray[i]{
			maxSumOfSubarray = sumOfSubarray[i]
		}
	}
	return maxSumOfSubarray
}

//Implementation with constant space complexity
func maximumSubarrayII(nums []int) int{
	maxSumOfSubarray := nums[0]
	j, k := nums[0], 0

	for i := 1; i < len(nums); i++{
		if j > 0{
			k = j + nums[i]
		}else{
			k = nums[i]
		}

		j = k

		if maxSumOfSubarray < k{
			maxSumOfSubarray = k
		}
	}

	return maxSumOfSubarray
}
