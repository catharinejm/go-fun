package file

import (
  "os"
  "syscall"
)

type File struct {
  fd int
  name string
}

func newFile(fd int, name string) *File {
  if fd < 0 {
    return nil
  }

  return &File{fd, name}
}

var (
  Stdin = newFile(0, "/dev/stdin")
  Stdout = newFile(1, "/dev/stdout")
  Stderr = newFile(2, "/dev/stderr")
)

func Open(name string, mode int, perm uint32) (file *File, err os.Error) {
  r, e := syscall.Open(name, mode, perm)
  if e != 0 {
    err = os.Errno(e)
  }
  return newFile(r, name), err
}
