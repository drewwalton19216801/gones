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

func New() *CPU6502 {
	return &CPU6502{}
}

func (c *CPU6502) Reset() {
	c.a = 0
	c.x = 0
	c.y = 0
	c.programCounter = 0
	c.stackPointer = 0
	c.status = 0
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
