package Longest_Palindromic_Substring

import "fmt"

//abbbd
//babad
func LongestPalindrome(s string) string{
	length := len(s)
	longestPalindromeLeft,longestPalindromeRight := length-1,length-1
	for i := 0; i <= length-2; i++{
		for j := 0; j <= length/2; j++{
			if (i-j >= 0 && i+j <= length-1) && (isPalindrome(i-j,i+j,s)) {
				if 2*j+1 > longestPalindromeRight-longestPalindromeLeft+1 {
					longestPalindromeLeft = i-j
					longestPalindromeRight = i+j
				}
			}else{
				break
			}
			//
			//The following code will cause logical error:
			//If current string is palindrome but it's short than the longest,the loop will break but not go on search.
			//
			//if (i-j >= 0 && i+j <= length-1) && (isPalindrome(i-j,i+j,s)) && (2*j+1 >= longestPalindromeRight-longestPalindromeLeft+1){
			//		longestPalindromeLeft = i-j
			//		longestPalindromeRight = i+j
			//}else{
			//	break
			//}
		}
		for j := 0; j <= length/2; j++{
			if (i-j >= 0 && i+j+1 <= length-1) && (isPalindrome(i-j,i+j+1,s)) {
				if 2*j+2 > longestPalindromeRight-longestPalindromeLeft+1 {
					longestPalindromeLeft = i-j
					longestPalindromeRight = i+j+1
				}
			}else{
				break
			}
			//
			//The following code will cause logical error:
			//If current string is palindrome but it's short than the longest,the loop will break but not go on search.
			//
			//if (i-j >= 0 && i+j+1 <= length-1) && (isPalindrome(i-j,i+j+1,s)) && (2*j+2 >= longestPalindromeRight-longestPalindromeLeft+1) {
			//		longestPalindromeLeft = i-j
			//		longestPalindromeRight = i+j+1
			//}else{
			//	break
			//}
		}
	}
	return s[longestPalindromeLeft:longestPalindromeRight+1]
}

func isPalindrome(left int, right int,s string) bool{
	for left < right{
		if s[left] != s[right]{
			return false
		}else{
			left++
			right--
		}
	}
	return true
}

//子问题：子串s[i:j]是一个palindrome，s[i:j+1]是不是也是palindrome
//子问题解决依赖于s[i+1:j]是不是palindrome并且s[i]==s[j+1],即 s[i+1:j] == true && s[i]==s[j+1]
//基础状态：s[i][i] = true, s[i][i+1] = s[i]==s[i+1]
func LongestPalindromeDynamicProgramming(s string) string{
	length := len(s)

	p := make([][]bool,length)
	for i,_ := range p{
		p[i] = make([]bool,length)
	}

	for i := 0; i < length; i++{
		p[i][i] = true
		if i+1 < length{
			if s[i] == s[i+1]{
				p[i][i+1] = true
			}else{
				p[i][i+1] = false
			}
		}
	}

	for i := 2; i < length; i++{
		k := i
		for j := 0; j < length-i; j++{
			//if k < length{
				if s[j] == s[k] && p[j+1][k-1] == true{
					p[j][k] = true
				}else{
					p[j][k] = false
				}
				k++
			//}else{
			//	break
			//}
		}
	}
	//for i := 0; i < length; i++{
	//	for j := i+2; j < length; j++{
	//		if s[i] == s[j] && p[i+1][j-1] == true{
	//			p[i][j] = true
	//		}else{
	//			p[i][j] = false
	//		}
	//	}
	//}
	fmt.Println(p)

	maxLen, right, left := 1, 0, 0
	for i := 0; i < length; i++{
		for j := i; j < length; j++{
			if p[i][j] == true && maxLen < j-i+1{
				maxLen = j-i+1
				right = j
				left = i
			}
		}
	}
	fmt.Println(left,right)
	return s[left:right+1]
}