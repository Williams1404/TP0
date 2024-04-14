package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

const ruta = "archivo1.in"
const ruta1 = "archivo2.in"

func main() {
	vector1 := CopiarArchivoEnArreglo(ruta)
	vector2 := CopiarArchivoEnArreglo(ruta1)
	res := ejercicios.Comparar(vector1, vector2)
	switch res {
	case 1:
		Ordenar_Imprimir(vector1)
	case 2:
		Ordenar_Imprimir(vector2)
	case 0:
		Ordenar_Imprimir(vector1)
	}
}

func Ordenar_Imprimir(vector []int) {
	vector = ejercicios.Seleccion(vector)
	for i := 0; i < len(vector); i++ {
		fmt.Println(vector[i])
	}
}

func CopiarArchivoEnArreglo(ruta string) []int {
	var arreglo []int
	archivo, err := os.Open(ruta)
	if err != nil {
		fmt.Printf("error %v al abrir el archivo %s", ruta, err)
	}
	defer archivo.Close()
	s := bufio.NewScanner(archivo)
	for s.Scan() {
		numero, err := strconv.Atoi(s.Text())
		if err != nil {
			fmt.Println(err)
		}
		arreglo = append(arreglo, numero)

	}
	return arreglo
}
