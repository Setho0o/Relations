package main

import "image/color"

func (g *Game) Circle(n *Node) {

}

func (g *Game) Square(n *Node) {
  p := n.box.Size()
  for i := range p.X {
    for j := range p.Y {
      g.s.Set(n.box.Min.X + i, n.box.Min.Y + j, n.c) 
    }
  }
  g.s.Set(n.x,n.y, color.Black) 
}
