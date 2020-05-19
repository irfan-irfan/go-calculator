package main

import "fmt"

// Simple calculator created using Go programming language.

func main() {

  var nomor1, nomor2 int
  var operator string
  var hasil int = 0

  fmt.Print("Enter first number : ")
  fmt.Scanln(&nomor1)
  fmt.Print("Enter second number : ")
  fmt.Scanln(&nomor2)
  fmt.Print("Enter operator (+, -, *,/, %) : ")
  fmt.Scanln(&operator)

  switch operator {
    case "+":
      hasil = nomor1 + nomor2
    case "-":
      hasil = nomor1 - nomor2
    case "*":
      hasil = nomor1 * nomor2
    case "/":
      hasil = nomor1 / nomor2
    case "%":
      hasil = nomor1 % nomor2
    default:
      fmt.Println("ERROR")
    }
    fmt.Printf("%d %s %d = %d\n", nomor1, operator, nomor2, hasil)
}
