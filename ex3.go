package main

import (
  "./file"
  "fmt"
  "os"
)

func main() {
  hello := []byte("hello, world\n")
  file.Stdout.Write(hello)
  f, err := file.Open("/does/not/exist", 0, 0)
  if f == nil {
    fmt.Printf("can't open file; err=%s\n", err.String())
    os.Exit(1)
  }
}
