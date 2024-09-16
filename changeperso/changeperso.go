package changeperso

package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
	"golang.org/x/term"
)

// Personnage représente un personnage avec un nom pour simplifier l'affichage
type Personnage struct {
	Nom string
}

// Game gère le jeu et les personnages
type Game struct {
	Personnages   []Personnage
	IndexActuel   int
}

func NewGame() *Game {
	return &Game{
		Personnages: []Personnage{
			{Nom: "Personnage 1"},
			{Nom: "Personnage 2"},
			{Nom: "Personnage 3"},
		},
		IndexActuel: 0,
	}
}

func (g *Game) ChangerPersonnage(direction string) {
	if direction == "droite" {
		g.IndexActuel = (g.IndexActuel + 1) % len(g.Personnages)
	} else if direction == "gauche" {
		g.IndexActuel = (g.IndexActuel - 1 + len(g.Personnages)) % len(g.Personnages)
	}
}

func (g *Game) AfficherPersonnage() {
	fmt.Printf("Personnage actuel : %s\n", g.Personnages[g.IndexActuel].Nom)
}

func main() {
	game := NewGame()
	
	// Désactiver le buffering de sortie pour que l'affichage soit immédiat
	fmt.Printf("\033[2J") // Efface l'écran
	fmt.Printf("\033[H")  // Déplace le curseur en haut à gauche
	
	// Détection des touches
	for {
		game.AfficherPersonnage()
		
		// Lire l'entrée de l'utilisateur
		fmt.Print("Utilisez 'gauche' ou 'droite' pour changer de personnage (ou 'exit' pour quitter) : ")
		var direction string
		fmt.Scanln(&direction)
		
		if direction == "exit" {
			break
		}
		
		game.ChangerPersonnage(direction)
		
		// Pause pour ne pas saturer la console
		time.Sleep(500 * time.Millisecond)
		// Effacer l'écran avant de réafficher le personnage
		fmt.Printf("\033[2J")
	}
}
