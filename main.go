package main

import rl "github.com/gen2brain/raylib-go/raylib"

const ROWS int32 = 12
const COLUMNS int32 = 12
const CELL_SIZE int32 = 50

func main() {
	const windowWidth int32 = COLUMNS * CELL_SIZE
	const windowHeight int32 = ROWS * CELL_SIZE

	rl.InitWindow(windowWidth, windowHeight, "Maze Generator")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.EndDrawing()
	}
}
