package engine

import (
	"fmt"
	"main/src/entity"
	"main/src/fight"
	"main/src/item"
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
	if rl.IsKeyPressed(rl.KeyK) {
		e.Camera.Zoom = 1
		e.StateMenu = SETTINGS
	}
}

func (e *Engine) SettingsLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyK) {
		e.Camera.Zoom = 2
		e.StateMenu = HOME
	}
	//Musique
	rl.UpdateMusicStream(e.Music)
}

func (e *Engine) CharactersLogic() {
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
		e.LoadCharacters()
	}

	if rl.IsKeyPressed(rl.KeyLeft) {
		if e.Player.Class > 0 {
			e.Player.Class--
		} else {
			e.Player.Class = 1
		}
		e.UnloadCharacters()
		e.LoadCharacters()
	}
}

func (e *Engine) InGameLogic() {

	if !e.Player.IsAlive {
		e.StateEngine = GAMEOVER
	}

	//dash
	dashSpeed := e.Player.Speed * 1.5
	dashDuration := 100 * time.Millisecond
	dashCooldown := 2 * time.Second
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

	if rl.IsKeyPressed(rl.KeyTab) {
		e.StateEngine = INVENTAIRE
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
		e.Music = rl.LoadMusicStream("sounds/music/oiseau.mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)
}

func (e *Engine) TempleLogic() {

}

func (e *Engine) CheckCollisions() {

	e.MonsterCollisions()
}

func (e *Engine) CheckColliions() {
	e.PnjsColliions()
}

func (e *Engine) MonsterCollisions() {

	for i, monster := range e.Monsters {
		if monster.Position.X > e.Player.Position.X-40 &&
			monster.Position.X < e.Player.Position.X+40 &&
			monster.Position.Y > e.Player.Position.Y-40 &&
			monster.Position.Y < e.Player.Position.Y+40 {

			if monster.Name == "Maitre" {
				e.NormalTalk(monster, "La peur est un ennemi plus grand que celui que tu affronteras souviens-toi")
				if rl.IsKeyPressed(rl.KeyE) {
					fight.Fight(&e.Player, &e.Monsters[i])
				}
			}
			if monster.Name == "Ryuzo" {
				e.NormaTal(monster, "Vien te battre !")
				if rl.IsKeyPressed(rl.KeyE) {
					fight.Fight(&e.Player, &e.Monsters[i])
				}

			}
		}
	}
}

func (e *Engine) NormaTal(m entity.Monster, sentence string) {
	e.RenderDialo(m, sentence)
}

func (e *Engine) PnjsColliions() {

	for _, pnj := range e.Pnjs {
		if pnj.Position.X > e.Player.Position.X-40 &&
			pnj.Position.X < e.Player.Position.X+40 &&
			pnj.Position.Y > e.Player.Position.Y-40 &&
			pnj.Position.Y < e.Player.Position.Y+40 {

			if pnj.Name == "Garde 1" {
				e.NralTalkp(pnj, "Après cet entrainement trouver l'arbre \n     sacré gardé par le forgeron.")
			} else if pnj.Name == "Garde 2" {
				e.NormaleTlkp(pnj, "\n   Le maitre \n   vous attend.")
			 } else if pnj.Name == "Villageois" {
				e.NorlTalkp(pnj, "Salut cher voyageur.")
			 }else if pnj.Name == "SDF" {
				e.NormaTalkp(pnj, "Salut cher voyageur voici une arme pour ta quête !")
				if rl.IsKeyPressed(rl.KeyR) {
					e.Player.Inventaire = append(e.Player.Inventaire, item.Item{
						Name:         "Epée",
						Price:        1,
						IsConsumable: false,
						IsEquippable: true,
					})
				}
			}

			//quand tu parle a aucun pnj
		}
	}
}

func (e *Engine) TempleCollisions() {

}

func (e *Engine) NormalTalk(m entity.Monster, sentence string) {
	e.RenderDialog(m, sentence)
}
func (e *Engine) NralTalkp(p entity.Pnjs, sentence string) {
	e.RnderDialog(p, sentence)
}
func (e *Engine) NorlTalkp(p entity.Pnjs, sentence string) {
	e.endeDialog(p, sentence)
}
func (e *Engine) NormaleTlkp(p entity.Pnjs, sentence string) {
	e.RendeDialog(p, sentence)
}
func (e *Engine) NormaTalkp(p entity.Pnjs, sentence string) {
	e.RenerDialog(p, sentence)
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

func (e *Engine) GameOverLogic() {
	if rl.IsKeyPressed(rl.KeyT) {
		e.Player.IsAlive = true
		e.Player.Health = 100
		e.Player.Position = rl.NewVector2(100, 200)
		e.StateEngine = INGAME
	}
}

func (e *Engine) InventaireLogic() {
	if rl.IsKeyPressed(rl.KeyTab) {
		e.StateEngine = INGAME
	}
}
