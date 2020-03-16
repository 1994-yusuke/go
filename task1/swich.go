package main

import (
    "fmt"
)

func main() {
    a := "abc"
    switch{
    case len(a)== 1 || len(a) == 2:
        fmt.Println("length is 1 or 2")
    case len(a) == 3:
        fmt.Println("length is 3")
    case len(a) > 3:
        fmt.Println("length is over 3")

    }

}