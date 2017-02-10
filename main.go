package main

import (
	"fmt"

	"engo.io/ecs"
	"engo.io/engo"
)

type myScene struct{}

func (*myScene) Type() string { return "myGame" }

func (*myScene) Preload() {}

func (*myScene) Setup(*ecs.World) {}

func main() {
	opts := engo.RunOptions{
		Title:  "Hello World",
		Width:  800,
		Height: 800,
	}
	fmt.Println("Hello!")
	engo.Run(opts, &myScene{})
}
