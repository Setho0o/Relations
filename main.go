package main

import (
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
  X = 1920
  Y = 1080
  LayoutX = X/4
  LayoutY = Y/4

)


type Game struct{
  gr *Graph
  s *ebiten.Image
}

func (g *Game) Update() error {
  g.Keys()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
  g.s = screen
  g.Display()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return LayoutX, LayoutY
}

func main() {
	ebiten.SetWindowSize(X,Y)
  
  game := GameInit()
  game.gr.Spawn(100,100,5,square,color.White)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func GameInit() *Game {
  return &Game{
    gr: &Graph {
      all: []*Node{},
    },
  }
}

func (g *Game) Keys() {
  if ebiten.IsKeyPressed(ebiten.KeyQ) {
    os.Exit(0)
  }
}

func(g *Game) Display() {
  //collison
  for _,node := range g.gr.all {
    node.Display(g)
  }
}
