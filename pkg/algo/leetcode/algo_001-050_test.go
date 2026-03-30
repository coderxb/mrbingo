package leetcode

import (
	"reflect"
	"testing"
)

/* 
	TestTwoSum tests the twoSum function in algo_001.go 
	Test cases: 
		$go test -v -run TestTwoSum
*/
func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "基本用例",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "目标值在中间",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "相同元素",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "负数",
			nums:   []int{-1, -2, -3, -4, -5},
			target: -8,
			want:   []int{2, 4},
		},
		{
			name:   "含零元素",
			nums:   []int{0, 4, 3, 0},
			target: 0,
			want:   []int{0, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

/*
	TestAddTwoNumbers tests the addTwoNumbers function in algo_002.go
	Test cases:
		$go test -v -run TestAddTwoNumbers
*/
func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   []int
		l2   []int
		want []int
	}{
		{
			name: "基本用例: 342+465=807",
			l1:   []int{2, 4, 3},
			l2:   []int{5, 6, 4},
			want: []int{7, 0, 8},
		},
		{
			name: "两个零",
			l1:   []int{0},
			l2:   []int{0},
			want: []int{0},
		},
		{
			name: "不等长链表: 99+1=100",
			l1:   []int{9, 9},
			l2:   []int{1},
			want: []int{0, 0, 1},
		},
		{
			name: "全进位: 999+1=1000",
			l1:   []int{9, 9, 9},
			l2:   []int{1},
			want: []int{0, 0, 0, 1},
		},
		{
			name: "长链表相加: 9999999+9999=10009998",
			l1:   []int{9, 9, 9, 9, 9, 9, 9},
			l2:   []int{9, 9, 9, 9},
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
		{
			name: "两个空链表",
			l1:   nil,
			l2:   nil,
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := sliceToList(tt.l1)
			l2 := sliceToList(tt.l2)
			got := listToSlice(addTwoNumbers(l1, l2))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers(%v, %v) = %v, want %v", tt.l1, tt.l2, got, tt.want)
			}
		})
	}
}
