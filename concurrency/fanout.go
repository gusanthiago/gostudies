package main

import(
    "math/rand"
    "time"
    "fmt"
)

func sleep() {
    time.Sleep(
        time.Duration(rand.Intn(300)) * time.Millisecond,
    )
}

func producer(ch chan<- int) {
    
    // var n int
    for {
        sleep()
        
        n := rand.Intn(100)

        fmt.Printf("-> %d\n", n)
        ch <- n
    }
}

func consumer(ch <- chan int, name string) {
    for n := range ch {
        fmt.Printf("Consumer %s <- %d\n", name, n)
    }
}

func fanOut(channelA <- chan int, channelB, channelC chan <- int) {
    for n := range channelA {
        if n < 50 {
            channelB <- n
        } else {
            channelC <- n
        }
    }
}


func main() {
    channelA := make(chan int)
    channelB := make(chan int)
    channelC := make(chan int)
    
    go producer(channelA)
    go consumer(channelB, "B")
    go consumer(channelC, "C")

    fanOut(channelA, channelB, channelC)
}