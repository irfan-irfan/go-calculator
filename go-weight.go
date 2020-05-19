package main

import "fmt"

// Conversion from g (grams) to kg, hg, dag, dg ,cg and mg
// using Go programming language.

func main () {

  var g float64
  var baseunit string
  var result float64

  fmt.Print("Converter for grams (g) to\n")
  fmt.Print("kg, hg, dag, dg, cg and mg\n")
  fmt.Print("\n")
  fmt.Print("Enter grams (g) : \n")
  fmt.Scanln(&g)
  fmt.Print("\n")
  fmt.Print("Enter base unit \n")
  fmt.Print("(kg, hg, dag, dg, cg and mg : )\n")
  fmt.Scanln(&baseunit)
  fmt.Print("\n")

  switch baseunit {
    case "kg":
      result = g * 0.001
    case "hg":
      result = g * 0.01
    case "dag":
      result = g * 0.1
    case "dg":
      result = g * 10
    case "cg":
      result = g * 100
    case "mg":
      result = g * 1000
    default:
      fmt.Print("ERROR")
  }
  fmt.Printf("%v g equals %v %v.\n", g, result, baseunit)
}
