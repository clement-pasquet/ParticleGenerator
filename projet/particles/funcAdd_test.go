package particles


import (
	"project-particles/config"
	"testing"
	"container/list"
	// "time"
	// "math/rand"
)



func Test_setColor(t *testing.T) {
	config.General.Flag = 1
	var maParticule Particle = Particle{
		PositionX: float64(1),
		PositionY: float64(1),
		Rotation:0.5,
		ScaleX:    config.General.ScaleX, ScaleY: config.General.ScaleY, //Partie à remplacer
		ColorRed: config.General.ColorRed / 255, ColorGreen: config.General.ColorGreen / 255, ColorBlue: config.General.ColorBlue / 255,
		Opacity: config.General.Opacity,
		SpeedX:  1,
		SpeedY:  1,
		LifeSpan:1,
	}
	setColor(&maParticule)
	if maParticule.PositionX <float64(config.General.WindowSizeX)/3{
		if maParticule.ColorRed != 0{
			t.Errorf("ATTENTION, LA COULEUR GAUCHE N'EST PAS BLEU")
		}
		if maParticule.ColorGreen != 0 {
			t.Errorf("ATTENTION, LA COULEUR GAUCHE N'EST PAS BLEU")
			
		}
		if maParticule.ColorBlue != 255 {
			t.Errorf("ATTENTION, LA COULEUR GAUCHE N'EST PAS BLEU")
			
		}
	}

	if maParticule.PositionX <(float64(config.General.WindowSizeX)/3)*2{
		if maParticule.ColorRed != 255{
			t.Errorf("ATTENTION, LA COULEUR DU MILIEU N'EST PAS BLANC")
		}
		if maParticule.ColorGreen != 255 {
			t.Errorf("ATTENTION, LA COULEUR DU MILIEU N'EST PAS BLANC")
			
		}
		if maParticule.ColorBlue != 255 {
			t.Errorf("ATTENTION, LA COULEUR DU MILIEU N'EST PAS BLANC")
			
		}
	}

	if maParticule.PositionX <float64(config.General.WindowSizeX){
		if maParticule.ColorRed != 255{
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
	myParticule,ok := createNParticles(1,l).Front().Value.(*Particle)
	if !ok{
		t.Error("La particule n'as pas été bien créé !")
	}
	if config.General.RandomSpawn == true && float64(myParticule.LifeSpan)/60 >= config.General.LifeSpanMax+1 {
		t.Error("ATTENTION, IL EXISTE AU MOINS UNE PARTICULE PLUS VIEUX QUE", config.General.LifeSpanMax,"SECONDES")
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
