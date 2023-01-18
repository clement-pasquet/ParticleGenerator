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

// Test les puissances de 2
func Test_puissance2(t *testing.T) {
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

// Test la fonction setcolor qui s'occupe entre autre de générer des drapeaux
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
	if maParticule.PositionX < float64(config.General.WindowSizeX)/3 { //Vérifie si le premier tiers gauche de l'écran est bleu ou non (pour faire le drapeau francais)
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

	if maParticule.PositionX < (float64(config.General.WindowSizeX)/3)*2 { //Vérifie si le deuxième tiers gauche de l'écran est blanc ou non
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

	if maParticule.PositionX < float64(config.General.WindowSizeX) { //Vérifie si le troisième tiers gauche de l'écran est rouge ou non
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

// Test la durée de vie d'une particule pour vérifier qu'elle est bien supprimé ou non
func Test_LifeSpanIsTooAged(t *testing.T) { //V
	var myList *list.List = list.New()
	s := &System{Content: myList}
	createNParticles(5, s)
	var e *list.Element = s.Content.Front()
	var myParticule *Particle = e.Value.(*Particle)
	myParticule.LifeSpan = 250000
	LifeSpanIsTooAged(e, myParticule, s)

	var myList2 *list.List = list.New()
	s2 := &System{Content: myList2}
	createNParticles(5, s2)
	var e2 *list.Element = s2.Content.Front()
	var myParticule2 *Particle = e2.Value.(*Particle)
	myParticule2.LifeSpan = 250000

	if myParticule.IsInLife { //regarde si la variable "IsInLife" a bien été mis à false, autrement dit, que la particule est bien considérée morte
		t.Error("La particule n'as pas été bien créé !")
	}
	if myParticule.Opacity != 0 { //regarde si l'opacité d'une particule en dehors de l'écran est bien a 0
		t.Errorf("La Particule crée a beau être en dehors de l'écran, son opacité n'est pas à 0")
	}
	if myParticule == myParticule2 { //Regarde si une particule soumise à IsOutOfView et une qui ne l'est pas sont identiques, si c'est le cas, c'est surement que la fonction de marche pas bien
		t.Errorf("IsOutOfView ne met pas la particule à la fin de la liste ou ne fonctionne pas")
	}
	if e == e2 { //même principe avec des particules
		t.Errorf("IsOutOfView ne met pas la particule à la fin de la liste ou ne fonctionne pas")
	}

}

func Test_createNParticles(t *testing.T) {
	var a float64 = 100
	var myList *list.List = list.New()
	var myList2 *list.List = list.New()
	s := &System{Content: myList}
	createNParticles(int(a), s)

	if myList.Len() == myList2.Len() { //Vérifie si CreateNParticules crée belle et bien une liste remplie de particule
		t.Errorf("ATTENTION, createNParticles ne crée pas le bon nombre de particules")
	}
	if myList.Len() != int(a) || myList2.Len() == int(a) { //Vérifie si createNParticles crée un nombre correct de particule
		t.Errorf("ATTENTION, createNParticles ne crée pas le bon nombre de particules")
	}

}

// Vérifie si une particule qui est en dehors de l'écran est belle et bien supprimée
func Test_IsOutOfView(t *testing.T) {
	var myList *list.List = list.New()
	s := &System{Content: myList}
	createNParticles(10, s)
	var e *list.Element = s.Content.Front()
	var myParticule *Particle = e.Value.(*Particle)
	myParticule.PositionX = float64(config.General.WindowSizeX) + config.General.Margin + 500
	myParticule.PositionY = float64(config.General.WindowSizeY) + config.General.Margin + 500

	eOld := s.Content.Front()

	IsOutOfView(e, myParticule, s)

	eNew := s.Content.Front()

	if myParticule.Opacity != 0 { //regarde si l'opacité d'une particule en dehors de l'écran est bien a 0
		t.Errorf("La Particule crée a beau être en dehors de l'écran, son opacité n'est pas à 0")
	}
	if myParticule.IsInLife { //vérifie si la particule a la variable "inlife" à true, si c'est le cas la fonction ne marche pas
		t.Errorf("La Variable IsInLife indique que la particule n'est pas morte")
	}
	if eOld == eNew { //Check si une particule a bien été mise au "fond" de la liste, autrement dit qu'elle est considérée morte
		t.Errorf("IsOutOfView ne met pas la particule à la fin de la liste ou ne fonctionne pas")
	}
}

// Vérifie si une particule qui touche un mur change le signe de la vitesse ou non
func Test_collisionWall(t *testing.T) {
	fmt.Println("")
	var myList *list.List = list.New()
	s := &System{Content: myList}
	createNParticles(3, s)
	var e *list.Element = s.Content.Front()
	var myParticule *Particle = e.Value.(*Particle)

	myParticule.PositionX = -100
	myParticule.SpeedX = -100

	fmt.Print(myParticule, e, myList, s)

	if myParticule.PositionX == 100 || myParticule.SpeedX == -95 {
		t.Errorf("La particule ne change pas correctement de position/vitesse en fc de sa position !")
	}

}

func Test_update(t *testing.T) {
	var myList *list.List = list.New() //Crée une premiere liste avec son système, son élément et sa particule.
	s := &System{Content: myList}      //Ils sont faits pour être modifiés par la fonction update
	var e *list.Element = s.Content.Front()
	createNParticles(2, s)
	var myParticule *Particle = s.Content.Front().Value.(*Particle)

	s.Update()

	var myList2 *list.List = list.New() //Crée une deuxième liste avec son  système, son élément et sa particule
	s2 := &System{Content: myList2}     //ils ne sont pas faits pour être mis à jour et serviront d'éléments de comparaisons avec la 1e liste
	var e2 *list.Element = s.Content.Front()
	createNParticles(3, s2)
	var myParticule2 *Particle = s2.Content.Front().Value.(*Particle)

	if e == e2 { //Vérifie si les deux éléments ont la même adresse mémoire
		t.Errorf("La fonction update n'a pas changé les arguments d'une particule")
	}
	if myList == nil && myList2 == nil { //Regarde si les deux listes sont nil, si c'était le cas, ca voudrait dire que update n'a pas bien fonctionné
		t.Errorf("Les listes de particules utilisés sont vides")
	}
	if myParticule == myParticule2 { //Vérifie si les deux particules sont les mêmes, si c'était le cas, ca voudrait dire que update n'a pas bien fonctionné
		t.Errorf("La fonction update n'a pas changé les arguments d'une particule")
	}
}

func Test_getImg(t *testing.T) {

}
