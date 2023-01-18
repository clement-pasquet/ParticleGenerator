package main

import (
	"fmt"
	"image/color"
	"log"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particles"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// Draw se charge d'afficher à l'écran l'état actuel du système de particules
// g.system. Elle est appelée automatiquement environ 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction pourra être légèrement modifiée quand
// c'est précisé dans le sujet.
var ArrierePlan *ebiten.Image
var PlayButton *ebiten.Image
var SlideRouge *ebiten.Image
var SlideVert *ebiten.Image
var Entry *ebiten.Image
var mplusNormalFont font.Face
var EntryFocus bool //Variable permettant de vérifier si une entrée est déjà en train d'être utilisé.

// Initialisation des variables qui serviront plus tard à modifier la caractéristique éponyme de config.json.
var EntryVelocity EntryText
var EntryWindowsSizeX EntryText
var EntryWindowsSizeY EntryText
var EntryInitNumParticles EntryText
var EntrySpawnRate EntryText
var EntryGravity EntryText
var EntryMargin EntryText
var EntryFlag EntryText
var EntrySizeShape EntryText
var EntryGeneratorShape EntryText
var EntryParticleSize EntryText
var EntryLifeSpan EntryText
var EntryCustomImageSRC EntryText

var IsInitialised bool

type EntryText struct {
	isBeingModified bool
	posX, posY      float64
	text            string
	oldText         string
}

func (g *game) Draw(screen *ebiten.Image) {
	if !config.HasGameStarted {
		if !IsInitialised {
			EntryVelocity = EntryText{isBeingModified: false, posX: 805, posY: 228, text: fmt.Sprintf("%f", config.General.Velocity), oldText: fmt.Sprintf("%f", config.General.Velocity)}
			EntryWindowsSizeX = EntryText{isBeingModified: false, posX: 215, posY: 179, text: strconv.Itoa(config.General.WindowSizeX), oldText: strconv.Itoa(config.General.WindowSizeX)}
			EntryWindowsSizeY = EntryText{isBeingModified: false, posX: 380, posY: 179, text: strconv.Itoa(config.General.WindowSizeY), oldText: strconv.Itoa(config.General.WindowSizeY)}
			EntryInitNumParticles = EntryText{isBeingModified: false, posX: 430, posY: 230, text: strconv.Itoa(config.General.InitNumParticles), oldText: strconv.Itoa(config.General.InitNumParticles)}
			EntrySpawnRate = EntryText{isBeingModified: false, posX: 285, posY: 303, text: fmt.Sprintf("%f", config.General.SpawnRate), oldText: fmt.Sprintf("%f", config.General.SpawnRate)}
			EntryGravity = EntryText{isBeingModified: false, posX: 158, posY: 327, text: fmt.Sprintf("%f", config.General.Gravity), oldText: fmt.Sprintf("%f", config.General.Gravity)}
			EntryMargin = EntryText{isBeingModified: false, posX: 130, posY: 361, text: fmt.Sprintf("%f", config.General.Margin), oldText: fmt.Sprintf("%f", config.General.Margin)}
			EntryFlag = EntryText{isBeingModified: false, posX: 149, posY: 389, text: strconv.Itoa(config.General.Flag), oldText: strconv.Itoa(config.General.Flag)}
			EntrySizeShape = EntryText{isBeingModified: false, posX: 374, posY: 553, text: fmt.Sprintf("%f", config.General.SizeShape), oldText: fmt.Sprintf("%f", config.General.SizeShape)}
			EntryGeneratorShape = EntryText{isBeingModified: false, posX: 358, posY: 660, text: config.General.GeneratorShape, oldText: config.General.GeneratorShape}
			EntryParticleSize = EntryText{isBeingModified: false, posX: 939, posY: 174, text: fmt.Sprintf("%f", config.General.ScaleX), oldText: fmt.Sprintf("%f", config.General.ScaleX)}
			EntryLifeSpan = EntryText{isBeingModified: false, posX: 943, posY: 331, text: fmt.Sprintf("%f", config.General.LifeSpanMax), oldText: fmt.Sprintf("%f", config.General.LifeSpanMax)}
			EntryCustomImageSRC = EntryText{isBeingModified: false, posX: 680, posY: 585, text: config.General.CustomImageSRC, oldText: config.General.CustomImageSRC}
			IsInitialised = true
		}
		options := ebiten.DrawImageOptions{}
		options.GeoM.Rotate(0)
		options.GeoM.Scale(1, 1)
		options.GeoM.Translate(0, 0)
		options.ColorM.Scale(1, 1, 1, 1)
		screen.DrawImage(ArrierePlan, &options)
		CreateMenu(screen, g)
	} else {
		var i int = 0
		living := true

		for e := g.system.Content.Front(); e != nil && living; e = e.Next() {
			i++
			p, ok := e.Value.(*particles.Particle)
			if ok {
				if p.IsInLife {
					options := ebiten.DrawImageOptions{}
					options.GeoM.Rotate(p.Rotation)
					options.GeoM.Scale(p.ScaleX, p.ScaleY)
					options.GeoM.Translate(p.PositionX, p.PositionY)
					options.ColorM.Scale(p.ColorRed, p.ColorGreen, p.ColorBlue, p.Opacity)
					screen.DrawImage(assets.ParticleImage, &options)

				}

			}
		}
		if config.General.Debug { //Sert à activer le mode Débug : affiche les FPS et le Nb de Part. Mortes et Vivantes
			ebitenutil.DebugPrint(screen, fmt.Sprint("FPS : ", ebiten.CurrentTPS()))
			ebitenutil.DebugPrintAt(screen, fmt.Sprint("Nb Part. Mortes : ", g.system.NbParticulesMortes), 0, 20)
			ebitenutil.DebugPrintAt(screen, fmt.Sprint("Nb Part. Vivantes : ", g.system.Content.Len()-g.system.NbParticulesMortes), 0, 40)
		}
	}

}

// Ouvre les images et s'occupent de leurs erreurs; récupère la fonte utilisé pour les textes.
func GetBtn() {
	var err error

	SlideRouge, _, err = ebitenutil.NewImageFromFile("assets/ProjetParticule_Slider_Off.png") //Ligne qui "ouvre" une image grâce à ebiten.
	if err != nil {                                                                           //Si il y a une erreur durant l'ouverture, crée une erreur.
		log.Fatal("Problem while loading particle image: ", err)
	}
	SlideVert, _, err = ebitenutil.NewImageFromFile("assets/ProjetParticule_Slider_On.png")
	if err != nil {
		log.Fatal("Problem while loading particle image: ", err)
	}
	Entry, _, err = ebitenutil.NewImageFromFile("assets/ProjetParticule_ChampTexte1.png")
	if err != nil {
		log.Fatal("Problem while loading particle image: ", err)
	}
	PlayButton, _, err = ebitenutil.NewImageFromFile("assets/Projet_Particule_boutonJouer.png")
	if err != nil {
		log.Fatal("Problem while loading particle image: ", err)
	}
	ArrierePlan, _, err = ebitenutil.NewImageFromFile("assets/ProjetParticule_ArrierePlanFlou(1).png")
	if err != nil {
		log.Fatal("Problem while loading particle image: ", err)
	}
	//Recupere la font
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    14,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

// Fonction principale de l'interface graphique. Gère les collisions des boutons & les affiche.
func CreateMenu(screen *ebiten.Image, g *game) {
	mx, my := ebiten.CursorPosition()

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && mx >= 200 && float64(mx) < 200+float64(SlideRouge.Bounds().Dx())*0.25 && my >= 208 && float64(my) < 208+float64(SlideRouge.Bounds().Dy())*0.25 {
		if config.General.Debug {
			config.General.Debug = false
		} else {
			config.General.Debug = true
		}
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && mx >= 250 && float64(mx) < 250+float64(SlideRouge.Bounds().Dx())*0.25 && my >= 270 && float64(my) < 270+float64(SlideRouge.Bounds().Dy())*0.25 {
		if config.General.RandomSpawn {
			config.General.RandomSpawn = false
		} else {
			config.General.RandomSpawn = true
		}
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && mx >= 1050 && float64(mx) < 1050+float64(SlideRouge.Bounds().Dx())*0.25 && my >= 267 && float64(my) < 267+float64(SlideRouge.Bounds().Dy())*0.25 {
		if config.General.Collision {
			config.General.Collision = false
		} else {
			config.General.Collision = true
		}
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && mx >= 918 && float64(mx) < 918+float64(SlideRouge.Bounds().Dx())*0.25 && my >= 500 && float64(my) < 500+float64(SlideRouge.Bounds().Dy())*0.25 {
		if config.General.CustomImageBool {
			config.General.CustomImageBool = false
		} else {
			config.General.CustomImageBool = true
		}
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && mx >= 405 && float64(mx) < 405+float64(SlideRouge.Bounds().Dx())*0.25 && my >= 605 && float64(my) < 605+float64(SlideRouge.Bounds().Dy())*0.25 {
		if config.General.MouseClick {
			config.General.MouseClick = false
		} else {
			config.General.MouseClick = true
		}
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && mx >= 1045 && float64(mx) < 1045+float64(SlideRouge.Bounds().Dx()) && my >= 650 && float64(my) < 650+float64(SlideRouge.Bounds().Dy())*0.5 {
		g.system = particles.NewSystem()
		ebiten.SetWindowTitle(config.General.WindowTitle)
		ebiten.SetWindowSize(config.General.WindowSizeX, config.General.WindowSizeY)

		config.HasGameStarted = true

	}

	options := ebiten.DrawImageOptions{}

	options.ColorM.Scale(1, 1, 1, 1)
	options.GeoM.Translate(1045*2, 650*2)
	options.GeoM.Scale(0.5, 0.5)
	screen.DrawImage(PlayButton, &options)

	ShowButton(screen, 200, 208, config.General.Debug)           // Debug
	ShowButton(screen, 250, 270, config.General.RandomSpawn)     // random Spawn
	ShowButton(screen, 1050, 267, config.General.Collision)      //Colision
	ShowButton(screen, 918, 500, config.General.CustomImageBool) //Afficher image
	ShowButton(screen, 405, 605, config.General.MouseClick)      //deplacement du générateur

	//Les Entry. Cette partie sert à gérer ce qui arrive quand on clique sur un champ texte de l'interface graphique.
	if !EntryFocus {
		//Velocity
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && mx >= 805 && float64(mx) < 805+float64(SlideRouge.Bounds().Dx())*0.5 && my >= 228 && float64(my) < 228+float64(SlideRouge.Bounds().Dy())*0.25 {
			EntryVelocity.isBeingModified = true
			EntryFocus = true
		}
		//WindowsSizeX
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && mx >= int(EntryWindowsSizeX.posX) && float64(mx) < EntryWindowsSizeX.posX+float64(SlideRouge.Bounds().Dx())*0.5 && float64(my) >= EntryWindowsSizeX.posY && float64(my) < EntryWindowsSizeX.posY+float64(SlideRouge.Bounds().Dy())*0.25 {
			EntryWindowsSizeX.isBeingModified = true
			EntryFocus = true
		}
		//WindowsSizeY
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && mx >= int(EntryWindowsSizeY.posX) && float64(mx) < EntryWindowsSizeY.posX+float64(SlideRouge.Bounds().Dx())*0.5 && float64(my) >= EntryWindowsSizeY.posY && float64(my) < EntryWindowsSizeY.posY+float64(SlideRouge.Bounds().Dy())*0.25 {
			EntryWindowsSizeY.isBeingModified = true
			EntryFocus = true
		}
		//InitParticle
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && mx >= int(EntryInitNumParticles.posX) && float64(mx) < EntryInitNumParticles.posX+float64(SlideRouge.Bounds().Dx())*0.5 && float64(my) >= EntryInitNumParticles.posY && float64(my) < EntryInitNumParticles.posY+float64(SlideRouge.Bounds().Dy())*0.25 {
			EntryInitNumParticles.isBeingModified = true
			EntryFocus = true
		}
		//SpawnRate
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && mx >= int(EntrySpawnRate.posX) && float64(mx) < EntrySpawnRate.posX+float64(SlideRouge.Bounds().Dx())*0.5 && float64(my) >= EntrySpawnRate.posY && float64(my) < EntrySpawnRate.posY+float64(SlideRouge.Bounds().Dy())*0.25 {
			EntrySpawnRate.isBeingModified = true
			EntryFocus = true
		}
		//Gravity
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && mx >= int(EntryGravity.posX) && float64(mx) < EntryGravity.posX+float64(SlideRouge.Bounds().Dx())*0.5 && float64(my) >= EntryGravity.posY && float64(my) < EntryGravity.posY+float64(SlideRouge.Bounds().Dy())*0.25 {
			EntryGravity.isBeingModified = true
			EntryFocus = true
		}
		//Margin
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && mx >= int(EntryMargin.posX) && float64(mx) < EntryMargin.posX+float64(SlideRouge.Bounds().Dx())*0.5 && float64(my) >= EntryMargin.posY && float64(my) < EntryMargin.posY+float64(SlideRouge.Bounds().Dy())*0.25 {
			EntryMargin.isBeingModified = true
			EntryFocus = true
		}
		//Flag
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && mx >= int(EntryFlag.posX) && float64(mx) < EntryFlag.posX+float64(SlideRouge.Bounds().Dx())*0.5 && float64(my) >= EntryFlag.posY && float64(my) < EntryFlag.posY+float64(SlideRouge.Bounds().Dy())*0.25 {
			EntryFlag.isBeingModified = true
			EntryFocus = true
		}
		//SizeShape
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && mx >= int(EntrySizeShape.posX) && float64(mx) < EntrySizeShape.posX+float64(SlideRouge.Bounds().Dx())*0.5 && float64(my) >= EntrySizeShape.posY && float64(my) < EntrySizeShape.posY+float64(SlideRouge.Bounds().Dy())*0.25 {
			EntrySizeShape.isBeingModified = true
			EntryFocus = true
		}
		//GeneratorShape
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && mx >= int(EntryGeneratorShape.posX) && float64(mx) < EntryGeneratorShape.posX+float64(SlideRouge.Bounds().Dx())*0.5 && float64(my) >= EntryGeneratorShape.posY && float64(my) < EntryGeneratorShape.posY+float64(SlideRouge.Bounds().Dy())*0.25 {
			EntryGeneratorShape.isBeingModified = true
			EntryFocus = true
		}
		//ParticleSize
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && mx >= int(EntryParticleSize.posX) && float64(mx) < EntryParticleSize.posX+float64(SlideRouge.Bounds().Dx())*0.5 && float64(my) >= EntryParticleSize.posY && float64(my) < EntryParticleSize.posY+float64(SlideRouge.Bounds().Dy())*0.25 {
			EntryParticleSize.isBeingModified = true
			EntryFocus = true
		}
		//LifeSpan
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && mx >= int(EntryLifeSpan.posX) && float64(mx) < EntryLifeSpan.posX+float64(SlideRouge.Bounds().Dx())*0.5 && float64(my) >= EntryLifeSpan.posY && float64(my) < EntryLifeSpan.posY+float64(SlideRouge.Bounds().Dy())*0.25 {
			EntryLifeSpan.isBeingModified = true
			EntryFocus = true
		}
		//CustomImageSRC
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && mx >= int(EntryCustomImageSRC.posX) && float64(mx) < EntryCustomImageSRC.posX+float64(SlideRouge.Bounds().Dx())*0.5 && float64(my) >= EntryCustomImageSRC.posY && float64(my) < EntryCustomImageSRC.posY+float64(SlideRouge.Bounds().Dy())*0.25 {
			EntryCustomImageSRC.isBeingModified = true
			EntryFocus = true
		}

	}

	ShowEntry(screen, &EntryWindowsSizeX)
	config.General.WindowSizeX, _ = strconv.Atoi(EntryWindowsSizeX.oldText)

	ShowEntry(screen, &EntryWindowsSizeY)
	config.General.WindowSizeY, _ = strconv.Atoi(EntryWindowsSizeY.oldText)

	ShowEntry(screen, &EntryInitNumParticles)
	config.General.InitNumParticles, _ = strconv.Atoi(EntryInitNumParticles.oldText)

	ShowEntry(screen, &EntrySpawnRate)
	config.General.SpawnRate, _ = strconv.ParseFloat(EntrySpawnRate.oldText, 64)

	ShowEntry(screen, &EntryGravity)
	config.General.Gravity, _ = strconv.ParseFloat(EntryGravity.oldText, 64)

	ShowEntry(screen, &EntryMargin)
	config.General.Margin, _ = strconv.ParseFloat(EntryMargin.oldText, 64)

	ShowEntry(screen, &EntryFlag)
	config.General.Flag, _ = strconv.Atoi(EntryFlag.oldText)

	ShowEntry(screen, &EntrySizeShape)
	config.General.SizeShape, _ = strconv.ParseFloat(EntrySizeShape.oldText, 64)

	ShowEntry(screen, &EntryGeneratorShape)
	config.General.GeneratorShape = EntryGeneratorShape.oldText

	ShowEntry(screen, &EntryParticleSize)
	config.General.ScaleX, _ = strconv.ParseFloat(EntryParticleSize.oldText, 64)

	ShowEntry(screen, &EntryParticleSize)
	config.General.ScaleY, _ = strconv.ParseFloat(EntryParticleSize.oldText, 64)

	ShowEntry(screen, &EntryLifeSpan)
	config.General.LifeSpanMax, _ = strconv.ParseFloat(EntryLifeSpan.oldText, 64)

	ShowEntry(screen, &EntryCustomImageSRC)
	config.General.CustomImageSRC = EntryCustomImageSRC.oldText

	if config.General.Debug { //Sert à activer le mode Débug : affiche les FPS et le Nb de Part. Mortes et Vivantes
		ebitenutil.DebugPrint(screen, fmt.Sprint("FPS : ", ebiten.CurrentTPS()))
	}
}

// Affiche un bouton à deux états : allumés et verts ou éteint et rouge.
func ShowButton(screen *ebiten.Image, posX float64, posY float64, etat bool) {
	options := ebiten.DrawImageOptions{}

	options.ColorM.Scale(1, 1, 1, 1)
	options.GeoM.Translate(posX*4, posY*4)
	options.GeoM.Scale(0.25, 0.25)
	if etat {
		screen.DrawImage(SlideVert, &options)

	} else {
		screen.DrawImage(SlideRouge, &options)
	}

}

// Permet d'actualiser l'affichage des champs-texte.
func ShowEntry(screen *ebiten.Image, entry *EntryText) {
	options := ebiten.DrawImageOptions{}

	options.ColorM.Scale(1, 1, 1, 1)
	options.GeoM.Translate(entry.posX*2, entry.posY*4)
	options.GeoM.Scale(0.5, 0.25)
	screen.DrawImage(Entry, &options)

	if entry.isBeingModified {
		a := entry.text
		particles.KeyboardInput()

		fmt.Println(len(entry.text) - len(a))

		entry.text = particles.KeyStringNew
		if len(entry.text) > 0 {
			if entry.text[len(entry.text)-1] == '#' { //Si l'utilisateur a appuyé sur "Entrée", alors la valeur qu'il a inscrit est conservé
				entry.text = particles.KeyStringNew[:len(entry.text)-1]
				particles.KeyStringNew = ""
				EntryFocus = false
				entry.isBeingModified = false

			} else if entry.text[len(entry.text)-1] == '^' { //Si l'utilisateur a appuyé sur "Echap", alors on remet le texte d'origine du champ-texte.
				entry.text = entry.oldText
				particles.KeyStringNew = ""
				EntryFocus = false
				entry.isBeingModified = false
				return

			}
		}

		text.Draw(screen, entry.text, mplusNormalFont, int(entry.posX)+10, int(entry.posY)+18, color.Black)

	} else {
		entry.oldText = entry.text
		text.Draw(screen, entry.text, mplusNormalFont, int(entry.posX)+10, int(entry.posY)+18, color.Black)

	}
}
