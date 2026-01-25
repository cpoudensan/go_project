# Application Client–Serveur – Calcul du plus court chemin (Go)

## Description
Ce projet est une application client–serveur en Go permettant de calculer le plus court chemin et la distance minimale entre deux villes.

Le serveur implémente l’algorithme de Dijkstra sur un graphe pondéré représentant des villes et les distances entre elles.  
Le client envoie des requêtes textuelles au serveur, qui renvoie le chemin optimal et la distance correspondante.

Le serveur est capable de gérer plusieurs clients simultanément grâce à l’utilisation des goroutines.

---

## Comment exécuter le programme

Le serveur doit être lancé AVANT le client!

### Étape 1 : Lancer le serveur
Ouvrir un terminal, se placer dans le dossier `server`(à l'intérieur du dossier cmd), puis exécuter :

go run .

Dans un autre terminal, se placer dans le dossier `client`(à l'intérieur du dossier cmd), puis exécuter : 

go run .
Puis entrer une ligne du type: ROUTE Paris Bordeaux