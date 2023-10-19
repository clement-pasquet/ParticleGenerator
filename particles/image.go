package particles

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//Cette fonction retourne l'objet ebiten.Image d'une image passé en argument grâce à sa source
func getImg(src string) *ebiten.Image {
	var err error
	var monImage *ebiten.Image
	monImage, _, err = ebitenutil.NewImageFromFile(src)
	if err != nil {
		log.Fatal("Problem while loading particle image: ", err)
	}
	return monImage
}

//Permet de créer une liste de liste de liste répertoriant les couleur de chaques pixel d'une image
func getListColor(src string) [][][]float64 {
	var ls [][][]float64
	var img *ebiten.Image = getImg(src)
	var maxX, maxY int = img.Size()
	for y := 1; y <= maxY; y++ {
		var lsY [][]float64 = [][]float64{}
		for x := 1; x <= maxX; x++ {
			var r, g, b, a uint32 = img.At(x, y).RGBA()
			var lsRgba []float64 = []float64{float64(r), float64(g), float64(b), float64(a)}
			lsY = append(lsY, lsRgba)
		}
		ls = append(ls, lsY)
	}
	return ls

}
