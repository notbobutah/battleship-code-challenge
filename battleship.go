package main

import (
	"fmt"
	"math/rand"
	"time"
)

// following vars should be self explanitory
var intBoardSize = 10
var intTotalShips = 5
var intPlayer1ShipsLeft = 5
var intPlayer2ShipsLeft = 5
var arrPlayer1Board [10][10]int
var arrPlayer2Board [10][10]int
const emptycell = " "
const shipcell = "O"
const shipwreckcell = "+"
const missedcell = "-"

func main() {
	title()
	rand.Seed(time.Now().UnixNano())
	initBoard()
	randomShips(1)
	randomShips(2)
	printBoard(1)
	printBoard(2)
}
func title() {
	fmt.Println("+++++++++++++++++++++++++++")
	fmt.Println("        Battleship")
	fmt.Println("---------------------------")
}

//Initialized both boards with zero values
func initBoard() {
	i := 0
	j := 0
	intPlayer1ShipsLeft = 5
	intPlayer2ShipsLeft = 5

	for i = 0; i < intBoardSize; i++ {
		for j = 0; j < intBoardSize; j++ {
			arrPlayer1Board[i][j] = 0
			arrPlayer2Board[i][j] = 0
		}
	}
}

func randomShips(player int){
	intShipCounter := 0
	intRandomShipatX := 0
	intRandomShipatY := 0
	for {
		intRandomShipatX = rand.Intn(10)
		intRandomShipatY = rand.Intn(10)
		if(player ==1) {
			if arrPlayer1Board[intRandomShipatX][intRandomShipatY] == 0 {
				arrPlayer1Board[intRandomShipatX][intRandomShipatY] = 1
				intShipCounter = intShipCounter + 1
			}
		} else {
			if arrPlayer2Board[intRandomShipatX][intRandomShipatY] == 0 {
				arrPlayer2Board[intRandomShipatX][intRandomShipatY] = 1
				intShipCounter = intShipCounter + 1
			}
		}
		if intShipCounter == 5 {
			break
		}
	}
}

func printBoard(player int) {
	i := 0
	j := 0
	var arrBoard *[10][10]int
	if (player == 1) {
		arrBoard = &arrPlayer1Board;
	} else {
		arrBoard = &arrPlayer2Board;
	}

	intShipCounter := 0
	sChar := 'A'
	for i = 0; i < intBoardSize; i++ {
		fmt.Print(string(sChar))
		sChar = sChar + 1
		for j = 0; j < intBoardSize; j++ {
			switch arrBoard[i][j] {
			case 0:
				fmt.Print(emptycell)
			case 1:
				fmt.Print(shipcell)
				intShipCounter = intShipCounter + 1
			case 2:
				fmt.Print(shipwreckcell)
			case 3:
				fmt.Print(missedcell)
			}
		}
		fmt.Println("")
	}
	fmt.Println("Ships are O:", intShipCounter)
}