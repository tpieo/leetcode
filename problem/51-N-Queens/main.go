// this is a slice trap, 
// if you pass a slice to a function and append it,
// slice will not change, though array add a element.
// because slice is a temporary variable like pointer.

var ans [][]string  

func solveNQueens(n int) [][]string {
	column := make([]bool, n)             // from top to bottom
	rightDiagonal := make([]bool, 2*n-1)  // from top to right
	leftDiagonal := make([]bool, 2*n-1)   // from top to left
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	ans = make([][]string, 0, n)
	dfs(0, n, board, column, rightDiagonal, leftDiagonal)
	return ans
}

func dfs(row, n int, board [][]byte, column, rightDiagonal, leftDiagonal []bool) {
	if row == n {
		ans = append(ans, makeBoard(board))
		return
	}
	for i := 0; i < n; i++ {
		if column[i] || rightDiagonal[n-1-row+i] || leftDiagonal[i+row] {
			continue
		}
		board[row][i] = 'Q'
		column[i] = true
		rightDiagonal[n-1-row+i] = true
		leftDiagonal[i+row] = true
		dfs(row+1, n, board, column, rightDiagonal, leftDiagonal)
		board[row][i] = '.'
		column[i] = false
		rightDiagonal[n-1-row+i] = false
		leftDiagonal[i+row] = false
	}
}

// change byte to string for result format.
// string is immutable, byte[] is changable.
func makeBoard(board [][]byte) (newBoard []string) {
	newBoard = make([]string, len(board))
	for i := 0; i < len(board); i++ {
		newBoard[i] = string(board[i])
	}
	return newBoard
}