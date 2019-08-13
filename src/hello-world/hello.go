package main

import "fmt"
import "greet"

func main() {
    fmt.Printf("hello, world 4\n")
    fmt.Printf(greet.Morning)
    a,b := greet.GetGreeings()

    fmt.Printf("1:%s 2:%s", a, b)
}
