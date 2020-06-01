package main

import (
	"Tic-Tac-Toe/components"
	"Tic-Tac-Toe/service"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var board *components.Board
	for {
		flag := 0
		fmt.Println("Enter the size of the board: ")
		sizeNow, _ := reader.ReadString('\n')
		sizeNow = strings.TrimSpace(sizeNow)
		size, err := strconv.Atoi(sizeNow)
		// fmt.Println(size)
		if err != nil || size > 4 || size < 2 {
			fmt.Println("Please enter a valid number")
			flag = 1
		}
		if flag == 0 {
			board = components.NewBoard(uint8(size))
			break
		}
	}
	// fmt.Println(size)
	var player1 *components.Player
	fmt.Println("Enter the Name of first player: ")
	name1, _ := reader.ReadString('\n')
	name1 = strings.TrimSpace(name1)
	fmt.Println("Select mark 'X' or 'O'.")
	for {
		mark1 := ""
		flag := 0
		mark1, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			flag = 1
		}
		mark1 = strings.TrimSpace(mark1)
		// fmt.Println(mark1, mark1 == components.XMark)
		if mark1 != components.XMark && mark1 != components.OMark {
			fmt.Println("Please enter a valid mark 'X' or 'O'.")
			flag = 1
		}
		if flag == 0 {
			player1 = components.NewPlayer(name1, mark1)
			break
		}
	}

	var player2 *components.Player
	fmt.Println("Enter the Name of second player : ")
	name2, _ := reader.ReadString('\n')
	name2 = strings.TrimSpace(name2)
	mark2 := ""
	if player1.Mark == components.XMark {
		mark2 = components.OMark
	} else {
		mark2 = components.XMark
	}

	player2 = components.NewPlayer(name2, mark2)
	fmt.Println("")
	fmt.Println("----------Information of players----------")
	fmt.Printf("\tPlayer 1: Name:%s\t Mark:%s\n", player1.Name, player1.Mark)
	fmt.Printf("\tPlayer 2: Name:%s\t Mark:%s\n", player2.Name, player2.Mark)
	fmt.Println("")

	ourBoardService := service.NewBoardService(board)
	ourResultService := service.NewResultService(ourBoardService)
	ourGameService := service.NewGameService(ourResultService, [2]*components.Player{player1, player2})
	// fmt.Println(board, size)
	for {
		var result service.Result
		for {
			flag := 0
			fmt.Println(player1.Name, "enter your position:")
			indexNow, _ := reader.ReadString('\n')
			indexNow = strings.TrimSpace(indexNow)
			index1, err := strconv.Atoi(indexNow)
			if err != nil {
				fmt.Println(err)
				flag = 1
			}
			// fmt.Println(reflect.TypeOf(index1))
			index := uint8(index1)
			result, err = ourGameService.Play(uint8(index))
			if err != nil {
				fmt.Println(err)
				flag = 1
			}
			if flag == 0 {
				break
			}
		}
		// fmt.Println("gege")
		fmt.Println(ourGameService.PrintBoard())
		fmt.Println("")
		if result.Win == true {
			fmt.Println("----Winner----")
			fmt.Printf("Hurray %s won\n", player1.Name)
			fmt.Println("Thank you for playing this game.")
			break
		}
		if result.Draw == true {
			fmt.Println("----DRAW----")
			fmt.Println("Thank you for playing this game.")
			break
		}
		for {
			flag := 0
			fmt.Println(player2.Name, "enter your position:")
			indexNow, err := reader.ReadString('\n')
			indexNow = strings.TrimSpace(indexNow)
			index, err := strconv.Atoi(indexNow)
			if err != nil {
				fmt.Println(err)
				flag = 1
			}
			// fmt.Println("hdh")
			result, err = ourGameService.Play(uint8(index))
			if err != nil {
				fmt.Println(err)
				flag = 1
			}
			if flag == 0 {
				break
			}
		}
		fmt.Println(ourGameService.PrintBoard())
		fmt.Println("")
		if result.Win == true {
			fmt.Println("----Winner----")
			fmt.Printf("Hurray %s won\n", player2.Name)
			fmt.Println("Thank you for playing this game.")
			break
		}
		if result.Draw == true {
			fmt.Println("----DRAW----")

			fmt.Println("Thank you for playing this game.")
			break
		}
	}
}


