package particles

import (
	"container/list"
	"math"

	//"fmt"
	"strings"

	"math/rand"
	"project-particles/assets"
	"project-particles/config"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var ImageColorls [][][]float64
var KeyStringNew string

// Fais une puissance carré
func puissance2(a float64) float64 {
	return a * a
}

// Fonction qui crée un nb de particule dans un système ( avec recyclage )
func createNParticles(nb int, s *System) {

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < nb; i++ {
		var posX float64 = float64(config.General.SpawnX)
		var posY float64 = float64(config.General.SpawnY)

		Opacity := config.General.Opacity
		var speedX float64 = rand.Float64() * config.General.Velocity
		var speedY float64 = rand.Float64() * config.General.Velocity
		var signe []int = []int{-1, 1}
		speedX = speedX * float64(signe[rand.Intn(2)])
		speedY = speedY * float64(signe[rand.Intn(2)])
		var angleP float64
		if config.General.GeneratorShape == "triangle" {
			var angle float64 = float64(rand.Intn(360))
			angleP = (angle * math.Pi) / 180
			n := 3
			if n > 2 {
				mul_angle := 0
				for angle-float64((360/n)*mul_angle) > float64(((360 / n) / 2)) {
					mul_angle = mul_angle + 1
				}
				speedX = speedX / math.Cos(((angle-float64(360/n)*float64(mul_angle))*math.Pi)/180)
				speedY = speedY / math.Cos(((angle-float64(360/n)*float64(mul_angle))*math.Pi)/180)

			}
		} else if config.General.GeneratorShape == "rectangle" {
			config.General.RandomSpawn = true
			config.General.Gravity = 0.03
			config.General.SpawnRate = 20
			config.General.ColorBlue = 250 - rand.Float64()*30
			config.General.ColorRed = 255 - rand.Float64()*30
			config.General.ColorGreen = 255 - rand.Float64()*30
			config.General.Opacity = rand.Float64()*0.8 + 0.2

			posY = float64((config.General.WindowSizeY / 2) - 100 - rand.Intn(50))
			posX = float64((config.General.WindowSizeX / 2) - 100 - rand.Intn(50))
		} else if config.General.GeneratorShape == "carre" {
			angle := float64(rand.Intn(180))
			nb_seg := 3
			mul_angle := 0
			for angle-float64((360/nb_seg)*mul_angle) > float64(((360 / nb_seg) / 2)) {
				mul_angle = mul_angle + 1
			}
			posX = float64(config.General.SpawnX) + (rand.Float64()*config.General.Velocity)/math.Cos(((angle-float64(360/nb_seg)*float64(mul_angle))*math.Pi)/180)
			posY = float64(config.General.SpawnY) + (rand.Float64()*config.General.Velocity)/math.Sin(((angle-float64(360/nb_seg)*float64(mul_angle))*math.Pi)/180)

			config.General.RandomSpawn = false
			config.General.ColorRed = rand.Float64() * 255
			config.General.ColorGreen = rand.Float64() * 255
			config.General.ColorBlue = rand.Float64() * 255
			speedX = rand.Float64() * config.General.Velocity * math.Cos(angle)
			speedY = rand.Float64() * config.General.Velocity * math.Sin(angle)
			//posX = float64(config.General.SpawnX)
			//posY = float64(config.General.SpawnY)

		} else if config.General.GeneratorShape == "eclipse" {
			Opacity = 0
			posX = float64(config.General.SpawnX)
			posY = float64(config.General.SpawnY)

		} else if config.General.GeneratorShape == "cercle" {
			rand.Seed(time.Now().UnixNano())
			angle := (rand.Float64() * 12 * math.Pi) / (rand.Float64() * 6)
			posX = float64(config.General.SpawnX + rand.Intn(int(config.General.SizeShape)))
			posY = float64(config.General.SpawnY) + float64(rand.Intn(int(config.General.SizeShape)))
			// posX = float64(config.General.SpawnX) + config.General.SizeShape*math.Cos(angle)
			// posY = float64(config.General.SpawnY) + config.General.SizeShape*math.Sin(angle)
			speedX = rand.Float64() * config.General.Velocity * math.Cos(angle)
			speedY = rand.Float64() * config.General.Velocity * math.Sin(angle)
			config.General.ColorRed = rand.Float64() * 255
			config.General.ColorGreen = rand.Float64() * 255
			config.General.ColorBlue = rand.Float64() * 255

		} else {
			var speedX float64 = rand.Float64() * config.General.Velocity
			var speedY float64 = rand.Float64() * config.General.Velocity
			var signe []int = []int{-1, 1}
			speedX = speedX * float64(signe[rand.Intn(2)])
			speedY = speedY * float64(signe[rand.Intn(2)])
		}

		if config.General.RandomSpawn {
			posX = float64(rand.Intn(config.General.WindowSizeX))
			posY = float64(rand.Intn(config.General.WindowSizeY))

		}

		if s.NbParticulesMortes > 0 {

			s.NbParticulesMortes--
			var particuleDerniereMorte *list.Element = s.Content.Back()

			var maParticule *Particle = particuleDerniereMorte.Value.(*Particle)

			maParticule.PositionX = float64(posX)
			maParticule.PositionY = float64(posY)
			maParticule.ScaleX = config.General.ScaleX
			maParticule.ScaleY = config.General.ScaleY
			maParticule.ColorRed = config.General.ColorRed / 255
			maParticule.ColorGreen = config.General.ColorGreen / 255
			maParticule.ColorBlue = config.General.ColorBlue / 255
			maParticule.Opacity = Opacity
			maParticule.SpeedX = speedX
			maParticule.SpeedY = speedY
			maParticule.LifeSpan = 0
			maParticule.IsInLife = true
			maParticule.Angle = angleP
			s.Content.MoveToFront(particuleDerniereMorte)

		} else {
			p := &Particle{
				PositionX: float64(posX),
				PositionY: float64(posY),
				ScaleX:    config.General.ScaleX, ScaleY: config.General.ScaleY,
				ColorRed: config.General.ColorRed / 255, ColorGreen: config.General.ColorGreen / 255, ColorBlue: config.General.ColorBlue / 255,
				Opacity:  Opacity,
				SpeedX:   speedX,
				SpeedY:   speedY,
				IsInLife: true,
				Angle:    angleP,
			}
			s.Content.PushFront(p)

		}
	}

}

// Permet de gérer la forme "eclipse" du GeneratorShape de config.json
func shapePropriete(p *Particle) {
	if config.General.GeneratorShape == "eclipse" {
		// if (p.PositionX <= float64(config.General.SpawnX) && p.PositionX <= float64(config.General.SpawnX)+50) || (p.PositionY <= float64(config.General.SpawnY) && p.PositionY <= float64(config.General.SpawnY)+50) {
		// 	p.Opacity = 0
		// } else {
		// 	p.Opacity = 1

		// }

		var distanceA float64 = math.Sqrt(puissance2((p.PositionX - float64(config.General.SpawnX))) + puissance2((p.PositionY - float64(config.General.SpawnY))))
		if distanceA < config.General.SizeShape*10 {
			p.Opacity = 0

		} else {
			if distanceA < config.General.SizeShape*10+(config.General.SizeShape*20)/100 {
				p.Opacity = p.Opacity - 0.1

			} else {
				p.Opacity = 1
			}
		}
	}

}

// Gère la collision des particules avec les murs
func collisionWall(p *Particle) {
	var x, y int = assets.ParticleImage.Size()
	//

	if p.PositionX <= 0 {
		if p.SpeedX < 0 {
			p.SpeedX = -p.SpeedX
			p.PositionX = p.PositionX + 5
		}
	}
	if (p.PositionX + float64(x)*config.General.ScaleX) > float64(config.General.WindowSizeX) {
		if p.SpeedX > 0 {
			p.SpeedX = -p.SpeedX
			p.PositionX = p.PositionX - 10

		}
	}
	if (p.PositionY + float64(y)*config.General.ScaleY) > float64(config.General.WindowSizeY) {
		if p.SpeedY > 0 {
			p.SpeedY = -p.SpeedY
			p.PositionY = p.PositionY - 5
		}
	}
	if p.PositionY <= 0 {
		if p.SpeedY < 0 {
			p.SpeedY = -p.SpeedY
			p.PositionY = p.PositionY + 5
		}
	}

}

// Gère la collision des particules entre elles
func collisionBetweenParticles(e *list.Element) {
	if e != nil {
		if e.Next() != nil {
			var sizeX, sizeY int = assets.ParticleImage.Size()
			var maParticule *Particle = e.Value.(*Particle)
			var maParticuleSuivante *Particle = e.Next().Value.(*Particle)
			if maParticule.PositionX <= maParticuleSuivante.PositionX && maParticule.PositionX+float64(sizeX)*config.General.ScaleX > maParticuleSuivante.PositionX {
				//Si en Haut
				if (maParticuleSuivante.PositionY + float64(sizeY)*config.General.ScaleY) == maParticule.PositionY {
					maParticule.SpeedY = maParticule.SpeedY * -1
					maParticuleSuivante.SpeedY = maParticuleSuivante.SpeedY * -1
				}
				//Si en bas
				if (maParticuleSuivante.PositionY) == maParticule.PositionY+float64(sizeY)*config.General.ScaleY {
					maParticule.SpeedY = maParticule.SpeedY * -1
					maParticuleSuivante.SpeedY = maParticuleSuivante.SpeedY * -1
				}
			}
			if maParticule.PositionY <= maParticuleSuivante.PositionY && maParticule.PositionY+float64(sizeY)*config.General.ScaleY > maParticuleSuivante.PositionY {
				//Si à Gauche
				if (maParticuleSuivante.PositionX + float64(sizeX)*config.General.ScaleX) == maParticule.PositionX {
					maParticule.SpeedX = maParticule.SpeedX * -1
					maParticuleSuivante.SpeedX = maParticuleSuivante.SpeedX * -1
				}
				//Si à Droite
				if (maParticuleSuivante.PositionX) == maParticule.PositionX+float64(sizeX)*config.General.ScaleX {
					maParticule.SpeedX = maParticule.SpeedX * -1
					maParticuleSuivante.SpeedX = maParticuleSuivante.SpeedX * -1
				}
			}
		}
	}

}

// Fonction setcolor qui s'occupe, entre autre, de générer des drapeaux.
func setColor(p *Particle) {

	if config.General.CustomImageBool {
		//config.General.SpawnRate = 1000
		config.General.ColorBlue = 0
		config.General.ColorGreen = 0
		config.General.ColorRed = 0

		if len(ImageColorls) < 1 {
			//src := "assets/" + config.General.CustomImageSRC
			ImageColorls = getListColor(config.General.CustomImageSRC)
		}
		var lenX, lenY int = len(ImageColorls[0]), len(ImageColorls)

		if p.PositionX > 0 && int(p.PositionX) < lenX && p.PositionY > 0 && int(p.PositionY) < lenY {

			var color []float64 = ImageColorls[int(p.PositionY)][int(p.PositionX)]

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

	var sizeX int = config.General.WindowSizeX
	var _ int = config.General.WindowSizeY

	switch param {
	case 1:
		//FRANCE
		if p.PositionX < float64(sizeX)/3 {
			p.ColorRed = France[0][0] / 255
			p.ColorGreen = France[0][1] / 255
			p.ColorBlue = France[0][2] / 255
		} else if p.PositionX < (float64(sizeX)/3)*2 {
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

// Supprime les particules une fois qu'elles sont en dehors de l'écran et de la margin
func IsOutOfView(e *list.Element, p *Particle, s *System) {
	if p.PositionX < -config.General.Margin || p.PositionX > float64(config.General.WindowSizeX)+config.General.Margin || p.PositionY > float64(config.General.WindowSizeY)+config.General.Margin {
		s.Content.MoveToBack(e)
		p.IsInLife = false
		setInvisible(p) //met la particule invisible et en dehors de l'écran
	}
}

// Supprime la particule si  elle est vieille de plus de X secondes
func LifeSpanIsTooAged(ele *list.Element, p *Particle, s *System) {
	if config.General.LifeSpanMax <= float64(p.LifeSpan)/60 {
		s.Content.MoveToBack(ele)
		p.IsInLife = false

		setInvisible(p) //met la particule invisible et en dehors de l'écran

	}
}

// Met l'opacité de la particule à 0 et la met en dehors de l'écran
func setInvisible(p *Particle) {
	p.Opacity = 0
	p.PositionX = float64(config.General.WindowSizeX) + 500
}

func KeyboardInput() {

	var key string
	if inpututil.IsKeyJustPressed(ebiten.KeyDigit0) || inpututil.IsKeyJustPressed(ebiten.Key0) || inpututil.IsKeyJustPressed(ebiten.KeyNumpad0) {
		key = "0"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDigit1) || inpututil.IsKeyJustPressed(ebiten.Key1) || inpututil.IsKeyJustPressed(ebiten.KeyNumpad1) {
		key = "1"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDigit2) || inpututil.IsKeyJustPressed(ebiten.Key2) || inpututil.IsKeyJustPressed(ebiten.KeyNumpad2) {
		key = "2"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDigit3) || inpututil.IsKeyJustPressed(ebiten.Key3) || inpututil.IsKeyJustPressed(ebiten.KeyNumpad3) {
		key = "3"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDigit4) || inpututil.IsKeyJustPressed(ebiten.Key4) || inpututil.IsKeyJustPressed(ebiten.KeyNumpad4) {
		key = "4"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDigit5) || inpututil.IsKeyJustPressed(ebiten.Key5) || inpututil.IsKeyJustPressed(ebiten.KeyNumpad5) {
		key = "5"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDigit6) || inpututil.IsKeyJustPressed(ebiten.Key6) || inpututil.IsKeyJustPressed(ebiten.KeyNumpad6) {
		key = "6"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDigit7) || inpututil.IsKeyJustPressed(ebiten.Key7) || inpututil.IsKeyJustPressed(ebiten.KeyNumpad7) {
		key = "7"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDigit8) || inpututil.IsKeyJustPressed(ebiten.Key8) || inpututil.IsKeyJustPressed(ebiten.KeyNumpad8) {
		key = "8"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDigit9) || inpututil.IsKeyJustPressed(ebiten.Key9) || inpututil.IsKeyJustPressed(ebiten.KeyNumpad9) {
		key = "9"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		key = "q"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyB) {
		key = "b"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		key = "c"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		key = "d"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		key = "e"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		key = "f"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyG) {
		key = "g"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		key = "h"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyI) {
		key = "i"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyJ) {
		key = "j"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyK) {
		key = "k"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyL) {
		key = "l"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyM) {
		key = "m"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyN) {
		key = "n"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyO) {
		key = "o"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		key = "p"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		key = "a"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		key = "r"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		key = "s"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyT) {
		key = "t"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyU) {
		key = "u"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyV) {
		key = "v"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		key = "w"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyX) {
		key = "x"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyY) {
		key = "y"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		key = "z"
	}

	for x := 'a'; x < 'z'; x++ {
		if key == string(x) {
			if ebiten.IsKeyPressed(ebiten.KeyShiftLeft) || ebiten.IsKeyPressed(ebiten.KeyShiftRight) {
				key = strings.ToUpper(key)
			}
			KeyStringNew = KeyStringNew + key
			break
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		key = "enter"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		key = "escape"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		key = "backspace"
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyComma) {
		key = "."
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyNumpadDivide) {
		key = "/"
	}
	if len(ebiten.InputChars()) > 0 {
		if ebiten.InputChars()[len(ebiten.InputChars())-1] == '/' {
			key = "/"
		}
	}

	if key == "/" || key == "0" || key == "1" || key == "2" || key == "3" || key == "4" || key == "5" || key == "6" || key == "7" || key == "8" || key == "9" || key == "." {
		KeyStringNew = KeyStringNew + key
	}
	if key == "backspace" { //si on appuie sur la touche retour arrière ( la flèche qui pointe vers la gauche), on supprime le dernier élément du string
		if len(KeyStringNew) >= 1 {
			KeyStringNew = string(KeyStringNew[:len(KeyStringNew)-1])
		}
	}
	if key == "escape" {
		KeyStringNew = KeyStringNew + "^"
	}
	if key == "enter" {
		KeyStringNew = KeyStringNew + "#"
	}

}
