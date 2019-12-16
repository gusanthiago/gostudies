package main

import(
    "fmt"
    "math/rand"
    "time"
)

func processTest(ch chan int) {
    n := rand.Intn(3000)
    time.Sleep(time.Duration(n) * time.Millisecond)
    ch <- n
}

func main() {
    ch := make(chan int)
    go processTest(ch)

    fmt.Println("waiting for processor")
    n := <- ch
    fmt.Printf("Process took %dms\n", n)
}