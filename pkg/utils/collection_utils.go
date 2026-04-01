package utils

import "math/rand"

// BuildRandomArray 随机生成一个数组并返回, 数组大小为size, 数值范围在min和max之间, 默认10个元素
func BuildRandomArray(size int, min int, max int) []int {
	if size <= 0 {
		size = 10
	}
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = min + rand.Intn(max-min+1)
	}
	return arr
}

// BuildRandomSortedArray 随机生成一个有序的数组并返回, 数组大小为size, 数值范围在min和max之间, 默认10个元素
func BuildRandomSortedArray(size int, min int, max int) []int {
	arr := BuildRandomArray(size, min, max)
	// 对数组进行排序
	for i := 0; i < size-1; i++ {
		for j := 0; j < size-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// PrintAny 打印任意类型的值和类型
func PrintAny[T any](value T)  {
	fmt.Printf("Value: %v, Type: %T\n", value, value)
}