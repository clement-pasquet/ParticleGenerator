package config

// Config définit les champs qu'on peut trouver dans un fichier de config.
// Dans le fichier les champs doivent porter le même nom que dans le type si
// dessous, y compris les majuscules. Tous les champs doivent obligatoirement
// commencer par des majuscules, sinon il ne sera pas possible de récupérer
// leurs valeurs depuis le fichier de config.
// Vous pouvez ajouter des champs et ils seront automatiquement lus dans le
// fichier de config. Vous devrez le faire plusieurs fois durant le projet.
type Config struct {
	WindowTitle              string //Nom du programme
	WindowSizeX, WindowSizeY int    //Résolution du programme, actuellement 1280*720
	ParticleImage            string //Lien de l'image utilisé pour les particules
	Debug                    bool   //Permet Principalement d'afficher les FPS
	InitNumParticles         int    //Nombre de particule qui s'afficheront au démarrage
	RandomSpawn              bool
	SpawnX, SpawnY           int //Définit le centre de l'écran
	SpawnRate                float64
	ScaleX                   float64
	ScaleY                   float64
	Opacity                  float64
	ColorRed                 float64
	ColorBlue                float64
	ColorGreen               float64
	Velocity                 float64 //Gere la vitesse horizontale et verticale
}

var General Config
