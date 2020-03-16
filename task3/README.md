# Task 3

## JSON

Go で JSON を扱う際には `encoding/json` パッケージを使います。

* Go の値を JSON に変換する
* 構造体などを JSON のオブジェクトや配列に変換できる

### JSON Encode

`json.Encoder` 型を使用する。

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	p := &Person{Name: "foo", Age: 31}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(p); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}
```

### JSON Decode

`json.Decoder` 型を使用する。

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	p := &Person{Name: "foo", Age: 31}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(p); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())

	var p2 Person
	dec := json.NewDecoder(&buf)
	if err := dec.Decode(&p2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(p2)
}
```

## CSV to JSON

以下の CSV ファイルを JSON に変換するプログラムを作成してください。

* [sample.csv](./sample.csv)

```csv
year,name,company
2014,Swift,Apple
2011,Kotlin,JetBrains
2011,Dart,Google
2010,Rust,Mozilla
2009,Go,Google
2000,C#,Microsoft
```

実行例:

```console
$ csv2json sample.csv
[{"year":"2014","name":"Swift","company":"Apple"},{"year":"2011","name":"Kotlin","company":"JetBrains"},{"year":"2011","name":"Dart","company":"Google"},{"year":"2010","name":"Rust","company":"Mozilla"},{"year":"2009","name":"Go","company":"Google"},{"year":"2000","name":"C#","company":"Microsoft"}]
```

整形後の JSON:

```json
[
    {
        "year": "2014",
        "name": "Swift",
        "company": "Apple"
    },
    {
        "year": "2011",
        "name": "Kotlin",
        "company": "JetBrains"
    },
    {
        "year": "2011",
        "name": "Dart",
        "company": "Google"
    },
    {
        "year": "2010",
        "name": "Rust",
        "company": "Mozilla"
    },
    {
        "year": "2009",
        "name": "Go",
        "company": "Google"
    },
    {
        "year": "2000",
        "name": "C#",
        "company": "Microsoft"
    }
]
```

> 📝File:
>
> このリポジトリに `./task3/csv2json` ディレクトリを作成し、そこにプログラムを配置してください
