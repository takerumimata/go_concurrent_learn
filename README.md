# go_concurrent_learn
go言語を用いた並行処理について学ぶために書いたコードをpushするためのリポジトリ
以下のサイトを参考にいろいろ勉強したい
https://qiita.com/tenntenn/items/0e33a4959250d1a55045

# ゴルーチンgoroutin
ゴルーチンはGoのプログラムにおいて最も基本的な構成単位。全てのGoプログラムには最低一つの御ルーチンがある。それがメインゴルーチン。
単純に言えば、御ルーチンは他のコードに対し、並行に実行している関数のこと。
goキーワードを関数呼び出しの前に置くことで、簡単に起動できる。
```
func main() {
    go sayHello()
    // 他の処理を続ける
}

func sayHello() {
    fmt.Println("hello!")
}
```

無名関数でも同様に起動することが可能。
```
go func() {
    fmt.Println("hello!")
} () // 即値で実行するので()を後ろにつける
// 他の処理を続ける
```
他にも、変数に関数を代入することでも同様の処理が可能
```
sayHello := func() {
    fmt.Println("hello")
}
go sayHello()
// 他の処理を続ける
```
