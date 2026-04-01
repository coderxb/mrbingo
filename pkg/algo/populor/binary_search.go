/**
二分查找, 未找到返回-1
说明：
	前提条件：数组必须是有序的（升序或降序）
	时间复杂度：O(log n)
	空间复杂度：O(1)
**/
package populor


// 二分查找, 未找到返回-1
func binarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		// 计算中间位置, 避免溢出
		mid := low + (high - low) / 2
		// 比较中间元素与目标值
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}