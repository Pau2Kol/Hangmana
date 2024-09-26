package main

import (
	"bufio"
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
	rdmnbr := rand.Intn(len(lines))                               //choisi nombre aléatoire dans la limite
	selecmot := strings.ToUpper(strings.TrimSpace(lines[rdmnbr])) //met le mot en maj

	return strings.Split(selecmot, "") //divise le mot en mettant des espaces
}

func pendu(mot []string) {
	motC := strings.Join(mot, " ")
	fmt.Println(motC) // Print le mot caché faut penser a l'enlever c pour les tests
	motshown := strings.Split(motC, "")
	luse := []string{}
	for i, v := range motshown {
		if v >= "A" && v <= "Z" { //change le mot en _ sauf rdmindex
			motshown[i] = "_"
		}
	}
	motref := string(motC)
	for i := 0; len(mot)/2-1 > i; i++ {
		rdmindex := rdm(motshown)
		motshown[rdmindex] = string(motC[rdmindex])
	}
	fmt.Println("Bonne chance t'a 10 essais sinon: rm -rf / ")

	for i := 10; i > 0; {
		fmt.Println(strings.Join(motshown, "")) //faut supprimer debug
		guess := input(mot)
		if guess == strings.Join(mot,""){ //c vrmnt de la merde 4 ligne parce que j'ai la flemme si guess = mot a trouver
			 welive()
			 return
		}
		luse = append(luse, guess)                 //Prend l'input de l'user
		if !veriflettre(motref, guess, motshown) { // motshown == string[] / motC et motref == string
			i--
			printlependu(i)
		
		}
		if compare(motshown, motref) {
			welive()
			return
		}
		fmt.Print("\nLettre(s) déjà utilisés",luse,"\n")

	}
	fmt.Println("Nan le niveau c'est grave la")
}

func printlependu(i int) {
	fmt.Printf("Pas présent ou déjà mis, il te reste %d essais\n", i)
	file, err := os.Open("dic/hangman.txt")
	if err != nil {
		fmt.Println("Error opening hangman.txt:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	startLine := (9 - i) * 8 // José 8 ligne
	for scanner.Scan() {
		if lineCount >= startLine && lineCount < startLine+7 {
			fmt.Println(scanner.Text())
		}
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("frr le fichier iléou", err)
	}
}

func veriflettre(motref string, guess string, motshown []string) bool {
	c := false
	for i := range motref {
		if string(motref[i]) == guess && guess != motshown[i] {
			motshown[i] = string(motref[i])
			c = true
		}

	}
	return c
}

func rdm(motshown []string) int {
	rdmindex := rand.Intn(len(motshown))
	for motshown[rdmindex] != "_" {
		rdmindex = rand.Intn(len(motshown) - 1)
	}
	fmt.Println(rdmindex) //debug a enlever
	return rdmindex
}

func input(mot []string) string {
	var guess string
	fmt.Print("\nMot ou lettre :")
	fmt.Scanln(&guess)
	guess = strings.ToUpper(guess)
	if guess >= "A" && guess <= "Z" || guess == strings.Join(mot,"") {
		return guess
	}
	return input(mot)

}

func compare(motshown []string, motref string) bool { // aucun intêret j'ai fait une fonction pour une ligne
	return strings.Join(motshown, "") == motref
}
func welive() {
	str := `⠀⠀⠀⠀⠀⢀⡤⠖⠒⠢⢄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
	⠀⠀⠀⠀⠀⠀⠀⠀⡴⠃⠀⠀⠀⠀⠀⠙⢦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
	⠀⠀⠀⠀⠀⠀⠀⣰⠁⠀⠀⠀⠀⠀⠀⠀⠈⠳⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
	⠀⠀⠀⠀⠀⠀⡰⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠹⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
	⠀⠀⠀⠀⣠⠞⠁⠀⠀⠀⠀⠀⠀⠀⠂⠀⠤⠤⡀⠈⠳⣄⠀⠀⠀⠀⠀⠀⠀⠀
	⠀⠀⣠⠞⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠑⢄⠀⠀⠀⠀⠀⠀
	⢠⠞⠁⠀⣀⣠⣤⠤⠤⠤⠤⢤⣤⠤⠤⠤⠤⣤⣀⣀⡀⠀⠀⠀⠑⢤⠀⠀⠀⠀
	⣣⠔⠚⠻⣄⣡⣞⣄⣠⣆⠀⢼⣼⣄⣀⣀⣠⣆⠜⡘⡻⠟⠙⣲⠦⣈⢳⡀⠀⠀
	⡇⠒⢲⡤⡜⠉⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠉⠙⠛⠤⣖⠬⠓⠂⠉⣿⠇⠀⠀
	⠙⠲⠦⠬⣧⡀⠀⠀⠀⠀⠀⣠⣿⣿⣷⡄⠀⠀⠀⠀⠀⣞⠀⢀⣲⠖⠋⠀⠀⠀
	⠀⠀⠀⠀⠘⣟⢢⠃⠀⠀⠀⠉⠙⠻⠛⠁⠀⠀⠀⢀⡜⠒⢋⡝⠁⢀⣀⣤⠂⠀
	⠀⠀⠀⠀⠀⡇⠷⠆⠶⠖⠀⠀⠀⠀⠀⠀⠀⠀⣠⠮⠤⠟⠉⠀⢰⠱⡾⣧⠀⠀
	⠀⠀⠀⠀⠀⠹⢄⣀⣀⠀⠀⠀⠀⠀⠀⣀⡤⠚⠁⠀⢠⣤⡀⣼⢾⠀⠀⡟⠀⠀
	⠀⠀⠀⠀⠀⠀⠀⠀⠙⠛⠛⠒⡏⠀⡡⠣⢖⣯⠶⢄⣀⣿⡾⠋⢸⢀⡶⠿⠲⡀
	⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡰⣹⠃⣀⣤⠞⠋⠀⠉⠢⣿⣿⡄⠀⣿⠏⠀⠀⠐⢣
	⠀⠀⠀⠀⠀⠀⠀⠀⣠⠞⢱⢡⡾⠋⠀⠀⢀⡐⣦⣀⠈⠻⣇⢸⢁⣤⡙⡆⠈⡏
	⠀⠀⠀⠀⠀⠀⣠⠎⢁⠔⡳⡟⠀⠐⠒⠒⠋⠀⠠⡯⠙⢧⡈⠻⣮⠯⣥⠧⠞⠁
	⠀⠀⠀⣀⠴⠋⠀⢶⠋⢸⡝⠀⠀⠀⠀⠀⠀⠀⠀⣸⢦⠀⠙⡆⠘⠦⢄⡀⠀⠀
	⠀⠀⣸⠅⢀⡤⢺⢸⠀⢸⡃⠤⠀⠀⠀⠀⣀⡤⢚⣋⣿⢄⡀⢇⡀⠀⠀⣝⡶⠀
	⠀⠀⢿⠀⡏⠀⠘⠞⠀⢸⡵⣦⠤⠤⠖⣿⠥⠞⠉⠀⢸⠖⠁⠀⠙⠢⣑⠶⣽⢂
	⠀⠀⠸⠤⠃⠀⠀⠀⠀⠀⠉⢳⠂⠈⡽⠁⠀⠀⠀⢀⡼⠒⠓⢤⠀⠀⠀⠙⠚⠛
	⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠓⡎⠀⠀⠀⠀⢠⠎⣠⠀⠀⠈⢳⠀⠀⠀⠀⠀
	⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡇⠀⠀⢸⡶⠗⠋⣱⠄⠀⠀⠀⣧⠀⠀⠀⢀
	⠀⠀⠀⠀⠀⠀⠀⣀⠴⠒⠒⠦⣤⣷⠂⢀⡸⠁⠀⡼⠁⠀⠀⠀⠈⢺⠀⠀⠀⠀
	⠀⠀⠀⠀⠀⢠⠋⢀⣀⡀⠀⠀⠀⠀⠀⠈⡇⠀⠀⠙⠢⠤⠤⣄⡤⠼⠀⠀⠀⠀
	⠀⠀⠀⠀⠀⠀⠑⢦⣄⣉⣑⠢⠄⠀⠀⠀⡇`

	fmt.Print(str, "we live we love")
}
