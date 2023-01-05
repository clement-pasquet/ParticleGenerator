package particles

import (
	"container/list"
	"fmt"
	"project-particles/config"
	"testing"
	// "time"
	// "math/rand"
	// "project-particles"
)

func Test_setColor(t *testing.T) {
	config.General.Flag = 1
	var maParticule Particle = Particle{
		PositionX: float64(1),
		PositionY: float64(1),
		Rotation:  0.5,
		ScaleX:    config.General.ScaleX, ScaleY: config.General.ScaleY, //Partie à remplacer
		ColorRed: config.General.ColorRed / 255, ColorGreen: config.General.ColorGreen / 255, ColorBlue: config.General.ColorBlue / 255,
		Opacity:  config.General.Opacity,
		SpeedX:   1,
		SpeedY:   1,
		LifeSpan: 1,
	}
	setColor(&maParticule)
	if maParticule.PositionX < float64(config.General.WindowSizeX)/3 {
		if maParticule.ColorRed != 0 {
			t.Errorf("ATTENTION, LA COULEUR GAUCHE N'EST PAS BLEU")
		}
		if maParticule.ColorGreen != 0 {
			t.Errorf("ATTENTION, LA COULEUR GAUCHE N'EST PAS BLEU")

		}
		if maParticule.ColorBlue != 255 {
			t.Errorf("ATTENTION, LA COULEUR GAUCHE N'EST PAS BLEU")

		}
	}

	if maParticule.PositionX < (float64(config.General.WindowSizeX)/3)*2 {
		if maParticule.ColorRed != 255 {
			t.Errorf("ATTENTION, LA COULEUR DU MILIEU N'EST PAS BLANC")
		}
		if maParticule.ColorGreen != 255 {
			t.Errorf("ATTENTION, LA COULEUR DU MILIEU N'EST PAS BLANC")

		}
		if maParticule.ColorBlue != 255 {
			t.Errorf("ATTENTION, LA COULEUR DU MILIEU N'EST PAS BLANC")

		}
	}

	if maParticule.PositionX < float64(config.General.WindowSizeX) {
		if maParticule.ColorRed != 255 {
			t.Errorf("ATTENTION, LA COULEUR DU MILIEU N'EST PAS BLANC")
		}
		if maParticule.ColorGreen != 0 {
			t.Errorf("ATTENTION, LA COULEUR DU MILIEU N'EST PAS BLANC")

		}
		if maParticule.ColorBlue != 0 {
			t.Errorf("ATTENTION, LA COULEUR DU MILIEU N'EST PAS BLANC")

		}
	}

}

func Test_LifeSpanIsTooAged(t *testing.T) {
	l := list.New()
	myParticule, ok := createNParticles(1, l).Front().Value.(*Particle)
	if !ok {
		t.Error("La particule n'as pas été bien créé !")
	}
	if config.General.RandomSpawn == true && float64(myParticule.LifeSpan)/60 >= config.General.LifeSpanMax+1 {
		t.Error("ATTENTION, IL EXISTE AU MOINS UNE PARTICULE PLUS VIEUX QUE", config.General.LifeSpanMax, "SECONDES")
	}

}

func Test_createNParticles(t *testing.T) {
	var a float64 = 100
	var myList *list.List = list.New()
	var myList2 *list.List = list.New()
	myList2 = createNParticles(int(a), myList2)
	fmt.Println(myList, myList2)

	if myList.Len() == myList2.Len() {
		t.Errorf("ATTENTION, createNParticles ne crée pas le bon nombre de particules")
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
