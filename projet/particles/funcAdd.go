package particles

import (
	"container/list"
	"fmt"
	"math"
	"math/rand"
	"project-particles/config"
	"time"
)

var gigachadLs [][][]float64

func puissance2(a float64) float64 {
	return a * a
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

		if config.General.GeneratorShape == "rectangle" {
			config.General.RandomSpawn = true
			config.General.Gravity = 0.03
			config.General.SpawnRate = 20
			config.General.ColorBlue = 250 - rand.Float64()*30
			config.General.ColorRed = 255 - rand.Float64()*30
			config.General.ColorGreen = 255 - rand.Float64()*30
			config.General.Opacity = rand.Float64()*0.8 + 0.2

			posY = float64((config.General.WindowSizeY / 2) - 100 - rand.Intn(50))
			posX = float64((config.General.WindowSizeX / 2) - 100 - rand.Intn(50))
		}
		if config.General.GeneratorShape == "carre" {
			config.General.RandomSpawn = false
			config.General.Gravity = 0.08
			config.General.SpawnRate = 2
			config.General.Velocity = 3
			config.General.ScaleX = 0.05
			config.General.ScaleY = 0.05
			config.General.ColorRed = rand.Float64() * 255
			config.General.ColorGreen = rand.Float64() * 255
			config.General.ColorBlue = rand.Float64() * 255

			posX = float64((config.General.WindowSizeX / 2) - 100 + rand.Intn(200))
			posY = float64((config.General.WindowSizeY / 2) - 100 + rand.Intn(200))

		}
		if config.General.GeneratorShape == "rond" {
			config.General.ParticleImage = "assets/particle3.png"
			config.General.RandomSpawn = false
			config.General.SpawnRate = 30
			config.General.ColorRed = 200 + float64(rand.Intn(50))
			config.General.ColorGreen = 200 + float64(rand.Intn(50))
			config.General.ColorBlue = 200 + float64(rand.Intn(50))
			if rand.Intn(2) == 0 {
				config.General.Velocity = float64(rand.Intn(150))
			}
			if rand.Intn(2) == 1 {
				config.General.Velocity = -float64(rand.Intn(150))
			}

			fmt.Println("C'EST ROND")

			posX = float64(config.General.WindowSizeX) / 2
			posY = float64(config.General.WindowSizeY) / 2

		}

		if config.General.RandomSpawn {
			posX = float64(rand.Intn(config.General.WindowSizeX))
			posY = float64(rand.Intn(config.General.WindowSizeY))
		}

		if particulesMortes.Len() > 0 {

			var pMortesFront *list.Element = particulesMortes.Front()
			particulesMortes.Remove(pMortesFront)

			var maParticule *Particle = pMortesFront.Value.(*Particle)

			maParticule.PositionX = float64(posX)
			maParticule.PositionY = float64(posY)
			maParticule.ScaleX = config.General.ScaleX
			maParticule.ScaleY = config.General.ScaleY //Partie à remplacer
			maParticule.ColorRed = config.General.ColorRed / 255
			maParticule.ColorGreen = config.General.ColorGreen / 255
			maParticule.ColorBlue = config.General.ColorBlue / 255
			maParticule.Opacity = config.General.Opacity
			maParticule.SpeedX = speedX
			maParticule.SpeedY = speedY
			maParticule.LifeSpan = 0
			l.PushFront(maParticule)

		} else {
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
	}

	return l
}

func IsOutOfView(e *list.Element, p *Particle, l *list.List) bool {
	if p.PositionX < -config.General.Margin || p.PositionX > float64(config.General.WindowSizeX)+config.General.Margin || p.PositionY > float64(config.General.WindowSizeY)+config.General.Margin {
		setInvisible(p)
		go l.Remove(e)
		return true
	}
	return false
}

func LifeSpanIsTooAged(ele *list.Element, p *Particle, l *list.List) bool {
	if config.General.RandomSpawn {

		if config.General.LifeSpanMax <= float64(p.LifeSpan)/60 { //Supprime la particule si  elle est vieille de plus de 2 secondes
			setInvisible(p) //met la particule invisible et en dehors de l'écran
			go l.Remove(ele)
			return true
		}
	}
	return false
}

func setInvisible(p *Particle) {
	p.Opacity = 0
	p.PositionX = float64(config.General.WindowSizeX) + 500
}

func setColor(p *Particle) {

	if config.General.CustomImageBool {
		config.General.SpawnRate = 1000
		config.General.ColorBlue = 0
		config.General.ColorGreen = 0
		config.General.ColorRed = 0

		if len(gigachadLs) < 1 {
			//src := "assets/" + config.General.CustomImageSRC
			gigachadLs = getListColor(config.General.CustomImageSRC)
		}
		var lenX, lenY int = len(gigachadLs[0]), len(gigachadLs)

		if p.PositionX > 0 && int(p.PositionX) < lenX && p.PositionY > 0 && int(p.PositionY) < lenY {

			var color []float64 = gigachadLs[int(p.PositionY)][int(p.PositionX)]

			var r, g, b, _ uint32 = uint32(color[0]), uint32(color[1]), uint32(color[2]), uint32(color[3])
			p.ColorRed = float64(r) / 65535
			p.ColorGreen = float64(g) / 65535
			p.ColorBlue = float64(b) / 65535
		}

	}
	var param int = config.General.Flag

	//color
	var b1 []float64 = []float64{0, 0, 255}
	var b2 []float64 = []float64{0, 43, 127}
	var b3 []float64 = []float64{0, 91, 187}
	var blanc []float64 = []float64{255, 255, 255}
	var r1 []float64 = []float64{255, 0, 0}
	var vert1 []float64 = []float64{0, 255, 0}
	var jaune1 []float64 = []float64{255, 255, 0}
	var orange1 []float64 = []float64{255, 229, 2}
	var orange2 []float64 = []float64{255, 160, 10}
	var noir1 []float64 = []float64{75, 75, 75}
	//Country
	var France [][]float64 = [][]float64{b1, blanc, r1}
	var Italie [][]float64 = [][]float64{vert1, blanc, r1}
	var Irelande [][]float64 = [][]float64{vert1, blanc, orange1}
	var Belgique [][]float64 = [][]float64{noir1, jaune1, r1}
	var Roumanie [][]float64 = [][]float64{b2, orange1, r1}
	var Inde [][]float64 = [][]float64{orange2, blanc, vert1}
	var Allemagne [][]float64 = [][]float64{noir1, r1, jaune1}
	var Japon [][]float64 = [][]float64{blanc, r1}
	var Russie [][]float64 = [][]float64{blanc, b1, r1}
	var Ukraine [][]float64 = [][]float64{b3, jaune1}

	switch param {
	case 1:
		//FRANCE
		if p.PositionX < 427 {
			p.ColorRed = France[0][0] / 255
			p.ColorGreen = France[0][1] / 255
			p.ColorBlue = France[0][2] / 255
		} else if p.PositionX < 853 {
			p.ColorRed = France[1][0] / 255
			p.ColorGreen = France[1][1] / 255
			p.ColorBlue = France[1][2] / 255
		} else {
			p.ColorRed = France[2][0] / 255
			p.ColorGreen = France[2][1] / 255
			p.ColorBlue = France[2][2] / 255
		}

	case 2:
		//ITALIE
		if p.PositionX < 427 {
			p.ColorRed = Italie[0][0] / 255
			p.ColorGreen = Italie[0][1] / 255
			p.ColorBlue = Italie[0][2] / 255
		} else if p.PositionX < 853 {
			p.ColorRed = Italie[1][0] / 255
			p.ColorGreen = Italie[1][1] / 255
			p.ColorBlue = Italie[1][2] / 255
		} else {
			p.ColorRed = Italie[2][0] / 255
			p.ColorGreen = Italie[2][1] / 255
			p.ColorBlue = Italie[2][2] / 255
		}

	case 3:
		//Irelande
		if p.PositionX < 427 {
			p.ColorRed = Irelande[0][0] / 255
			p.ColorGreen = Irelande[0][1] / 255
			p.ColorBlue = Irelande[0][2] / 255
		} else if p.PositionX < 853 {
			p.ColorRed = Irelande[1][0] / 255
			p.ColorGreen = Irelande[1][1] / 255
			p.ColorBlue = Irelande[1][2] / 255
		} else {
			p.ColorRed = Irelande[2][0] / 255
			p.ColorGreen = Irelande[2][1] / 255
			p.ColorBlue = Irelande[2][2] / 255
		}

	case 4:
		//Belgique
		if p.PositionX < 427 {
			p.ColorRed = Belgique[0][0] / 255
			p.ColorGreen = Belgique[0][1] / 255
			p.ColorBlue = Belgique[0][2] / 255
		} else if p.PositionX < 853 {
			p.ColorRed = Belgique[1][0] / 255
			p.ColorGreen = Belgique[1][1] / 255
			p.ColorBlue = Belgique[1][2] / 255
		} else {
			p.ColorRed = Belgique[2][0] / 255
			p.ColorGreen = Belgique[2][1] / 255
			p.ColorBlue = Belgique[2][2] / 255
		}
	case 5:
		//Roumanie
		if p.PositionX < 427 {
			p.ColorRed = Roumanie[0][0] / 255
			p.ColorGreen = Roumanie[0][1] / 255
			p.ColorBlue = Roumanie[0][2] / 255
		} else if p.PositionX < 853 {
			p.ColorRed = Roumanie[1][0] / 255
			p.ColorGreen = Roumanie[1][1] / 255
			p.ColorBlue = Roumanie[1][2] / 255
		} else {
			p.ColorRed = Roumanie[2][0] / 255
			p.ColorGreen = Roumanie[2][1] / 255
			p.ColorBlue = Roumanie[2][2] / 255
		}
	case 6:
		//Inde
		if p.PositionY < 240 {
			p.ColorRed = Inde[0][0] / 255
			p.ColorGreen = Inde[0][1] / 255
			p.ColorBlue = Inde[0][2] / 255
		} else if p.PositionY < 480 {
			p.ColorRed = Inde[1][0] / 255
			p.ColorGreen = Inde[1][1] / 255
			p.ColorBlue = Inde[1][2] / 255
		} else {
			p.ColorRed = Inde[2][0] / 255
			p.ColorGreen = Inde[2][1] / 255
			p.ColorBlue = Inde[2][2] / 255
		}
		var a float64 = math.Sqrt(puissance2((p.PositionX - float64(config.General.WindowSizeX)/2)) + puissance2((p.PositionY - float64(config.General.WindowSizeY)/2)))

		if a > 80 && a < 100 {
			p.ColorRed = b1[0] / 255
			p.ColorGreen = b1[1] / 255
			p.ColorBlue = b1[2] / 255
		}

		if a < 22 {
			p.ColorRed = b1[0] / 255
			p.ColorGreen = b1[1] / 255
			p.ColorBlue = b1[2] / 255
		}

	case 7:
		//Allemagne
		if p.PositionY < 240 {
			p.ColorRed = Allemagne[0][0] / 255
			p.ColorGreen = Allemagne[0][1] / 255
			p.ColorBlue = Allemagne[0][2] / 255
		} else if p.PositionY < 480 {
			p.ColorRed = Allemagne[1][0] / 255
			p.ColorGreen = Allemagne[1][1] / 255
			p.ColorBlue = Allemagne[1][2] / 255
		} else {
			p.ColorRed = Allemagne[2][0] / 255
			p.ColorGreen = Allemagne[2][1] / 255
			p.ColorBlue = Allemagne[2][2] / 255
		}
	case 8:
		//JAPON
		var a float64 = math.Sqrt(puissance2((p.PositionX - float64(config.General.WindowSizeX)/2)) + puissance2((p.PositionY - float64(config.General.WindowSizeY)/2)))
		p.ColorRed = Japon[0][0] / 255
		p.ColorGreen = Japon[0][1] / 255
		p.ColorBlue = Japon[0][2] / 255
		if a < 150 {
			p.ColorRed = Japon[1][0] / 255
			p.ColorGreen = Japon[1][1] / 255
			p.ColorBlue = Japon[1][2] / 255
		}
	case 9:
		//RUSSIE
		if p.PositionY < 240 {
			p.ColorRed = Russie[0][0] / 255
			p.ColorGreen = Russie[0][1] / 255
			p.ColorBlue = Russie[0][2] / 255
		} else if p.PositionY < 480 {
			p.ColorRed = Russie[1][0] / 255
			p.ColorGreen = Russie[1][1] / 255
			p.ColorBlue = Russie[1][2] / 255
		} else {
			p.ColorRed = Russie[2][0] / 255
			p.ColorGreen = Russie[2][1] / 255
			p.ColorBlue = Russie[2][2] / 255
		}

	case 10:
		//UKRAINE
		if p.PositionY < float64(config.General.WindowSizeY)/2 {
			p.ColorRed = Ukraine[0][0] / 255
			p.ColorGreen = Ukraine[0][1] / 255
			p.ColorBlue = Ukraine[0][2] / 255
		} else {
			p.ColorRed = Ukraine[1][0] / 255
			p.ColorGreen = Ukraine[1][1] / 255
			p.ColorBlue = Ukraine[1][2] / 255
		}
	}
}
