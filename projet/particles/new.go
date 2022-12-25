package particles

import (
	"container/list"
	"fmt"
	"math/rand"
	"project-particles/config"
	"strconv"
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
		var posX float64 = float64(config.General.SpawnX)
		var posY float64 = float64(config.General.SpawnY)
		var speedX float64 = rand.Float64() * config.General.Velocity
		var speedY float64 = rand.Float64() * config.General.Velocity
		var s string = strconv.FormatInt(int64(i), 2)
		var ns int = len(s)
		if ns < 2 {
			s = "00"
		} else {
			s = s[ns-2:]
		}
		if string(s[0]) == "1" {
			speedX = speedX * -1
		}
		if string(s[1]) == "1" {
			speedY = speedY * -1
		}

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
			SpeedX:  speedX,
			SpeedY:  speedY,
		})
	}

	return System{Content: l}
}
