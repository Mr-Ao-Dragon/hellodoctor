package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"

	"github.com/Mr-Ao-Dragon/hellodoctor/tool/gen"
)

func GenAndPrint() error {
	// preGenTime := strconv.FormatInt(time.Now().UnixNano(), 10)
	token, err := gen.Token(16)
	// postGenTime := strconv.FormatInt(time.Now().UnixNano(), 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
	return err
}

// func main() {
//	//go func() {
//	//	err := GenAndPrint(16)
//	//	if err != nil {
//	//		panic(err)
//	//	}
//	//}()
//	err := GenAndPrint(9)
//	if err != nil {
//		panic(err)
//	}
// }

var sum = 0

func main() {
	// GenNum := int64(0)
	defer ants.Release()
	runTimes := 4096
	// Use the common pool.
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		err := GenAndPrint()
		if err != nil {
			return
		}
		wg.Done()
	}
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = ants.Submit(syncCalculateSum)
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")

	// Use the pool with a function,
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	p, _ := ants.NewPoolWithFunc(16, func(i interface{}) {
		err := GenAndPrint()
		// GenNum++
		// fmt.Println("GenNum: ", GenNum)
		if err != nil {
			return
		}
		wg.Done()
	})
	defer p.Release()
	// Submit tasks one by one.
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)

	// Use the MultiPool and set the capacity of the 10 goroutine pools to unlimited.
	// If you use -1 as the pool size parameter, the size will be unlimited.
	// There are two load-balancing algorithms for pools: ants.RoundRobin and ants.LeastTasks.
	mp, _ := ants.NewMultiPool(10, -1, ants.RoundRobin)
	defer func(mp *ants.MultiPool, timeout time.Duration) {
		err := mp.ReleaseTimeout(timeout)
		if err != nil {
		}
	}(mp, 5*time.Second)
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = mp.Submit(syncCalculateSum)
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", mp.Running())
	fmt.Printf("finish all tasks.\n")

	// Use the MultiPoolFunc and set the capacity of 10 goroutine pools to (runTimes/10).
	mpf, _ := ants.NewMultiPoolWithFunc(10, runTimes/10, func(i interface{}) {
		// myFunc(i)
		err := GenAndPrint()
		if err != nil {
			return
		}
		wg.Done()
	}, ants.LeastTasks)
	defer func(mpf *ants.MultiPoolWithFunc, timeout time.Duration) {
		err := mpf.ReleaseTimeout(timeout)
		if err != nil {
			return
		}
	}(mpf, 5*time.Second)
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = mpf.Invoke(int32(i))
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", mpf.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)
}
