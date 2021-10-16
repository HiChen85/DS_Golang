package stack

// 迷宫老鼠程序

//func createMaze(row, column int) [][]int {
//	rand.Seed(time.Now().UnixNano())
//
//	maze := make([][]int, row)
//	for i := 0; i < len(maze); i++ {
//		maze[i] = make([]int, column)
//	}
//
//	for i := 0; i < row; i++ {
//		randomNum := column/5 + rand.Intn(column/3)
//		for j := 0; j < randomNum; j++ {
//			if i == 0 || i == 1 {
//				index := 2 + rand.Intn(column-2)
//				maze[i][index] = 1
//				if maze[i][index] == 1 {
//					index = index/2 + rand.Intn(column/4) + 1
//					maze[i][index] = 1
//				}
//			} else if i == row-1 || i == row-2 {
//				index := rand.Intn(column - 3)
//				maze[i][index] = 1
//				if maze[i][index] == 1 {
//					index = index/2 + rand.Intn(column/4) + 1
//					maze[i][index] = 1
//				}
//			} else {
//				index := rand.Intn(column)
//				maze[i][index] = 1
//				if maze[i][index] == 1 {
//					index = index/2 + rand.Intn(column/4) + 1
//					maze[i][index] = 1
//				}
//			}
//		}
//	}
//	for i := 0; i < row; i++ {
//		fmt.Println(maze[i])
//	}
//	return maze
//}

//func createMaze(n int) [][]int {
//	maze := make([][]int, n)
//	for i := 0; i < len(maze); i++ {
//		maze[i] = make([]int, n)
//	}
//}
//
//func TestMaze() {
//	row := 10
//	//column := 10
//	createMaze(row)
//
//}
