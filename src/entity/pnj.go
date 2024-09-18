package entity

import (

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Pnjs struct {
	Name string
	Position  rl.Vector2
	Health    int
	Money     int

	IsAlive bool

	Sprite rl.Texture2D
}
