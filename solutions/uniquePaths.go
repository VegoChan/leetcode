//leetcode problem:Unique Paths
package main



//到点[m,n]的所有路径数目，依赖于前面点的路径数目，有重叠子问题，选用动态规划
//子问题：到[x,y]的所有路径数目，即dp[x,y]
//子问题决策：如何到达位置[x,y]?从位置[x-1,y]下移一格/从[x,y-1]位置右移一格
//dp[x,y]=dp[x,y-1]+dp[x-1,y]

//optimize
//use one dimensional slice instead of two dimensional slice
func uniquePaths(m int, n int) int{
	dp := make([]int, n)

	for i,_ := range dp{
		dp[i] = 1
	}

	for i := 1; i < m; i++{
		for j := 1; j < n; j++{
			dp[j] += dp[j-1]
		}
	}

	return dp[n-1]
}


