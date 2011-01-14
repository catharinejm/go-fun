package main

import "flag"

var w = flag.Int("w", 800, "width")
var h = flag.Int("h", 600, "height")

func main() {
  flag.Parse()
  width := *w
  height := *h
}

type Vector struct {
  X, Y, Z float64
}

func (p *Vector) Normalize() {
  len := p.Len()
  p.X /= len;
  p.Y /= len;
  p.Z /= len;
}
