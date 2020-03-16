# Task 1

## Print

Hello World と出力するプログラムを作成してください

> 📝FileName: `hello.go`

## 変数

3 の 3 乗を計算し, 出力するプログラムを作成してください

> 📝FileName: `power.go`

## if

変数 `a` が奇数の場合は "odd", 偶数の場合は "even" と出力するプログラムを作成してください

> 📝FileName: `if.go`

```go
func main() {
	a := 1
	// Please EDIT
}
```

## switch

変数 `a` の文字数が 1 または 2 文字なら "length is 1 or 2", 3 文字なら "length is 3", それ以上なら "length is over 3" と出力するプログラムを作成してください

> 📝FileName: `switch.go`

```go
func main() {
	a := "abc"
	// Please EDIT
}
```

## for

1 から 10 までの数字を 1 行ずつ出力するプログラムを作成してください

> 📝FileName: `for.go`

## FizzBuzz

以下の要件を満たすプログラムを作成してください

* 1 から 100 までの数を出力する
* 3 の倍数の時には数の代わりに "Fizz" と出力する
* 5 の倍数の時には数の代わりに "Buzz" と出力する
* 3 と 5 両方の倍数の時には数の代わりに "FizzBuzz" と出力する

> 📝FileName: `fizzbuzz.go`

## おみくじ

以下の要件を満たすプログラムを作成してください

* 1 から 6 までの数字をランダムに出力し, その次の行に運勢を出力する
* 数字の運勢のマッピングは以下とする
  * 6: 大吉
  * 5: 中吉
  * 4: 小吉
  * 3: 吉
  * 2: 凶
  * 1: 大凶

> 📝FileName: `omikuji.go`

実行例:

```shell
$ go run omikuzi.go
6
大吉
```

> 📝 ランダムな数字の出し方
> 
> `math/rand` パッケージを使う
