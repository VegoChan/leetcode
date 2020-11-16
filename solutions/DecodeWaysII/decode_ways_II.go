package DecodeWaysII

const divisor = 1000000007
func NumDecodings(s string) int{
	length := len(s)
	num := make([]int,length)

	if s[0] >= '1' && s[0] <= '9'{
		num[0] = 1
	}else if s[0] == '*'{
		num[0] = 9
	}

	if length == 1{
		return num[0]
	}

	if s[1] >= '1' && s[1] <= '9'{
		num[1] = num[0]
	}else if s[1] == '*' {
		num[1] = 9*num[0]
	}

	if s[0] == '1'{
		if s[1] >= '0' && s[1] <= '9'{
			num[1] += num[0]
		}else if s[1] == '*'{
			num[1] += 9*num[0]
		}
	}else if s[0] == '2'{
		if s[1] >= '0' && s[1] <= '6'{
			num[1] += num[0]
		}else if s[1] == '*'{
			num[1] += 6*num[0]
		}
	}else if s[0] == '*'{
		if s[1] >= '0' && s[1] <= '6'{
			num[1] += 2
		}else if s[1] >= '7' && s[1] <= '9'{
			num[1] += 1
		}else if s[1] == '*'{
			num[1] += 15
		}
	}

	for i := 2; i < length; i++{
		if s[i] >= '1' && s[i] <= '9'{
			num[i] = num[i-1]%divisor
		}else if s[i] == '*' {
			num[i] = 9*num[i-1]%divisor
		}

		if s[i-1] == '1'{
			if s[i] >= '0' && s[i] <= '9'{
				num[i] = (num[i]+num[i-2])%divisor
			}else if s[i] == '*'{
				num[i] = (num[i]+9*num[i-2])%divisor
			}
		}else if s[i-1] == '2'{
			if s[i] >= '0' && s[i] <= '6'{
				num[i] = (num[i]+num[i-2])%divisor
			}else if s[i] == '*'{
				num[i] = (num[i]+6*num[i-2])%divisor
			}
		}else if s[i-1] == '*'{
			if s[i] >= '0' && s[i] <= '6'{
				num[i] = (num[i]+2*num[i-2])%divisor
			}else if s[i] >= '7' && s[i] <= '9'{
				num[i] = (num[i]+num[i-2])%divisor
			}else if s[i] == '*'{
				num[i] = (num[i]+15*num[i-2])%divisor
			}
		}
	}

	return num[length-1]

}

//问题：解码字符串，有多少种解法
//子问题：解码子串s[i]有多少种解法num[i]
//关于第i个字符解码问题：单独解码；和前一个字符s[i-1]一起解码
//num[i]=num{单独解码}+num{组合解码}

func NumDecodingsWithConstantSpaceComplexity(s string) int{
	first,second,third := 0,0,0
	length := len(s)

	//initialization first,second
	if s[0] == '*'{
		first = 9
	}else if s[0] >= '1' && s[0] <= '9'{
		first = 1
	}

	if length == 1{
		return first
	}

	if s[1] == '*'{
		second = first*9
	}else if s[1] >= '1' && s[1] <= '9'{
		second = first
	}

	if s[0] == '1'{
		if s[1] >= '0' && s[1] <= '9'{
			second += 1
		}else if s[1] == '*'{
			second += 9
		}
	}else if s[0] == '2'{
		if s[1] >= '0' && s[1] <= '6'{
			second += 1
		}else if s[1] == '*'{
			second += 6
		}
	}else if s[0] == '*'{
		if s[1] >= '0' && s[1] <= '6'{
			second += 2
		}else if s[1] >= '7' && s[1] <= '9'{
			second += 1
		}else if s[1] == '*'{
			second += 15
		}
	}

	if length == 2{
		return second
	}

	for i := 2; i < length; i++{
		third = 0

		if s[i] >= '1' && s[i] <= '9'{
			third = second
		}else if s[i] == '*'{
			third = 9*second
		}

		if s[i-1] == '1'{
			if s[i] >= '0' && s[i] <= '9'{
				third = (first + third)%divisor
			}else if s[i] == '*'{
				third = (third + 9*first)%divisor
			}
		}else if s[i-1] == '2'{
			if s[i] >= '0' && s[i] <= '6'{
				third = (first + third)%divisor
			}else if s[i] == '*'{
				third = (third + 6*first)%divisor
			}
		}else if s[i-1] == '*'{
			if s[i] >= '0' && s[i] <= '6'{
				third = (2*first + third)%divisor
			}else if s[i] >= '7' && s[i] <= '9'{
				third = (third + first)%divisor
			}else if s[i] == '*'{
				third = (third + (9+6)*first)%divisor
			}
		}

		first = second
		second = third
	}

	return third
}
