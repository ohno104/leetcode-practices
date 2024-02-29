package main

// Time Complixity: O(N)
// Space Complixity: O(N)
func generateMatrix(n int) [][]int {
	sx, sy := 0, 0 // 記錄loop起始位置
	loop := n / 2  // 正方形
	center := n / 2
	count := 1  // 計算數值
	offset := 1 // 保持左閉右開
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	for loop > 0 {
		x, y := sx, sy

		for x < n-offset {
			res[y][x] = count
			x++
			count++
		}

		for y < n-offset {
			res[y][x] = count
			y++
			count++
		}

		for x > sx {
			res[y][x] = count
			x--
			count++
		}

		for y > sy {
			res[y][x] = count
			y--
			count++
		}

		sx += 1
		sy += 1
		offset += 1
		loop--
	}

	if n%2 == 1 {
		res[center][center] = n * n
	}

	return res
}
