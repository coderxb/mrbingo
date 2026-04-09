package populor

import (
	"testing"
)

/*
ExecMode: go test -v -run Test_selection_sort 
*/
func Test_selection_sort(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "随机数组",
			nums:     []int{64, 25, 12, 22, 11},
			expected: []int{11, 12, 22, 25, 64},
		},
		{
			name:     "已经排序的数组",
			nums:     []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "逆序数组",	
			nums:     []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "包含重复元素的数组",
			nums:     []int{3, 1, 2, 3, 1},
			expected: []int{1, 1, 2, 3, 3},
		},
		{
			name:     "单元素数组",
			nums:     []int{42},
			expected: []int{42},
		},
		{
			name:     "空数组",
			nums:     []int{},
			expected: []int{},
		},
	}	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			selection_sort(tt.nums)
			for i := range tt.nums {
				if tt.nums[i] != tt.expected[i] {
					t.Errorf("排序失败: got %v, want %v", tt.nums, tt.expected)
					break
				}
			}
		})
	}
}