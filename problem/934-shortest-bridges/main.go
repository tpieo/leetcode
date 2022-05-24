func shortestBridge(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    direction := []int{-1, 0, 1, 0, -1}

    // find an islang, set the found islang to 2
    q := make([][]int, 0)
    found := false
    for i := 0; i < m; i++ {
        if found {
            break
        }
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                found = true
                k := 0
                q = append(q, []int{i, j})
                grid[i][j] = 2
                for k < len(q) {
                    x, y := q[k][0], q[k][1]
                    k++
                    for d := 0; d < 4; d++ {
                        r, c := x + direction[d], y + direction[d+1]
                        if 0 <= r && r < m && 0 <= c && c < n && grid[r][c] == 1 {
                            q = append(q, []int{r, c})
                            grid[r][c] = 2
                        }
                    }
                }
                break
            }
        }
    }

    level := 0
    for len(q) > 0 {
        size := len(q)
        for size > 0 {
            x, y := q[0][0], q[0][1]
            q = q[1:]
            size--
            for k := 0; k < 4; k++ {
                r, c := x+direction[k], y+direction[k+1]
                if 0 <= r && r < m && 0 <= c && c < n {
                    if grid[r][c] == 2 {
                        continue
                    }
                    if grid[r][c] == 1 {
                        return level
                    }
                    q = append(q, []int{r, c})
                    grid[r][c] = 2
                }
            }
        }
        level++
    }
    return level
}