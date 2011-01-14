package main

import (
  "flag"
  "math"
)

var w = flag.Int("w", 800, "width")
var h = flag.Int("h", 600, "height")

type Vector struct {
  X, Y, Z float64
}

func (p *Vector) Len() float64 {
  return math.Sqrt(p.X*p.X + p.Y*p.Y + p.Z*p.Z)
}

func (p *Vector) Normalize() {
  len := p.Len()
  p.X /= len;
  p.Y /= len;
  p.Z /= len;
}

func distanceToSphere(p Vector) float64 {
  return p.Len() - 1.0
}

type Sphere struct {
  Position Vector
  Radius float64
}

func (s Sphere) Distance(p Vector) float64 {
  p = p.TranslateBy(s.Position)

  return p.Len() - s.Radius
}

type DistanceFieldObject interface {
  Distance(p Vector) float64
}

func main() {
  flag.Parse()

  field := DistanceFieldObject{Vector{400, 300, 100}, 100}
}
