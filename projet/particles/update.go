package particles

import (
	"container/list"
	"fmt"
)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
func (s *System) Update() {
	var ele *Particle
	var myList *list.List = s.Content
	fmt.Println(myList.Len())
	for e := myList.Front(); e != nil; e = e.Next() {
		ele = e.Value.(*Particle)
		ele.PositionX = ele.PositionX + ele.SpeedX
		ele.PositionY = ele.PositionY + ele.SpeedY
	}

}
