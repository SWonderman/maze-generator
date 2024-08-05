package main

import (
    "math/rand"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const ROWS int32 = 20
const COLUMNS int32 = 40
const CELL_SIZE int32 = 25

const WALL_CHAR byte = 'x'
const PASSAGE_CHAR byte = '+'

func generateWallFilledMatrix() *[][]byte {
	matrix := make([][]byte, ROWS)
	for i := range matrix {
		matrix[i] = make([]byte, COLUMNS)
	}

	for r := range ROWS {
		for c := range COLUMNS {
			matrix[r][c] = WALL_CHAR
		}
	}

	return &matrix
}

func drawReadyPassageNodes(row int32, column int32, readyToDrawPassages []*GridNode, passagesCount int) {
	for idx, gridNode := range readyToDrawPassages {
		passageColor := rl.LightGray
		if idx == len(readyToDrawPassages)-1 && passagesCount != len(readyToDrawPassages) {
			passageColor = rl.Red
		}

		if row == int32(gridNode.Row) && column == int32(gridNode.Column) {
			rl.DrawRectangle(column*CELL_SIZE, row*CELL_SIZE, CELL_SIZE, CELL_SIZE, passageColor)
		}
	}
}

func main() {
	const windowWidth int32 = COLUMNS * CELL_SIZE
	const windowHeight int32 = ROWS * CELL_SIZE

	rl.InitWindow(windowWidth, windowHeight, "Maze Generator")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	matrix := generateWallFilledMatrix()
	passages := RunMazeGeneratingPrims(*generateWallFilledMatrix(), &GridNode{rand.Intn(int(ROWS)), rand.Intn(int(COLUMNS))}, WALL_CHAR)

	visitedIdx := 0
	fillInterval := float32(0.01)
	intervalAccumulator := float32(0.0)
	readyToDrawPassages := []*GridNode{}

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()
		intervalAccumulator += dt

		if visitedIdx < len(passages) && intervalAccumulator > fillInterval {
			readyToDrawPassages = append(readyToDrawPassages, passages[visitedIdx])
			visitedIdx += 1
			intervalAccumulator = 0
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		for r := range ROWS {
			for c := range COLUMNS {
				cellColor := rl.Pink
				cell := (*matrix)[r][c]
				switch cell {
				case WALL_CHAR:
					cellColor = rl.Black
				}

				rl.DrawRectangle(c*CELL_SIZE, r*CELL_SIZE, CELL_SIZE, CELL_SIZE, cellColor)

				if len(readyToDrawPassages) > 0 {
					drawReadyPassageNodes(r, c, readyToDrawPassages, len(passages))
				}
			}
		}
		rl.EndDrawing()
	}
}
