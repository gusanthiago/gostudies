package main

import "fmt"

type List struct {
    head *Node
    tail *Node
}

func (l *List) First() *Node {
    return l.head
}

func (l *List) Last() *Node {
    return l.tail
}

func (l *List) Push(value int) {
    node := &Node{value: value}
    if l.head == nil {
        l.head = node
    } else {
        l.tail.next = node
        node.prev = l.tail
    }
    l.tail = node
}

func (n *Node) Next() *Node {
    return n.next
}

func (n *Node) Prev() *Node {
    return n.prev
}

type Node struct {
    value int
    next *Node
    prev *Node
}

func main() {
    l := &List{}
    l.Push(1)
    l.Push(2)
    l.Push(3)
    n := l.First()

    for n.Next() != nil {
        fmt.Println(n.value)
        n = n.Next()  
    }
    fmt.Println(n.value)


    for n.Prev() != nil {
        fmt.Println(n.value)
        n = n.Prev()  
    }
    fmt.Println(n.value)
  
}