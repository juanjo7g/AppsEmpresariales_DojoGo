package main

import "fmt"

func main() {
	y := []int{
	  	48,96,86,68,
  		57,82,63,70,
  		37,34,83,27,
  		19,97, 9,17,
		}
	fmt.Println("El número menor es ", retorna_menor(y))
}

func retorna_menor(x []int) int {
	menor := x[0]
	for i := 0; i < len(x); i++ {
		if menor > x[i] {
			menor = x[i]		
		}
	}
	return menor
}