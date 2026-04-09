package populor

/**
排序算法
- 1.选择排序
- 2.插入排序
- 3.冒泡排序
- 4.快速排序
- 5.归并排序
- 6.堆排序
- 7.希尔排序
- 8.计数排序
- 9.桶排序
- 10.基数排序
**/


/** 
	选择排序 
	- 时间复杂度为 𝑂(𝑛^2)、非自适应排序
	- 空间复杂度为 𝑂(1)、原地排序 
	- 稳定性：不稳定排序
**/
func selection_sort(nums []int) {
	size := len(nums)
	// 外层循环控制排序趟数, 内层循环控制每趟比较次数
	for i :=0; i < size; i++ {
		// 寻找最小元素的索引
		minIndex := i
		// 内层循环寻找最小元素
		for j := i + 1; j < size; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		// 交换值
		if minIndex != i {
			nums[i], nums[minIndex] = nums[minIndex], nums[i]
		}
	}
}

/**
	插入排序 
	- 时间复杂度为 𝑂(𝑛^2)、自适应排序
	- 空间复杂度为 𝑂(1)、原地排序 
	- 稳定性：稳定排序
**/
func insertion_sort(nums []int) {
	size := len(nums)
	// 外层循环控制排序趟数, 内层循环控制每趟比较次数
	for i := 1; i < size; i++ {
		// 选择当前元素作为key
		key := nums[i]
		// 内层循环将key与前面已经排序的元素进行比较，并将比key大的元素向后移动
		j := i - 1
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}
}
