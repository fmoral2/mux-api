package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"
)

func memoryUsage() uint64 {
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	return m.Alloc
}

func main() {
	memBefore := memoryUsage()

	// http.ListenAndServe("localhost:6060", nil)
	// f, err := os.Create("traceFile.out")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// var counter int
	// var lock sync.Mutex
	// wg := &sync.WaitGroup{}
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		lock.Lock()
	// 		defer lock.Unlock()
	// 		fmt.Println("counter go routine", counter)
	// 		counter++
	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()
	// fmt.Println(counter)

	testCases := []int{0, 2, 5, 7, 9, 12}
	var wg sync.WaitGroup
	ch := make(chan int)
	// var mutex sync.Mutex

	for i, v := range testCases {
		wg.Add(1)
		fmt.Printf("TESTCASE %v \n Index: %v\n", v, i)

		go func(i, v int) {
			defer wg.Done()

			startTime := time.Now()
			fmt.Printf("\nStart of GO ROUTINE ˜˜˜˜˜˜˜ running with testcase: %v \nat time: %v \n with addres: \n %v", v, startTime, &v)

			if v&1 != 0 {
				fmt.Printf("\nResult: \n Odd number: %d\n", v)
			} else {
				ch <- v
				fmt.Printf("\nResult: \n Even number: %d\n", v)
			}
			endTime := time.Now()
			fmt.Printf("\nEnd of GO ROUTINE running with testcase: %v \nat time: %v\n", v, endTime)
			fmt.Printf("\nGO ROUTINE execution time: %v\n", endTime.Sub(startTime))
		}(i, v)

		go func() {
			time.Sleep(1 * time.Second)
			for i := range ch {
				fmt.Println("\nvalue on channel", i)
			}
		}()

	}

	// NETWORK
	go func() {
		wg.Add(1)
		defer wg.Done()

		startTime := time.Now()
		fmt.Println("\nStart of network call GO ROUTINE")

		resp, err := http.Get("https://google.com")
		if err != nil {
			fmt.Println("Error making network call:", err)
			return
		}
		defer resp.Body.Close()

		reader := bufio.NewReader(resp.Body)
		_, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading first line of response body:", err)
			return
		}

		// fmt.Println("First line of response from Google:\n", firstLine)

		endTime := time.Now()
		fmt.Println("\nEnd of network call GO ROUTINE")
		fmt.Printf("\nGO ROUTINE for network call  execution time: %v\n", endTime.Sub(startTime))
	}()

	// os call
	go func() {
		wg.Add(1)
		defer wg.Done()

		startTime := time.Now()
		fmt.Println("\nStart of OS call simulation GO ROUTINE")

		res, err := os.ReadFile("report.json")
		if err != nil {
			fmt.Println("Error reading file:", err)
		}
		fmt.Printf("\n result from report file read:\n %v\n", string(res))

		endTime := time.Now()
		fmt.Println("\nEnd of OS call simulation GO ROUTINE")
		fmt.Printf("\nGO ROUTINE for os call execution time: %v\n", endTime.Sub(startTime))
	}()

	wg.Wait()
	close(ch)

	memAfter := memoryUsage()
	fmt.Printf("Memory usage before function: %v bytes\n", memBefore)
	fmt.Printf("Memory usage after function: %v bytes\n", memAfter)
	fmt.Printf("Memory used by function: %v bytes\n", memAfter-memBefore)

}
