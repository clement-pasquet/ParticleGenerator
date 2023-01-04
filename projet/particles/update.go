package particles

import (
	"container/list"
	"math/rand"
	"project-particles/config"

	"time"
)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
var nb float64


//Idée :Recycler les particules : Clément(Ibrahim)
func (s *System) Update() {/*
	var c1 []float64= []float64{222, 18, 18}
	var c2 []float64=[]float64{219, 134, 15}
	var c3 []float64=[]float64{231, 242, 12}
	var color [][]float64 = [][]float64{c1,c2,c3}*/
	var maParticule *Particle 
	var myList *list.List = s.Content // Ma Liste de particules présent à l'écran
	for e := myList.Front(); e != nil; e = e.Next() {
		maParticule = e.Value.(*Particle)

		maParticule.PositionX = maParticule.PositionX + maParticule.SpeedX 
		maParticule.PositionY = maParticule.PositionY + maParticule.SpeedY + float64(maParticule.LifeSpan)*0.05
		maParticule.LifeSpan = maParticule.LifeSpan +1 //Augmente le compteur de durée de vie d'1
		maParticule.Rotation = maParticule.Rotation-5

		if maParticule.PositionX <427{
			maParticule.ColorRed=0
			maParticule.ColorGreen=0
			maParticule.ColorBlue =255

		}else if maParticule.PositionX <853{
			maParticule.ColorRed=255
			maParticule.ColorGreen=255
			maParticule.ColorBlue =255

		}else{
			maParticule.ColorRed=255
			maParticule.ColorGreen=0
			maParticule.ColorBlue =0
			
			//52, 235, 94
		}
/*
		couleur := rand.Intn(3)

		if couleur == 0 && maParticule.ColorRed <= 1{
			maParticule.ColorRed = maParticule.ColorRed+float64(rand.Intn(100))/10000
		} else {
			maParticule.ColorRed = rand.Float64()
		}
		if couleur == 1 && maParticule.ColorGreen <=1 {
			maParticule.ColorGreen = maParticule.ColorGreen+float64(rand.Intn(100))/10000
		}else {
			maParticule.ColorGreen = rand.Float64()
		}
		if couleur == 2 && maParticule.ColorBlue <= 1 {
			maParticule.ColorBlue = maParticule.ColorBlue+float64(rand.Intn(100))/10000
		}else {
			maParticule.ColorBlue = rand.Float64()
		}
		*/

		
		if config.General.RandomSpawn == true{
			if maParticule.LifeSpan/60 >= 2{ //Supprime la particule sur elle est vieille de plus de 2 secondes
				myList.Remove(e)
				maParticule.Opacity=0
				maParticule.PositionX = 2000
				
				
			}
		}
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

func createNParticles(nb int, l *list.List) *list.List { 
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < nb; i++ {
		var posX float64 = float64(config.General.SpawnX)
		var posY float64 = float64(config.General.SpawnY)
		var speedX float64 = rand.Float64() * config.General.Velocity
		var speedY float64 = rand.Float64() * config.General.Velocity
		var signe []int = []int{-1, 1}
		speedX = speedX * float64(signe[rand.Intn(2)])
		speedY = speedY * -1

		if config.General.RandomSpawn {
			posX = float64(rand.Intn(config.General.WindowSizeX))
			posY = float64(rand.Intn(config.General.WindowSizeY))

		}
		l.PushFront(&Particle{
			PositionX: float64(posX),
			PositionY: float64(posY),
			ScaleX:    config.General.ScaleX, ScaleY: config.General.ScaleY, //Partie à remplacer
			ColorRed: config.General.ColorRed / 255, ColorGreen: config.General.ColorGreen / 255, ColorBlue: config.General.ColorBlue / 255,
			Opacity: config.General.Opacity,
			SpeedX:  speedX,
			SpeedY:  speedY,
		})
	}
	
	return l
}
