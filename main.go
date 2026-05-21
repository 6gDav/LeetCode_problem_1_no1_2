package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	inputs := []struct {
		nameofbenchmark string
		nums            []int
		target          int
	}{
		{
			"[Small Input Test]",
			[]int{2, 7, 11, 15},
			9,
		},
		{
			"[Stress Test - 10 000]",
			func() []int {
				res := make([]int, 0, 10_000)
				for i := 1; i <= 9_999; i++ {
					res = append(res, i)
				}
				res = append(res, 1)
				return res
			}(),
			2,
		},
	}

	fmt.Println("=== Benchmarking ===")

	for _, test := range inputs {
		var m1, m2 runtime.MemStats
		runtime.GC()
		runtime.ReadMemStats(&m1)
		start := time.Now()

		_ = twoSum(test.nums, test.target)

		elapsed := time.Since(start)
		runtime.ReadMemStats(&m2)

		fmt.Printf("\n%s\n", test.nameofbenchmark)

		fmt.Printf("Execution Time: %v\n", elapsed)
		fmt.Printf("Memory Delta: %v bytes\n", m2.Alloc-m1.Alloc)
		fmt.Printf("Current Memory: %v bytes\n", m2.Sys)
	}
}
