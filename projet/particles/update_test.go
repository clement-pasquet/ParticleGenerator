package particles

import (
	"container/list"
	"project-particles/config"
	"testing"
)

func Test_Update(t *testing.T) { //test permettant de vérifier si les particules bougent au lancement d'update
	var l *list.List = list.New()
	config.General.Velocity = 3
	createNParticles(1, l)
	var s *System = &System{Content: l}
	var particule *Particle = l.Front().Value.(*Particle)
	particule.SpeedX = 5
	particule.SpeedY = 4
	var posX float64 = particule.PositionX
	var posY float64 = particule.PositionY
	s.Update()
	if s.Content.Front().Value.(*Particle).PositionX == posX || s.Content.Front().Value.(*Particle).PositionY == posY {
		t.Errorf("La position de la particule n'as pas bougé")
	}

}
