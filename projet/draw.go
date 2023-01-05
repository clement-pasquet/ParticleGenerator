package main

import (
	"fmt"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particles"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Draw se charge d'afficher à l'écran l'état actuel du système de particules
// g.system. Elle est appelée automatiquement environ 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction pourra être légèrement modifiée quand
// c'est précisé dans le sujet.

func (g *game) Draw(screen *ebiten.Image) {

	for e := g.system.Content.Front(); e != nil; e = e.Next() {
		p, ok := e.Value.(*particles.Particle)
		if ok {
			options := ebiten.DrawImageOptions{}
			options.GeoM.Rotate(p.Rotation)
			options.GeoM.Scale(p.ScaleX, p.ScaleY)
			options.GeoM.Translate(p.PositionX, p.PositionY)
			options.ColorM.Scale(p.ColorRed, p.ColorGreen, p.ColorBlue, p.Opacity)
			screen.DrawImage(assets.ParticleImage, &options)
		}
	}
	if config.General.Debug {
		ebitenutil.DebugPrint(screen, fmt.Sprint(ebiten.CurrentTPS()))

	}
}

/*rand.Seed(time.Now().UnixNano())
for e := g.system.Content.Front(); e != nil; e = e.Next() {
	p, ok := e.Value.(*particles.Particle)
	if ok {
		options := ebiten.DrawImageOptions{}
		options.GeoM.Rotate(float64(rand.Intn(50)))
		options.GeoM.Scale(float64(rand.Intn(10)),float64(rand.Intn(10)))
		options.GeoM.Translate(p.PositionX+float64(rand.Intn(1280)-640), p.PositionY+float64(rand.Intn(720)-360))
		options.ColorM.Scale(float64(rand.Intn(100)), float64(rand.Intn(100)), float64(rand.Intn(100)), float64(rand.Intn(10))/10)
		screen.DrawImage(assets.ParticleImage, &options)
	}
}

}*/
