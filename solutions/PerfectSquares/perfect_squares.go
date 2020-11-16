package PerfectSquares

import "math"

//子问题：数字i的最少平方数有多少个num[i]
//num[i] = min{num[k]+num[i-k]}, 0 < k <= i
func NumSquares(n int) int {
	num := make([]int, n+1)
	num[1] = 1

	for i := 2; i <= n; i++{
		if int(math.Sqrt(float64(i)))*int(math.Sqrt(float64(i))) == i{
			num[i] = 1
		}else{
			num[i] = num[i-1]+num[1]
			for j := 2; j <= i/2; j++{
				num[i] = min(num[i],num[j]+num[i-j])
			}
		}
	}
	return num[n]
}

func min(x int, y int) int{
	if x > y{
		return y
	}
	return x
}
