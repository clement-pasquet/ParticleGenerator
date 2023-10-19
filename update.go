package main

import (
	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Update se charge d'appeler la fonction Update du système de particules
// g.system. Elle est appelée automatiquement exactement 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction ne devrait pas être modifiée sauf
// pour les deux dernières extensions.
func (g *game) Update() error {
	if config.HasGameStarted {
		g.system.Update()

		if config.General.MouseClick {
			if inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft) > 0 {
				config.General.SpawnX, config.General.SpawnY = ebiten.CursorPosition()

			}
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyNumpadAdd) {
			config.General.SpawnRate = config.General.SpawnRate + 10
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyNumpadSubtract) {
			config.General.SpawnRate = config.General.SpawnRate - 10
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyL) {
			config.General.ScaleX = config.General.ScaleX * 2
			config.General.ScaleY = config.General.ScaleY * 2

		}
		if inpututil.IsKeyJustPressed(ebiten.KeyM) {
			config.General.ScaleX = config.General.ScaleX / 2
			config.General.ScaleY = config.General.ScaleY / 2

		}

	} else {
		g.system.Update()
	}
	return nil

}
