package particles

import (
	"container/list"
	"math/rand"
	"project-particles/config"
	"time"
)

func puissance2(a float64) float64 {
	return a * a
}

func createNParticles(nb int, l *list.List) *list.List {

	/*
		nb : un entier représentant le nombre de particule souhaitant être créée
		l : une liste de particule
		Cette fonction sert à ajouter nb nouvelles particules à la liste l */

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < nb; i++ {
		var posX float64 = float64(config.General.SpawnX)
		var posY float64 = float64(config.General.SpawnY)
		var speedX float64 = rand.Float64() * config.General.Velocity
		var speedY float64 = rand.Float64() * config.General.Velocity
		var signe []int = []int{-1, 1}
		speedX = speedX * float64(signe[rand.Intn(2)])
		speedY = speedY * float64(signe[rand.Intn(2)])

		if config.General.RandomSpawn { //Appelle la fonction RandomSpawn si le paramètre "RandomSpawn" de config.json est mis à true
			posX = float64(rand.Intn(config.General.WindowSizeX))
			posY = float64(rand.Intn(config.General.WindowSizeY))
		}

		p := &Particle{
			PositionX: float64(posX),
			PositionY: float64(posY),
			ScaleX:    config.General.ScaleX, ScaleY: config.General.ScaleY,
			//Pour pouvoir écrire nos couleur RGB avec des valeurs comprisent entre 0 et 255, nous les divisons par 255 pour avoir une valeur entre 0 et 1
			ColorRed: config.General.ColorRed / 255, ColorGreen: config.General.ColorGreen / 255, ColorBlue: config.General.ColorBlue / 255,
			Opacity: config.General.Opacity,
			SpeedX:  speedX,
			SpeedY:  speedY,
		}
		l.PushFront(p)
	}
	return l
}
