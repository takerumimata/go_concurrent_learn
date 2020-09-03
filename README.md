# Golangによる並行処理を学ぶ
go言語を用いた並行処理について学ぶために書いたコードをpushするためのリポジトリ  
以下のサイトを参考にいろいろ勉強したい  
https://qiita.com/tenntenn/items/0e33a4959250d1a55045

# ゴルーチンgoroutin
ゴルーチンはGoのプログラムにおいて最も基本的な構成単位。全てのGoプログラムには最低一つのゴルーチンがある。それがメインゴルーチン。
単純に言えば、ゴルーチンは他のコードに対し、並行に実行している関数のこと。
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
同期のためのコーディングについて学ぶ。syncパッケージには低水準のメモリアクセス同期に関する並行処理のためのプリミティブが詰まっている。重要なのは以下
- WaitGroup
- Mutex(RWMutex)
- Cond
- Once
- Pool

## WaitGroup: 合流地点を作る
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

## Mutex(RWMutex): データをロックする
クリティカルセクションを作ると聞いて、ピンと来るだろうか。OSが排他制御を行う際に利用している仕組みとしてMutexやセマフォがある。平たく言うと、ある変数（あるいは任意のデータ）へのアクセス権を占有することで、他の処理が今扱っているデータを改変してしまわない（読み込むことも禁止することもできる）ことを保証する仕組みである。データを利用する際にデータをlockし、処理が終わればunlockする。基本的な概念はこれだけである。  
A tour of Goよりコードを拝借する。
```go
// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
```
構造体をまず定義し、mutexの機構と連想配列を用意している。goroutinとして走らせるインクリメントだけを行う単純な関数Incと、配列に格納されている値を参照し返すValue関数を定義し、それぞれ処理を走らせる時にmutexを利用しメモリへのアクセスを制限している。  
以下のサイトで実際に実行してみると、valueとして1000が返ってくることがわかる。  
https://go-tour-jp.appspot.com/concurrency/9
- - - 
# チャネル(channel)
チャネルを使うときは値をchan型の変数に渡し、プログラムの別の場所からその値を読み込む。
