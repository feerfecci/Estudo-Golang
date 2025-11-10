package main

import (
	"fmt"
	"sync"
	"time"
)

//SYNC
// func main() {
// 	var wg sync.WaitGroup

// 	wg.Add(3)

// 	go callDatabase(&wg)
// 	go callAPI(&wg)
// 	go processInternal(&wg)

// 	// time.Sleep(2 * time.Second)
// 	wg.Wait()
// }

// func callDatabase(wg *sync.WaitGroup) {
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("finalizado data")
// 	wg.Done()
// }

// func callAPI(wg *sync.WaitGroup) {
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("finalizado api")
// 	wg.Done()
// }

// func processInternal(wg *sync.WaitGroup) {
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("finalizado internal")
// 	wg.Done()
// }

// MUTEX

func main() {
	var m sync.Mutex
	i := 0
	for x := 0; x < 10000; x++ {
		// go ChangeNumber(&i, x)
		go func() {
			m.Lock()
			i++
			m.Unlock()
		}()
	}
	// go ChangeNumber(&i, 5)
	// go ChangeNumber(&i, 10)
	// go ChangeNumber(&i, 20)
	// wg.Wait()
	time.Sleep(time.Second * 2)
	fmt.Println(i)

}

func ChangeNumber(i *int, newNumber int) {
	*i = newNumber
}
