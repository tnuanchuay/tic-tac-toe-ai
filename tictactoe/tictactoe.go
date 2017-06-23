package tictactoe

import (
	"fmt"
)

var playBoard [3][3]int

type gameTree struct{
	board	[3][3]int
	tree	[9]*gameTree
}

var count = 0
var pickToWin = -1
var levelToWin = 99;
func PlayerPick(row, col int) int{
	count = 0
	pickToWin = -1

	fmt.Println(row, col)
	if playBoard[row][col] == 0{
		playBoard[row][col] = 1
		AI()
		aiPick(pickToWin)
		fmt.Println("AI select", pickToWin, "on level", levelToWin)
		return pickToWin
	}



	return -1
}

func aiPick(elem int){
	if(elem == 0){
		playBoard[0][0] = 2
	}
	if(elem == 1){
		playBoard[1][0] = 2
	}
	if(elem == 2){
		playBoard[2][0] = 2
	}
	if(elem == 3){
		playBoard[0][1] = 2
	}
	if(elem == 4){
		playBoard[1][1] = 2
	}
	if(elem == 5){
		playBoard[2][1] = 2
	}
	if(elem == 6){
		playBoard[0][2] = 2
	}
	if(elem == 7){
		playBoard[1][2] = 2
	}
	if(elem == 8){
		playBoard[2][2] = 2
	}
}

func AI(){
	var tree gameTree
	tree.board = playBoard
	recursiveFx(&tree, 2, 0, 0)
	fmt.Println(count, "ways to win")
}

func Win(board [3][3]int, level, root int){
	count++
	if level < levelToWin{
		pickToWin = root
		levelToWin = level
	}

	for i := 0; i < 3; i++{
		for j := 0; j < 3; j++{
			fmt.Print(board[j][i])
		}
		fmt.Println()
	}
	fmt.Println(level)
}

func recursiveFx(gt *gameTree, mode, level, root int){
	available := 0

	if isWin(gt.board, 1){
		return
	}

	if isWin(gt.board, 2){
		return
	}

	for i := 0; i < 3; i++{
		for j := 0; j < 3; j++{
			if gt.board[j][i] == 0{
				available++
			}
		}
	}

	if available == 0{
		return
	}

	nextMode := mode
	if mode == 1{
		nextMode = 2
	}else{
		nextMode = 1
	}

	for i := 0; i < len(gt.tree); i++{
		gt.tree[i] = &gameTree{}
		gt.tree[i].board = gt.board
		if i == 0{
			if gt.tree[i].board[0][0] == 0{
				gt.tree[i].board[0][0] = mode
				if(!isWin(gt.tree[i].board, 2)){
					if(level == 0){
						recursiveFx(gt.tree[i], nextMode, level+1, i)
					}else{
						recursiveFx(gt.tree[i], nextMode, level+1, root)
					}

				}else{
					Win(gt.tree[i].board, level+1, i)
				}
			}
		}else if i == 1{
			if gt.tree[i].board[1][0] == 0 {
				gt.tree[i].board[1][0] = mode
				if(!isWin(gt.tree[i].board, 2)){
					if(level == 0){
						recursiveFx(gt.tree[i], nextMode, level+1, i)
					}else{
						recursiveFx(gt.tree[i], nextMode, level+1, root)
					}
				}else{
					Win(gt.tree[i].board, level+1, i)
				}
			}
		}else if i == 2{
			if gt.tree[i].board[2][0] == 0 {
				gt.tree[i].board[2][0] = mode
				if(!isWin(gt.tree[i].board, 2)){
					if(level == 0){
						recursiveFx(gt.tree[i], nextMode, level+1, i)
					}else{
						recursiveFx(gt.tree[i], nextMode, level+1, root)
					}
				}else{
					Win(gt.tree[i].board, level+1, i)
				}
			}
		}else if i == 3{
			if gt.tree[i].board[0][1] == 0 {
				gt.tree[i].board[0][1] = mode
				if(!isWin(gt.tree[i].board, 2)){
					if(level == 0){
						recursiveFx(gt.tree[i], nextMode, level+1, 3)
					}else{
						recursiveFx(gt.tree[i], nextMode, level+1, root)
					}
				}else{
					Win(gt.tree[i].board, level+1, i)
				}
			}
		}else if i == 4{
			if gt.tree[i].board[1][1] == 0 {
				gt.tree[i].board[1][1] = mode
				if(!isWin(gt.tree[i].board, 2)){
					if(level == 0){
						recursiveFx(gt.tree[i], nextMode, level+1, i)
					}else{
						recursiveFx(gt.tree[i], nextMode, level+1, root)
					}
				}else{
					Win(gt.tree[i].board, level+1, i)
				}
			}
		}else if i == 5{
			if gt.tree[i].board[2][1] == 0 {
				gt.tree[i].board[2][1] = mode
				if(!isWin(gt.tree[i].board, 2)){
					if(level == 0){
						recursiveFx(gt.tree[i], nextMode, level+1, 5)
					}else{
						recursiveFx(gt.tree[i], nextMode, level+1, root)
					}
				}else{
					Win(gt.tree[i].board, level+1, i)
				}
			}
		}else if i == 6{
			if gt.tree[i].board[0][2] == 0 {
				gt.tree[i].board[0][2] = mode
				if(!isWin(gt.tree[i].board, 2)){
					if(level == 0){
						recursiveFx(gt.tree[i], nextMode, level+1, i)
					}else{
						recursiveFx(gt.tree[i], nextMode, level+1, root)
					}
				}else{
					Win(gt.tree[i].board, level+1, i)
				}
			}
		}else if i == 7{
			if gt.tree[i].board[1][2] == 0 {
				gt.tree[i].board[1][2] = mode
				if(!isWin(gt.tree[i].board, 2)){
					if(level == 0){
						recursiveFx(gt.tree[i], nextMode, level+1, i)
					}else{
						recursiveFx(gt.tree[i], nextMode, level+1, root)
					}
				}else{
					Win(gt.tree[i].board, level+1, i)
				}
			}
		}else if i == 8{
			if gt.tree[i].board[2][2] == 0 {
				gt.tree[i].board[2][2] = mode
				if(!isWin(gt.tree[i].board, 2)){
					if(level == 0){
						recursiveFx(gt.tree[i], nextMode, level+1, i)
					}else{
						recursiveFx(gt.tree[i], nextMode, level+1, root)
					}
				}else{
					Win(gt.tree[i].board, level+1, i)
				}
			}
		}
	}
}

func isWin(board [3][3]int, player int)bool{
	if(board[0][0] == player && board[0][1] == player && board[0][2] == player){
		return true
	}

	if(board[1][0] == player && board[1][1] == player && board[1][2] == player){
		return true
	}

	if(board[2][0] == player && board[2][1] == player && board[2][2] == player){
		return true
	}

	if(board[0][0] == player && board[1][0] == player && board[2][0] == player){
		return true
	}

	if(board[0][1] == player && board[1][1] == player && board[2][1] == player){
		return true
	}

	if(board[0][2] == player && board[1][2] == player && board[2][2] == player){
		return true
	}

	if(board[0][0] == player && board[1][0] == player && board[2][0] == player){
		return true
	}

	if(board[0][0] == player && board[1][1] == player && board[2][2] == player){
		return true
	}

	if(board[2][0] == player && board[1][1] == player && board[0][2] == player){
		return true
	}

	return false
}