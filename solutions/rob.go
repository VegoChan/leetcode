//leetcode题目：打家劫舍
package main



//子问题：偷到第i户的最大金额（与偷盗前i户不同），即dp[i]
//子问题转化为对第i户的决策问题，偷，只有一个选项，子问题定义决定
//偷的不同方式：
//1.均匀跳跃偷（1 3 5 7.../2 4 6...）：从i-2结尾到i,dp[i] = dp[i-2]+nums[i]
//2.非均匀跳跃偷（1 3 6...）把i-1的结尾去了，偷i,dp[i] = dp[i-1]-nums[i-1]+nums[i]
//dp[i] = max{dp[i-2]+nums[i],dp[i-1]-nums[i-1]+nums[i]}
func rob(nums []int) int{
	length := len(nums)
	if length == 0{
		return 0
	}
	if length == 1{
		return nums[0]
	}

	dp := make([]int, length)
	dp[0] = nums[0]
	dp[1] = nums[1]

	for i := 2; i < length; i++{
		dp[i] = Max(dp[i-2]+nums[i],dp[i-1]-nums[i-1]+nums[i])
	}

	max := dp[0]

	for i := 0; i < length; i++{
		if max < dp[i]{
			max = dp[i]
		}
	}

	return max
}

//子问题：前i家能偷到的最大金额，即dp[i]
//子问题转化为对第i家的决策问题，偷/不偷
//偷：dp[i] = dp[i-2]+nums[i]
//不偷：dp[i] = dp[i-1]
//dp[i]=max{dp[i-1],dp[i-2]+nums[i]}
func robII(nums []int) int{
	length := len(nums)
	if length == 0{
		return 0
	}
	if length == 1{
		return nums[0]
	}

	dp := make([]int, length)
	dp[0] = nums[0]
	dp[1] = Max(nums[0],nums[1])

	for i := 2; i < length; i++{
		dp[i] = Max(dp[i-1],dp[i-2]+nums[i])
	}

	return dp[length-1]
}

//Implementaion with constant space complexity
func robIII(nums []int) int{
	length := len(nums)
	if length == 0{
		return 0
	}
	if length == 1{
		return nums[0]
	}
	if length == 2{
		return Max(nums[0],nums[1])
	}
	var a,b,c int
	a = nums[0]
	b = Max(nums[0],nums[1])

	for i := 2; i < length; i++{
		c = Max(b,nums[i]+a)
		a = b
		b = c
	}

	return c
}
func Max(x,y int) int{
	if x > y{
		return x
	}
	return y
}
