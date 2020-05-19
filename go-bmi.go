package main

import "fmt"

// BMI using Metric and Imperial
// Written in Go programming language.

func main (){

  var heightm float64
  var weightm float64
  var heighti float64
  var weighti float64
  var bmim float64
  var bmii float64
  var choose string

  // Metric or imperial choice.
  fmt.Print("Welcome to the BMI Calculator.\n")
  fmt.Print("1 for Metric / 2 for Imperial : \n")
  fmt.Scanln(&choose)
  fmt.Println("")

  switch choose {

  case "1":

  // Enter height and weight for metric.
    fmt.Print("Enter height in cm :\n")
    fmt.Scanln(&heightm)
    fmt.Println("")
    fmt.Print("Enter weight in kg :\n")
    fmt.Scanln(&weightm)
    fmt.Println("")

  // BMI formula for metric.
    bmim = (weightm * 10000) / (heightm * heightm)

  // Print metric BMI
    fmt.Printf("BMI in metric is %v\n", bmim)

  // If statements for BMI classification (for metric).
    if bmim < 18.5{
      fmt.Printf("Underweight\n")
    }else if bmim >= 18.5 && bmim <= 24.9{
      fmt.Printf("Normal\n")
    }else if bmim >= 25.0 && bmim <= 29.9{
      fmt.Printf("Overweight\n")
    }else if bmim >= 30.0 && bmim <= 34.9{
      fmt.Printf("Obesity class 1\n")
    }else if bmim >= 35.0 && bmim <= 39.9{
      fmt.Printf("Obesity class 2\n")
    }else{
      fmt.Printf("Obesity class 3\n")
    }

    fmt.Println("")

  case "2":

  // Enter height and weight for imperial.
    fmt.Print("Enter height in inches :\n")
    fmt.Scanln(&heighti)
    fmt.Println("")
    fmt.Print("Enter weight in lbs :\n")
    fmt.Scanln(&weighti)
    fmt.Println("")

  // BMI formula for imperial.
    bmii = (weighti * 703) / (heighti * heighti)

  // Print imperial BMI.
    fmt.Printf("BMI in imperial is %v\n", bmii)

  // If statements for BMI classification (for imperial).
    if bmii < 18.5{
      fmt.Printf("Underweight\n")
    }else if bmii >= 18.5 && bmii <= 24.9{
      fmt.Printf("Normal\n")
    }else if bmii >= 25.0 && bmii <= 29.9{
      fmt.Printf("Overweight\n")
    }else if bmii >= 30.0 && bmii <= 34.9{
      fmt.Printf("Obesity class 1\n")
    }else if bmii >= 35.0 && bmii <= 39.9{
      fmt.Printf("Obesity class 2\n")
    }else{
      fmt.Printf("Obesity class 3\n")
    }

    fmt.Println("")

  default:

  // Error message printing.
    fmt.Println("ERROR")
  }
}
