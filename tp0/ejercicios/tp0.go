package ejercicios

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	*x, *y = *y, *x

}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {
	if len(vector) == 0 {
		return -1
	}
	var mayor int = vector[0]
	posicion := 0
	for i := 0; i < len(vector); i++ {
		if mayor < vector[i] {
			mayor = vector[i]
			posicion = i
		}
	}
	return posicion

}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.
func Comparar(vector1, vector2 []int) int {
	min := len(vector1)
	if len(vector2) < min {
		min = len(vector2)
	}
	for i := 0; i < min; i++ {
		if vector1[i] != vector2[i] {
			if CheckMayor(vector1[i], vector2[i]) {
				return 1
			}
			return -1
		}
	}
	if len(vector1) != len(vector2) {
		if CheckMayor(len(vector1), len(vector2)) {
			return 1
		}
		return -1
	}
	return 0
}

// funcion auxiliar para comparar vectores
func CheckMayor(vector1, vector2 int) bool {
	return vector1 > vector2
}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) []int {
	n := len(vector) - 1
	for n > 0 {
		max := Maximo(vector[:n+1])
		Swap(&vector[max], &vector[n])
		n -= 1
	}
	return vector
}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func Suma(vector []int) int {
	indice := 0
	return _Suma(&vector, &indice)
}
func _Suma(vector *[]int, indice *int) int {
	if len(*vector) == *indice {
		return 0
	}
	Actual := (*vector)[*indice]
	*indice++
	return Actual + _Suma(vector, indice)
}

// EsCadenaCapicua devuelve si la cadena es un palíndromo. Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func EsCadenaCapicua(cadena string) bool {
	return _capicua(cadena, 0, len(cadena)-1)
}

func _capicua(cadena string, inicio int, fin int) bool {
	if inicio >= fin {
		return true
	}
	if cadena[inicio] != cadena[fin] {
		return false
	}
	return _capicua(cadena, inicio+1, fin-1)

}
