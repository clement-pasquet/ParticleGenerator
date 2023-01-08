package particles

import (
	"container/list"
	"testing"
	// "time"
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
