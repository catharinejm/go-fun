package main

import (
  "./sort"
)

func ints() {
  data := []int{74, 59, 238, -784, 905, 0, 0, 42, 7586, -5467984, 7586}
  a := srt.IntArray(data)
  sort.Sort(a)
  if !sort.IsSorted(a) {
    panic("fail")
  }
}
