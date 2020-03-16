# Task 2

## slice

3 つの変数だけを使用するように以下のプログラムを修正してください。

```go
package main

import "fmt"

func main() {
	n1 := 1
	n2 := 11
	n3 := 111
	n4 := 1111

	sum := n1 + n2 + n3 + n4
	fmt.Println(sum)
}
```

> 📝FileName: `slice.go`

## map

A, B, C の 3 人のプレイヤーがいます。
関数 `game()` はプレイヤー名と得点を返却する関数です。
`game()` を 10 回実行して、それぞれのプレイヤーごとの得点を集計し、
合計得点が最も高かったプレイヤー名とその得点を出力してください。

```go
package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// TODO
}

func game() (user string, score int) {
	switch a := rand.Intn(3); {
	case a == 0:
		return "A", 10
	case a == 1:
		return "B", 10
	default:
		return "C", 10
	}
}
```

> 📝FileName: `map.go`

```shell
$ go run map.go
user: B
score: 60
```

## method

`Inc()` メソッドは自身に 1 を加算するメソッドです。
ですが、下のプログラムだと正しく動作しません。
原因を考え正しく動作する用に修正を行ってください。

```go
package main

import "fmt"

type Int int

func (n Int) Inc() {
	n++
}

func main() {
	var n Int
	fmt.Println(n)
	n.Inc()
	fmt.Println(n)
}
```

> 📝FileName: `method.go`

期待する実行例:

```shell
$ go run method.go
0
1
```

## command-line tool

以下の要件を満たす cat コマンドを作成してください。

* Linux コマンドの cat と同様にファイルの中身を表示する
* `-n` オプションでファイルの行数を表示する
* 複数のファイルを表示できるようにする
  * 例: cat [file1] [file2] ...
* ファイルが存在しない場合はエラーメッセージを表示する

実行例 (必ずしも例と同じように表示する必要はありません):

```shell
$ go run cat.go hello.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
}
```

```shell
$ go run cat.go -n hello.go
 1 package main
 2 
 3 import (
 4 	"fmt"
 5 )
 6 
 7 func main() {
 8 	fmt.Println("Hello, playground")
 9 }
10 
```

```shell
$ go run cat.go -n hello1.go hello2.go
=== hello1.go ===
 1 package main
 2 
 3 import (
 4 	"fmt"
 5 )
 6 
 7 func main() {
 8 	fmt.Println("Hello, 1")
 9 }
10 
=== hello2.go ===
 1 package main
 2 
 3 import (
 4 	"fmt"
 5 )
 6 
 7 func main() {
 8 	fmt.Println("Hello, 2")
 9 }
10 
```

> 📝File:
>
> このリポジトリに `./task2/cat` ディレクトリを作成し、そこにプログラムを配置してください
