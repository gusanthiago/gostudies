package main

import (
    "fmt"
    "time"
    "math/rand"
)

func echoWorker(in, out chan int) {
    for {
        n := <- in
     
        time.Sleep(
            time.Duration(rand.Intn(3000)) * time.Millisecond,
        )

        out <- n
    }
}

func producer(ch chan <- int) {
    i := 0
    for {
        fmt.Printf("-> Send job: %d\n", i)
        ch <- i
        i++
    }
}


func main() {

    in := make(chan int)
    out := make(chan int)

    for i := 0; i < 2; i++ {
       go echoWorker(in, out)
    }
    go producer(in)    
    
    for n := range out {
        fmt.Printf("<- Recv job: %d\n", n)
    }


}