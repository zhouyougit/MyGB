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
}

//NOP
func (cpu *Cpu)opCode00() byte {
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