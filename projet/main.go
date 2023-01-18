package main

import (
	"container/list"
	"log"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particles"

	"github.com/hajimehoshi/ebiten/v2"
)

// main est la fonction principale du projet. Elle commence par lire le fichier
// de configuration, puis elle charge en mémoire l'image d'une particule. Elle
// initialise ensuite la fenêtre d'affichage, puis elle crée un système de
// particules encapsulé dans un "game" et appelle la fonction RunGame qui se
// charge de faire les mise-à-jour (Update) et affichages (Draw) de manière
// régulière.
func main() {
	config.Get("config.json")
	assets.Get()
	GetBtn()

	ebiten.SetWindowTitle("Menu")
	ebiten.SetWindowSize(1280, 720)
	var l *list.List = list.New()

	g := game{system: particles.System{Content: l}} //system: particles.NewSystem()}

	err := ebiten.RunGame(&g)
	if err != nil {
		log.Print(err)
	}

}
