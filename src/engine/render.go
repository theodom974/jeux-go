package engine

import (
	"main/src/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) Rendering() {
	rl.DrawTexture(e.BackgroundHome, 0, 0, rl.White)
}

func (e *Engine) charactersRendering() {
	rl.ClearBackground(rl.Black)
	//rl.DrawTexturePro(e.Background, rl.NewRectangle(0, 0, 1414, 2000), rl.NewRectangle(700, 0, 700, 1080), rl.NewVector2(0, 0), 0, rl.White)
	rl.DrawText(">", 1500, 500, 200, rl.White)
	switch e.Player.Class {
	case entity.SAMOURAI:
		rl.DrawTexturePro(e.Background, rl.NewRectangle(0, 0, 1414, 2000), rl.NewRectangle(700, 0, 700, 1080), rl.NewVector2(0, 0), 0, rl.White)
		rl.DrawTexturePro(
			e.Player.Sprite,
			rl.NewRectangle(0, 0, 50, 70),
			rl.NewRectangle(950, 200, 400, 560),
			rl.Vector2{X: 0, Y: 0},
			0,
			rl.White,
		)
	case entity.NINJA:
		rl.DrawTexturePro(e.Background, rl.NewRectangle(0, 0, 1414, 2000), rl.NewRectangle(700, 0, 700, 1080), rl.NewVector2(0, 0), 0, rl.White)
		rl.DrawTexturePro(
			e.Player.Sprite,
			rl.NewRectangle(0, 0, 100, 100),
			rl.NewRectangle(650, 150, 800, 800),
			rl.Vector2{X: 0, Y: 0},
			0,
			rl.White,
		)
	}
}

func (e *Engine) HomeRendering() {
	rl.ClearBackground(rl.Black)
	rl.DrawTexture(e.Home, 0, 0, rl.White)

	rl.DrawText("Home Menu", int32(rl.GetScreenWidth())/2-rl.MeasureText("Home Menu", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.Red)
	rl.DrawText("[Enter] to Play", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Enter] to Play", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.Red)
	rl.DrawText("[Esc] to Quit", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to Quit", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.Red)
	rl.DrawText("[u] to an characters", int32(rl.GetScreenWidth())/2-rl.MeasureText("[u] to an characters", 20)/2, int32(rl.GetScreenHeight())/2+50, 20, rl.Red)

}

func (e *Engine) InGameRendering() {
	rl.ClearBackground(rl.Gray)

	rl.BeginMode2D(e.Camera) // On commence le rendu camera

	e.RenderMap()

	e.RenderMonsters()
	e.RenderPlayer()

	rl.EndMode2D() // On finit le rendu camera

	// Ecriture fixe (car pas affectée par le mode camera)
	rl.DrawText("Playing", int32(rl.GetScreenWidth())/2-rl.MeasureText("Playing", 40)/2, int32(rl.GetScreenHeight())/2-350, 40, rl.RayWhite)
	rl.DrawText("[P] or [Esc] to Pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] or [Esc] to Pause", 20)/2, int32(rl.GetScreenHeight())/2-300, 20, rl.RayWhite)

}

func (e *Engine) PauseRendering() {
	rl.DrawTexture(rl.LoadTexture("textures/image_pause.jpg"), 0, 0, rl.White)

	rl.DrawText("Paused", int32(rl.GetScreenWidth())/2-rl.MeasureText("Paused", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.Red)
	rl.DrawText("[P] or [Esc] to resume", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] or [Esc] to resume", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.RayWhite)
	rl.DrawText("[Q] to Quit", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to Quit", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.RayWhite)
	rl.DrawText("[u] to an characters", int32(rl.GetScreenWidth())/2-rl.MeasureText("[u] to an characters", 20)/2, int32(rl.GetScreenHeight())/2+50, 20, rl.RayWhite)

	rl.EndDrawing()
}

func (e *Engine) RenderPlayer() {
	switch e.Player.Class {
	case entity.NINJA :
		rl.DrawTexturePro(
			e.Player.Sprite,
			rl.NewRectangle(0, 0, 100, 100),
			rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 150, 150),
			rl.Vector2{X: 0, Y: 0},
			0,
			rl.White,
		)
	case entity.SAMOURAI :
		rl.DrawTexturePro(
			e.Player.Sprite,
			rl.NewRectangle(0, 0, 50, 70),
			rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 70, 90),
			rl.Vector2{X: 0, Y: 0},
			0,
			rl.White,
		)
	}
	// rl.DrawTexturePro(
	// 	e.Player.Sprite,
	// 	rl.NewRectangle(-20, -5, 100, 100),
	// 	rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 150, 150),
	// 	rl.Vector2{X: 0, Y: 0},
	// 	0,
	// 	rl.White,
	// )
}

func (e *Engine) RenderMonsters() {
	for _, monster := range e.Monsters {
		rl.DrawTexturePro(
			monster.Sprite,
			rl.NewRectangle(0, 0, 100, 100),
			rl.NewRectangle(monster.Position.X, monster.Position.Y, 150, 150),
			rl.Vector2{X: 0, Y: 0},
			0,
			rl.White,
		)
	}
}

func (e *Engine) RenderDialog(m entity.Monster, sentence string) {
	rl.BeginMode2D(e.Camera)

	rl.DrawText(
		sentence,
		int32(m.Position.X),
		int32(m.Position.Y)+50,
		10,
		rl.RayWhite,
	)

	rl.EndMode2D()
}
