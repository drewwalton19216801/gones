package cpu

// AddressingMode is an enum for addressing modes
type AddressingMode uint8

// Addressing modes
const (
	Implicit AddressingMode = iota
	Accumulator
	Immediate
	ZeroPage
	ZeroPageX
	ZeroPageY
	Relative
	Absolute
	AbsoluteX
	AbsoluteY
	Indirect
	IndexedIndirect
	IndirectIndexed
)

// String returns the string representation of the addressing mode
func (mode AddressingMode) String() string {
	switch mode {
	case Implicit:
		return "Implicit"
	case Accumulator:
		return "Accumulator"
	case Immediate:
		return "Immediate"
	case ZeroPage:
		return "ZeroPage"
	case ZeroPageX:
		return "ZeroPageX"
	case ZeroPageY:
		return "ZeroPageY"
	case Relative:
		return "Relative"
	case Absolute:
		return "Absolute"
	case AbsoluteX:
		return "AbsoluteX"
	case AbsoluteY:
		return "AbsoluteY"
	case Indirect:
		return "Indirect"
	case IndexedIndirect:
		return "IndexedIndirect"
	case IndirectIndexed:
		return "IndirectIndexed"
	default:
		return "Unknown"
	}
}

// --- Addressing modes ---
// Addressing mode functions calculate the effective address of an instruction,
// which is the address of the data that the instruction will operate on.
// Functions return the number of extra cycles that may be required to fetch
// the data.

func (c *CPU6502) implicit() int {
	// Implicit mode does nothing, but is implemented here for completeness
	return 0
}

func (c *CPU6502) accumulator() int {
	// Accumulator mode does nothing, but is implemented here for completeness
	return 0
}

// immediate sets the absolute address to the current program counter
// and increments the program counter.
//
// It returns the number of extra cycles used.
func (c *CPU6502) immediate() int {
	extraCyclesUsed := 0
	c.absoluteAddress = c.programCounter
	c.programCounter++

	return extraCyclesUsed
}

// zeroPage calculates the address from the zero page.
//
// It does not take any parameters.
// It returns 0, as it does not use any extra cycles.
func (c *CPU6502) zeroPage() int {
	extraCyclesUsed := 0
	c.absoluteAddress = uint16(c.bus.Read(c.programCounter) & 0xFF)
	c.programCounter++

	return extraCyclesUsed
}

// zeroPageX calculates the absolute address by adding the value in register X to the value read from the bus at the program counter.
//
// It then performs a bitwise AND operation to ensure the address stays within the zero page range.
// The program counter is incremented.
// The function returns the number of extra cycles used.
func (c *CPU6502) zeroPageX() int {
	extraCyclesUsed := 0
	c.absoluteAddress = uint16(c.bus.Read(c.programCounter)&0xFF) + uint16(c.x)
	c.absoluteAddress &= 0xFF

	c.programCounter++

	return extraCyclesUsed
}

// zeroPageY calculates the absolute address by adding the value in register Y to the value read from the bus at the program counter.
//
// It then performs a bitwise AND operation to ensure the address stays within the zero page range.
// The program counter is incremented.
// The function returns the number of extra cycles used.
func (c *CPU6502) zeroPageY() int {
	extraCyclesUsed := 0
	c.absoluteAddress = uint16(c.bus.Read(c.programCounter)&0xFF) + uint16(c.y)
	c.absoluteAddress &= 0xFF

	c.programCounter++

	return extraCyclesUsed
}

// relative performs a relative addressing mode operation in the CPU6502.
//
// It reads the value from the memory location pointed to by the program counter,
// and stores it in the temporary variable c.temp. Then, it sets the relativeAddress
// field of the CPU6502 struct to the value of c.temp.
//
// After that, it increments the program counter by 2 to point to the next instruction.
//
// The function does not use any extra cycles, so it returns 0.
func (c *CPU6502) relative() int {
	extraCyclesUsed := 0
	c.temp = uint16(c.bus.Read(c.programCounter))
	c.relativeAddress = c.temp

	c.programCounter += 2

	return extraCyclesUsed
}

// absolute calculates the absolute address.
//
// No parameters.
// Returns 0, as it does not use any extra cycles.
func (c *CPU6502) absolute() int {
	extraCyclesUsed := 0
	c.absoluteAddress = (uint16(c.bus.Read(c.programCounter)) << 8) | uint16(c.bus.Read(c.programCounter))

	c.programCounter += 2

	return extraCyclesUsed
}

// absoluteX performs an operation with the absolute addressing mode,
// using the value of the 'x' register as an offset. It calculates the
// absolute address by combining the value read from the program counter
// and the value of the 'x' register. If the operation crosses a page
// boundary, it adds an extra cycle. It returns the number of extra cycles
// used.
//
// No parameters.
// Returns an integer representing the number of extra cycles used.
func (c *CPU6502) absoluteX() int {
	extraCyclesUsed := 0
	c.absoluteAddress = (uint16(c.bus.Read(c.programCounter)) << 8) | uint16(c.bus.Read(c.programCounter))
	c.absoluteAddress += uint16(c.x)

	c.programCounter += 2

	// If page boundary is crossed, add an extra cycle
	if (c.absoluteAddress & 0xFF00) != ((c.absoluteAddress - uint16(c.x)) & 0xFF00) {
		extraCyclesUsed += 1
	}

	return extraCyclesUsed
}

// absoluteY calculates the absolute address by adding the value of the Y register to the
// address obtained from the program counter and the bus. It also checks if the page boundary
// is crossed and adds an extra cycle if necessary.
//
// It does not take any parameters.
// It returns an integer representing the number of extra cycles used.
func (c *CPU6502) absoluteY() int {
	extraCyclesUsed := 0
	c.absoluteAddress = (uint16(c.bus.Read(c.programCounter)) << 8) | uint16(c.bus.Read(c.programCounter))
	c.absoluteAddress += uint16(c.y)

	c.programCounter += 2

	// If page boundary is crossed, add an extra cycle
	if (c.absoluteAddress & 0xFF00) != ((c.absoluteAddress - uint16(c.y)) & 0xFF00) {
		extraCyclesUsed += 1
	}

	return extraCyclesUsed
}

// indirect performs an indirect addressing mode.
//
// It reads the value at the memory location specified by the programCounter and combines it with the next byte to form a 16-bit address.
// If the low byte of the resulting address is 0xFF, it fetches the high byte from the next memory location and combines it with the low byte to form the final address.
// Otherwise, it fetches the high byte from the memory location following the next byte and combines it with the low byte to form the final address.
//
// This simulates a bug in the original 6502.
//
// The programCounter is incremented by 1.
//
// Returns the number of extra cycles used.
func (c *CPU6502) indirect() int {
	extraCyclesUsed := 0
	c.temp = (uint16(c.bus.Read(c.programCounter)) << 8) | uint16(c.bus.Read(c.programCounter))

	if c.temp&0x00FF == 0x00FF {
		c.absoluteAddress = uint16(c.bus.Read(c.temp&0xFF00))<<8 | uint16(c.bus.Read(c.temp))
	} else {
		c.absoluteAddress = uint16(c.bus.Read(c.temp+1))<<8 | uint16(c.bus.Read(c.temp))
	}

	c.programCounter++

	return extraCyclesUsed
}

// indexedIndirect calculates the absolute address by adding the value in register X to the value read from the bus at the program counter.
//
// This function performs an indexed indirect addressing mode operation for the CPU6502.
// It reads the value at memory address (temp + x) and stores it in the absoluteAddress variable.
// The programCounter is then incremented.
// The function returns the number of extra cycles used.
func (c *CPU6502) indexedIndirect() int {
	extraCyclesUsed := 0
	c.temp = uint16(c.bus.Read(c.programCounter))
	c.absoluteAddress = uint16(c.bus.Read(uint16(c.temp) + uint16(c.x)&0xFF))

	c.programCounter++

	return extraCyclesUsed
}

// indirectIndexed calculates the indirect indexed addressing mode for the CPU6502.
//
// It retrieves the value at the address specified by the program counter and adds the value of the
// Y register to it. If the addition crosses a page boundary, an extra cycle is used. The function
// returns the number of extra cycles used.
func (c *CPU6502) indirectIndexed() int {
	extraCyclesUsed := 0
	c.temp = uint16(c.bus.Read(c.programCounter))
	c.absoluteAddress = uint16(c.bus.Read(c.temp)) + uint16(c.y)

	c.programCounter++

	// If page boundary is crossed, add an extra cycle
	if (c.absoluteAddress & 0xFF00) != ((c.absoluteAddress - uint16(c.y)) & 0xFF00) {
		extraCyclesUsed += 1
	}

	return extraCyclesUsed
}
