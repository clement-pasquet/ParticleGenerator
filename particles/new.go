package particles

import (
	"container/list"
	"fmt"
	"math/rand"
	"project-particles/config"
	"time"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.
func NewSystem() System {
	l := list.New()

	fmt.Println(l)

	rand.Seed(time.Now().UnixNano())
	var s System = System{Content: l}
	createNParticles(config.General.InitNumParticles,&s)

	return s
}
