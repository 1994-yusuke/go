package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	var n = flag.Bool("n", false, "通し番号を付与する")
	flag.Parse()
	var (
		files = flag.Args()
	)
	//通し番号の用意

	for _, name := range files {
		readFile(name, *n)
	}
}

func readFile(fileName string, o bool) {
	sf, err := os.Open(fileName)
	defer sf.Close()

	if err != nil {
		fmt.Fprintln(os.Stderr, "読み込みに失敗しました", err)
	} else {
		scanner := bufio.NewScanner(sf)
		for i := 0; scanner.Scan(); {
			i++
			if o {
				//オプションがある場合
				fmt.Printf("%v: ", i)
			}
			fmt.Println(scanner.Text())
		}
	}

}
