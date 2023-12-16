package cpu

type CPU6502 struct {
	A  byte
	X  byte
	Y  byte
	PC uint16
	SP byte
	P  byte
}

func New() *CPU6502 {
	return &CPU6502{}
}

func (c *CPU6502) Reset() {
	c.A = 0
	c.X = 0
	c.Y = 0
	c.PC = 0
	c.SP = 0
	c.P = 0
}
