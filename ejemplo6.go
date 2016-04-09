package main

import "fmt"

func main() {	

  appsEmp := "appsEmpresariales"

  for i := 1; i <= len(appsEmp); i++ {
    fmt.Println(appsEmp[0:i])
  }
  for i := 1; i <= len(appsEmp); i++ {
    fmt.Println(appsEmp[0:len(appsEmp)-i])
  }

}