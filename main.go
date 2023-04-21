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
    var first int
    fmt.Println("Enter first number: ")
    fmt.Scanln(&first)
    var second int
    fmt.Println("Enter second number: ")
    fmt.Scanln(&second)
    var third int
    fmt.Println("Enter third number: ")
    fmt.Scanln(&third)
    var fourth,fifth float64
    fmt.Println("Enter fourth number: ")
    fmt.Scanln(&fourth)
    fmt.Println("Enter fifth number: ")
    fmt.Scanln(&fifth)
    
    fmt.Println(calculator.Add(first,second,third))
    fmt.Println(calculator.Sum(first,second))
    fmt.Println(calculator.Diff(first,second))
    fmt.Println(calculator.Sub(first,second))
    fmt.Println(calculator.Mult(first,second))
    fmt.Println(calculator.Div(first,second))
    fmt.Println(calculator.Sqrt(fourth))
    fmt.Println(calculator.Pow(fourth,fifth))
    fmt.Println(calculator.Power(first,second))
    fmt.Println(calculator.Pow10(first))
    fmt.Println(calculator.Sincos(fourth))
    fmt.Println(calculator.Exp(fourth))
    fmt.Println(calculator.Log(fourth))
}