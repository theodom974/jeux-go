package engine

import (
	"main/src/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) Load() {
	// Chargement des textures du personnage
	e.Player.Sprite = rl.LoadTexture("textures/entities/ninja/ninja-Attack01.png")
	e.Background = rl.LoadTexture("textures/image_pause.jpg")
	e.LoadingScreen = rl.LoadTexture("textures/menu2.png")
	e.LoadingInventaire = rl.LoadTexture("textures/Inventaire.png")
}

func (e *Engine) Unload() {
	// On libère les textures chargées, le joueur, la map, les monstres, etc...
	rl.UnloadTexture(e.Player.Sprite)

	for _, sprite := range e.Sprites {
		rl.UnloadTexture(sprite)
	}

	for _, monster := range e.Monsters {
		rl.UnloadTexture(monster.Sprite)
	}

	for _, pnjs := range e.Pnjs {
		rl.UnloadTexture(pnjs.Sprite)
	}
}
func (e *Engine) UnloadCharacters() {
	rl.UnloadTexture(e.Player.Sprite)
}

func (e *Engine) loadCharacters() {
	if e.Player.Class == entity.NINJA {
		e.Player.Sprite = rl.LoadTexture("textures/entities/ninja/ninja-Attack01.png")
		e.Background = rl.LoadTexture("textures/entities/ninja/ninja.jpg")
	}
	if e.Player.Class == entity.SAMOURAI {
		e.Player.Sprite = rl.LoadTexture("textures/entities/Samourai/Samourai-Idle.png")
		e.Background = rl.LoadTexture("textures/entities/Samourai/Samourai.jpg")
	}
}