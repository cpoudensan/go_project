package main

import "fmt"

func main() {

	graph := map[string]map[string]int{

		// Île-de-France
		"Paris": {
			"Lille": 225,
			"Reims": 145,
			"Orleans": 130,
			"Rouen": 135,
		},
		"Versailles": {
			"Paris": 20,
			"Chartres": 90,
		},
		"Orleans": {
			"Paris": 130,
			"Tours": 120,
			"Bourges": 110,
		},

		// Nord
		"Lille": {
			"Paris": 225,
			"Amiens": 140,
		},
		"Amiens": {
			"Lille": 140,
			"Rouen": 120,
		},

		// Est
		"Reims": {
			"Paris": 145,
			"Metz": 190,
		},
		"Metz": {
			"Reims": 190,
			"Strasbourg": 170,
		},
		"Strasbourg": {
			"Metz": 170,
			"Mulhouse": 115,
		},
		"Mulhouse": {
			"Strasbourg": 115,
			"Besancon": 140,
		},
		"Besancon": {
			"Mulhouse": 140,
			"Dijon": 95,
		},
		"Dijon": {
			"Besancon": 95,
			"Lyon": 195,
		},

		// Centre / Ouest
		"Rouen": {
			"Paris": 135,
			"Amiens": 120,
			"Caen": 130,
		},
		"Caen": {
			"Rouen": 130,
			"Rennes": 180,
		},
		"Rennes": {
			"Caen": 180,
			"Nantes": 110,
		},
		"Nantes": {
			"Rennes": 110,
			"Angers": 95,
		},
		"Angers": {
			"Nantes": 95,
			"Tours": 120,
		},
		"Tours": {
			"Angers": 120,
			"Orleans": 120,
			"Poitiers": 105,
		},
		"Poitiers": {
			"Tours": 105,
			"Limoges": 130,
		},
		"Limoges": {
			"Poitiers": 130,
			"Clermont-Ferrand": 180,
		},

		// Sud / Est
		"Clermont-Ferrand": {
			"Limoges": 180,
			"Lyon": 165,
		},
		"Lyon": {
			"Dijon": 195,
			"Clermont-Ferrand": 165,
			"Valence": 105,
		},
		"Valence": {
			"Lyon": 105,
			"Avignon": 120,
		},
		"Avignon": {
			"Valence": 120,
			"Marseille": 100,
		},
		"Marseille": {
			"Avignon": 100,
			"Toulon": 65,
		},
		"Toulon": {
			"Marseille": 65,
			"Nice": 150,
		},
		"Nice": {
			"Toulon": 150,
		},
	}

	fmt.Println("Nombre de villes :", len(graph))
}
