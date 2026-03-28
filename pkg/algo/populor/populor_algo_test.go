package populor

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

/* 
ExecMode: go test -v -run Test_prime_solution_01 
*/
func Test_prime_solution_01(t *testing.T) {
	prime_solution_01()
}

/* 
ExecMode: go test -v -run Test_prime_solution_02 
*/
func Test_prime_solution_02(t *testing.T) {
	// 用 os.Pipe 捕获 solution_02 写入 os.Stdout 的输出
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	origStdout := os.Stdout
	os.Stdout = w

	prime_solution_02()

	w.Close()
	os.Stdout = origStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	got := buf.String()

	// 构造期望输出：2~100 以内的素数
	expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	var expBuf bytes.Buffer
	for _, p := range expected {
		fmt.Fprintf(&expBuf, "%d ", p)
	}

	if got != expBuf.String() {
		t.Errorf("solution_02 输出:\n%q\n期望:\n%q", got, expBuf.String())
	}
}

/* 
ExecMode: go test -v -run Test_factorial_recursion 
*/
func Test_factorial_recursion(t *testing.T) {
	var result = factorial_recursion(5)
	if result != 120 {
		t.Errorf("Expected 120, got %d", result)
	}
	var result2 = factorial_recursion(-1)
	if result2 != -1 {
		t.Errorf("Expected -1, got %d", result2)
	}
}

/*
ExecMode: go test -v -run Test_fibonacci_recursion
*/
func Test_fibonacci_recursion(t *testing.T) {
	var result = fibonacci_recursion(10)
	if result != 55 {
		t.Errorf("Expected 55, got %d", result)
	}
	var result2 = fibonacci_recursion(-1)
	if result2 != -1 {
		t.Errorf("Expected -1, got %d", result2)
	}
}