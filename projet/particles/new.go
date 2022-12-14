package particles

import (
	"container/list"
	"project-particles/config"
	"fmt"
	"math/rand"
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
		ScaleX: 10,
		ScaleY: 10,										//Partie à remplacer
		ColorRed: float64(rand.Intn(255)), ColorGreen: float64(rand.Intn(255)), ColorBlue: float64(rand.Intn(255)),
		Opacity: 1,
		SpawnRate:config.General.SpawnRate,
	})
	
	return System{Content: l}
}
