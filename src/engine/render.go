package engine

import (
	"fmt"
	"main/src/entity"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) SettingsRendering() {
	rl.DrawTexturePro(e.Background3, rl.NewRectangle(0, 0, 2000, 1414), rl.NewRectangle(0, 0, 1920, 1080), rl.NewVector2(0, 0), 0, rl.White)
}
func (e *Engine) LoreRendering() {
	rl.DrawTexturePro(e.Lore3, rl.NewRectangle(0, 0, 1980, 1080), rl.NewRectangle(0, 0, 1920, 1080), rl.NewVector2(0, 0), 0, rl.White)
}

func (e *Engine) CharactersRendering() {
	rl.ClearBackground(rl.Black)
	//rl.DrawTexturePro(e.Background, rl.NewRectangle(0, 0, 1414, 2000), rl.NewRectangle(0, 0, 700, 1080), rl.NewVector2(0, 0), 0, rl.White)
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
	//rl.DrawTexture(e.LoadingScreen, 0, 0, rl.White)
	rl.DrawTexturePro(e.LoadingScreen, rl.NewRectangle(float32(e.LoadingScreenSourceX), float32(e.LoadingScreenSourceY), 800, 450), rl.NewRectangle(0, 0, 1920, 1080), rl.NewVector2(0, 0), 0, rl.White)
	rl.DrawTexturePro(e.LoadingScreen, rl.NewRectangle(float32(e.LoadingScreenSourceX), float32(e.LoadingScreenSourceY), 800, 450), rl.NewRectangle(0, 0, 1920, 1080), rl.NewVector2(0, 0), 0, rl.White)

	rl.DrawText("Kenshin Sakura", int32(rl.GetScreenWidth())/5-rl.MeasureText("Kenshin Sakura", 80)/2, int32(rl.GetScreenHeight())/4-150, 80, rl.White)
	rl.DrawText("[Enter] play the game", int32(rl.GetScreenWidth())/8-rl.MeasureText("[Enter] to Play", 20)/2, int32(rl.GetScreenHeight())/4+100, 50, rl.White)
	rl.DrawText("[Esc] to quit", int32(rl.GetScreenWidth())/8-rl.MeasureText("[Esc] to Quit", 20)/2, int32(rl.GetScreenHeight())/4+400, 50, rl.White)
	rl.DrawText("[u] to select characters", int32(rl.GetScreenWidth())/8-rl.MeasureText("[u] to selet characters", 20)/2, int32(rl.GetScreenHeight())/4+200, 50, rl.White)
	rl.DrawText("[k] settings", int32(rl.GetScreenWidth())/8-rl.MeasureText("[k] settings", 20)/2, int32(rl.GetScreenHeight())/4+300, 50, rl.White)
}

func (e *Engine) InGameRendering() {
	rl.ClearBackground(rl.Gray)

	rl.BeginMode2D(e.Camera) // On commence le rendu camera

	e.RenderMap()

	e.RenderMonsters()
	e.RenderPlayer()
	e.RenderPnjs()

	if e.Player.Health <= 0 {
		e.GameOverRendering()
		return
	}

	rl.EndMode2D() // On finit le rendu camera

	// Ecriture fixe (car pas affectée par le mode camera)
	rl.DrawText("Playing", int32(rl.GetScreenWidth())/2-rl.MeasureText("Playing", 40)/2, int32(rl.GetScreenHeight())/2-350, 40, rl.RayWhite)
	rl.DrawText("[P] or [Esc] to Pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] or [Esc] to Pause", 20)/2, int32(rl.GetScreenHeight())/2-300, 20, rl.RayWhite)
	rl.DrawText(fmt.Sprintf("Health : %d ", e.Player.Health), int32(rl.GetScreenWidth())/26-rl.MeasureText("Health", 10)/2, int32(rl.GetScreenHeight())/20, 25, rl.Green)
	rl.DrawText(fmt.Sprintf("Energy : %d ", e.Player.Energy), int32(rl.GetScreenWidth())/26-rl.MeasureText("Health", 10)/2, int32(rl.GetScreenHeight())/12, 25, rl.Yellow)
	rl.DrawText(fmt.Sprintf("Health : %d ", e.Player.Health), int32(rl.GetScreenWidth())/26-rl.MeasureText("Health", 10)/2, int32(rl.GetScreenHeight())/20, 25, rl.Green)
	rl.DrawText(fmt.Sprintf("Energy : %d ", e.Player.Energy), int32(rl.GetScreenWidth())/26-rl.MeasureText("Health", 10)/2, int32(rl.GetScreenHeight())/12, 25, rl.Yellow)
	rl.DrawText(fmt.Sprintf("Health : %d ", e.Player.Health), int32(rl.GetScreenWidth())/26-rl.MeasureText("Health", 10)/2, int32(rl.GetScreenHeight())/20, 25, rl.Green)
	rl.DrawText(fmt.Sprintf("Energy : %d ", e.Player.Energy), int32(rl.GetScreenWidth())/26-rl.MeasureText("Health", 10)/2, int32(rl.GetScreenHeight())/12, 25, rl.Yellow)

}

func (e *Engine) PauseRendering() {
	rl.DrawTexture(e.Background, 0, 0, rl.White)
	rl.DrawTexture(e.Background, 0, 0, rl.White)

	rl.DrawText("Paused", int32(rl.GetScreenWidth())/2-rl.MeasureText("Paused", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.Red)
	rl.DrawText("[P] or [Esc] to resume", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] or [Esc] to resume", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.RayWhite)
	rl.DrawText("[Q] to Quit", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to Quit", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.RayWhite)

	rl.EndDrawing()
}

func (e *Engine) GameOverRendering() {
	rl.ClearBackground(rl.Red)

	rl.DrawText("[T] to respawn", int32(rl.GetScreenWidth())/2-rl.MeasureText("[T] to respawn", 20)/2, int32(rl.GetScreenHeight())/2+100, 30, rl.RayWhite)
	rl.DrawText("GAME OVER", int32(rl.GetScreenWidth())/2-rl.MeasureText("GAME OVER", 20)/2, int32(rl.GetScreenHeight())/2-200, 60, rl.RayWhite)
}


func (e *Engine) InventaireRendering() {
	rl.DrawTexturePro(e.LoadingInventaire, rl.NewRectangle(0, 0, 316, 149), rl.NewRectangle(0, 100, 500, 500), rl.NewVector2(0, 0), 0, rl.White)
}

func (e *Engine) RenderPlayer() {
	if e.Player.Class == entity.NINJA {
		rl.DrawTexturePro(
			e.Player.Sprite,
			rl.NewRectangle(80, 95, 100, 100),
			rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 150, 150),
			rl.Vector2{X: 0, Y: 0},
			0,
			rl.White,
		)
	}
	if e.Player.Class == entity.SAMOURAI {
		rl.DrawTexturePro(
			e.Player.Sprite,
			rl.NewRectangle(0, 0, 50, 50),
			rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 100, 100),
			rl.Vector2{X: 0, Y: 0},
			0,
			rl.White,
		)
	}
}

func (e *Engine) RenderMonsters() {
	for _, monster := range e.Monsters {
		if monster.IsAlive {
			if monster.IsAlive {
				rl.DrawTexturePro(
					monster.Sprite,
					rl.NewRectangle(0, 0, 100, 100),
					rl.NewRectangle(monster.Position.X, monster.Position.Y, 150, 150),
					rl.Vector2{X: 0, Y: 0},
					0,
					rl.White,
				)
				rl.DrawText(strconv.Itoa(monster.Health)+"/50", int32(monster.Position.X), int32(monster.Position.Y), 10, rl.White)
				rl.DrawText(strconv.Itoa(monster.Health)+"/50", int32(monster.Position.X), int32(monster.Position.Y), 10, rl.White)
			}

		}
		if monster.IsAlive {
			if monster.Name == "Ryuzo" {
				rl.DrawTexturePro(
					monster.Sprite,
					rl.NewRectangle(0, 0, 128, 128),
					rl.NewRectangle(monster.Position.X, monster.Position.Y, 100, 100),
					rl.Vector2{X: 0, Y: 0},
					0,
					rl.White,
				)
				rl.DrawText(strconv.Itoa(monster.Health)+"/100", int32(monster.Position.X), int32(monster.Position.Y), 10, rl.White)

			}
		}
		if monster.IsAlive {
			if monster.Name == "Maitre" {
				rl.DrawTexturePro(
					monster.Sprite,
					rl.NewRectangle(0, 0, 100, 100),
					rl.NewRectangle(monster.Position.X, monster.Position.Y, 150, 150),
					rl.Vector2{X: 0, Y: 0},
					0,
					rl.White,
				)
				rl.DrawText(strconv.Itoa(monster.Health)+"/60", int32(monster.Position.X), int32(monster.Position.Y), 10, rl.White)
			}
		}
	}
}

func (e *Engine) RenderPnjs() {
	for _, pnjs := range e.Pnjs {
		rl.DrawTexturePro(
			pnjs.Sprite,
			rl.NewRectangle(0, 0, 100, 100),
			rl.NewRectangle(pnjs.Position.X, pnjs.Position.Y, 150, 150),
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
		int32(m.Position.X)+25,
		int32(m.Position.Y)+50,
		10,
		rl.RayWhite,
	)

	rl.EndMode2D()
}
func (e *Engine) RenderDialo(m entity.Monster, sentence string) {
	rl.BeginMode2D(e.Camera)

	rl.DrawText(
		sentence,
		int32(m.Position.X)+25,
		int32(m.Position.Y)+95,
		10,
		rl.RayWhite,
	)

	rl.EndMode2D()
}

func (e *Engine) RnderDialog(p entity.Pnjs, sentence string) {
	rl.BeginMode2D(e.Camera)

	rl.DrawText(
		sentence,
		int32(p.Position.X)+50,
		int32(p.Position.Y)-2,
		10,
		rl.RayWhite,
	)
	rl.EndMode2D()
}
func (e *Engine) RendeDialog(p entity.Pnjs, sentence string) {
	rl.BeginMode2D(e.Camera)

	rl.DrawText(
		sentence,
		int32(p.Position.X)-25,
		int32(p.Position.Y)+50,
		10,
		rl.RayWhite,
	)
	rl.EndMode2D()
}

func (e *Engine) endeDialog(p entity.Pnjs, sentence string) {
	rl.BeginMode2D(e.Camera)

	rl.DrawText(
		sentence,
		int32(p.Position.X)-25,
		int32(p.Position.Y)+90,
		10,
		rl.RayWhite,
	)
	rl.EndMode2D()
}
func (e *Engine) RenerDialog(p entity.Pnjs, sentence string) {
	rl.BeginMode2D(e.Camera)

	rl.DrawText(
		sentence,
		int32(p.Position.X)-25,
		int32(p.Position.Y)+50,
		10,
		rl.RayWhite,
	)
	rl.EndMode2D()
	rl.EndMode2D()
	rl.EndMode2D()
}

