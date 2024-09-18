package engine

import (
	"main/src/entity"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 1400
	ScreenHeight = 800
)

func (e *Engine) Init() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "Arcadia")

	// Initialisation des variables de l'engine
	e.IsRunning = true
	e.Sprites = make(map[string]rl.Texture2D)

	e.LoadingScreen = rl.LoadTexture("textures/menu.jpg")

	// Initialisation des composants du jeu
	e.InitEntities()
	e.InitCamera()
	e.InitMusic()
	e.InitMap("textures/map/tilesets/map.json")

}

func (e *Engine) InitEntities() {

	e.Player = entity.Player{
		Position:  rl.Vector2{X: 100, Y: 200},
		Health:    100,
		Money:     1000,
		Speed:     2,
		Inventory: []item.Item{},

		IsAlive: true,

		Sprite: e.Player.Sprite,
	}

		e.Monsters = append(e.Monsters, entity.Monster{
			Name:     "claude",
			Position: rl.Vector2{X: 400, Y: 320},
			Health:   50,
			Damage:   3,
			Loot:     []item.Item{},
			Worth:    12,

			IsAlive: true,
			Sprite:  rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
		})

	e.Player.Money = 12

	e.Pnjs = append(e.Pnjs, entity.Pnjs{
		Name:     "Garde 1",
		Position: rl.Vector2{X: 975, Y: 748},
		Health:   100,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/Garde/Garde-H.png"),
	})
	e.Pnjs = append(e.Pnjs, entity.Pnjs{
		Name:     "Garde 2",
		Position: rl.Vector2{X: 975, Y: 675},
		Health:   100,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/Garde/Garde-H.png"),
	})
	e.Pnjs = append(e.Pnjs, entity.Pnjs{
		Name:     "Femme",
		Position: rl.Vector2{X: 140, Y: 1100},
		Health:   100,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/Garde/Gade-F.png"),
	})
	e.Pnjs = append(e.Pnjs, entity.Pnjs{
		Name:     "Villageois",
		Position: rl.Vector2{X: 140, Y: 195},
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
