package particles

import (
	"container/list"
	"project-particles/config"
	"testing"
	// "time"
	 "math/rand"
	// "project-particles"
)

func Test_createNParticles(t *testing.T) {
	var a float64 = 100
	var myList *list.List = list.New()
	var myList2 *list.List = list.New()
	myList2 = createNParticles(int(a), myList2)

	if myList.Len() == myList2.Len() {
		t.Errorf("ATTENTION, createNParticles ne crée pas le bon nombre de particules")
	}

}

func Test_puissance2(t *testing.T) { //Test d'une fonction qui calcule le carré d'une valeur a
	if puissance2(2) != 4 {
		t.Errorf("LA PUISSANCE DE 2 DE 2 EST 4")
	}
	if puissance2(0) != 0 {
		t.Errorf("LA PUISSANCE DE 2 DE 0 EST 0")
	}
	if puissance2(-41) != 1681 {
		t.Errorf("LA PUISSANCE DE 2 DE -41 EST 1681")
	}

}

func Test_RandomSpawn(t *testing.T) {
	for i:=0;i<1000;i++ {
		var posX float64 = float64(rand.Intn(500))
		var posY float64 = float64(rand.Intn(500))
		posX,posY = RandomSpawnFunc(posX, posY)

		if posX != float64(config.General.WindowSizeX/2) || posY != float64(config.General.WindowSizeY/2) {
			t.Errorf("RandomSpawn ne fonctionne pas")
		}
	}
}

/*
func TestPlusieurs(t *testing.T) {
	var x, y, z int
	copyptr(3, []*int{&x, &y, &z})
	if x != 3 {
		t.Errorf("Echec de la copie dans x")
	}
	if y != 3 {
		t.Errorf("Echec de la copie dans y")
	}
	if z != 3 {
		t.Errorf("Echec de la copie dans z")
	}
}*/
