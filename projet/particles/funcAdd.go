package particles

import ("math/rand"
	"container/list"
	"math"
	"time"
	"project-particles/config"
)

func puissance2(a float64) float64{
	return a*a
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

func LifeSpanIsTooAged(ele *list.Element,p *Particle,l *list.List) {
	if config.General.RandomSpawn == true{
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
		if config.General.LifeSpanMax<= float64(maParticule.LifeSpan)/60{ //Supprime la particule si  elle est vieille de plus de 2 secondes
			l.Remove(ele)
			p.Opacity=0
			p.PositionX = float64(config.General.WindowSizeX) + 500
		}
	}
}

func setColor(p *Particle) {
	
	var param int = config.General.Flag
	//color
	var b1 []float64= []float64{0, 0, 255}
	var blanc []float64=[]float64{255, 255,255}
	var r1 []float64=[]float64{255, 0, 0}
	var vert1 []float64=[]float64{0, 255, 0}
	var jaune1 []float64=[]float64{255, 255, 0}
	var orange1 []float64=[]float64{255, 128, 10}
	var orange2 []float64=[]float64{255, 160, 10}
	var noir1 []float64=[]float64{75,75,75}
	//Country
	var France [][]float64 = [][]float64{b1,blanc,r1}
	var Italie [][]float64 = [][]float64{vert1,blanc,r1}
	var Irelande [][]float64 = [][]float64{vert1,blanc,orange1}
	var Belgique [][]float64 = [][]float64{noir1,jaune1,r1}
	var Roumanie [][]float64 = [][]float64{b1,orange1,r1}
	var Inde [][]float64 = [][]float64{orange2,blanc,vert1}
	var Allemagne [][]float64 = [][]float64{noir1,r1,jaune1}


	switch param{
		case 1:
			//FRANCE
			if p.PositionX <427{ 
				p.ColorRed=France[0][0]/255
				p.ColorGreen=France[0][1]/255
				p.ColorBlue =France[0][2]/255
			}else if p.PositionX <853{
				p.ColorRed=France[1][0]/255
				p.ColorGreen=France[1][1]/255
				p.ColorBlue =France[1][2]/255
			}else{
				p.ColorRed=France[2][0]/255
				p.ColorGreen=France[2][1]/255
				p.ColorBlue =France[2][2]/255
			}
			
		case 2:
			//ITALIE
			if p.PositionX <427{
				p.ColorRed=Italie[0][0]/255
				p.ColorGreen=Italie[0][1]/255
				p.ColorBlue =Italie[0][2]/255
			}else if p.PositionX <853{
				p.ColorRed=Italie[1][0]/255
				p.ColorGreen=Italie[1][1]/255
				p.ColorBlue =Italie[1][2]/255
			}else{
				p.ColorRed=Italie[2][0]/255
				p.ColorGreen=Italie[2][1]/255
				p.ColorBlue =Italie[2][2]/255
			}

		case 3:
			//Irelande
			if p.PositionX <427{
				p.ColorRed=Irelande[0][0]/255
				p.ColorGreen=Irelande[0][1]/255
				p.ColorBlue =Irelande[0][2]/255
			}else if p.PositionX <853{
				p.ColorRed=Irelande[1][0]/255
				p.ColorGreen=Irelande[1][1]/255
				p.ColorBlue =Irelande[1][2]/255
			}else{
				p.ColorRed=Irelande[2][0]/255
				p.ColorGreen=Irelande[2][1]/255
				p.ColorBlue =Irelande[2][2]/255
			}
		
		case 4:
			//Belgique
			if p.PositionX <427{
				p.ColorRed=Belgique[0][0]/255
				p.ColorGreen=Belgique[0][1]/255
				p.ColorBlue =Belgique[0][2]/255
			}else if p.PositionX <853{
				p.ColorRed=Belgique[1][0]/255
				p.ColorGreen=Belgique[1][1]/255
				p.ColorBlue =Belgique[1][2]/255
			}else{
				p.ColorRed=Belgique[2][0]/255
				p.ColorGreen=Belgique[2][1]/255
				p.ColorBlue =Belgique[2][2]/255
			}
		case 5:
			//Roumanie
			if p.PositionX <427{
				p.ColorRed=Roumanie[0][0]/255
				p.ColorGreen=Roumanie[0][1]/255
				p.ColorBlue =Roumanie[0][2]/255
			}else if p.PositionX <853{
				p.ColorRed=Roumanie[1][0]/255
				p.ColorGreen=Roumanie[1][1]/255
				p.ColorBlue =Roumanie[1][2]/255
			}else{
				p.ColorRed=Roumanie[2][0]/255
				p.ColorGreen=Roumanie[2][1]/255
				p.ColorBlue =Roumanie[2][2]/255
			}
		case 6:
			//Inde
			if p.PositionY <240{
				p.ColorRed=Inde[0][0]/255
				p.ColorGreen=Inde[0][1]/255
				p.ColorBlue =Inde[0][2]/255
			}else if p.PositionY <480{
				p.ColorRed=Inde[1][0]/255
				p.ColorGreen=Inde[1][1]/255
				p.ColorBlue =Inde[1][2]/255
			}else{
				p.ColorRed=Inde[2][0]/255
				p.ColorGreen=Inde[2][1]/255
				p.ColorBlue =Inde[2][2]/255
			}
			var a float64 = math.Sqrt(puissance2((p.PositionX-float64(config.General.WindowSizeX)/2))+puissance2((p.PositionY-float64(config.General.WindowSizeY)/2)))
			
			if a >80 && a<100{
				p.ColorRed=b1[0]/255
				p.ColorGreen=b1[1]/255
				p.ColorBlue =b1[2]/255
			}

			if a<22{
				p.ColorRed=b1[0]/255
				p.ColorGreen=b1[1]/255
				p.ColorBlue =b1[2]/255
			}
			
			
		case 7:
			//Allemagne
			if p.PositionY <240{
				p.ColorRed=Allemagne[0][0]/255
				p.ColorGreen=Allemagne[0][1]/255
				p.ColorBlue =Allemagne[0][2]/255
			}else if p.PositionY <480{
				p.ColorRed=Allemagne[1][0]/255
				p.ColorGreen=Allemagne[1][1]/255
				p.ColorBlue =Allemagne[1][2]/255
			}else{
				p.ColorRed=Allemagne[2][0]/255
				p.ColorGreen=Allemagne[2][1]/255
				p.ColorBlue =Allemagne[2][2]/255
			}
		case 8:
			//JAPON
			var a float64 = math.Sqrt(puissance2((p.PositionX-float64(config.General.WindowSizeX)/2))+puissance2((p.PositionY-float64(config.General.WindowSizeY)/2)))
			p.ColorRed=blanc[0]/255
			p.ColorGreen=blanc[1]/255
			p.ColorBlue =blanc[2]/255
			if  a<150{
				p.ColorRed=r1[0]/255
				p.ColorGreen=r1[1]/255
				p.ColorBlue =r1[2]/255
			}		
	}		
}

func IsOutOfView(e *list.Element,p *Particle,l *list.List) bool {
	if p.PositionX < -config.General.Margin || p.PositionX > float64(config.General.WindowSizeX) + config.General.Margin || p.PositionY > float64(config.General.WindowSizeY) + config.General.Margin {
		
		
		p.PositionX = float64(config.General.WindowSizeX) + 500
		p.Opacity = 0
		go l.Remove(e) 
		return true
	}
	return false
}