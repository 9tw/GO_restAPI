package main

import "fmt"

func max(num []int) {
  total := 0
  for i := 0; i <= len(num)-1; i++ {
    total += num[i]
    fmt.Println(total)
  }
}

func main() {
  max([]int{1,2,3,4})
  max([]int{9,10,11,12})
}
