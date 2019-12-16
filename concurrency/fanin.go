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

func producer(ch chan<- int, name string) {
    for {
        sleep()
        
        n := rand.Intn(100)

        fmt.Printf("Channel %s -> %d\n", name, n)
        ch <- n
    }
}


func consumer(ch <- chan int) {
    for n := range ch {
        fmt.Printf("<- %d\n", n)
    }
}

func fanIn(channelA, channelB <- chan int, channelC chan <- int) {
    var n int
    for {
        select {
        case n = <- channelA:
            channelC <- n
        case n = <- channelB:
            channelC <- n
        }
    }
}


func main() {
    channelA := make(chan int)
    channelB := make(chan int)
    channelC := make(chan int)
    
    go producer(channelA, "A")
    go producer(channelB, "B")
    go consumer(channelC)

    fanIn(channelA, channelB, channelC)
}