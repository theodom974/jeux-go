package engine

import (
	"main/src/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type menu int

const (
	HOME     menu = iota
	SETTINGS menu = iota
	PLAY     menu = iota
)

type engine int

const (
	INGAME  engine = iota
	INFIGHT engine = iota
	PAUSE    engine = iota
	GAMEOVER engine = iota
)

type Engine struct {


	Player   entity.Player
	Monsters []entity.Monster
	Pnjs []entity.Pnjs

	Music       rl.Music
	MusicVolume float32

	Sprites map[string]rl.Texture2D

	Camera rl.Camera2D

	MapJSON MapJSON
	Background rl.Texture2D

	LoadingScreenCountFrame int
	LoadingScreen rl.Texture2D
	LoadingScreenSourceX int
	LoadingScreenSourceY int

	IsRunning   bool
	StateMenu   menu
	StateEngine engine
}
