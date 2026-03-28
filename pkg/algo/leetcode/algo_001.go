package leetcode

func twoSum(nums []int, target int) []int {
	// 使用哈希表存储值和对应的下标
	seen := make(map[int]int)
	
	for i, num := range nums {
		// 计算需要的补数
		complement := target - num
		
		// 检查补数是否已经在哈希表中
		if j, exists := seen[complement]; exists {
			return []int{j, i}
		}
		
		// 将当前数字和下标存入哈希表
		seen[num] = i
	}
	
	return []int{}
}