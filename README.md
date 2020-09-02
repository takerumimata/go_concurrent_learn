# Golangによる並行処理を学ぶ
go言語を用いた並行処理について学ぶために書いたコードをpushするためのリポジトリ  
以下のサイトを参考にいろいろ勉強したい  
https://qiita.com/tenntenn/items/0e33a4959250d1a55045

# ゴルーチンgoroutin
ゴルーチンはGoのプログラムにおいて最も基本的な構成単位。全てのGoプログラムには最低一つの御ルーチンがある。それがメインゴルーチン。
単純に言えば、御ルーチンは他のコードに対し、並行に実行している関数のこと。
goキーワードを関数呼び出しの前に置くことで、簡単に起動できる。
```go
func main() {
    go sayHello()
    // 他の処理を続ける
}

func sayHello() {
    fmt.Println("hello!")
}
```

無名関数でも同様に起動することが可能。
```go
go func() {
    fmt.Println("hello!")
} () // 即値で実行するので()を後ろにつける
// 他の処理を続ける
```
他にも、変数に関数を代入することでも同様の処理が可能
```go
sayHello := func() {
    fmt.Println("hello")
}
go sayHello()
// 他の処理を続ける
```
利用方法に関してはこれが全て。これ以降学べことは
- 正しい使用法
- 同期の方法
- まとめる方法
  
などで、使う分にはこれで使える。

# syncパッケージ
同期のためのコーディングについて学ぶ。syncパッケージには低水準のメモリアクセス同期に関する並行処理のためのプリミティブが詰まっている。

## WaitGroup
並行処理を行う上で、複数のタスクでその結果を気にしない、あるいは他に結果を収集する手段がある場合、それらの処理の合流地点として利用するのに使う。
```go
var wg sync.WaitGroup

wg.Add(1)
go func() {
    defer wg.Done()
    fmt.Println("1st goroutin sleeeping...")
    time.Sleep(1)
} ()

wg.Add(1)
go func() {
    defer wg.Done()
    fmt.Println("2nd goroutin sleeping...")
    time.Sleep(2)
} ()

wg.Wait()
fmt.Println("All goroutin complete!!")
```
基本的には、goroutineの起動前にAdd(1)、goroutineの終了前にDone()、同期待ちにWait()をかく。  
deferは上位ブロックの関数がreturnされるまで関数の実行を遅らせる。つまり上の例では、無名関数がreturnされるまで実行されない。これによって無名関数2つが並行処理を行うようになる。  
