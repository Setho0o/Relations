package main

import (
	"image"
	"image/color"
)

type Display int
const (
  circle Display = iota
  square 
)

type Graph struct {
  all []*Node
}

type Node struct {
  x int
  y int
  index int
  box image.Rectangle
  con []Node
  d Display 
  c color.Color
}

func(n *Node) Display(g *Game) {
  if n.index == -1 {
    return
  }
  switch n.d {
    case circle:
      g.Circle(n)
    case square:
      g.Square(n) 
  }
}

func (g *Graph) Spawn(x, y, size int, d Display, c color.Color) {
  g.all = append(g.all, &Node {
    x: x,
    y: y,
    index: len(g.all),
    box: image.Rect(x-size,y-size,x+size,y+size),
    con: []Node{},
    d: d,
    c: c,
  }) 
}

func (g *Graph) Kill(n *Node) {
  n.index = -1  // hard to delete from slice so just setting index to -1 for later
}


