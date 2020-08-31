// ゴルーチン生成前後でのメモリ消費量を計測する

package main

import (
	"fmt"
	"sync"
)

func main() {
	//go sayHello()
	// 他の処理を続ける
	var wg sync.WaitGroup
	salutation := "hello!"
	wg.Add(1)
	// {}の後ろに()をつけるのは無名関数だからだよ
	go func() {
		defer wg.Done()
		salutation = "welcome" // salutationの値を変更！
	}()
	wg.Wait()
	fmt.Println("salutation")
}

/*
このプログラムを実行すると、welcomeとコマンドラインに表示される。
つまり、ゴルーチンはそれが作られたアドレス空間と同じ空間で実行されることがわかる
*/
