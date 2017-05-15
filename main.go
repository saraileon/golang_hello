package main

import "fmt"

func main() {
  fmt.Println("hello, world")
  fmt.Println(42)
  // binary %b
  fmt.Printf("%d  - %b \n", 42, 42)
  // hexadecimal %x
  fmt.Printf("%d  - %x \n", 42, 42)
  fmt.Printf("%d  - %#X \n", 42, 42)
  // binary %b
  fmt.Printf("%d - %b \n", 394, 394)

  // for loop decimal, binary, hexadecimal
  for i := 0; i < 200; i ++ {
    fmt.Printf("%d \t %b \t %x \n", i, i, i)
  }
}
