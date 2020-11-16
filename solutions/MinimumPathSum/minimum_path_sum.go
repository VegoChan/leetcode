package MinimumPathSum

//子问题：到点[i,j]的最小路径和MPS[i,j]
//MPS[i,j]=min{MPS[i-1,j],MPS[i,j_1]}+p[i,j]
//two dimensional -> one dimensional

func MinPathSum(grid [][]int) int {
	gridILen := len(grid[0])
	gridLen := len(grid)

	mps := make([]int,gridILen)
	mps[0] = grid[0][0]
	for i := 1; i < gridILen; i++{
		mps[i] = grid[0][i]+mps[i-1]
	}

	for i := 1; i < gridLen; i++{
		mps[0] += grid[i][0]
		for j := 1; j < gridILen; j++{
			mps[j] = min(mps[j-1],mps[j])+grid[i][j]
		}
	}

	return mps[gridILen-1]
}

func min(x int, y int) int{
	if x > y {
		return y
	}
	return x
}