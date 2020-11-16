package main


func findLength(A []int, B []int) int {
	ans,k := 0,0

	for i := 0; i < len(A); i++{
		for j := 0; j < len(B); j++{
			k = 0
			for{
				if (i+k < len(A) && (j+k) < len(B))&&(A[i+k] == B[j+k]){
					k++
				}else{
					break
				}
			}
			ans = max(ans, k)
		}
	}
	return ans
}

func max(x int, y int) int{
	if x > y{
		return x
	}
	return y
}
