package parallel

import (
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

func TestParallel(t *testing.T) {
	for _, n := range []int{0, 1, 10, 100, 1000} {
		for _, p := range []int{1, 2, 4, 8, 16, 100} {
			assert.True(t, testParallelN(n, p), "test [parallel %d %d] failed", n, p)
		}
	}
}

func testParallelN(n, procs int) bool {
	data := make([]bool, n)
	before := runtime.GOMAXPROCS(0)
	runtime.GOMAXPROCS(procs)
	Parallel(0, n, procs, func(is <-chan int) {
		for i := range is {
			data[i] = true
		}
	})
	runtime.GOMAXPROCS(before)
	for i := 0; i < n; i++ {
		if !data[i] {
			return false
		}
	}
	return true
}
