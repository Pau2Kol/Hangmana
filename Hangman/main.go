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
	fmt.Print(mot)
	pendu(mot)
}

func mot() []string {
	//choisir un mot dans le txt
	//utiliser la function rdm de math pour avoir un mort aléatoire
	//mettre le mot en full maj + en Array
	//Donner le mot en full maj a la function main
	mot := []string{}
	fileIO, err := os.OpenFile("dic/words.txt", os.O_RDWR, 0600) // lit le fichier txt et vérifie les erreurs
	if err != nil {
		panic(err)
	}
	defer fileIO.Close()
	rawBytes, err := io.ReadAll(fileIO)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(rawBytes), "\n")

	rdmnbr := rand.Intn(84) //choisi un int random 0 < int < 84
	for i, line := range lines {
		if i == rdmnbr{ 
			mot = append(mot, line)
			break
		}
	}
	for i,v := range mot{
		mot[i] = strings.ToUpper(v)
	}
	return mot
}

func pendu(mot []string) {
	//vérifier si le mot contient la lettre
	//Print le pendu si oui
	//si non bah juste retourne le mot avec le nombre d'essai et la lettre en plus
}
