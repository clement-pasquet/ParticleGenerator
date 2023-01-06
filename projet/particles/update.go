package particles

import (
	"container/list"
	"math/rand"
	"project-particles/config"
)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
var nb float64
var particulesMortes *list.List = list.New()

func (s *System) Update() {
	var maParticule *Particle
	var myList *list.List = s.Content // Liste de particules présente à l'écran
	for e := myList.Front(); e != nil; e = e.Next() {
		maParticule = e.Value.(*Particle)
		maParticule.PositionX = maParticule.PositionX + maParticule.SpeedX
		maParticule.PositionY = maParticule.PositionY - maParticule.SpeedY
		maParticule.Rotation = maParticule.Rotation - float64(rand.Intn(12)/100)
	}

	//Partie permettant de gérer les nombres flottant du spawnRate (ex: si le SpawnRate = 0.5 une particule sera ajouté tous les deux appels à createNParticles)
	var a float64 = config.General.SpawnRate
	nb = nb + a - float64(int(a))
	if nb > 1 {
		nb = nb - 1
		a = a + 1
	}

	createNParticles(int(a), myList)
}
