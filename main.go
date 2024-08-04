package main

import (
	"fmt"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ROWS int32 = 9
const COLUMNS int32 = 9
const CELL_SIZE int32 = 150

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

func win() {
	const windowWidth int32 = COLUMNS * CELL_SIZE
	const windowHeight int32 = ROWS * CELL_SIZE

	rl.InitWindow(windowWidth, windowHeight, "Maze Generator")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	matrix := generateWallFilledMatrix()

	visitedIdx := 0
	fillInterval := float32(0.2)
	intervalAccumulator := float32(0.0)
	readyToDrawEdges := []*Edge{}

	resultMst := RunPrims(matrix, &GridNode{0, 0}, WALL_CHAR)
	if len(resultMst) != int(ROWS*COLUMNS)-1 {
		log.Fatalf("Incorrect result returned from Prims. Expected %d results but got %d", int(ROWS*COLUMNS)-1, len(resultMst))
	}

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()
		intervalAccumulator += fillInterval * dt

		if visitedIdx < len(resultMst) && intervalAccumulator > fillInterval {
			readyToDrawEdges = append(readyToDrawEdges, resultMst[visitedIdx])
			visitedIdx += 1
			intervalAccumulator = 0
		}

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

				if len(readyToDrawEdges) > 0 {
					for _, edge := range readyToDrawEdges {
						if r == int32(edge.To.Row) && c == int32(edge.To.Column) {
							rl.DrawLine(c*CELL_SIZE, r*CELL_SIZE, c*CELL_SIZE+CELL_SIZE, r*CELL_SIZE, rl.RayWhite)
							rl.DrawLine(c*CELL_SIZE, r*CELL_SIZE, c*CELL_SIZE, r*CELL_SIZE+CELL_SIZE, rl.RayWhite)
						}
					}
				}
			}
		}

		rl.EndDrawing()
	}
}

func main() {
	matrix := generateWallFilledMatrix()
    passages := RunMazeGeneratingPrims(matrix, &GridNode{0, 0}, WALL_CHAR)
    
    fmt.Println(len(passages))
    
    for r := range len(*matrix) {
        for c := range len((*matrix)[0]) {
            fmt.Print(string((*matrix)[r][c]))
        }
        fmt.Println()
    }
}
