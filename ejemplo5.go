package main

import "fmt"

func main() {
	
	fmt.Print("Ingrese su nombre: ")
  	var input1 string
  	fmt.Scanln(&input1)

  	fmt.Print("Ingrese su estatura: ")
  	var input2 float64
  	fmt.Scanf("%f",&input2)


  	p := Persona{input1, input2}

	fmt.Println(p.correr())
}

type Persona struct{
	nombre string
	estatura float64
}
func (p Persona)correr() string {
	return p.nombre + " corriendo"
}

