package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var n, in int
var matrice [][]int
var vizitatLat, coadaLat, parcLat []int
var vizitatAd, stivaAd []int

func citireMatriceAdiacenta() [][]int {
	file, err := ioutil.ReadFile("matriceAdiacenta.txt")

	if err != nil {
		fmt.Println("Eroare citire matrice de adiacenta.")
	}

	lines := strings.Split(string(file), "\r")
	n := len(lines)

	matrix := make([][]int, n)

	for i, line := range lines {
		matrix[i] = make([]int, n)

		vals := strings.Split(strings.TrimPrefix(line, "\n"), " ")

		for j, value := range vals {
			matrix[i][j], _ = strconv.Atoi(value)
		}
	}

	return matrix
}

func DFS(nod int) {
	/*Marcam nodul ca si vizitat si il adaugam in stiva*/
	vizitatAd[nod] = 1
	stivaAd = append(stivaAd, nod)

	for i := 0; i < n; i++ {
		/*Daca exista arc cu nodul curent si vecinul curent nu a fost vizitat*/
		if matrice[nod][i] == 1 && vizitatAd[i] == 0 {
			DFS(i)
		}
	}
}

func BFS(indexStart int) {
	/*Marcam nodul initial si il adaugam in coada*/
	vizitatLat[indexStart] = 1
	coadaLat = append(coadaLat, indexStart)

	for len(coadaLat) != 0 {
		index := coadaLat[0]

		for i := 0; i < n; i++ {
			/*Daca exista arc cu nodul curent si vecinul curent nu a fost vizitat*/
			if matrice[index][i] == 1 && vizitatLat[i]== 0{
				/*Adaugam nodul i in coada si il marcam ca si vizitat*/
				vizitatLat[i] = 1
				coadaLat = append(coadaLat,i)
			}
		}

		/*Adaugam nodul curent in lista de rezultate*/
		parcLat = append(parcLat, index)
		/*Eliminam nodul curent din coada*/
		coadaLat = coadaLat[1:]
	}
}

func printResult(ar []int) {
	for _, val := range ar {
		fmt.Print(fmt.Sprintf("%d ", val+1))
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	matrice = citireMatriceAdiacenta()
	n = len(matrice)

	/*Initializare vectori noduri vizitate pentru ambele parcurgeri*/
	vizitatAd = make([]int, n)
	vizitatLat = make([]int, n)

	/*Citire nod initial de unde incep parcurgerile*/
	fmt.Println("Introduceti nodul de pornire: ")
	text, _ := reader.ReadString('\n')
	indexStart, _ := strconv.Atoi(strings.TrimSuffix(text, "\n"))
	/*Nodurile sunt stocate in memorie de la pozitia 0*/
	indexStart--

	DFS(indexStart)
	fmt.Println("Parcurgere in adancime:")
	printResult(stivaAd)

	BFS(indexStart)
	fmt.Println("\nParcurgere in latime:")
	printResult(parcLat)

}
