package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalindromeNaive(s string) bool {
	lower := strings.ToLower(s)
	b := &strings.Builder{}
	for _, ch := range lower {
		if isAlnum(ch) {
			b.WriteRune(ch)
		}
	}
	finalStr := b.String()
	rev := Reverse(finalStr)
	return finalStr == rev
}

func isAlnum(r rune) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') || ('0' <= r && r <= '9')
}

func isPalindrome(s string) bool {
	leftIndex, rightIndex := 0, len(s)-1
	for leftIndex < rightIndex {
		leftChar := unicode.ToLower(rune(s[leftIndex]))
		rightChar := unicode.ToLower(rune(s[rightIndex]))
		switch {
		case leftChar == rightChar:
			leftIndex++
			rightIndex--
		case !isAlnum(leftChar):
			leftIndex++
		case !isAlnum(rightChar):
			rightIndex--
		default:
			if leftChar != rightChar {
				return false
			}
		}
	}
	return true
}

func main() {
	// fmt.Println(isPalindrome("A man, a plan, a canal: Panama") == true)
	// fmt.Println(isPalindrome("race a car") == false)
	fmt.Println(isPalindromeNaive("0P") == true)
}
