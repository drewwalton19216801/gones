package main

import (
	"fmt"

	cpu6502 "github.com/drewwalton19216801/gones/cpu"
)

func main() {
	cpu := cpu6502.New()
	cpu.Reset()
	fmt.Println("gones booting up...")
}
