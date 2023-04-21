package main

import (
    "fmt"
  //  "calculator"
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
    var example,ch string 
    fmt.Println("Enter check: ")
    fmt.Scanln(&example)
    ch  = example
    switch ch {
    case "OddEven":
        OddEvenFun(first)
    case "Fibnocci":
        fmt.Println(fibnocciFun(first))
    case "Prime":
        primeFun(first)
    default: 
        fmt.Println("Invalid input")
    }
 /*   var second int
    fmt.Println("Enter second number: ")
    fmt.Scanln(&second)*/
    /* var third int
    fmt.Println("Enter third number: ")
    fmt.Scanln(&third)
    var fourth,fifth float64
    fmt.Println("Enter fourth number: ")
    fmt.Scanln(&fourth)
    fmt.Println("Enter fifth number: ")
    fmt.Scanln(&fifth)
    
    fmt.Println(calculator.Add(first,second,third))
    fmt.Println(calculator.Diff(first,second))
    fmt.Println(calculator.Mult(first,second))
    fmt.Println(calculator.Div(first,second))
    fmt.Println(calculator.Sqrt(fourth))
    fmt.Println(calculator.Pow(fourth,fifth))
    fmt.Println(calculator.Pow10(first))
    fmt.Println(calculator.Sincos(fourth))
    fmt.Println(calculator.Exp(fourth))
    fmt.Println(calculator.Log(fourth)) */

}
func OddEvenFun(x int) {
    if x % 2 == 0 {
       fmt.Println(x, " is even")
    } else {
       fmt.Println(x, " is odd")
    }
} 
func primeFun(x int) {
    var flag int = 0
    if x == 0 || x ==1 {
        fmt.Println(x, "is not prime")
    } else {
        var i int =2
        for  ; i < x ; i++ {
            if x % i == 0 {
                flag = 1
                break
            }
        }
        if flag == 1 {
            fmt.Println(x, "is not prime")
        } else {
            fmt.Println(x, "is prime")
        }
    }
}

func fibnocciFun(x int) int{
    if x == 1 {
        return x
    } else if x == 2 {
        return x
    } else {
        return x * fibnocciFun(x -1)
    }
}  