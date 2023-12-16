package cpu

type CPU6502 struct {
	a              byte
	x              byte
	y              byte
	programCounter uint16
	stackPointer   byte
	status         byte

	bus Bus

	fetched         byte
	temp            uint16
	absoluteAddress uint16
	relativeAddress uint16
	addressingMode  AddressingMode
	opcode          byte
	cycles          int
}

const (
	RegA byte = iota
	RegX
	RegY
	RegSP
	RegP
)

const (
	FlagC byte = 1 << iota
	FlagZ
	FlagI
	FlagD
	FlagB
	FlagU
	FlagV
	FlagN
)

func New() *CPU6502 {
	return &CPU6502{}
}

func (c *CPU6502) Reset() {
	// Reset internal registers
	c.a = 0
	c.x = 0
	c.y = 0
	c.programCounter = (uint16(c.bus.Read(0xFFFC)) << 8) | uint16(c.bus.Read(0xFFFD))
	c.stackPointer = 0xFD
	c.status = 0 | FlagU

	// Clear internal variables
	c.relativeAddress = 0
	c.absoluteAddress = 0
	c.fetched = 0
	c.temp = 0
	c.addressingMode = Implicit
	c.opcode = 0

	// Reset takes 8 cycles
	c.cycles = 8
}

// getRegister returns the value of the specified register in the CPU6502 struct.
//
// reg: The register to retrieve the value from.
// Returns: The value of the specified register.
func (c *CPU6502) GetRegister(reg byte) byte {
	switch reg {
	case RegA:
		return c.a
	case RegX:
		return c.x
	case RegY:
		return c.y
	case RegSP:
		return c.stackPointer
	case RegP:
		return c.status
	default:
		return 0
	}
}

func (c *CPU6502) GetPC() uint16 {
	return c.programCounter
}
