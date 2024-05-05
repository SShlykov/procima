package parallel

import (
	"sync"
)

// Parallel обрабатывает задачи в procs несколько потоков
func Parallel(start, stop, procs int, fn func(<-chan int)) {
	count := stop - start
	if count < 1 {
		return
	}
	c := make(chan int, count)
	for i := start; i < stop; i++ {
		c <- i
	}
	close(c)

	var wg sync.WaitGroup
	for i := 0; i < procs; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fn(c)
		}()
	}
	wg.Wait()
}
