package main

import "fmt"

// Creating a pyramid using Go programming language.

func main() {

  var Rows int
  var Addition int = 0

  fmt.Print("Enter number of rows for the pyramid : \n")
  fmt.Scanln(&Rows)

// Creating space for object

  for i := 1; i <= Rows; i++ {
    Addition = 0
    for Space := 1; Space <= Rows-i; Space++ {
      fmt.Print("      ")
    }

// Generating object for pyramid

    for {
      fmt.Print("BLOCK ")
      Addition++
      if(Addition == 2*i-1){
        break
      }
    }

// Printing pyramid

    fmt.Println("")
  }
}
