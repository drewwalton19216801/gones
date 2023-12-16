package cpu

type Instruction uint8

// Legal instructions
const (
	_ Instruction = iota
	ADC
	AND
	ASL
	BCC
	BCS
	BEQ
	BIT
	BMI
	BNE
	BPL
	BRK
	BVC
	BVS
	CLC
	CLD
	CLI
	CLV
	CMP
	CPX
	CPY
	DEC
	DEX
	DEY
	EOR
	INC
	INX
	INY
	JMP
	JSR
	LDA
	LDX
	LDY
	LSR
	NOP
	ORA
	PHA
	PHP
	PLA
	PLP
	ROL
	ROR
	RTI
	RTS
	SBC
	SEC
	SED
	SEI
	STA
	STX
	STY
	TAX
	TAY
	TSX
	TXA
	TXS
	TYA
)

// InstructionNames is a map of instruction names
var InstructionNames = map[Instruction]string{
	ADC: "ADC",
	AND: "AND",
	ASL: "ASL",
	BCC: "BCC",
	BCS: "BCS",
	BEQ: "BEQ",
	BIT: "BIT",
	BMI: "BMI",
	BNE: "BNE",
	BPL: "BPL",
	BRK: "BRK",
	BVC: "BVC",
	BVS: "BVS",
	CLC: "CLC",
	CLD: "CLD",
	CLI: "CLI",
	CLV: "CLV",
	CMP: "CMP",
	CPX: "CPX",
	CPY: "CPY",
	DEC: "DEC",
	DEX: "DEX",
	DEY: "DEY",
	EOR: "EOR",
	INC: "INC",
	INX: "INX",
	INY: "INY",
	JMP: "JMP",
	JSR: "JSR",
	LDA: "LDA",
	LDX: "LDX",
	LDY: "LDY",
	LSR: "LSR",
	NOP: "NOP",
	ORA: "ORA",
	PHA: "PHA",
	PHP: "PHP",
	PLA: "PLA",
	PLP: "PLP",
	ROL: "ROL",
	ROR: "ROR",
	RTI: "RTI",
	RTS: "RTS",
	SBC: "SBC",
	SEC: "SEC",
	SED: "SED",
	SEI: "SEI",
	STA: "STA",
	STX: "STX",
	STY: "STY",
	TAX: "TAX",
	TAY: "TAY",
	TSX: "TSX",
	TXA: "TXA",
	TXS: "TXS",
	TYA: "TYA",
}

// InstructionInfo contains information about an instruction
type InstructionInfo struct {
	Instruction Instruction
	Opcode      uint8
	Mode        AddressingMode
	Cycles      uint8
	Execute     func(*CPU6502) int
}

var InstructionTable = map[uint8]InstructionInfo{
	0x69: {ADC, 0x69, Immediate, 2, (*CPU6502).adc},
	0x65: {ADC, 0x65, ZeroPage, 3, (*CPU6502).adc},
	0x75: {ADC, 0x75, ZeroPageX, 4, (*CPU6502).adc},
	0x6D: {ADC, 0x6D, Absolute, 4, (*CPU6502).adc},
	0x7D: {ADC, 0x7D, AbsoluteX, 4, (*CPU6502).adc},
	0x79: {ADC, 0x79, AbsoluteY, 4, (*CPU6502).adc},
	0x61: {ADC, 0x61, IndexedIndirect, 6, (*CPU6502).adc},
	0x71: {ADC, 0x71, IndirectIndexed, 5, (*CPU6502).adc},
	0x29: {AND, 0x29, Immediate, 2, (*CPU6502).and},
	0x25: {AND, 0x25, ZeroPage, 3, (*CPU6502).and},
	0x35: {AND, 0x35, ZeroPageX, 4, (*CPU6502).and},
	0x2D: {AND, 0x2D, Absolute, 4, (*CPU6502).and},
	0x3D: {AND, 0x3D, AbsoluteX, 4, (*CPU6502).and},
	0x39: {AND, 0x39, AbsoluteY, 4, (*CPU6502).and},
	0x21: {AND, 0x21, IndexedIndirect, 6, (*CPU6502).and},
	0x31: {AND, 0x31, IndirectIndexed, 5, (*CPU6502).and},
	0x0A: {ASL, 0x0A, Accumulator, 2, (*CPU6502).asl},
	0x06: {ASL, 0x06, ZeroPage, 5, (*CPU6502).asl},
	0x16: {ASL, 0x16, ZeroPageX, 6, (*CPU6502).asl},
	0x0E: {ASL, 0x0E, Absolute, 6, (*CPU6502).asl},
	0x1E: {ASL, 0x1E, AbsoluteX, 7, (*CPU6502).asl},
	0x90: {BCC, 0x90, Relative, 2, (*CPU6502).bcc},
	0xB0: {BCS, 0xB0, Relative, 2, (*CPU6502).bcs},
	0xF0: {BEQ, 0xF0, Relative, 2, (*CPU6502).beq},
	0x24: {BIT, 0x24, ZeroPage, 3, (*CPU6502).bit},
	0x2C: {BIT, 0x2C, Absolute, 4, (*CPU6502).bit},
	0x30: {BMI, 0x30, Relative, 2, (*CPU6502).bmi},
	0xD0: {BNE, 0xD0, Relative, 2, (*CPU6502).bne},
	0x10: {BPL, 0x10, Relative, 2, (*CPU6502).bpl},
	0x00: {BRK, 0x00, Implicit, 7, (*CPU6502).brk},
	0x50: {BVC, 0x50, Relative, 2, (*CPU6502).bvc},
	0x70: {BVS, 0x70, Relative, 2, (*CPU6502).bvs},
	0x18: {CLC, 0x18, Implicit, 2, (*CPU6502).clc},
	0xD8: {CLD, 0xD8, Implicit, 2, (*CPU6502).cld},
	0x58: {CLI, 0x58, Implicit, 2, (*CPU6502).cli},
	0xB8: {CLV, 0xB8, Implicit, 2, (*CPU6502).clv},
	0xC9: {CMP, 0xC9, Immediate, 2, (*CPU6502).cmp},
	0xC5: {CMP, 0xC5, ZeroPage, 3, (*CPU6502).cmp},
	0xD5: {CMP, 0xD5, ZeroPageX, 4, (*CPU6502).cmp},
	0xCD: {CMP, 0xCD, Absolute, 4, (*CPU6502).cmp},
	0xDD: {CMP, 0xDD, AbsoluteX, 4, (*CPU6502).cmp},
	0xD9: {CMP, 0xD9, AbsoluteY, 4, (*CPU6502).cmp},
	0xC1: {CMP, 0xC1, IndexedIndirect, 6, (*CPU6502).cmp},
	0xD1: {CMP, 0xD1, IndirectIndexed, 5, (*CPU6502).cmp},
	0xE0: {CPX, 0xE0, Immediate, 2, (*CPU6502).cpx},
	0xE4: {CPX, 0xE4, ZeroPage, 3, (*CPU6502).cpx},
	0xEC: {CPX, 0xEC, Absolute, 4, (*CPU6502).cpx},
	0xC0: {CPY, 0xC0, Immediate, 2, (*CPU6502).cpy},
	0xC4: {CPY, 0xC4, ZeroPage, 3, (*CPU6502).cpy},
	0xCC: {CPY, 0xCC, Absolute, 4, (*CPU6502).cpy},
	0xC6: {DEC, 0xC6, ZeroPage, 5, (*CPU6502).dec},
	0xD6: {DEC, 0xD6, ZeroPageX, 6, (*CPU6502).dec},
	0xCE: {DEC, 0xCE, Absolute, 6, (*CPU6502).dec},
	0xDE: {DEC, 0xDE, AbsoluteX, 7, (*CPU6502).dec},
	0xCA: {DEX, 0xCA, Implicit, 2, (*CPU6502).dex},
	0x88: {DEY, 0x88, Implicit, 2, (*CPU6502).dey},
	0x49: {EOR, 0x49, Immediate, 2, (*CPU6502).eor},
	0x45: {EOR, 0x45, ZeroPage, 3, (*CPU6502).eor},
	0x55: {EOR, 0x55, ZeroPageX, 4, (*CPU6502).eor},
	0x4D: {EOR, 0x4D, Absolute, 4, (*CPU6502).eor},
	0x5D: {EOR, 0x5D, AbsoluteX, 4, (*CPU6502).eor},
	0x59: {EOR, 0x59, AbsoluteY, 4, (*CPU6502).eor},
	0x41: {EOR, 0x41, IndexedIndirect, 6, (*CPU6502).eor},
	0x51: {EOR, 0x51, IndirectIndexed, 5, (*CPU6502).eor},
	0xE6: {INC, 0xE6, ZeroPage, 5, (*CPU6502).inc},
	0xF6: {INC, 0xF6, ZeroPageX, 6, (*CPU6502).inc},
	0xEE: {INC, 0xEE, Absolute, 6, (*CPU6502).inc},
	0xFE: {INC, 0xFE, AbsoluteX, 7, (*CPU6502).inc},
	0xE8: {INX, 0xE8, Implicit, 2, (*CPU6502).inx},
	0xC8: {INY, 0xC8, Implicit, 2, (*CPU6502).iny},
	0x4C: {JMP, 0x4C, Absolute, 3, (*CPU6502).jmp},
	0x6C: {JMP, 0x6C, Indirect, 5, (*CPU6502).jmp},
	0x20: {JSR, 0x20, Absolute, 6, (*CPU6502).jsr},
	0xA9: {LDA, 0xA9, Immediate, 2, (*CPU6502).lda},
	0xA5: {LDA, 0xA5, ZeroPage, 3, (*CPU6502).lda},
	0xB5: {LDA, 0xB5, ZeroPageX, 4, (*CPU6502).lda},
	0xAD: {LDA, 0xAD, Absolute, 4, (*CPU6502).lda},
	0xBD: {LDA, 0xBD, AbsoluteX, 4, (*CPU6502).lda},
	0xB9: {LDA, 0xB9, AbsoluteY, 4, (*CPU6502).lda},
	0xA1: {LDA, 0xA1, IndexedIndirect, 6, (*CPU6502).lda},
	0xB1: {LDA, 0xB1, IndirectIndexed, 5, (*CPU6502).lda},
	0xA2: {LDX, 0xA2, Immediate, 2, (*CPU6502).ldx},
	0xA6: {LDX, 0xA6, ZeroPage, 3, (*CPU6502).ldx},
	0xB6: {LDX, 0xB6, ZeroPageY, 4, (*CPU6502).ldx},
	0xAE: {LDX, 0xAE, Absolute, 4, (*CPU6502).ldx},
	0xBE: {LDX, 0xBE, AbsoluteY, 4, (*CPU6502).ldx},
	0xA0: {LDY, 0xA0, Immediate, 2, (*CPU6502).ldy},
	0xA4: {LDY, 0xA4, ZeroPage, 3, (*CPU6502).ldy},
	0xB4: {LDY, 0xB4, ZeroPageX, 4, (*CPU6502).ldy},
	0xAC: {LDY, 0xAC, Absolute, 4, (*CPU6502).ldy},
	0xBC: {LDY, 0xBC, AbsoluteX, 4, (*CPU6502).ldy},
	0x4A: {LSR, 0x4A, Accumulator, 2, (*CPU6502).lsr},
	0x46: {LSR, 0x46, ZeroPage, 5, (*CPU6502).lsr},
	0x56: {LSR, 0x56, ZeroPageX, 6, (*CPU6502).lsr},
	0x4E: {LSR, 0x4E, Absolute, 6, (*CPU6502).lsr},
	0x5E: {LSR, 0x5E, AbsoluteX, 7, (*CPU6502).lsr},
	0xEA: {NOP, 0xEA, Implicit, 2, (*CPU6502).nop},
	0x09: {ORA, 0x09, Immediate, 2, (*CPU6502).ora},
	0x05: {ORA, 0x05, ZeroPage, 3, (*CPU6502).ora},
	0x15: {ORA, 0x15, ZeroPageX, 4, (*CPU6502).ora},
	0x0D: {ORA, 0x0D, Absolute, 4, (*CPU6502).ora},
	0x1D: {ORA, 0x1D, AbsoluteX, 4, (*CPU6502).ora},
	0x19: {ORA, 0x19, AbsoluteY, 4, (*CPU6502).ora},
	0x01: {ORA, 0x01, IndexedIndirect, 6, (*CPU6502).ora},
	0x11: {ORA, 0x11, IndirectIndexed, 5, (*CPU6502).ora},
	0x48: {PHA, 0x48, Implicit, 3, (*CPU6502).pha},
	0x08: {PHP, 0x08, Implicit, 3, (*CPU6502).php},
	0x68: {PLA, 0x68, Implicit, 4, (*CPU6502).pla},
	0x28: {PLP, 0x28, Implicit, 4, (*CPU6502).plp},
	0x2A: {ROL, 0x2A, Accumulator, 2, (*CPU6502).rol},
	0x26: {ROL, 0x26, ZeroPage, 5, (*CPU6502).rol},
	0x36: {ROL, 0x36, ZeroPageX, 6, (*CPU6502).rol},
	0x2E: {ROL, 0x2E, Absolute, 6, (*CPU6502).rol},
	0x3E: {ROL, 0x3E, AbsoluteX, 7, (*CPU6502).rol},
	0x6A: {ROR, 0x6A, Accumulator, 2, (*CPU6502).ror},
	0x66: {ROR, 0x66, ZeroPage, 5, (*CPU6502).ror},
	0x76: {ROR, 0x76, ZeroPageX, 6, (*CPU6502).ror},
	0x6E: {ROR, 0x6E, Absolute, 6, (*CPU6502).ror},
	0x7E: {ROR, 0x7E, AbsoluteX, 7, (*CPU6502).ror},
	0x40: {RTI, 0x40, Implicit, 6, (*CPU6502).rti},
	0x60: {RTS, 0x60, Implicit, 6, (*CPU6502).rts},
	0xE9: {SBC, 0xE9, Immediate, 2, (*CPU6502).sbc},
	0xE5: {SBC, 0xE5, ZeroPage, 3, (*CPU6502).sbc},
	0xF5: {SBC, 0xF5, ZeroPageX, 4, (*CPU6502).sbc},
	0xED: {SBC, 0xED, Absolute, 4, (*CPU6502).sbc},
	0xFD: {SBC, 0xFD, AbsoluteX, 4, (*CPU6502).sbc},
	0xF9: {SBC, 0xF9, AbsoluteY, 4, (*CPU6502).sbc},
	0xE1: {SBC, 0xE1, IndexedIndirect, 6, (*CPU6502).sbc},
	0xF1: {SBC, 0xF1, IndirectIndexed, 5, (*CPU6502).sbc},
	0x38: {SEC, 0x38, Implicit, 2, (*CPU6502).sec},
	0xF8: {SED, 0xF8, Implicit, 2, (*CPU6502).sed},
	0x78: {SEI, 0x78, Implicit, 2, (*CPU6502).sei},
	0x85: {STA, 0x85, ZeroPage, 3, (*CPU6502).sta},
	0x95: {STA, 0x95, ZeroPageX, 4, (*CPU6502).sta},
	0x8D: {STA, 0x8D, Absolute, 4, (*CPU6502).sta},
	0x9D: {STA, 0x9D, AbsoluteX, 5, (*CPU6502).sta},
	0x99: {STA, 0x99, AbsoluteY, 5, (*CPU6502).sta},
	0x81: {STA, 0x81, IndexedIndirect, 6, (*CPU6502).sta},
	0x91: {STA, 0x91, IndirectIndexed, 6, (*CPU6502).sta},
	0x86: {STX, 0x86, ZeroPage, 3, (*CPU6502).stx},
	0x96: {STX, 0x96, ZeroPageY, 4, (*CPU6502).stx},
	0x8E: {STX, 0x8E, Absolute, 4, (*CPU6502).stx},
	0x84: {STY, 0x84, ZeroPage, 3, (*CPU6502).sty},
	0x94: {STY, 0x94, ZeroPageX, 4, (*CPU6502).sty},
	0x8C: {STY, 0x8C, Absolute, 4, (*CPU6502).sty},
	0xAA: {TAX, 0xAA, Implicit, 2, (*CPU6502).tax},
	0xA8: {TAY, 0xA8, Implicit, 2, (*CPU6502).tay},
	0xBA: {TSX, 0xBA, Implicit, 2, (*CPU6502).tsx},
	0x8A: {TXA, 0x8A, Implicit, 2, (*CPU6502).txa},
	0x9A: {TXS, 0x9A, Implicit, 2, (*CPU6502).txs},
	0x98: {TYA, 0x98, Implicit, 2, (*CPU6502).tya},
}
