package main

import "fmt"
import "montecarlo"

func main() {
  c := &Card { val: 1, color: 2 }
  fmt.Println("Card is %d : %d", c.val, c.color)
}