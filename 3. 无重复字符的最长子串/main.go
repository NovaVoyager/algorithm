package main

//https://leetcode.cn/problems/longest-substring-without-repeating-characters/
func main() {

}

func lengthOfLongestSubstring(s string) int {
	left, right, length := 0, 0, 0
	m := make(map[uint8]int, 0)
	for right < len(s) {
		c, ok := m[s[right]]
		if ok && c >= left { //出现重复字符串
			//移动左边下标到重复字符位置的后一位，让left、right之间没有重复字符
			left = c + 1
		}
		m[s[right]] = right
		right++
		length = max(length, right-left)
	}

	return length
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
