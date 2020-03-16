package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    for i := 1; i <= 3; i++ {
        a := rand.Intn(6)
        if a == 1{
        fmt.Println(a)
        fmt.Println("大凶")
        break
        }else if a == 2{
        fmt.Println(a)
        fmt.Println("凶")
        break
        }else if a == 3 {
        fmt.Println(a)
        fmt.Println("吉")
        break
        }else if a == 4 {
        fmt.Println(a)
        fmt.Println("小吉")
        break
        }else if a == 5{
        fmt.Println(a)
        fmt.Println("中吉")
        break
        }else if a == 6{
        fmt.Println(a)
        fmt.Println("大吉")
        break
        }else{
        continue
        }
    }
}