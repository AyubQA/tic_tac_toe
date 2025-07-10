package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type GameState int
type BoardField int

const (
	empty BoardField = iota
	cross
	nought
)

const (
	playing GameState = iota
	draw
	crossWin
	noughtWin
	quit
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	boardSize := 3 // Размер поля
	state := playing
	currentPlayer := cross // текущий игрок

	for {
		fmt.Print("Enter the size of the boards (3-9): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}

		input = strings.TrimSpace(input)

		size, err := strconv.Atoi(input)
		if err != nil {
			size = boardSize
		}

		if size < 3 || size > 9 {
			fmt.Println("Invalid input. Please try again.")
			continue
		}
		boardSize = size
		break
	}

	// Инициализация доски
	board := make([][]BoardField, boardSize)
	for i := range boardSize {
		board[i] = make([]BoardField, boardSize)
	}

	// Вывод в терминал состояния игрового поля
	fmt.Print(" ")
	for i := range boardSize {
		fmt.Printf("%d ", i+1) //Вывод номера столбца
	}

	fmt.Println()
	for i := range boardSize {
		fmt.Printf("%d ", i+1) // вывод номера строки
		for j := range boardSize {
			switch board[i][j] {
			case empty:
				fmt.Print(". ")
			case cross:
				fmt.Print("X ")
			case nought:
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
	//Завершение вывода в терминал

	// Основной игровой цикл
	for state == playing {
		playerSymbol := "X"
		if currentPlayer == nought {
			playerSymbol = "0"
		}
		fmt.Printf("%s turn, Enter row and column (e.g. 1 2): ",
			playerSymbol)

		validInput := false

		for !validInput {
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Invalid input. Please try again.")
				continue
			}

			input = strings.TrimSpace(input)
			if input == "q" {
				state = quit
				break
			}

			parts := strings.Fields(input)
			if len(parts) != 2 {
				fmt.Println("Invalid input. Please try again.")
				continue
			}

			row, err1 := strconv.Atoi(parts[0])
			col, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil {
				fmt.Println("Invalid input. Please try again.")
				continue
			}

			if row < 1 || row > boardSize || col < 1 || col > boardSize {
				fmt.Println("Invalid input. Please try again.")
				continue
			}

			// Приведение к 0-индексации
			row--
			col--

			if board[row][col] != empty {
				fmt.Println("This cell is already occupied!")
				continue
			}

			//Выполнения хода
			board[row][col] = currentPlayer

			//Проверка выигрыша по строкам и столбцам
			winFound := false
			for i := range boardSize {
				rowWin := true
				colWin := true
				for j := range boardSize {
					if board[i][j] == currentPlayer {
						rowWin = false
					}
					if board[j][i] == currentPlayer {
						colWin = false
					}
				}
				if rowWin || colWin {
					winFound = true
					break
				}
			}

			// Проверка главной диагонали
			if !winFound {
				diagWin := true
				for i := range boardSize {
					if board[i][i] != currentPlayer {
						diagWin = false
						break
					}
				}
				if diagWin {
					winFound = true
				}
			}

			// Проверка обратной диагонали
			if !winFound {
				antiDiagWin := true
				for i := range boardSize {
					if board[i][boardSize-i-1] != currentPlayer {
						antiDiagWin = false
						break
					}
				}
				if antiDiagWin {
					winFound = true
				}
			}

			if winFound {
				if currentPlayer == cross {
					state = crossWin
				} else {
					state = noughtWin
				}
			} else {
				// Проверка на ничью (заполнена ли доска)
				full := true
				for i := range boardSize {
					for j := range boardSize {
						if board[i][j] == empty {
							full = false
							break
						}
					}
					if !full {
						break
					}
				}
				if full {
					state = draw
				}
			}

			// Вывод текущего состояния доски
			fmt.Print(" ")
			for i := range boardSize {
				fmt.Printf("%d ", i+1)
			}
			fmt.Println()
			for i := range boardSize {
				fmt.Printf("%d ", i+1)
				for j := range boardSize {
					switch board[i][j] {
					case empty:
						fmt.Print(". ")
					case cross:
						fmt.Print("X ")
					case nought:
						fmt.Print("0 ")
					}
				}
				fmt.Println()
			}
			//Завершение вывода в терминал
		}

		// Вывод сообщения о результате, если игра окончена
		if state == crossWin {
			fmt.Println("X wins!")
		} else if state == noughtWin {
			fmt.Println("0 wins!")
		} else if state == draw {
			fmt.Println("Its a Draw!")
		} else {
			// Переключение игрока, если игра продолжается
			if currentPlayer == cross {
				currentPlayer = nought
			} else {
				currentPlayer = cross
			}
		}
		validInput = false
	}
}
