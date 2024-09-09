package main

import "fmt"


func ImpotsRevenus(nbParts float32, revenus float32)float32{

	var revenusNpParts float32 = revenus / nbParts

	var res float32

	if revenusNpParts <= 11294{
		res = 0
	}else if revenusNpParts > 11294 && revenusNpParts <= 28797{
		res = ((revenusNpParts - 11294) * 0.11)-(873-((revenusNpParts - 11294) * 0.11)*0.4525)
		fmt.Println(873-((revenusNpParts - 11294) * 0.11)*0.4525)
	}else if revenusNpParts > 28797 && revenusNpParts <= 82341{
		res = (28797-11294)*0.11 + ((revenusNpParts - 28797) * 0.3)
	}else if revenusNpParts > 82341 && revenusNpParts <= 177106{
		res = (28797-11294)*0.11 + ((82341 - 28797)*0.3) + ((revenusNpParts - 82341)*0.41)
	}else{
		res = (28797-11294)*0.11 + ((82341 - 28797)*0.3) + ((177106 - 82341)*0.41) + ((revenusNpParts - 177106)*0.45)
	}
	fmt.Printf("Nombre de parts : %.0f. \nRevenus imposable annuel : %.0f.\nImpôts sur le revenu: %.2f.\n", nbParts, revenusNpParts, res)
	return res
}

func ImpotsSocietes(){


}

func Menu(){
	var abattement float32 = 1
	var personne string
	fmt.Print("Êtes-vous une société ou une personne?\n1: Société\n2: Personne\n")
	fmt.Scan(&personne)
	switch personne {
	case "1":
		var statut string
		fmt.Print("1: Auto-Entreprise\n2: EURL\n3: SASU\nEntrez le numéro corresondant à votre statut juridique.\n")
		fmt.Scan(&statut)
	
		switch statut{
		case "1":
			var nbParts float32
			var administrationF string
			var ca float32
			var cotisations float32
			var fraischambre float32
			var contribution float32
			var totalcot float32
			fmt.Print("1: Aritsanale\n2: Commerciale.\n3: Libérale.\nEntrez le numéro correspondant à votre administration fiscale.\n")
			fmt.Scan(&administrationF)
			fmt.Print("Quel est votre nombre de parts?\n")
			fmt.Scan(&nbParts)
			fmt.Print("Quel est votre chiffre d'affaire annuel?\n")
			fmt.Scan(&ca)
			switch administrationF {
			case "1":
				var typeactivite string
				fmt.Print("1: Vente de biens\n2: Prestation de service\nEntre le numéro correspondant à votre type d'activité.\n")
				fmt.Scan(&typeactivite)
				switch typeactivite {
				case "1":
					cotisations, fraischambre, contribution = 12.3, 0.22, 0.3
					totalcot = CotisationsContributions(ca, cotisations, fraischambre, contribution)
					abattement = 0.29
				case "2":
					cotisations, fraischambre, contribution = 21.2, 0.48, 0.3
					totalcot = CotisationsContributions(ca, cotisations, fraischambre, contribution)
					abattement = 0.50
				}
			case "2":
				var typeactivite string
				fmt.Print("1: Vente de biens\n2: Prestation de service\nEntre le numéro correspondant à votre type d'activité.\n")
				fmt.Scan(&typeactivite)
				switch typeactivite {
				case "1":
					cotisations, fraischambre, contribution = 12.3, 0.04, 0.1
					totalcot = CotisationsContributions(ca, cotisations, fraischambre, contribution)
					abattement = 0.29
				case "2":
					cotisations, fraischambre, contribution = 21.2, 0.04, 0.1
					totalcot = CotisationsContributions(ca, cotisations, fraischambre, contribution)
					abattement = 0.50
				}
			case "3":
				cotisations, fraischambre, contribution = 23.1, 0, 0.2
				totalcot = CotisationsContributions(ca, cotisations, fraischambre, contribution)
				abattement = 0.66
			}
			res := ImpotsRevenus(nbParts, ca*abattement)
			fmt.Printf("Total de cotisations et contributions: %.0f\n", totalcot)
			totalrevenus := ca-totalcot-res
			fmt.Printf("Revenus total après impôts: %.0f\n", totalrevenus)
		case "2":
			//EURL
				//1 IR 2 IS
		case "3":
			//SASU
		}
	case "2":
		var nbParts float32
		var revenus float32
		fmt.Print("Entrez votre nombre  de parts: ")
		fmt.Scan(&nbParts)
		fmt.Print("Entrez vos revenus annuel après charges (réels ou abattement automatique de 10%): ")
		fmt.Scan(&revenus)
		ImpotsRevenus(nbParts, revenus)
	}
}

func CotisationsContributions(CA float32, cotisation float32, fdc float32, contribution float32) float32{
	return CA*(cotisation/100)+ CA*(fdc/100) + CA*(contribution/100)
}