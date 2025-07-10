package main

import "tic_tac_toe/game"

func main() {
	for {
		if game.InitBoard() {
			break
		}
	}
	game.Play()
}
