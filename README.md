# Gestion de Particule

## Le Concept

Le but est de créer un générateur de particule.

Ce générateur de particule pourra servir à créer **différents effets complexes** (explosions, feu, pluies ...) et sera complété par différentes "extensions" tel que l'ajout de la gravité, d'une "durée de vie" à chaque particule, d'un "générateur de forme" pour générer des particules sur une forme carré, rectangulaire, ronde; de collisions et de plusieurs optimisation mémoire.


## Utiliser notre programme

Pour lancer notre programme, il suffit de lancer le fichier "project-particles" sous Linux avec la commande : 
```sh 
./gestion-Particule-08-01/project-particles
```
    
Pour modifier des variables (vitesse, taille de l'écran, taille, _SpawnRate_ ...) il faut éditer le fichier **config.json** :
```json
{
	"WindowTitle": "Release Projet Particules - Bryan & Clément",
	"WindowSizeX": 1280,
    "WindowSizeY": 720,
	"ParticleImage": "assets/particle2.png", 
	"Debug": true,
	"InitNumParticles": 200,
	"RandomSpawn": false,
	"SpawnX": 640,
	"SpawnY": 360,
	"SpawnRate":0.9,
	"ColorRed": 200,
	"ColorGreen": 200 ,
	"ColorBlue": 200,
	"ScaleX": 0.32,
	"ScaleY": 0.32,
	"Opacity": 1,
	"Velocity": 5
}
```
### **Afficher les IPS**

Le paramètre **Debug** sert à afficher les Images Par Secondes :
```json
    "Debug": true
```
### **Activer ou Désactiver l'Apparition Aléatoire**
Le paramètre **RandomSpawn** sert à faire apparaitre les particules en un point (le centre) ou les faire apparaître de manière aléatoire :
```json
    "RandomSpawn": false
```
### **Changer les propriétés des particules**
Ces propriétés regissent l'apparence des particules, mais pour l'instant, ces paramètres ne changent pas en cours de route.
```json
	"ColorRed": 200,
	"ColorGreen": 200 ,
	"ColorBlue": 200,
	"ScaleX": 0.32,
	"ScaleY": 0.32,
	"Opacity": 1,
	"Velocity": 5
```
## Tester le Projet

Pour vérifier un projet, il est important d'écrire des tests et de les valider.
Pour faire cela, il faut faire la commande, sous Linux : 
```sh
go test ./gestion-Particule-08-01/particles
```

## Structure du Projet
Le projet est contenu dans **3 dossiers** :

Le premier dossier contient l'image que chaque particule utilise, et le code pour l'ouvrir.

* assets :
* * getassets.go
* * particle2.png
* * particle3.png  

Le deuxième dossier contient **getconfig.go** dont le but est de lire le fichier **config.json** et le fichier **type.go** dont le but est de définir la structure du fichier json.
* config :
*   * getconfig.go
*   * type.go

Et finalement le dossier particles contient le fichier **new.go** qui crée les premières particules à apparaitre, **type.go** qui définit la structure particule, **update.go** qui est appelé 60 fois par seconde et qui met à jour les propriétés des particules et finalement **funcAdd.go** et **funcAdd_test.go** qui sont de *nouveaux fichiers* et contiennent une partie des fonctions utilisés dans d'autres fichiers.
* particles :
*   * funcAdd_test.go
*   * funcAdd.go
*   * new.go
*   * type.go
*   * update.go
