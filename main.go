package main

import (
    "fmt"
    "calculator"
)
func main() {
    /*
    var a int64 = calculator.Add(2,3)
    var numFloat float64
	numFloat = float64(a)
    fmt.Printf("%.2f", numFloat)*/
    fmt.Println(calculator.Add(2,3,5))
    fmt.Println(calculator.Diff(2,3))
    fmt.Println(calculator.Mult(2,3))
    fmt.Println(calculator.Div(2,3))
    fmt.Println(calculator.Sqrt(9))
    fmt.Println(calculator.Pow(2,3))
    fmt.Println(calculator.Pow10(2))
    fmt.Println(calculator.Sincos(0))
    fmt.Println(calculator.Exp(2))
    fmt.Println(calculator.Log(4))
}