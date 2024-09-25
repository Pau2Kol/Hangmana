package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
)	

func main() {
	mot := mot()
	pendu(mot)
}

func mot() []string {
	fileIO, err := os.OpenFile("dic/words.txt", os.O_RDWR, 0600) //lit le fichier ta capté 
	if err != nil {
		panic(err)
	}
	defer fileIO.Close()

	rawBytes, err := io.ReadAll(fileIO)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(rawBytes), "\n")
	rdmnbr := rand.Intn(len(lines)) //choisi nombre aléatoire dans la limite
	selecmot := strings.ToUpper(strings.TrimSpace(lines[rdmnbr])) //met le mot en maj


	return strings.Split(selecmot, "") //divise le mot en mettant des espaces
}

func pendu(mot []string) {
	motC := strings.Join(mot, " ")
	fmt.Println(motC) // Print le mot caché faut penser a l'enlever c pour les tests
	caca := strings.Split(motC, "")

	rdmindex := rand.Intn(len(mot)-1)
	if caca[rdmindex] == " "{	//la je prend un index aléatoire qui n'est pas espace comme lettre de départ
		rdmindex += 1
	}
	for i,v := range caca{
		if v >= "A" && v <= "Z" && i != rdmindex{ //change le mot en _ sauf rdmindex
			caca[i] = "_"
		}
	}
	fmt.Println("Good luck, you have 10 attempts.")
	var lettre string
	for i := 10; i >= 0 ; i--{
		fmt.Println(strings.Join(caca, ""))	
		if !veriflettre(lettre){
			printlependu(i)
		}else{
			i++
		}
		//le plan c'était de demander une lettre puis vérifier si elle est est dans le mot avec une loop pour chaque caractère
		//et print le pendu plus message essai remaining si ct faux sauf que euuuuuuuuuh
		//vazy prendre un argument de la console ca ma casser les couilles
		//Bref faut: demander lettre -> update le mot si vrai print pendu avec la fonction printlependu si faux 
		//A FINIR
	}
}

func printlependu(i int){
	//bah elle doit faire son nom quoi 
}

func veriflettre(lettre string) bool{
	return false
}