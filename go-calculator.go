package main

import "fmt"

func main() {

  var nomor1, nomor2 int
  var operator string

  fmt.Print("Masukkan angka pertama : ")
  fmt.Scanln(&nomor1)
  fmt.Print("Masukkan angka kedua : ")
  fmt.Scanln(&nomor2)
  fmt.Print("Masukkan operasi (+, -, /, *, %) : ")
  fmt.Scanln(&operator)

  hasil := 0

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
    fmt.Printf("%d %s %d = %d", nomor1, operator, nomor2, hasil)
}
