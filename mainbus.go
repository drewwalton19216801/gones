package main

import (
	"fmt"

	cpu "github.com/drewwalton19216801/gones/cpu"
)

// Implements the Bus interface found in cpu/bus.go
type MainBus struct {
	cpu *cpu.CPU6502

	// Cartridge
	cartridge *Cartridge

	// 2K of RAM
	mem [2048]byte

	systemClockCounter uint32 // System clock counter
}

func NewBus(cpu *cpu.CPU6502) *MainBus {
	return &MainBus{
		cpu: cpu,
	}
}

func (b *MainBus) Read(addr uint16) byte {
	data := uint8(0)
	if b.cartridge.cpuRead(addr, &data) {
		// Cartridge space
	} else if addr <= 0x1FFF {
		// System RAM address range
		data = b.mem[addr&0x07FF]
	} else if addr >= 0x2000 && addr <= 0x3FFF {
		// PPU address range, log for now
		fmt.Printf("PPU read: 0x%04X = 0x%02X\n", addr, data)
	}
	return byte(data)
}

func (b *MainBus) Write(addr uint16, data byte) {
	if b.cartridge.cpuWrite(addr, data) {
		// The cartridge "sees all" and has the facility to veto
		// the propagation of the bus transaction if it requires.
		// This allows the cartridge to map any address to some
		// other data, including the facility to divert transactions
		// with other physical devices. The NES does not do this
		// but I figured it might be quite a flexible way of adding
		// "custom" hardware to the NES in the future!
	} else if addr <= 0x1FFF {
		// System RAM address range
		b.mem[addr&0x07FF] = data
	} else if addr >= 0x2000 && addr <= 0x3FFF {
		// PPU address range, log for now
		fmt.Printf("PPU write: 0x%04X = 0x%02X\n", addr, data)
	}
}

func (b *MainBus) insertCartridge(cartridge *Cartridge) {
	b.cartridge = cartridge
}

func (b *MainBus) Reset() {
	b.cpu.Reset()
	b.systemClockCounter = 0
}
