package engine

import (
	"fmt"
	"main/src/entity"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) Rendering() {
	rl.DrawTexture(rl.LoadTexture("textures/menup.png"), 0, 0, rl.White)
}

func (e *Engine) HomeRendering() {
	//rl.DrawTexture(e.LoadingScreen, 0, 0, rl.White)
	rl.DrawTexturePro(e.LoadingScreen, rl.NewRectangle(float32(e.LoadingScreenSourceX), float32(e.LoadingScreenSourceY), 800, 450), rl.NewRectangle(0, 0, 1920, 1080), rl.NewVector2(0,0), 0, rl.White)

	rl.DrawText("Kenshin Sakura", int32(rl.GetScreenWidth())/5-rl.MeasureText("Kenshin Sakura", 80)/2, int32(rl.GetScreenHeight())/4-150, 80, rl.White)
	rl.DrawText("[Enter] Pour Jouer", int32(rl.GetScreenWidth())/8-rl.MeasureText("[Enter] to Play", 20)/2, int32(rl.GetScreenHeight())/4+100, 50, rl.White)
	rl.DrawText("[Esc] Pour Quitter", int32(rl.GetScreenWidth())/8-rl.MeasureText("[Esc] to Quit", 20)/2, int32(rl.GetScreenHeight())/2, 50, rl.White)
	
}

func (e *Engine) InGameRendering() {
	rl.ClearBackground(rl.Gray)

	rl.BeginMode2D(e.Camera) // On commence le rendu camera

	e.RenderMap()

	e.RenderMonsters()
	e.RenderPlayer()
	e.RenderPnjs()

	rl.EndMode2D() // On finit le rendu camera

	// Ecriture fixe (car pas affectée par le mode camera)
	rl.DrawText("Playing", int32(rl.GetScreenWidth())/2-rl.MeasureText("Playing", 40)/2, int32(rl.GetScreenHeight())/2-350, 40, rl.RayWhite)
	rl.DrawText("[P] or [Esc] to Pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] or [Esc] to Pause", 20)/2, int32(rl.GetScreenHeight())/2-300, 20, rl.RayWhite)
	rl.DrawText(fmt.Sprintf("Health : %d ", e.Player.Health), int32(rl.GetScreenWidth())/26-rl.MeasureText("Health", 10)/2,int32(rl.GetScreenHeight())/20, 25, rl.Green)
	rl.DrawText(fmt.Sprintf("Energy : %d ", e.Player.Energy), int32(rl.GetScreenWidth())/26-rl.MeasureText("Health", 10)/2,int32(rl.GetScreenHeight())/12, 25, rl.Yellow)

}

func (e *Engine) PauseRendering() {
	rl.DrawTexture(e.Background , 0, 0, rl.White)

	rl.DrawText("Paused", int32(rl.GetScreenWidth())/2-rl.MeasureText("Paused", 40)/2, int32(rl.GetScreenHeight())/2-150, 40, rl.Red)
	rl.DrawText("[P] or [Esc] to resume", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] or [Esc] to resume", 20)/2, int32(rl.GetScreenHeight())/2, 20, rl.RayWhite)
	rl.DrawText("[Q] to Quit", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to Quit", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.RayWhite)

	rl.EndDrawing()
}

func (e *Engine) RenderPlayer() {

	rl.DrawTexturePro(
		e.Player.Sprite,
		rl.NewRectangle(80, 95, 100, 100),
		rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 150, 150),
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)

}

func (e *Engine) RenderMonsters() {
	for _, monster := range e.Monsters {
	  if monster.IsAlive {
			rl.DrawTexturePro(
				monster.Sprite,
				rl.NewRectangle(0, 0, 100, 100),
				rl.NewRectangle(monster.Position.X, monster.Position.Y, 150, 150),
				rl.Vector2{X: 0, Y: 0},
				0,
				rl.White,
			)
			rl.DrawText(strconv.Itoa(monster.Health) + "/50", int32(monster.Position.X), int32(monster.Position.Y), 10, rl.White )
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
		int32(m.Position.X),
		int32(m.Position.Y)+50,
		10,
		rl.RayWhite,
	)

	rl.EndMode2D()
}

func (e *Engine) RendrDialog(p entity.Pnjs, sentence string) {
	rl.BeginMode2D(e.Camera)

	rl.DrawText(
		sentence,
		int32(p.Position.X),
		int32(p.Position.Y)+50,
		10,
		rl.RayWhite,
	)
rl.EndMode2D()	
}