package CountSortedVewelStrings

func countVowelStrings(n int) int {
	num := [6]int{0,1,1,1,1,1}
	res := 5
	for i := 2; i <= n; i++{
		for j := 1; j < 6; j++{
			num[j] += j
			res += num[j]*(6-j)
		}
	}

	return res
}
