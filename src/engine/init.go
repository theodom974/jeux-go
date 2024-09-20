package engine

import (
	"main/src/entity"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 1920
	ScreenHeight = 1080
)

func (e *Engine) Init() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "Arcadia")

	// Initialisation des variables de l'engine
	e.IsRunning = true
	e.Sprites = make(map[string]rl.Texture2D)

	e.LoadingScreen = rl.LoadTexture("textures/menu.jpg")
	e.Player.Class = entity.NINJA
	e.Background = rl.LoadTexture("textures/entities/ninja/ninja.jpg")

	// Initialisation des composants du jeu
	e.InitEntities()
	e.InitCamera()
	e.InitMusic()
	e.InitMap("textures/map/tilesets/map.json")
	e.LoadingScreenSourceX = 0
	e.LoadingScreenSourceY = 0

}

func (e *Engine) InitEntities() {

	e.Player = entity.Player{
		Position:  rl.Vector2{X: 100, Y: 200},
		Health:    100,
		Money:     1000,
		Speed:     2,
		Energy:    50,
		Inventaire: []item.Item{},

		IsAlive: true,

		Sprite: e.Player.Sprite,
	}

		e.Monsters = append(e.Monsters, entity.Monster{
			Name:     "Maitre",
			Position: rl.Vector2{X: 1270, Y: 900},
			Health:   60,
			Damage:   3,
			Loot:     []item.Item{},
			Worth:    12,

			IsAlive: true,
			Sprite:  rl.LoadTexture("textures/entities/Villageois/Villageois-M.png"),
		})
		e.Monsters = append(e.Monsters, entity.Monster{
			Name: "Ryuzo",
			Position: rl.Vector2{X: 140, Y: 1120},
			Health: 100,
			Damage: 25,

			IsAlive: true,
			Sprite: rl.LoadTexture("textures/entities/mechant/Idle.png"),
		})

	e.Player.Money = 12

	e.Pnjs = append(e.Pnjs, entity.Pnjs{
		Name:     "Garde 1",
		Position: rl.Vector2{X: 1310, Y: 810},
		Health:   100,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/Garde/Garde-H.png"),
	})
	e.Pnjs = append(e.Pnjs, entity.Pnjs{
		Name:     "Garde 2",
		Position: rl.Vector2{X: 1200, Y: 810},
		Health:   100,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/Garde/Garde-H.png"),
	})
	e.Pnjs = append(e.Pnjs, entity.Pnjs{
		Name:     "Villageois",
		Position: rl.Vector2{X: 160, Y: 195},
		Health:   100,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/Villageois/villageois.png"),
	})
	e.Pnjs = append(e.Pnjs, entity.Pnjs{
		Name:     "SDF",
		Position: rl.Vector2{X: 1250, Y: 45},
		Health:   100,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/Villageois/Villageois-S.png"),
	})
}

func (e *Engine) InitCamera() {
	e.Camera = rl.NewCamera2D( //Camera vide, a changer dans chaque logique de scene
		rl.NewVector2(0, 0),
		rl.NewVector2(0, 0),
		0.0,
		2.0,
	)
}

func (e *Engine) InitMusic() {
	rl.InitAudioDevice()

	e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-08-Egress.mp3")

	rl.PlayMusicStream(e.Music)
}

