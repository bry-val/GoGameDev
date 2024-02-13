package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	SCREEN_WIDTH  = 1080
	SCREEN_HEIGHT = 480
)

var (
	running  = true
	bkgColor = rl.NewColor(140, 211, 196, 255)
)

func drawScene() {

}

func input() {

}

func update() {
	running = !rl.WindowShouldClose()

}

func render() {
	rl.BeginDrawing()

	rl.ClearBackground(bkgColor)

	rl.EndDrawing()
}

func init() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Sproutlands 3D")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)
}

func quit() {
	rl.CloseWindow()
}
func main() {

	for running {
		input()
		update()
		render()
	}
	quit()
}
