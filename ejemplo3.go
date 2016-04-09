package main

import "fmt"

func main() {
	y := []int{
	  	48,96,86,68,
  		57,82,63,70,
  		37,34,83,27,
  		19,97, 9,17,
		}
	menor := y[0]
	for i := 0; i < len(y); i++ {
		if menor > y[i] {
			menor = y[i]		
		}
	}
	fmt.Println("El número menor es ", menor)

}

/*EJERCICIO: Escribir un programa que encuentre el numero menor en el 
siguiente arreglo y lo imprima:

x := []int{
  48,96,86,68,
  57,82,63,70,
  37,34,83,27,
  19,97, 9,17,
}*/