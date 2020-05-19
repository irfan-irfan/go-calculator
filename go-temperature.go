package main

import "fmt"

// Temperature conversion from Celsius
// to Fahrenheit, Kelvin, Rankine, and Reaumur
// using Go programming language.

func main(){

  var Celsius float64
  var Result float64
  var operator string

  fmt.Print("Enter Celsius temperature : \n")
  fmt.Scanln(&Celsius)
  fmt.Print("Convert to : Fahrenheit, Kelvin, Rankine, Reaumur\n")
  fmt.Print("Typing the scale is case-sensitive here, (e.g : Reaumur)\n")
  fmt.Print("Type scale : \n")
  fmt.Scanln(&operator)

  switch operator {
    case "Fahrenheit" :
      Result = (Celsius * 1.8) + 32
    case "Kelvin" :
      Result = Celsius + 273.15
    case "Rankine" :
      Result = (Celsius * 1.8) + 32 + 459.67
    case "Reaumur" :
      Result = Celsius * 0.8
    default :
      fmt.Println("ERROR")
  }
  fmt.Printf("%v degrees in Celsius is %v degrees in %v.\n", Celsius, Result, operator)
}
