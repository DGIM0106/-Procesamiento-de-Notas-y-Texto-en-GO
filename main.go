package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	bidi := [6][4]float64{
		{9.5, 8.9, 8.6, 9.0},
		{7.5, 6.8, 8.0, 9.2},
		{9.3, 9.7, 8.5, 8.4},
		{6.5, 7.8, 8.9, 9.1},
		{8.0, 8.3, 7.5, 6.9},
		{9.9, 10.0, 8.7, 9.8},
	}
	sumanot, promedio := []float64{}, []float64{}
	fmt.Println("******Ejercicio 1******")
	fmt.Println("******Nota mas alta y baja del estudiante******")
	//suma de de notas por estudiante; alta y baja
	for i := 0; i < len(bidi); i++ {
		var sumaind, compaM, compam float64
		compaM = 0
		compam = 1000
		for j := 0; j < len(bidi[i]); j++ {
			sumaind += bidi[i][j]
			if bidi[i][j] > compaM {
				compaM = bidi[i][j]
			}
			if bidi[i][j] < compam {
				compam = bidi[i][j]
			}
		}
		fmt.Println("La nota mas alta del estudiante", i+1, "es:", compaM)
		fmt.Println("La nota mas baja del estudiante", i+1, "es:", compam)
		sumanot = append(sumanot, sumaind)

	}

	//sacar el promedio por estudante
	for i := 0; i < len(sumanot); i++ {
		prome := sumanot[i] / float64(4)
		promedio = append(promedio, prome)
	}
	//print promedio individual
	fmt.Println("******Promedio por estudiante******")
	for i := 0; i < len(promedio); i++ {
		fmt.Println("La nota promediada del estudiante", i+1, "es de:", promedio[i])
	}
	var sum float64
	for i := 0; i < len(promedio); i++ {
		sum += promedio[i]
	}
	proGen := sum / float64(6)
	fmt.Println("******Promedio general******")
	fmt.Println("Promedio general del curso es de:", proGen)

	fmt.Println("******Ejercicio 2******")
	//scaner con bufio
	fmt.Print("Ingrese el texto: ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	txtstr := scan.Text()

	//separar guardar, limpiar y ordenar las palabras
	var txtsep []string
	sep := strings.Split(txtstr, "/n")
	txtsep = append(txtsep, sep...)
	txtsep = strings.Fields(txtstr)

	//crear map y saber las veces que se repite
	contador := make(map[string]int)
	for _, palabra := range txtsep {
		contador[palabra]++
	}

	//lista ordenada alfa
	fmt.Println("Lista de palabras orenadas")
	for i := 0; i < len(txtsep); i++ {
		fmt.Println(txtsep[i])
	}

	//veces que se repite una palabra
	for palabra, veces := range contador {
		fmt.Println("La palabra \""+palabra+"\" se repite:", veces)
	}
}
