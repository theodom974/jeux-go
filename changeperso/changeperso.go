package changeperso_test

// Liste des images des personnages
Characters := []*ebiten.Image{char1, char2}

// Index du personnage actuellement sélectionné
CurrentIndex := 0

if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
    CurrentIndex = (CurrentIndex + 1) % len(Characters)
}
if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
    CurrentIndex = (CurrentIndex - 1 + len(Characters)) % len(Characters)
}

func (g *Game) Draw(screen *ebiten.Image) {
    screen.Fill(color.RGBA{0, 0, 0, 255}) // Efface l'écran avec une couleur noire
    
    // Dessiner le personnage actuellement sélectionné
    currentCharacter := Characters[CurrentIndex]
    op := &ebiten.DrawImageOptions{}
    op.GeoM.Translate(100, 100) // Positionner le personnage
    screen.DrawImage(currentCharacter, op)
}
