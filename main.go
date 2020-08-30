// goruituinのおべんきょ
package main

import (
	"fmt"
	"sync"
)

func main() {
	//go sayHello()
	// 他の処理を続ける
	var wg sync.WaitGroup
	sayHello := func() {
		defer wg.Done()
		fmt.Println("hello!!")
	}

	wg.Add(1)
	go sayHello()
	wg.Wait()
}

