package entity

import (
	"fmt"
	"main/src/item"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type class int

const (
	NINJA class = iota
	SAMOURAI class = iota
	MAITRE class = iota
	RYUZO class = iota
)

type Player struct {
	Position  rl.Vector2
	Health    int
	Money     int
	Speed     float32
	Inventaire []item.Item
	NormalSpeed float32  
    Dashing     bool      
    LastDash    time.Time 
	Energy    int

	Class class
	IsAlive bool

	Sprite rl.Texture2D
}

func (p *Player) Attack(m *Monster) {
	m.Health -= 15
}

func (p *Player) ToString() {
	fmt.Printf(`
	Joueur:
		Vie: %d,
		Argent: %d,
		Inventaire: %+v
	
	\n`, p.Health, p.Money, p.Inventaire)
}
