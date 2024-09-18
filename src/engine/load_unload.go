package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) Load() {
	// Chargement des textures du personnage
	e.Player.Sprite = rl.LoadTexture("textures/entities/ninja/ninja-Attack01.png")
	e.Background = rl.LoadTexture("textures/image_pause.png")
	e.LoadingScreen = rl.LoadTexture("textures/menu2.png")
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
