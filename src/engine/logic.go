package engine

import (
	"fmt"
	"main/src/entity"
	"main/src/fight"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) HomeLogic() {
	e.LoadingScreenCountFrame++
	fmt.Println(e.LoadingScreenCountFrame)
	if e.LoadingScreenCountFrame%4 == 1 {
		if e.LoadingScreenSourceX == 15200 {
			e.LoadingScreenSourceX = 0
		} else {
			e.LoadingScreenSourceX += 800
		}
	}

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/bleachost.mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)

	//Menus
	if rl.IsKeyPressed(rl.KeyEnter) {
		e.StateMenu = PLAY
		rl.StopMusicStream(e.Music)
		e.StateEngine = INGAME

	}
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.IsRunning = false
	}
	if rl.IsKeyPressed(rl.KeyU) {
		e.StateMenu = CHARACTERS
	}
}

func (e *Engine) SettingsLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyB) {
		e.StateMenu = HOME
	}
	//Musique
	rl.UpdateMusicStream(e.Music)
}

func (e *Engine) charactersLogic() {
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.StateMenu = HOME
	}

	if rl.IsKeyPressed(rl.KeyRight) {
		if e.Player.Class < 1 {
			e.Player.Class++
		} else {
			e.Player.Class = 0
		}
		e.UnloadCharacters()
		e.loadCharacters()
	}

	if rl.IsKeyPressed(rl.KeyLeft) {
		if e.Player.Class > 0 {
			e.Player.Class--
		} else {
			e.Player.Class = 1
		}
		e.UnloadCharacters()
		e.loadCharacters()
	}
}

func (e *Engine) InGameLogic() {
	//dash
	dashSpeed := e.Player.Speed * 1.5
	dashDuration := 100 * time.Millisecond
	dashCooldown := 5 * time.Second
	now := time.Now()

	if (rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeySpace)) && now.Sub(e.Player.LastDash) > dashCooldown {
		e.Player.Dashing = true
		e.Player.LastDash = now
	}

	if e.Player.Dashing {
		e.Player.Speed = dashSpeed

		if now.Sub(e.Player.LastDash) > dashDuration {
			e.Player.Dashing = false
			e.Player.Speed = 2
		}
	}
	// Mouvement
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		e.Player.Position.Y -= e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		e.Player.Position.Y += e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		e.Player.Position.X -= e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		e.Player.Position.X += e.Player.Speed
	}

	// Camera
	e.Camera.Target = rl.Vector2{X: e.Player.Position.X + 70, Y: e.Player.Position.Y + 70}
	e.Camera.Offset = rl.Vector2{X: ScreenWidth / 2, Y: ScreenHeight / 2}

	// Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = PAUSE
	}

	e.CheckCollisions()
	e.CheckColliions()

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/Vent2.mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)
}

func (e *Engine) CheckCollisions() {

	e.MonsterCollisions()
}

func (e *Engine) CheckColliions() {
	e.PnjsColliions()
}

func (e *Engine) MonsterCollisions() {

	for i, monster := range e.Monsters {
		if monster.Position.X > e.Player.Position.X-20 &&
			monster.Position.X < e.Player.Position.X+20 &&
			monster.Position.Y > e.Player.Position.Y-20 &&
			monster.Position.Y < e.Player.Position.Y+20 {

			if monster.Name == "claude" {
				e.NormalTalk(monster, "Bonjour")
				if rl.IsKeyPressed(rl.KeyE) {
					fight.Fight(&e.Player, &e.Monsters[i])
				}
			}
		} else {
			//...
		}
	}
}

func (e *Engine) PnjsColliions() {

	for _, pnj := range e.Pnjs {
		if pnj.Position.X > e.Player.Position.X-20 &&
			pnj.Position.X < e.Player.Position.X+20 &&
			pnj.Position.Y > e.Player.Position.Y-20 &&
			pnj.Position.Y < e.Player.Position.Y+20 {

			if pnj.Name == "Garde 1" {
				e.NoralTalkp(pnj, "Salut cher voyageur")
			} else if pnj.Name == "Garde 2" {
				e.NoralTalkp(pnj, "jte baise")
			} else {
				//quand tu parle a aucun pnj
			}
		}
	}
}

func (e *Engine) NormalTalk(m entity.Monster, sentence string) {
	e.RenderDialog(m, sentence)
}
func (e *Engine) NoralTalkp(p entity.Pnjs, sentence string) {
	e.RendrDialog(p, sentence)
}

func (e *Engine) PauseLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = INGAME
	}
	if rl.IsKeyPressed(rl.KeyA) {
		e.StateMenu = HOME
		rl.StopMusicStream(e.Music)
	}

	//Musique
	rl.UpdateMusicStream(e.Music)
}
