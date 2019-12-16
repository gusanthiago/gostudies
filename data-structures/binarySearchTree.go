package main

import "math"

type Node struct {
    value int
    left *Node
    right *Node
}

func (n *Node) Insert(value int) *Node {
    if n == nil {
        return &Node{value: value}
    }
    if value == n.value {
        return n
    }
    if value < n.value {
        if n.left == nil {
            n.left = &Node{value: value}
        } else {
            n.left.Insert(value)
        }
    } else {
        if n.right == nil {
            n.right = &Node{value: value}
        } else {
            n.right.Insert(value)
        }
    }
    return n
}

func (root *Node) Height() int {
    if root == nil {
        return -1
    }
    return int(
        math.Max(
            float64(root.left.Height()),
            float64(root.right.Height())) + 1)
}

func (n *Node) IsExist(value int) bool {
    if n == nil {
        return false
    }
    if n.value == value {
        return true
    }
    if n.value > value {
        return n.left.IsExist(value)
    } else {
        return n.right.IsExist(value)
    }
}

func printNode(n *Node) {
    if n == nil {
        return
    }
    printNode(n.left)
    printNode(n.right)
    println(n.value)
}


func main() {
    t := &Node{value: 10}

    t.Insert(10).
    Insert(5).
    Insert(7).
    Insert(20).
    Insert(30).
    Insert(-2)

    printNode(t)
    println(t.IsExist(10))
    println(t.IsExist(1))
    println(t.IsExist(-2))

    println("Height", t.Height())
}