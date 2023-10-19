package particles

import (
	"container/list"

	//"math/rand"

	"project-particles/config"
	// "fmt"
)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
var spawnRateInteger float64

//var spawnRateValue float64 = config.General.SpawnRate

func (s *System) Update() {

	if config.HasGameStarted {
		var maParticule *Particle
		var myList *list.List = s.Content // Ma Liste de particules présent à l'écran
		var i int = 0
		var living bool = true
		var temp *list.Element
		for e := myList.Front(); e != nil && living; e = temp {
			temp = e.Next()
			i++
			maParticule, _ = e.Value.(*Particle)
			if maParticule.IsInLife {
				maParticule.PositionX = maParticule.PositionX + maParticule.SpeedX
				maParticule.PositionY = maParticule.PositionY + maParticule.SpeedY + float64(maParticule.LifeSpan)*config.General.Gravity //+ math.Sin(maParticule.SpeedY) //+ config.General.SizeShape*math.Sin(maParticule.Angle)
				if config.General.Collision {
					collisionWall(maParticule)
				}

				IsOutOfView(e, maParticule, s)
				if config.General.LifeSpanMax != 0 {
					LifeSpanIsTooAged(e, maParticule, s) //Enleve les particules trop vieilles
				}

				maParticule.LifeSpan++ //Augmente le compteur de durée de vie d'1

				maParticule.Rotation = 0
				shapePropriete(maParticule)

				setColor(maParticule) //Sert à afficher telle ou telle drapeau en fonction de la valeur de "flag" de config.json
			} else {
				living = false
				s.NbParticulesMortes = s.Content.Len() - i
			}

		}
		//Partie permettant de gérer les nombres flottant du spawnRate
		var spawnRate float64 = config.General.SpawnRate
		spawnRateInteger = spawnRateInteger + spawnRate - float64(int(spawnRate))
		if spawnRateInteger > 1 {
			spawnRateInteger = spawnRateInteger - 1
			spawnRate = spawnRate + 1
		}
		createNParticles(int(spawnRate), s)
	}

}
