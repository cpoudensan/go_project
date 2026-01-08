package main

var Graph = map[string]map[string]int{
	"Paris": {
		"Lille": 225, "Rennes": 350, "Nantes": 385,
		"Strasbourg": 490, "Dijon": 315, "Lyon": 465,
	},
	"Lille": {
		"Paris": 225, "Metz": 280,
	},
	"Rennes": {
		"Paris": 350, "Nantes": 110,
	},
	"Nantes": {
		"Rennes": 110, "Paris": 385, "Bordeaux": 335,
	},
	"Bordeaux": {
		"Nantes": 335, "Toulouse": 245,
	},
	"Toulouse": {
		"Bordeaux": 245, "Montpellier": 240,
	},
	"Montpellier": {
		"Toulouse": 240, "Nîmes": 55, "Béziers": 75,
	},
	"Nîmes": {
		"Montpellier": 55, "Avignon": 40,
	},
	"Béziers": {
		"Montpellier": 75, "Perpignan": 120,
	},
	"Perpignan": {
		"Béziers": 120,
	},
	"Avignon": {
		"Nîmes": 40, "Valence": 135, "Aix-en-Provence": 85,
	},
	"Aix-en-Provence": {
		"Avignon": 85, "Marseille": 30,
	},
	"Marseille": {
		"Aix-en-Provence": 30, "Toulon": 65,
	},
	"Toulon": {
		"Marseille": 65, "Cannes": 120,
	},
	"Cannes": {
		"Toulon": 120, "Antibes": 15, "Nice": 35,
	},
	"Antibes": {
		"Cannes": 15,
	},
	"Nice": {
		"Cannes": 35,
	},
	"Valence": {
		"Avignon": 135, "Lyon": 105, "Grenoble": 95,
	},
	"Lyon": {
		"Paris": 465, "Valence": 105, "Saint-Étienne": 65,
		"Chambéry": 110,
	},
	"Saint-Étienne": {
		"Lyon": 65, "Clermont-Ferrand": 145,
	},
	"Clermont-Ferrand": {
		"Saint-Étienne": 145,
	},
	"Grenoble": {
		"Valence": 95, "Chambéry": 60, "Annecy": 105,
	},
	"Chambéry": {
		"Lyon": 110, "Grenoble": 60, "Annecy": 55,
	},
	"Annecy": {
		"Chambéry": 55, "Grenoble": 105,
	},
	"Dijon": {
		"Paris": 315, "Besançon": 95, "Nancy": 200,
	},
	"Besançon": {
		"Dijon": 95,
	},
	"Nancy": {
		"Dijon": 200, "Metz": 55, "Strasbourg": 150,
	},
	"Metz": {
		"Lille": 280, "Nancy": 55,
	},
	"Strasbourg": {
		"Paris": 490, "Nancy": 150,
	},
}
