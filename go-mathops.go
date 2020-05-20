package main

// Simple math operations using "math" package

import (
  "fmt"
  "math"
)

func main(){

  var number1 float64
  var number2 float64
  var result float64

  fmt.Printf("Enter number 1 : \n")
  fmt.Scanln(&number1)
  fmt.Printf("Enter number 2 : \n")
  fmt.Scanln(&number2)

  // Power
  result = math.Pow(number1, 2)
  fmt.Printf("%v^%d = %v\n", number1, 2, result)

  // Sine
  result = math.Sin(number1)
  fmt.Printf("Sin(%v) = %v \n", number1, result)

  // Cosine
  result = math.Cos(number2)
  fmt.Printf("Cos(%v) = %v \n", number2, result)

  // Square root
  result = math.Sqrt(number1 * number2)
  fmt.Printf("Sqrt(%v * %v) = %v \n", number1, number2, result)
}
