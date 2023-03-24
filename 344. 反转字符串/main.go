package main

import "fmt"

//https://leetcode.cn/problems/reverse-string/
func main() {
	str := "hello"
	bstr := []byte(str)
	reverseString(bstr)
	fmt.Println(string(bstr))
}

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
