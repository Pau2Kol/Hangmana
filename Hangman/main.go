package main

import (
	"fmt"
	"io"
	"math/rand/v2"
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
	fileIO, err := os.OpenFile("dic/words.txt", os.O_RDWR, 0600)
	if err != nil {
		panic(err)
	}
	defer fileIO.Close()
	rawBytes, err := io.ReadAll(fileIO)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(rawBytes), "\n")
	for i, line := range lines {
		if i == rand.IntN(84) {
			mot = append(mot, line)
			break
		}
	}
	return mot
}

func pendu(mot []string) {
	//vérifier si le mot contient la lettre
	//Print le pendu si oui
	//si non bah juste retourne le mot avec le nombre d'essai et la lettre en plus
}
