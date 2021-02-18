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
  max([]int{5,6,7,8})
<<<<<<< HEAD
=======
  max([]int{9,10,11,12})
>>>>>>> 91c1d6b20e63e6e9f38c7b6ac2a47cfe32a1aa22
}
