package particles

import (
	"container/list"
	"project-particles/config"

)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
var nb float64


//Idée :Recycler les particules : Clément(Ibrahim)
func (s *System) Update() {
	
	var maParticule *Particle 
	var myList *list.List = s.Content // Ma Liste de particules présent à l'écran
	for e := myList.Front(); e != nil; e = e.Next() {
		maParticule = e.Value.(*Particle)
		maParticule.PositionX = maParticule.PositionX + maParticule.SpeedX 
		maParticule.PositionY = maParticule.PositionY + maParticule.SpeedY + float64(maParticule.LifeSpan)*0.05 
		maParticule.LifeSpan++ //Augmente le compteur de durée de vie d'1
		maParticule.Rotation = maParticule.Rotation-0
		setColor(maParticule)
		LifeSpanIsTooAged(e,maParticule,myList)
	}	
	//Partie permettant de gérer les nombres flottant du spawnRate
	var a float64 = config.General.SpawnRate
	nb = nb + a - float64(int(a))
	if nb > 1 {
		nb = nb - 1
		a = a + 1
	}
	createNParticles(int(a), myList)

}