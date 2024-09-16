package changeperso

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"image/color"
	_ "image/png"
)

type Game struct {
	Characters   []*ebiten.Image // Liste des images des personnages
	CurrentIndex int              // Index du personnage actuellement sélectionné
}

func NewGame() (*Game, error) {
	// Charger les images des personnages
	char1, _, err := ebitenutil.NewImageFromFile("char1.png", ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	char2, _, err := ebitenutil.NewImageFromFile("char2.png", ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	
	// Créer le jeu avec les personnages chargés
	return &Game{
		Characters:   []*ebiten.Image{char1, char2},
		CurrentIndex: 0,
	}, nil
}

func (g *Game) Update() error {
	// Changer de personnage avec les touches gauche/droite
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.CurrentIndex = (g.CurrentIndex + 1) % len(g.Characters)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.CurrentIndex = (g.CurrentIndex - 1 + len(g.Characters)) % len(g.Characters)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255}) // Efface l'écran avec du noir
	
	// Dessiner le personnage actuellement sélectionné
	currentCharacter := g.Characters[g.CurrentIndex]
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(320-currentCharacter.Bounds().Dx()/2, 240-currentCharacter.Bounds().Dy()/2) // Centrer le personnage
	screen.DrawImage(currentCharacter, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func main() {
	game, err := NewGame()
	if err != nil {
		panic(err)
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Character Selection Menu")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
