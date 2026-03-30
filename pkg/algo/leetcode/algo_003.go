/**
给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。
示例 1:
输入: s = "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。注意 "bca" 和 "cab" 也是正确答案。
示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

提示：
0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成
**/

func lengthOfLongestSubstring(s string) int {
	// 使用哈希表存储字符和对应的下标
	charSet := make(map[byte]int)
	maxLength := 0
	left`` := 0

	for right := 0; right < len(s); right++ {
		char := s[right]
		// 如果字符已经存在，并且在当前窗口内，则更新窗口的起始位置
		if index, exists := charSet[char]; exists && index >= left {
			left = index + 1
		}
		// 更新字符的最新下标
		charSet[char] = right
		// 更新最大长度
		if right - left + 1 > maxLength {
			maxLength = right - left + 1
		}
	}
	return maxLength
}