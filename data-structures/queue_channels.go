package main

type Queue struct {
    items chan int
}

func (q *Queue) Push(number int) {
    q.items <- number
}

func (q *Queue) Dequeue() int {
    return <- q.items 
}

func main() {

    q := Queue{items: make(chan int, 10010)}
    for i := 0; i < 10000; i++ {
        q.Push(i)
    }

    println(q.Dequeue())
    q.Push(10)
    println(q.Dequeue())
    println(q.Dequeue())
    println(q.Dequeue())

}
