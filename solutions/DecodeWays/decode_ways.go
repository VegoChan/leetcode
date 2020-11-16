package DecodeWays

//子问题：前i个字符的解码方法数
//1-9:num[i-1]+1
//1,0-9:num[i-2]+1
//2,0-6:num[i-2]+1
//one dimentional array -> variable(linear->constant)
func numDecodings(s string) int {
	length := len(s)
	num := make([]int,length)
	num[0] = 1
	if s[0] >= '1' && s[1] <= '9'{
		num[1] = 1
	}
	if length < 2{
		return num[0]
	}
	if s[1] >= '1' && s[1] <= '9'{
		num[1] = num[0]
	}
	if s[0] == '1' && (s[1] >= '0' && s[1] <= '9'){
		num[1]++
	}else if s[0] == '2' && (s[1] >= '0' && s[1] <= '6') {
		num[1]++
	}

	for i := 2; i < length; i++{
		if s[i] >= '1' && s[i] <= '9'{
			num[i] = num[i-1]
		}
		if s[i-1] == '1' && (s[i] >= '0' && s[i] <= '9'){
			num[i] += num[i-2]
		}else if s[i-1] == '2' && (s[i] >= '0' && s[i] <= '6') {
			num[i] += num[i-2]
		}
	}

	return num[length-1]
}

func numDecodingsWithConstantSpaceComplexity(s string) int {
	first, second, third := 0,0,0
	length := len(s)

	if s[0] >= '1' && s[0] <= '9'{
		first = 1
	}

	if length == 1{
		return first
	}

	if s[1] >= '1' && s[1] <= '9'{
		second = first
	}
	if s[0] == '1' && (s[1] >= '0' && s[1] <= '9') {
		second++
	}else if s[0] == '2' && (s[1] >= '0' && s[1] <= '6'){
		second++
	}

	if length == 2{
		return second
	}

	for i := 2; i < length; i++{
		third = 0
		if s[i] >= '1' && s[i] <= '9'{
			third = second
		}
		if s[i-1] == '1' && (s[i] >= '0' && s[i] <= '9') {
			third += first
		}else if s[i-1] == '2' && (s[i] >= '0' && s[i] <= '6'){
			third += first
		}

		first = second
		second = third
	}

	return third
}


