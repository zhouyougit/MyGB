package gb

type OPCode struct {
	Func func(*Cpu) byte
	Cycles [2]int
	Length uint16
	Mnemonic string
}

var OPCodesMap = [0x100]OPCode {
	0x00: {
		Func: (*Cpu).opCode00,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "NOP",
	},
	0x3E: {
		Func: (*Cpu).opCode3E,
		Cycles: [2]int{8, 0},
		Length: 2,
		Mnemonic: "LD A,d8",
	},
	0xA0: {
		Func: (*Cpu).opCodeA0,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "AND B",
	},
	0xA1: {
		Func: (*Cpu).opCodeA1,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "AND C",
	},
	0xA2: {
		Func: (*Cpu).opCodeA2,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "AND D",
	},
	0xA3: {
		Func: (*Cpu).opCodeA3,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "AND E",
	},
	0xA4: {
		Func: (*Cpu).opCodeA4,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "AND H",
	},
	0xA5: {
		Func: (*Cpu).opCodeA5,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "AND L",
	},
	0xA6: {
		Func: (*Cpu).opCodeA6,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "AND (HL)",
	},
	0xA7: {
		Func: (*Cpu).opCodeA7,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "AND A",
	},
	0xA8: {
		Func: (*Cpu).opCodeA8,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "XOR B",
	},
	0xA9: {
		Func: (*Cpu).opCodeA9,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "XOR C",
	},
	0xAA: {
		Func: (*Cpu).opCodeAA,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "XOR D",
	},
	0xAB: {
		Func: (*Cpu).opCodeAB,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "XOR E",
	},
	0xAC: {
		Func: (*Cpu).opCodeAC,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "XOR H",
	},
	0xAD: {
		Func: (*Cpu).opCodeAD,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "XOR L",
	},
	0xAE: {
		Func: (*Cpu).opCodeAE,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "XOR (HL)",
	},
	0xAF: {
		Func: (*Cpu).opCodeAF,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "XOR A",
	},
	0xB0: {
		Func: (*Cpu).opCodeB0,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "OR B",
	},
	0xB1: {
		Func: (*Cpu).opCodeB1,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "OR C",
	},
	0xB2: {
		Func: (*Cpu).opCodeB2,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "OR D",
	},
	0xB3: {
		Func: (*Cpu).opCodeB3,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "OR E",
	},
	0xB4: {
		Func: (*Cpu).opCodeB4,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "OR H",
	},
	0xB5: {
		Func: (*Cpu).opCodeB5,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "OR L",
	},
	0xB6: {
		Func: (*Cpu).opCodeB6,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "OR (HL)",
	},
	0xB7: {
		Func: (*Cpu).opCodeB7,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "OR A",
	},
	0xC2: {
		Func: (*Cpu).opCodeC2,
		Cycles: [2]int{16, 12},
		Length: 3,
		Mnemonic: "JP NZ,a16",
	},
	0xC3: {
		Func: (*Cpu).opCodeC3,
		Cycles: [2]int{16, 0},
		Length: 3,
		Mnemonic: "JP a16",
	},
	0xE0: {
		Func: (*Cpu).opCodeE0,
		Cycles: [2]int{12, 0},
		Length: 2,
		Mnemonic: "LDH (a8),A",
	},
	0xE6: {
		Func: (*Cpu).opCodeE6,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "AND d8",
	},
	0xEE: {
		Func: (*Cpu).opCodeEE,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "XOR d8",
	},
	0xF0: {
		Func: (*Cpu).opCodeF0,
		Cycles: [2]int{12, 0},
		Length: 2,
		Mnemonic: "LDH A,(a8)",
	},
	0xF3: {
		Func: (*Cpu).opCodeF3,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "DI",
	},
	0xF6: {
		Func: (*Cpu).opCodeF6,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "OR d8",
	},
	0xFB: {
		Func: (*Cpu).opCodeFB,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "EI",
	},
}

// NOP
func (cpu *Cpu)opCode00() byte {
	return 0
}

// LD A,d8
func (cpu *Cpu)opCode3E() byte {
	cpu.reg.A = cpu.gb.Mem.Read(cpu.reg.PC)
	cpu.reg.PC++
	return 0
}

// AND B
func (cpu *Cpu)opCodeA0() byte {
	cpu.reg.A = cpu.reg.A & cpu.reg.B
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	cpu.setFlagH()
	return 0
}

// AND C
func (cpu *Cpu)opCodeA1() byte {
	cpu.reg.A = cpu.reg.A & cpu.reg.C
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	cpu.setFlagH()
	return 0
}

// AND D
func (cpu *Cpu)opCodeA2() byte {
	cpu.reg.A = cpu.reg.A & cpu.reg.D
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	cpu.setFlagH()
	return 0
}

// AND E
func (cpu *Cpu)opCodeA3() byte {
	cpu.reg.A = cpu.reg.A & cpu.reg.E
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	cpu.setFlagH()
	return 0
}

// AND H
func (cpu *Cpu)opCodeA4() byte {
	cpu.reg.A = cpu.reg.A & cpu.reg.H
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	cpu.setFlagH()
	return 0
}

// AND L
func (cpu *Cpu)opCodeA5() byte {
	cpu.reg.A = cpu.reg.A & cpu.reg.L
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	cpu.setFlagH()
	return 0
}

// AND (HL)
func (cpu *Cpu)opCodeA6() byte {
	cpu.reg.A = cpu.reg.A & cpu.gb.Mem.Read(cpu.getHL())
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	cpu.setFlagH()
	return 0
}

// AND A
func (cpu *Cpu)opCodeA7() byte {
	cpu.reg.A = cpu.reg.A & cpu.reg.A
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	cpu.setFlagH()
	return 0
}

// XOR B
func (cpu *Cpu)opCodeA8() byte {
	cpu.reg.A = cpu.reg.A ^ cpu.reg.B
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	return 0
}

// XOR C
func (cpu *Cpu)opCodeA9() byte {
	cpu.reg.A = cpu.reg.A ^ cpu.reg.C
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	return 0
}

// XOR D
func (cpu *Cpu)opCodeAA() byte {
	cpu.reg.A = cpu.reg.A ^ cpu.reg.D
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	return 0
}

// XOR E
func (cpu *Cpu)opCodeAB() byte {
	cpu.reg.A = cpu.reg.A ^ cpu.reg.E
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	return 0
}

// XOR H
func (cpu *Cpu)opCodeAC() byte {
	cpu.reg.A = cpu.reg.A ^ cpu.reg.H
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	return 0
}

// XOR L
func (cpu *Cpu)opCodeAD() byte {
	cpu.reg.A = cpu.reg.A ^ cpu.reg.L
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	return 0
}

// XOR (HL)
func (cpu *Cpu)opCodeAE() byte {
	cpu.reg.A = cpu.reg.A ^ cpu.gb.Mem.Read(cpu.getHL())
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	return 0
}

// XOR A
func (cpu *Cpu)opCodeAF() byte {
	cpu.reg.A = cpu.reg.A ^ cpu.reg.A
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	return 0
}

// OR B
func (cpu *Cpu)opCodeB0() byte {
	cpu.reg.A = cpu.reg.A | cpu.reg.B
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	return 0
}

// OR C
func (cpu *Cpu)opCodeB1() byte {
	cpu.reg.A = cpu.reg.A | cpu.reg.C
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	return 0
}

// OR D
func (cpu *Cpu)opCodeB2() byte {
	cpu.reg.A = cpu.reg.A | cpu.reg.D
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	return 0
}

// OR E
func (cpu *Cpu)opCodeB3() byte {
	cpu.reg.A = cpu.reg.A | cpu.reg.E
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	return 0
}

// OR H
func (cpu *Cpu)opCodeB4() byte {
	cpu.reg.A = cpu.reg.A | cpu.reg.H
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	return 0
}

// OR L
func (cpu *Cpu)opCodeB5() byte {
	cpu.reg.A = cpu.reg.A | cpu.reg.L
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	return 0
}

// OR (HL)
func (cpu *Cpu)opCodeB6() byte {
	cpu.reg.A = cpu.reg.A | cpu.gb.Mem.Read(cpu.getHL())
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	return 0
}

// OR A
func (cpu *Cpu)opCodeB7() byte {
	cpu.reg.A = cpu.reg.A | cpu.reg.A
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	return 0
}

// JP a16
func (cpu *Cpu)opCodeC2() byte {
	if !cpu.getFlagZ() {
		cpu.reg.PC = cpu.gb.Mem.ReadUint16(cpu.reg.PC)
		return 0
	}
	cpu.reg.PC += 2
	return 1
}

// JP a16
func (cpu *Cpu)opCodeC3() byte {
	cpu.reg.PC = cpu.gb.Mem.ReadUint16(cpu.reg.PC)
	return 0
}

// LDH (a8),A
func (cpu *Cpu)opCodeE0() byte {
	addr := cpu.gb.Mem.Read(cpu.reg.PC)
	cpu.reg.PC++
	cpu.gb.Mem.Write(0xFF00 + uint16(addr), cpu.reg.A)
	return 0
}

// AND d8
func (cpu *Cpu)opCodeE6() byte {
	arg := cpu.gb.Mem.Read(cpu.reg.PC)
	cpu.reg.PC++
	cpu.reg.A = cpu.reg.A & arg
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	cpu.setFlagH()
	return 0
}

// XOR d8
func (cpu *Cpu)opCodeEE() byte {
	arg := cpu.gb.Mem.Read(cpu.reg.PC)
	cpu.reg.PC++
	cpu.reg.A = cpu.reg.A ^ arg
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	return 0
}

// LDH A,(a8)
func (cpu *Cpu)opCodeF0() byte {
	cpu.reg.A = cpu.gb.Mem.Read(0xFF00 + uint16(cpu.reg.PC))
	cpu.reg.PC++
	return 0
}

// DI
func (cpu *Cpu)opCodeF3() byte {
	cpu.setFlagIME(false)
	return 0
}

// OR d8
func (cpu *Cpu)opCodeF6() byte {
	arg := cpu.gb.Mem.Read(cpu.reg.PC)
	cpu.reg.PC++
	cpu.reg.A = cpu.reg.A | arg
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	return 0
}

// EI
func (cpu *Cpu)opCodeFB() byte {
	cpu.setFlagIME(true)
	return 0
}