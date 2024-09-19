package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (engine *Engine) Run() {
	rl.SetTargetFPS(60)

	for engine.IsRunning {

		rl.BeginDrawing()

		switch engine.StateMenu {
		case HOME:
			engine.HomeRendering()
			engine.HomeLogic()

		case CHARACTERS:
			engine.CharactersRendering()
			engine.CharactersLogic()

		case SETTINGS:
			engine.SettingsLogic()
			engine.SettingsRendering()

		case PLAY:
			switch engine.StateEngine {
			case INGAME:
				engine.InGameRendering()
				engine.InGameLogic()
			case PAUSE:
				engine.PauseRendering()
				engine.PauseLogic()

			case GAMEOVER:
				//...
			
			case INVENTAIRE:
				engine.InventaireRendering()
				engine.InventaireLogic()

			}
		}
		rl.EndDrawing()
	}
	
}
