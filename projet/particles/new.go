package particles

import (
	"container/list"
	"fmt"
	"project-particles/config"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.
func NewSystem() System {
	l := list.New()

	fmt.Println(l)

	//
	l.PushFront(&Particle{
		PositionX: float64(config.General.WindowSizeX) / 2,
		PositionY: float64(config.General.WindowSizeY) / 2,
		ScaleX:    10, ScaleY: 10, //Partie à remplacer
		ColorRed: config.General.ColorRed / 255, ColorGreen: config.General.ColorGreen / 255, ColorBlue: config.General.ColorBlue / 255,
		Opacity:   1,
		SpawnRate: config.General.SpawnRate,
	})

	return System{Content: l}
}
