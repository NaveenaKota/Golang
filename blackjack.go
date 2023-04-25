package main

import "fmt"
var dealerResult, firstResult, SecondResult int
var strPlayers = [3]string {"dealer", "player1", "player2"}
func main() {
	blackjack()
	if ( dealerResult > 21 && dealerResult < 17)  && (firstResult > 21 && firstResult < 17) && (SecondResult >21 && SecondResult < 17) {
		fmt.Println("No Wins")
	} else if (dealerResult ==21) {
		fmt.Println(strPlayers[0], "is Winner")	
	} else if (firstResult == 21) {
		fmt.Println(strPlayers[1], "is Winner")
	} else if (SecondResult == 21 ) {
		fmt.Println(strPlayers[2], "is Winner")
	} else if dealerResult > 21 && firstResult > 21 {
		fmt.Println(strPlayers[2], "is Winner")
	} else if dealerResult > 21 && SecondResult > 21 {
		fmt.Println(strPlayers[1], "is Winner")
	} else if SecondResult >21 && firstResult > 21 {
		fmt.Println(strPlayers[0], "is Winner")
	} else {
		if dealerResult > firstResult && dealerResult > SecondResult {
			if dealerResult > 21{
				if firstResult > SecondResult {
					fmt.Println(strPlayers[1], "is Winner")
				} else {
					fmt.Println(strPlayers[2], "is Winner")
				}
			} else {
				if dealerResult > 17 && dealerResult < 21 {
					fmt.Println(strPlayers[0], "is Winner")
				} else {
					fmt.Println("No Wins")
				}
			}
		} else if firstResult > dealerResult && firstResult > SecondResult {
			if firstResult > 21 {
				if dealerResult > SecondResult {
					fmt.Println(strPlayers[0], "is Winner")
				} else {
					fmt.Println(strPlayers[2], "is Winner")
				}
			} else {
				if firstResult > 17 && firstResult < 21 {
					fmt.Println(strPlayers[1], "is Winner")
				} else {
					fmt.Println("No Wins")
				}
			}
		} else {
			if SecondResult > 21 {
				if dealerResult > firstResult {
					fmt.Println(strPlayers[0], "is Winner")					
				} else {
					fmt.Println(strPlayers[1], "is Winner")
				}
			} else {
				if SecondResult > 17 && SecondResult < 21 {
					fmt.Println(strPlayers[1], "is Winner")
				} else {
					fmt.Println("No Wins")
				}
			}

		}
	}
}
func blackjack () {
	for i := 0; i < len(strPlayers); i++ {
		var first,second int
        fmt.Println(strPlayers[i],"Enter first card number: ")
        fmt.Scanln(&first)
        fmt.Println(strPlayers[i], "Enter second card number: ")
        fmt.Scanln(&second)
		sum := first + second
		if strPlayers[i] == "dealer" {
			dealerResult =sum
		} else if strPlayers[i] == "player1" {
			firstResult = sum
		} else {
			SecondResult =sum
		}
	}
}
