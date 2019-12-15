package main

type Queue struct {
    items[] int
}

func (q *Queue) Push(number int) {
    q.items = append(q.items, number)
}

func (q *Queue) Dequeue() int {
    if len(q.items) == 0 {
        return -1
    }
    item, items := q.items[0], q.items[1:]
    q.items = items
    return item
}



func main() {

    q := Queue{}
    for i := 0; i < 10000; i++ {
        q.Push(i)
    }


    println(q.Dequeue())
    q.Push(10)
    println(q.Dequeue())
    println(q.Dequeue())
    println(q.Dequeue())
}
