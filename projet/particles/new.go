package particles

import (
	"container/list"
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


	rand.Seed(time.Now().UnixNano())
	for i := 0; i < config.General.InitNumParticles; i++ {
		var posX float64 = float64(config.General.SpawnX)
		var posY float64 = float64(config.General.SpawnY)
		var speedX float64 = rand.Float64() * config.General.Velocity
		var speedY float64 = rand.Float64() * config.General.Velocity
		var signe []int = []int{-1, 1}
		speedX = speedX * float64(signe[rand.Intn(2)]) //Donne à la vitesse en X et en Y une valeur aléatoire
		speedY = speedY * float64(signe[rand.Intn(2)])
		

		l.PushFront(&Particle{
			PositionX: float64(posX),
			PositionY: float64(posY),
			Rotation:  0.5,
			ScaleX:    config.General.ScaleX, ScaleY: config.General.ScaleY, 
			//Comme nos couleurs de notre config.json sont en 8 bits, nous les divisons par 255 pour avoir une valeur entre 0 et 1
			ColorRed: config.General.ColorRed / 255, ColorGreen: config.General.ColorGreen / 255, ColorBlue: config.General.ColorBlue / 255,
			Opacity: config.General.Opacity, 
			SpeedX:  speedX,
			SpeedY:  speedY,
		})
	}

	createNParticles(config.General.InitNumParticles, l)

	return System{Content: l}
}
