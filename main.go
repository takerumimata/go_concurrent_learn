// goruituinのおべんきょ
package main

import "fmt"

func main() {
	go sayHello()
	// 他の処理を続ける
}

func sayHello() {
	fmt.Println("hello!")
}
