package main

import "fmt"

func hello() {
  println("hello")
}

func sum (x, y int) int {
  return x + y
}

func main() {

  hello()
  
  productClosure := func(x, y int) int {
    return x*y
  }

  for a,b := 0,0; a < 3; a,b = a + 1, b + 1 {
    fmt.Printf("Sum %d  + %d = %d\n", a, b, sum(a, b))
    fmt.Printf("Product %d * %d = %d\n", a, b, productClosure(a, b))
  }

}