package main

type Stack struct {
    items[] int
}

func (s *Stack) Push(item int) {
    s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
    left := len(s.items)
    if left == 0 {
        return nil
    }
    item, items := s.items[left-1], s.items[0:left-1] 
    s.items = items
    return item
}

func main() {

    s := Stack{}

    for i := 0; i < 4; i++ {
        s.Push(i)
    }

    println(s.Pop())

}
