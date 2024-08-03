package main

import rl "github.com/gen2brain/raylib-go/raylib"

const ROWS int32 = 12
const COLUMNS int32 = 12
const CELL_SIZE int32 = 50

const WALL_CHAR byte = 'x'

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

func main() {
	const windowWidth int32 = COLUMNS * CELL_SIZE
	const windowHeight int32 = ROWS * CELL_SIZE

	rl.InitWindow(windowWidth, windowHeight, "Maze Generator")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	matrix := generateWallFilledMatrix()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		for r := range ROWS {
			for c := range COLUMNS {
				cellColor := rl.White
				cell := (*matrix)[r][c]
				switch cell {
				case WALL_CHAR:
					cellColor = rl.DarkBlue
				}

				rl.DrawRectangle(c*CELL_SIZE, r*CELL_SIZE, CELL_SIZE, CELL_SIZE, cellColor)
			}
		}

		rl.EndDrawing()
	}
}
