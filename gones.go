package main

import (
	"fmt"

	cpu6502 "github.com/drewwalton19216801/gones/cpu"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 680
	screenHeight = 480
	pixelWidth   = 2
	pixelHeight  = 2
)

func main() {
	cpu := cpu6502.New()
	mainbus := NewBus(cpu)
	cpu.ConnectBus(mainbus)
	cart := NewCartridge("test.nes")
	if cart.ImageValid() {
		mainbus.insertCartridge(cart)
	} else {
		fmt.Println("Failed to load cartridge")
		return
	}
	cpu.Reset()

	rl.InitWindow(680, 480, "Gones")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText(fmt.Sprintf("A: %d", cpu.GetRegister(cpu6502.RegA)), 10, 10, 20, rl.Red)
		rl.EndDrawing()
	}
}
