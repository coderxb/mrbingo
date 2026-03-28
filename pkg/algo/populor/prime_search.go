package populor

import "fmt"

/**
判断一个数是否是素数（只有被1和自身整除的数）
**/
func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

/** 输出2~100以内的素数 - 方案一 **/
func prime_solution_01()  {
	for i := 2; i <= 100; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
}

/** 输出2~100以内的素数 - 方案二 **/
func prime_solution_02()  {
	// 从2开始，依次判断每个数是否是素数
	for i := 2; i <= 100; i++ {
		// 判断i是否是素数
		prime := true
		for j := 2; j < i/j; j++ {
			// 如果i被j整除过了，说明i不是素数
			if i%j == 0 {
				prime = false
				break
			}
			// 如果j已经大于i/j了，说明i没有被2~i/j之间的数整除过了，说明i是素数
			if prime {
				fmt.Printf("%d ", i)
			}
		}
	}
}