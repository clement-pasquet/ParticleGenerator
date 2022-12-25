package particles

import (
	"container/list"
	"fmt"
	"math/rand"
	"project-particles/config"
)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
var nb float64

func (s *System) Update() {
	var ele *Particle
	var myList *list.List = s.Content
	fmt.Println(myList.Len())
	for e := myList.Front(); e != nil; e = e.Next() {
		ele = e.Value.(*Particle)
		ele.PositionX = ele.PositionX + ele.SpeedX
		ele.PositionY = ele.PositionY + ele.SpeedY
	}
	var a float64 = config.General.SpawnRate
	nb = nb + a - float64(int(a))
	if nb > 1 {
		nb = nb - 1
		a = a + 1
	}
	createNParticles(int(a), myList)

}

func createNParticles(nb int, l *list.List) *list.List {
	for i := 0; i < nb; i++ {
		var posX float64 = float64(config.General.SpawnX)
		var posY float64 = float64(config.General.SpawnY)
		var speedX float64 = rand.Float64() * config.General.Velocity
		var speedY float64 = rand.Float64() * config.General.Velocity
		var signe []int = []int{-1, 1}
		speedX = speedX * float64(signe[rand.Intn(2)])
		speedY = speedY * float64(signe[rand.Intn(2)])

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
	return l
}
