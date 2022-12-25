package particles

import (
	"container/list"
	"fmt"
	"math/rand"
	"project-particles/config"
	"time"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.
func NewSystem() System {
	l := list.New()

	fmt.Println(l)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < config.General.InitNumParticles; i++ {
		var posX float64 = float64(config.General.WindowSizeX)
		var posY float64 = float64(config.General.WindowSizeY)
		if config.General.RandomSpawn {
			posX = float64(rand.Intn(config.General.WindowSizeX))
			posY = float64(rand.Intn(config.General.WindowSizeY))

		}
		l.PushFront(&Particle{
			PositionX: float64(posX),
			PositionY: float64(posY),
			ScaleX:    config.General.ScaleX, ScaleY: config.General.ScaleY, //Partie à remplacer
			ColorRed: config.General.ColorRed / 255, ColorGreen: config.General.ColorGreen / 255, ColorBlue: config.General.ColorBlue / 255,
			Opacity: config.General.Opacity,
			SpeedX:  config.General.Velocity,
			SpeedY:  config.General.Velocity,
		})
	}

	return System{Content: l}
}
