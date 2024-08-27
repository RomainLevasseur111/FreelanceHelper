package main

import "fmt"


func ImpotsRevenus(){
	var nbParts float32
	var revenus float32
	fmt.Print("Entrez votre nombre  de parts: ")
	fmt.Scan(&nbParts)
	fmt.Print("Entrez vos revenus annuel: ")
	fmt.Scan(&revenus)

	var revenusNpParts float32 = revenus / nbParts
	var abattement float32 = revenusNpParts *0.90

	var res float32

	if abattement <= 11294{
		res = 0
	}else if abattement > 11294 && abattement <= 28797{
		res = (abattement - 11294) * 0.11
	}else if abattement > 28797 && abattement <= 82341{
		res = (28797-11294)*0.11 + ((abattement - 28797) * 0.3)
	}else if abattement > 82341 && abattement <= 177106{
		res = (28797-11294)*0.11 + ((82341 - 28797)*0.3) + ((abattement - 82341)*0.41)
	}else{
		res = (28797-11294)*0.11 + ((82341 - 28797)*0.3) + ((177106 - 82341)*0.41) + ((abattement - 177106)*0.45)
	}
	fmt.Printf("Nombre de parts : %.0f. \nRevenus imposable annuel : %.0f.\nRevenus imposable après abattement automatique de 10%%: %.2f.\nImpôts sur le revenu: %.2f.\n", nbParts, revenusNpParts, abattement, res)
}

func ImpotsIS(){

	var statut string
	fmt.Print("1: Entreprise individuelle\n2:EIRL\n3:EURL\n4:SASU\nEntrez le numéro corresondant à votre statut juridique.")
	fmt.Scan(&statut)

	switch statut{
	case "1":
		
	}
}