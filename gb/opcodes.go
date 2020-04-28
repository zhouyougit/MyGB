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
	0x01: {
		Func: (*Cpu).opCode01,
		Cycles: [2]int{12, 0},
		Length: 3,
		Mnemonic: "LD BC,d16",
	},
	0x02: {
		Func: (*Cpu).opCode02,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "LD (BC),A",
	},
	0x06: {
		Func: (*Cpu).opCode06,
		Cycles: [2]int{8, 0},
		Length: 2,
		Mnemonic: "LD B,d8",
	},
	0x08: {
		Func: (*Cpu).opCode08,
		Cycles: [2]int{20, 0},
		Length: 3,
		Mnemonic: "LD (a16),SP",
	},
	0x0A: {
		Func: (*Cpu).opCode0A,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "LD A,(BC)",
	},
	0x0E: {
		Func: (*Cpu).opCode0E,
		Cycles: [2]int{8, 0},
		Length: 2,
		Mnemonic: "LD C,d8",
	},
	0x11: {
		Func: (*Cpu).opCode11,
		Cycles: [2]int{12, 0},
		Length: 3,
		Mnemonic: "LD DE,d16",
	},
	0x12: {
		Func: (*Cpu).opCode12,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "LD (DE),A",
	},
	0x16: {
		Func: (*Cpu).opCode16,
		Cycles: [2]int{8, 0},
		Length: 2,
		Mnemonic: "LD D,d8",
	},
	0x1A: {
		Func: (*Cpu).opCode1A,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "LD A,(DE)",
	},
	0x1E: {
		Func: (*Cpu).opCode1E,
		Cycles: [2]int{8, 0},
		Length: 2,
		Mnemonic: "LD E,d8",
	},
	0x21: {
		Func: (*Cpu).opCode21,
		Cycles: [2]int{12, 0},
		Length: 3,
		Mnemonic: "LD HL,d16",
	},
	0x22: {
		Func: (*Cpu).opCode22,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "LD (HL+),A",
	},
	0x26: {
		Func: (*Cpu).opCode26,
		Cycles: [2]int{8, 0},
		Length: 2,
		Mnemonic: "LD H,d8",
	},
	0x2A: {
		Func: (*Cpu).opCode2A,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "LD A,(HL+)",
	},
	0x2E: {
		Func: (*Cpu).opCode2E,
		Cycles: [2]int{8, 0},
		Length: 2,
		Mnemonic: "LD L,d8",
	},
	0x31: {
		Func: (*Cpu).opCode31,
		Cycles: [2]int{12, 0},
		Length: 3,
		Mnemonic: "LD SP,d16",
	},
	0x32: {
		Func: (*Cpu).opCode32,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "LD (HL-),A",
	},
	0x36: {
		Func: (*Cpu).opCode36,
		Cycles: [2]int{8, 0},
		Length: 2,
		Mnemonic: "LD (HL),d8",
	},
	0x3A: {
		Func: (*Cpu).opCode3A,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "LD A,(HL-)",
	},
	0x3E: {
		Func: (*Cpu).opCode3E,
		Cycles: [2]int{8, 0},
		Length: 2,
		Mnemonic: "LD A,d8",
	},
	0x40: {
		Func: (*Cpu).opCode40,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD B,B",
	},
	0x41: {
		Func: (*Cpu).opCode41,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD B,C",
	},
	0x42: {
		Func: (*Cpu).opCode42,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD B,D",
	},
	0x43: {
		Func: (*Cpu).opCode43,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD B,E",
	},
	0x44: {
		Func: (*Cpu).opCode44,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD B,H",
	},
	0x45: {
		Func: (*Cpu).opCode45,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD B,L",
	},
	0x46: {
		Func: (*Cpu).opCode46,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "LD B,(HL)",
	},
	0x47: {
		Func: (*Cpu).opCode47,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD B,A",
	},
	0x48: {
		Func: (*Cpu).opCode48,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD C,B",
	},
	0x49: {
		Func: (*Cpu).opCode49,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD C,C",
	},
	0x4A: {
		Func: (*Cpu).opCode4A,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD C,D",
	},
	0x4B: {
		Func: (*Cpu).opCode4B,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD C,E",
	},
	0x4C: {
		Func: (*Cpu).opCode4C,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD C,H",
	},
	0x4D: {
		Func: (*Cpu).opCode4D,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD C,L",
	},
	0x4E: {
		Func: (*Cpu).opCode4E,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "LD C,(HL)",
	},
	0x4F: {
		Func: (*Cpu).opCode4F,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD B,A",
	},
	0x50: {
		Func: (*Cpu).opCode50,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD D,B",
	},
	0x51: {
		Func: (*Cpu).opCode51,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD D,C",
	},
	0x52: {
		Func: (*Cpu).opCode52,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD D,D",
	},
	0x53: {
		Func: (*Cpu).opCode53,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD D,E",
	},
	0x54: {
		Func: (*Cpu).opCode54,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD D,H",
	},
	0x55: {
		Func: (*Cpu).opCode55,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD D,L",
	},
	0x56: {
		Func: (*Cpu).opCode56,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "LD D,(HL)",
	},
	0x57: {
		Func: (*Cpu).opCode57,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD D,A",
	},
	0x58: {
		Func: (*Cpu).opCode58,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD E,B",
	},
	0x59: {
		Func: (*Cpu).opCode59,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD E,C",
	},
	0x5A: {
		Func: (*Cpu).opCode5A,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD E,D",
	},
	0x5B: {
		Func: (*Cpu).opCode5B,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD E,E",
	},
	0x5C: {
		Func: (*Cpu).opCode5C,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD E,H",
	},
	0x5D: {
		Func: (*Cpu).opCode5D,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD E,L",
	},
	0x5E: {
		Func: (*Cpu).opCode5E,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "LD E,(HL)",
	},
	0x5F: {
		Func: (*Cpu).opCode5F,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD E,A",
	},
	0x60: {
		Func: (*Cpu).opCode60,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD H,B",
	},
	0x61: {
		Func: (*Cpu).opCode61,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD H,C",
	},
	0x62: {
		Func: (*Cpu).opCode62,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD H,D",
	},
	0x63: {
		Func: (*Cpu).opCode63,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD H,E",
	},
	0x64: {
		Func: (*Cpu).opCode64,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD H,H",
	},
	0x65: {
		Func: (*Cpu).opCode65,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD H,L",
	},
	0x66: {
		Func: (*Cpu).opCode66,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "LD H,(HL)",
	},
	0x67: {
		Func: (*Cpu).opCode67,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD H,A",
	},
	0x68: {
		Func: (*Cpu).opCode68,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD L,B",
	},
	0x69: {
		Func: (*Cpu).opCode69,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD L,C",
	},
	0x6A: {
		Func: (*Cpu).opCode6A,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD L,D",
	},
	0x6B: {
		Func: (*Cpu).opCode6B,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD L,E",
	},
	0x6C: {
		Func: (*Cpu).opCode6C,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD L,H",
	},
	0x6D: {
		Func: (*Cpu).opCode6D,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD L,L",
	},
	0x6E: {
		Func: (*Cpu).opCode6E,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "LD L,(HL)",
	},
	0x6F: {
		Func: (*Cpu).opCode6F,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD L,A",
	},
	0x70: {
		Func: (*Cpu).opCode70,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "LD (HL),B",
	},
	0x71: {
		Func: (*Cpu).opCode71,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "LD (HL),C",
	},
	0x72: {
		Func: (*Cpu).opCode72,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "LD (HL),D",
	},
	0x73: {
		Func: (*Cpu).opCode73,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "LD (HL),E",
	},
	0x74: {
		Func: (*Cpu).opCode74,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "LD (HL),H",
	},
	0x75: {
		Func: (*Cpu).opCode75,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "LD (HL),L",
	},
	0x77: {
		Func: (*Cpu).opCode77,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "LD (HL),A",
	},
	0x78: {
		Func: (*Cpu).opCode78,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD A,B",
	},
	0x79: {
		Func: (*Cpu).opCode79,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD A,C",
	},
	0x7A: {
		Func: (*Cpu).opCode7A,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD A,D",
	},
	0x7B: {
		Func: (*Cpu).opCode7B,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD A,E",
	},
	0x7C: {
		Func: (*Cpu).opCode7C,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD A,H",
	},
	0x7D: {
		Func: (*Cpu).opCode7D,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD A,L",
	},
	0x7E: {
		Func: (*Cpu).opCode7E,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "LD A,(HL)",
	},
	0x7F: {
		Func: (*Cpu).opCode7F,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "LD A,A",
	},
	0x80: {
		Func: (*Cpu).opCode80,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "ADD A,B",
	},
	0x81: {
		Func: (*Cpu).opCode81,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "ADD A,C",
	},
	0x82: {
		Func: (*Cpu).opCode82,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "ADD A,D",
	},
	0x83: {
		Func: (*Cpu).opCode83,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "ADD A,E",
	},
	0x84: {
		Func: (*Cpu).opCode84,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "ADD A,H",
	},
	0x85: {
		Func: (*Cpu).opCode85,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "ADD A,L",
	},
	0x86: {
		Func: (*Cpu).opCode86,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "ADD A,(HL)",
	},
	0x87: {
		Func: (*Cpu).opCode87,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "ADD A,A",
	},
	0x88: {
		Func: (*Cpu).opCode88,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "ADC A,B",
	},
	0x89: {
		Func: (*Cpu).opCode89,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "ADC A,C",
	},
	0x8A: {
		Func: (*Cpu).opCode8A,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "ADC A,D",
	},
	0x8B: {
		Func: (*Cpu).opCode8B,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "ADC A,E",
	},
	0x8C: {
		Func: (*Cpu).opCode8C,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "ADC A,H",
	},
	0x8D: {
		Func: (*Cpu).opCode8D,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "ADC A,L",
	},
	0x8E: {
		Func: (*Cpu).opCode8E,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "ADC A,(HL)",
	},
	0x8F: {
		Func: (*Cpu).opCode8F,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "ADC A,A",
	},
	0x90: {
		Func: (*Cpu).opCode90,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "SUB B",
	},
	0x91: {
		Func: (*Cpu).opCode91,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "SUB C",
	},
	0x92: {
		Func: (*Cpu).opCode92,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "SUB D",
	},
	0x93: {
		Func: (*Cpu).opCode93,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "SUB E",
	},
	0x94: {
		Func: (*Cpu).opCode94,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "SUB H",
	},
	0x95: {
		Func: (*Cpu).opCode95,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "SUB L",
	},
	0x96: {
		Func: (*Cpu).opCode96,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "SUB (HL)",
	},
	0x97: {
		Func: (*Cpu).opCode97,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "SUB A",
	},
	0x98: {
		Func: (*Cpu).opCode98,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "SBC A,B",
	},
	0x99: {
		Func: (*Cpu).opCode99,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "SBC A,C",
	},
	0x9A: {
		Func: (*Cpu).opCode9A,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "SBC A,D",
	},
	0x9B: {
		Func: (*Cpu).opCode9B,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "SBC A,E",
	},
	0x9C: {
		Func: (*Cpu).opCode9C,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "SBC A,H",
	},
	0x9D: {
		Func: (*Cpu).opCode9D,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "SBC A,L",
	},
	0x9E: {
		Func: (*Cpu).opCode9E,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "SBC A,(HL)",
	},
	0x9F: {
		Func: (*Cpu).opCode9F,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "SBC A,A",
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
	0xC1: {
		Func: (*Cpu).opCodeC1,
		Cycles: [2]int{12, 0},
		Length: 1,
		Mnemonic: "POP BC",
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
	0xC5: {
		Func: (*Cpu).opCodeC5,
		Cycles: [2]int{16, 0},
		Length: 1,
		Mnemonic: "PUSH BC",
	},
	0xC6: {
		Func: (*Cpu).opCodeC6,
		Cycles: [2]int{8, 0},
		Length: 2,
		Mnemonic: "ADD A,d8",
	},
	0xCE: {
		Func: (*Cpu).opCodeCE,
		Cycles: [2]int{8, 0},
		Length: 2,
		Mnemonic: "ADC A,d8",
	},
	0xD1: {
		Func: (*Cpu).opCodeD1,
		Cycles: [2]int{12, 0},
		Length: 1,
		Mnemonic: "POP DE",
	},
	0xD5: {
		Func: (*Cpu).opCodeD5,
		Cycles: [2]int{16, 0},
		Length: 1,
		Mnemonic: "PUSH DE",
	},
	0xD6: {
		Func: (*Cpu).opCodeD6,
		Cycles: [2]int{8, 0},
		Length: 2,
		Mnemonic: "SUB d8",
	},
	0xDE: {
		Func: (*Cpu).opCodeDE,
		Cycles: [2]int{8, 0},
		Length: 2,
		Mnemonic: "SBC A,d8",
	},
	0xE0: {
		Func: (*Cpu).opCodeE0,
		Cycles: [2]int{12, 0},
		Length: 2,
		Mnemonic: "LDH (a8),A",
	},
	0xE1: {
		Func: (*Cpu).opCodeE1,
		Cycles: [2]int{12, 0},
		Length: 1,
		Mnemonic: "POP HL",
	},
	0xE2: {
		Func: (*Cpu).opCodeE2,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "LD (C),A",
	},
	0xE5: {
		Func: (*Cpu).opCodeE5,
		Cycles: [2]int{16, 0},
		Length: 1,
		Mnemonic: "PUSH HL",
	},
	0xE6: {
		Func: (*Cpu).opCodeE6,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "AND d8",
	},
	0xEA: {
		Func: (*Cpu).opCodeEA,
		Cycles: [2]int{16, 0},
		Length: 3,
		Mnemonic: "LD (a16),A",
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
	0xF1: {
		Func: (*Cpu).opCodeF1,
		Cycles: [2]int{12, 0},
		Length: 1,
		Mnemonic: "POP AF",
	},
	0xF2: {
		Func: (*Cpu).opCodeF2,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "LD A,(C)",
	},
	0xF3: {
		Func: (*Cpu).opCodeF3,
		Cycles: [2]int{4, 0},
		Length: 1,
		Mnemonic: "DI",
	},
	0xF5: {
		Func: (*Cpu).opCodeF5,
		Cycles: [2]int{16, 0},
		Length: 1,
		Mnemonic: "PUSH AF",
	},
	0xF6: {
		Func: (*Cpu).opCodeF6,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "OR d8",
	},
	0xF8: {
		Func: (*Cpu).opCodeF8,
		Cycles: [2]int{12, 0},
		Length: 2,
		Mnemonic: "LD HL,SP+r8",
	},
	0xF9: {
		Func: (*Cpu).opCodeF9,
		Cycles: [2]int{8, 0},
		Length: 1,
		Mnemonic: "LD SP,HL",
	},
	0xFA: {
		Func: (*Cpu).opCodeFA,
		Cycles: [2]int{16, 0},
		Length: 3,
		Mnemonic: "LD A,(a16)",
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

// LD BC,d16
func (cpu *Cpu)opCode01() byte {
	arg := cpu.gb.Mem.ReadUint16(cpu.reg.PC)
	cpu.reg.PC += 2
	cpu.reg.setBC(arg)
	return 0
}

// LD (BC),A
func (cpu *Cpu)opCode02() byte {
	cpu.gb.Mem.Write(cpu.reg.getBC(), cpu.reg.A)
	return 0
}

// LD B,d8
func (cpu *Cpu)opCode06() byte {
	cpu.reg.B = cpu.gb.Mem.Read(cpu.reg.PC)
	cpu.reg.PC++
	return 0
}

// LD (a16),SP
func (cpu *Cpu)opCode08() byte {
	addr := cpu.gb.Mem.ReadUint16(cpu.reg.PC)
	cpu.reg.PC += 2
	cpu.gb.Mem.WriteUint16(addr, cpu.reg.SP)
	return 0
}

// LD A,(BC)
func (cpu *Cpu)opCode0A() byte {
	cpu.reg.A = cpu.gb.Mem.Read(cpu.reg.getBC())
	return 0
}

// LD C,d8
func (cpu *Cpu)opCode0E() byte {
	cpu.reg.C = cpu.gb.Mem.Read(cpu.reg.PC)
	cpu.reg.PC++
	return 0
}

// LD DE,d16
func (cpu *Cpu)opCode11() byte {
	arg := cpu.gb.Mem.ReadUint16(cpu.reg.PC)
	cpu.reg.PC += 2
	cpu.reg.setDE(arg)
	return 0
}

// LD (DE),A
func (cpu *Cpu)opCode12() byte {
	cpu.gb.Mem.Write(cpu.reg.getDE(), cpu.reg.A)
	return 0
}

// LD D,d8
func (cpu *Cpu)opCode16() byte {
	cpu.reg.D = cpu.gb.Mem.Read(cpu.reg.PC)
	cpu.reg.PC++
	return 0
}

// LD A,(DE)
func (cpu *Cpu)opCode1A() byte {
	cpu.reg.A = cpu.gb.Mem.Read(cpu.reg.getDE())
	return 0
}

// LD E,d8
func (cpu *Cpu)opCode1E() byte {
	cpu.reg.E = cpu.gb.Mem.Read(cpu.reg.PC)
	cpu.reg.PC++
	return 0
}

// LD HL,d16
func (cpu *Cpu)opCode21() byte {
	arg := cpu.gb.Mem.ReadUint16(cpu.reg.PC)
	cpu.reg.PC += 2
	cpu.reg.setHL(arg)
	return 0
}

// LD (HL+),A
func (cpu *Cpu)opCode22() byte {
	cpu.gb.Mem.Write(cpu.reg.getHL(), cpu.reg.A)
	cpu.reg.setHL(cpu.reg.getHL() + 1)
	return 0
}

// LD H,d8
func (cpu *Cpu)opCode26() byte {
	cpu.reg.H = cpu.gb.Mem.Read(cpu.reg.PC)
	cpu.reg.PC++
	return 0
}

// LD A,(HL+)
func (cpu *Cpu)opCode2A() byte {
	cpu.reg.A = cpu.gb.Mem.Read(cpu.reg.getHL())
	cpu.reg.setHL(cpu.reg.getHL() + 1)
	return 0
}

// LD L,d8
func (cpu *Cpu)opCode2E() byte {
	cpu.reg.L = cpu.gb.Mem.Read(cpu.reg.PC)
	cpu.reg.PC++
	return 0
}

// LD SP,d16
func (cpu *Cpu)opCode31() byte {
	arg := cpu.gb.Mem.ReadUint16(cpu.reg.PC)
	cpu.reg.PC += 2
	cpu.reg.SP = arg
	return 0
}

// LD (HL-),A
func (cpu *Cpu)opCode32() byte {
	cpu.gb.Mem.Write(cpu.reg.getHL(), cpu.reg.A)
	cpu.reg.setHL(cpu.reg.getHL() - 1)
	return 0
}

// LD (HL),d8
func (cpu *Cpu)opCode36() byte {
	arg := cpu.gb.Mem.Read(cpu.reg.PC)
	cpu.reg.PC++
	cpu.gb.Mem.Write(cpu.reg.getHL(), arg)
	return 0
}

// LD A,(HL-)
func (cpu *Cpu)opCode3A() byte {
	cpu.reg.A = cpu.gb.Mem.Read(cpu.reg.getHL())
	cpu.reg.setHL(cpu.reg.getHL() - 1)
	return 0
}

// LD A,d8
func (cpu *Cpu)opCode3E() byte {
	cpu.reg.A = cpu.gb.Mem.Read(cpu.reg.PC)
	cpu.reg.PC++
	return 0
}

// LD B,B
func (cpu *Cpu)opCode40() byte {
	//cpu.reg.B = cpu.reg.B
	return 0
}

// LD B,C
func (cpu *Cpu)opCode41() byte {
	cpu.reg.B = cpu.reg.C
	return 0
}

// LD B,D
func (cpu *Cpu)opCode42() byte {
	cpu.reg.B = cpu.reg.D
	return 0
}

// LD B,E
func (cpu *Cpu)opCode43() byte {
	cpu.reg.B = cpu.reg.E
	return 0
}

// LD B,H
func (cpu *Cpu)opCode44() byte {
	cpu.reg.B = cpu.reg.H
	return 0
}

// LD B,L
func (cpu *Cpu)opCode45() byte {
	cpu.reg.B = cpu.reg.L
	return 0
}

// LD B,(HL)
func (cpu *Cpu)opCode46() byte {
	cpu.reg.B = cpu.gb.Mem.Read(cpu.reg.getHL())
	return 0
}

// LD B,A
func (cpu *Cpu)opCode47() byte {
	cpu.reg.B = cpu.reg.A
	return 0
}

// LD C,B
func (cpu *Cpu)opCode48() byte {
	cpu.reg.C = cpu.reg.B
	return 0
}

// LD C,C
func (cpu *Cpu)opCode49() byte {
	//cpu.reg.C = cpu.reg.C
	return 0
}

// LD C,D
func (cpu *Cpu)opCode4A() byte {
	cpu.reg.C = cpu.reg.D
	return 0
}

// LD C,E
func (cpu *Cpu)opCode4B() byte {
	cpu.reg.C = cpu.reg.E
	return 0
}

// LD C,H
func (cpu *Cpu)opCode4C() byte {
	cpu.reg.C = cpu.reg.H
	return 0
}

// LD C,L
func (cpu *Cpu)opCode4D() byte {
	cpu.reg.C = cpu.reg.L
	return 0
}

// LD C,(HL)
func (cpu *Cpu)opCode4E() byte {
	cpu.reg.C = cpu.gb.Mem.Read(cpu.reg.getHL())
	return 0
}

// LD C,A
func (cpu *Cpu)opCode4F() byte {
	cpu.reg.C = cpu.reg.A
	return 0
}

// LD D,B
func (cpu *Cpu)opCode50() byte {
	cpu.reg.D = cpu.reg.B
	return 0
}

// LD D,C
func (cpu *Cpu)opCode51() byte {
	cpu.reg.D = cpu.reg.C
	return 0
}

// LD D,D
func (cpu *Cpu)opCode52() byte {
	//cpu.reg.D = cpu.reg.D
	return 0
}

// LD D,E
func (cpu *Cpu)opCode53() byte {
	cpu.reg.D = cpu.reg.E
	return 0
}

// LD D,H
func (cpu *Cpu)opCode54() byte {
	cpu.reg.D = cpu.reg.H
	return 0
}

// LD D,L
func (cpu *Cpu)opCode55() byte {
	cpu.reg.D = cpu.reg.L
	return 0
}

// LD D,(HL)
func (cpu *Cpu)opCode56() byte {
	cpu.reg.D = cpu.gb.Mem.Read(cpu.reg.getHL())
	return 0
}

// LD D,A
func (cpu *Cpu)opCode57() byte {
	cpu.reg.D = cpu.reg.A
	return 0
}

// LD E,B
func (cpu *Cpu)opCode58() byte {
	cpu.reg.E = cpu.reg.B
	return 0
}

// LD E,C
func (cpu *Cpu)opCode59() byte {
	cpu.reg.E = cpu.reg.C
	return 0
}

// LD E,D
func (cpu *Cpu)opCode5A() byte {
	cpu.reg.E = cpu.reg.D
	return 0
}

// LD E,E
func (cpu *Cpu)opCode5B() byte {
	//cpu.reg.E = cpu.reg.E
	return 0
}

// LD E,H
func (cpu *Cpu)opCode5C() byte {
	cpu.reg.E = cpu.reg.H
	return 0
}

// LD E,L
func (cpu *Cpu)opCode5D() byte {
	cpu.reg.E = cpu.reg.L
	return 0
}

// LD E,(HL)
func (cpu *Cpu)opCode5E() byte {
	cpu.reg.E = cpu.gb.Mem.Read(cpu.reg.getHL())
	return 0
}

// LD E,A
func (cpu *Cpu)opCode5F() byte {
	cpu.reg.E = cpu.reg.A
	return 0
}

// LD H,B
func (cpu *Cpu)opCode60() byte {
	cpu.reg.H = cpu.reg.B
	return 0
}

// LD H,C
func (cpu *Cpu)opCode61() byte {
	cpu.reg.H = cpu.reg.C
	return 0
}

// LD H,D
func (cpu *Cpu)opCode62() byte {
	cpu.reg.H = cpu.reg.D
	return 0
}

// LD H,E
func (cpu *Cpu)opCode63() byte {
	cpu.reg.H = cpu.reg.E
	return 0
}

// LD H,H
func (cpu *Cpu)opCode64() byte {
	//cpu.reg.H = cpu.reg.H
	return 0
}

// LD H,L
func (cpu *Cpu)opCode65() byte {
	cpu.reg.H = cpu.reg.L
	return 0
}

// LD H,(HL)
func (cpu *Cpu)opCode66() byte {
	cpu.reg.H = cpu.gb.Mem.Read(cpu.reg.getHL())
	return 0
}

// LD H,A
func (cpu *Cpu)opCode67() byte {
	cpu.reg.H = cpu.reg.A
	return 0
}

// LD L,B
func (cpu *Cpu)opCode68() byte {
	cpu.reg.L = cpu.reg.B
	return 0
}

// LD L,C
func (cpu *Cpu)opCode69() byte {
	cpu.reg.L = cpu.reg.C
	return 0
}

// LD L,D
func (cpu *Cpu)opCode6A() byte {
	cpu.reg.L = cpu.reg.D
	return 0
}

// LD L,E
func (cpu *Cpu)opCode6B() byte {
	cpu.reg.L = cpu.reg.E
	return 0
}

// LD L,H
func (cpu *Cpu)opCode6C() byte {
	cpu.reg.L = cpu.reg.H
	return 0
}

// LD L,L
func (cpu *Cpu)opCode6D() byte {
	//cpu.reg.L = cpu.reg.L
	return 0
}

// LD L,(HL)
func (cpu *Cpu)opCode6E() byte {
	cpu.reg.L = cpu.gb.Mem.Read(cpu.reg.getHL())
	return 0
}

// LD L,A
func (cpu *Cpu)opCode6F() byte {
	cpu.reg.L = cpu.reg.A
	return 0
}

// LD (HL),B
func (cpu *Cpu)opCode70() byte {
	cpu.gb.Mem.Write(cpu.reg.getHL(), cpu.reg.B)
	return 0
}

// LD (HL),C
func (cpu *Cpu)opCode71() byte {
	cpu.gb.Mem.Write(cpu.reg.getHL(), cpu.reg.C)
	return 0
}

// LD (HL),D
func (cpu *Cpu)opCode72() byte {
	cpu.gb.Mem.Write(cpu.reg.getHL(), cpu.reg.D)
	return 0
}

// LD (HL),E
func (cpu *Cpu)opCode73() byte {
	cpu.gb.Mem.Write(cpu.reg.getHL(), cpu.reg.E)
	return 0
}

// LD (HL),H
func (cpu *Cpu)opCode74() byte {
	cpu.gb.Mem.Write(cpu.reg.getHL(), cpu.reg.H)
	return 0
}

// LD (HL),L
func (cpu *Cpu)opCode75() byte {
	cpu.gb.Mem.Write(cpu.reg.getHL(), cpu.reg.L)
	return 0
}

// LD (HL),A
func (cpu *Cpu)opCode77() byte {
	cpu.gb.Mem.Write(cpu.reg.getHL(), cpu.reg.A)
	return 0
}

// LD A,B
func (cpu *Cpu)opCode78() byte {
	cpu.reg.A = cpu.reg.B
	return 0
}

// LD A,C
func (cpu *Cpu)opCode79() byte {
	cpu.reg.A = cpu.reg.C
	return 0
}

// LD A,D
func (cpu *Cpu)opCode7A() byte {
	cpu.reg.A = cpu.reg.D
	return 0
}

// LD A,E
func (cpu *Cpu)opCode7B() byte {
	cpu.reg.A = cpu.reg.E
	return 0
}

// LD A,H
func (cpu *Cpu)opCode7C() byte {
	cpu.reg.A = cpu.reg.H
	return 0
}

// LD A,L
func (cpu *Cpu)opCode7D() byte {
	cpu.reg.A = cpu.reg.L
	return 0
}

// LD A,(HL)
func (cpu *Cpu)opCode7E() byte {
	cpu.reg.A = cpu.gb.Mem.Read(cpu.reg.getHL())
	return 0
}

// LD A,A
func (cpu *Cpu)opCode7F() byte {
	//cpu.reg.A = cpu.reg.A
	return 0
}

// ADD A,B
func (cpu *Cpu)opCode80() byte {
	v1 := uint16(cpu.reg.A)
	v2 := uint16(cpu.reg.B)
	res := v1 + v2
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if byte(res) == 0 {
		cpu.setFlagZ()
	}
	carry := v1 ^ v2 ^ res
	if (carry & 0x10) > 0 {
		cpu.setFlagH()
	}
	if (carry & 0x100) > 0 {
		cpu.setFlagC()
	}
	return 0
}

// ADD A,C
func (cpu *Cpu)opCode81() byte {
	v1 := uint16(cpu.reg.A)
	v2 := uint16(cpu.reg.C)
	res := v1 + v2
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if byte(res) == 0 {
		cpu.setFlagZ()
	}
	carry := v1 ^ v2 ^ res
	if (carry & 0x10) > 0 {
		cpu.setFlagH()
	}
	if (carry & 0x100) > 0 {
		cpu.setFlagC()
	}
	return 0
}

// ADD A,D
func (cpu *Cpu)opCode82() byte {
	v1 := uint16(cpu.reg.A)
	v2 := uint16(cpu.reg.D)
	res := v1 + v2
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if byte(res) == 0 {
		cpu.setFlagZ()
	}
	carry := v1 ^ v2 ^ res
	if (carry & 0x10) > 0 {
		cpu.setFlagH()
	}
	if (carry & 0x100) > 0 {
		cpu.setFlagC()
	}
	return 0
}

// ADD A,E
func (cpu *Cpu)opCode83() byte {
	v1 := uint16(cpu.reg.A)
	v2 := uint16(cpu.reg.E)
	res := v1 + v2
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if byte(res) == 0 {
		cpu.setFlagZ()
	}
	carry := v1 ^ v2 ^ res
	if (carry & 0x10) > 0 {
		cpu.setFlagH()
	}
	if (carry & 0x100) > 0 {
		cpu.setFlagC()
	}
	return 0
}

// ADD A,H
func (cpu *Cpu)opCode84() byte {
	v1 := uint16(cpu.reg.A)
	v2 := uint16(cpu.reg.H)
	res := v1 + v2
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if byte(res) == 0 {
		cpu.setFlagZ()
	}
	carry := v1 ^ v2 ^ res
	if (carry & 0x10) > 0 {
		cpu.setFlagH()
	}
	if (carry & 0x100) > 0 {
		cpu.setFlagC()
	}
	return 0
}

// ADD A,L
func (cpu *Cpu)opCode85() byte {
	v1 := uint16(cpu.reg.A)
	v2 := uint16(cpu.reg.L)
	res := v1 + v2
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if byte(res) == 0 {
		cpu.setFlagZ()
	}
	carry := v1 ^ v2 ^ res
	if (carry & 0x10) > 0 {
		cpu.setFlagH()
	}
	if (carry & 0x100) > 0 {
		cpu.setFlagC()
	}
	return 0
}

// ADD A,(HL)
func (cpu *Cpu)opCode86() byte {
	v1 := uint16(cpu.reg.A)
	v2 := uint16(cpu.gb.Mem.Read(cpu.reg.getHL()))
	res := v1 + v2
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if byte(res) == 0 {
		cpu.setFlagZ()
	}
	carry := v1 ^ v2 ^ res
	if (carry & 0x10) > 0 {
		cpu.setFlagH()
	}
	if (carry & 0x100) > 0 {
		cpu.setFlagC()
	}
	return 0
}

// ADD A,A
func (cpu *Cpu)opCode87() byte {
	v1 := uint16(cpu.reg.A)
	res := v1 + v1
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if byte(res) == 0 {
		cpu.setFlagZ()
	}
	carry := v1 ^ v1 ^ res
	if (carry & 0x10) > 0 {
		cpu.setFlagH()
	}
	if (carry & 0x100) > 0 {
		cpu.setFlagC()
	}
	return 0
}

// ADC A,B
func (cpu *Cpu)opCode88() byte {
	v1 := uint16(cpu.reg.A)
	v2 := uint16(cpu.reg.B)
	carry := uint16(0)
	if cpu.getFlagC() {
		carry = uint16(1)
	}
	res := v1 + v2 + carry
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if byte(res) == 0 {
		cpu.setFlagZ()
	}
	if (v1 & 0x000F) + (v2 & 0x000F) + carry > 0x000F {
		cpu.setFlagH()
	}
	if res > 0x00FF {
		cpu.setFlagC()
	}
	return 0
}

// ADC A,C
func (cpu *Cpu)opCode89() byte {
	v1 := uint16(cpu.reg.A)
	v2 := uint16(cpu.reg.C)
	carry := uint16(0)
	if cpu.getFlagC() {
		carry = uint16(1)
	}
	res := v1 + v2 + carry
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if byte(res) == 0 {
		cpu.setFlagZ()
	}
	if (v1 & 0x000F) + (v2 & 0x000F) + carry > 0x000F {
		cpu.setFlagH()
	}
	if res > 0x00FF {
		cpu.setFlagC()
	}
	return 0
}

// ADC A,D
func (cpu *Cpu)opCode8A() byte {
	v1 := uint16(cpu.reg.A)
	v2 := uint16(cpu.reg.D)
	carry := uint16(0)
	if cpu.getFlagC() {
		carry = uint16(1)
	}
	res := v1 + v2 + carry
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if byte(res) == 0 {
		cpu.setFlagZ()
	}
	if (v1 & 0x000F) + (v2 & 0x000F) + carry > 0x000F {
		cpu.setFlagH()
	}
	if res > 0x00FF {
		cpu.setFlagC()
	}
	return 0
}

// ADC A,E
func (cpu *Cpu)opCode8B() byte {
	v1 := uint16(cpu.reg.A)
	v2 := uint16(cpu.reg.E)
	carry := uint16(0)
	if cpu.getFlagC() {
		carry = uint16(1)
	}
	res := v1 + v2 + carry
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if byte(res) == 0 {
		cpu.setFlagZ()
	}
	if (v1 & 0x000F) + (v2 & 0x000F) + carry > 0x000F {
		cpu.setFlagH()
	}
	if res > 0x00FF {
		cpu.setFlagC()
	}
	return 0
}

// ADC A,H
func (cpu *Cpu)opCode8C() byte {
	v1 := uint16(cpu.reg.A)
	v2 := uint16(cpu.reg.H)
	carry := uint16(0)
	if cpu.getFlagC() {
		carry = uint16(1)
	}
	res := v1 + v2 + carry
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if byte(res) == 0 {
		cpu.setFlagZ()
	}
	if (v1 & 0x000F) + (v2 & 0x000F) + carry > 0x000F {
		cpu.setFlagH()
	}
	if res > 0x00FF {
		cpu.setFlagC()
	}
	return 0
}

// ADC A,L
func (cpu *Cpu)opCode8D() byte {
	v1 := uint16(cpu.reg.A)
	v2 := uint16(cpu.reg.L)
	carry := uint16(0)
	if cpu.getFlagC() {
		carry = uint16(1)
	}
	res := v1 + v2 + carry
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if byte(res) == 0 {
		cpu.setFlagZ()
	}
	if (v1 & 0x000F) + (v2 & 0x000F) + carry > 0x000F {
		cpu.setFlagH()
	}
	if res > 0x00FF {
		cpu.setFlagC()
	}
	return 0
}

// ADC A,(HL)
func (cpu *Cpu)opCode8E() byte {
	v1 := uint16(cpu.reg.A)
	v2 := uint16(cpu.gb.Mem.Read(cpu.reg.getHL()))
	carry := uint16(0)
	if cpu.getFlagC() {
		carry = uint16(1)
	}
	res := v1 + v2 + carry
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if byte(res) == 0 {
		cpu.setFlagZ()
	}
	if (v1 & 0x000F) + (v2 & 0x000F) + carry > 0x000F {
		cpu.setFlagH()
	}
	if res > 0x00FF {
		cpu.setFlagC()
	}
	return 0
}

// ADC A,A
func (cpu *Cpu)opCode8F() byte {
	v1 := uint16(cpu.reg.A)
	carry := uint16(0)
	if cpu.getFlagC() {
		carry = uint16(1)
	}
	res := v1 + v1 + carry
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if byte(res) == 0 {
		cpu.setFlagZ()
	}
	if (v1 & 0x000F) + (v1 & 0x000F) + carry > 0x000F {
		cpu.setFlagH()
	}
	if res > 0x00FF {
		cpu.setFlagC()
	}
	return 0
}

// SUB B
func (cpu *Cpu)opCode90() byte {
	v1 := int16(cpu.reg.A)
	v2 := int16(cpu.reg.B)
	res := v1 - v2
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	cpu.setFlagN()
	if (v1 & 0xF) - (v2 & 0xF) < 0 {
		cpu.setFlagH()
	}
	if res < 0 {
		cpu.setFlagC()
	}
	return 0
}

// SUB C
func (cpu *Cpu)opCode91() byte {
	v1 := int16(cpu.reg.A)
	v2 := int16(cpu.reg.C)
	res := v1 - v2
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	cpu.setFlagN()
	if (v1 & 0xF) - (v2 & 0xF) < 0 {
		cpu.setFlagH()
	}
	if res < 0 {
		cpu.setFlagC()
	}
	return 0
}

// SUB D
func (cpu *Cpu)opCode92() byte {
	v1 := int16(cpu.reg.A)
	v2 := int16(cpu.reg.D)
	res := v1 - v2
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	cpu.setFlagN()
	if (v1 & 0xF) - (v2 & 0xF) < 0 {
		cpu.setFlagH()
	}
	if res < 0 {
		cpu.setFlagC()
	}
	return 0
}

// SUB E
func (cpu *Cpu)opCode93() byte {
	v1 := int16(cpu.reg.A)
	v2 := int16(cpu.reg.E)
	res := v1 - v2
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	cpu.setFlagN()
	if (v1 & 0xF) - (v2 & 0xF) < 0 {
		cpu.setFlagH()
	}
	if res < 0 {
		cpu.setFlagC()
	}
	return 0
}

// SUB H
func (cpu *Cpu)opCode94() byte {
	v1 := int16(cpu.reg.A)
	v2 := int16(cpu.reg.H)
	res := v1 - v2
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	cpu.setFlagN()
	if (v1 & 0xF) - (v2 & 0xF) < 0 {
		cpu.setFlagH()
	}
	if res < 0 {
		cpu.setFlagC()
	}
	return 0
}

// SUB L
func (cpu *Cpu)opCode95() byte {
	v1 := int16(cpu.reg.A)
	v2 := int16(cpu.reg.L)
	res := v1 - v2
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	cpu.setFlagN()
	if (v1 & 0xF) - (v2 & 0xF) < 0 {
		cpu.setFlagH()
	}
	if res < 0 {
		cpu.setFlagC()
	}
	return 0
}

// SUB (HL)
func (cpu *Cpu)opCode96() byte {
	v1 := int16(cpu.reg.A)
	v2 := int16(cpu.gb.Mem.Read(cpu.reg.getHL()))
	res := v1 - v2
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	cpu.setFlagN()
	if (v1 & 0xF) - (v2 & 0xF) < 0 {
		cpu.setFlagH()
	}
	if res < 0 {
		cpu.setFlagC()
	}
	return 0
}

// SUB A
func (cpu *Cpu)opCode97() byte {
	v1 := int16(cpu.reg.A)
	v2 := int16(cpu.reg.A)
	res := v1 - v2
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	cpu.setFlagN()
	if (v1 & 0xF) - (v2 & 0xF) < 0 {
		cpu.setFlagH()
	}
	if res < 0 {
		cpu.setFlagC()
	}
	return 0
}

// SBC A,B
func (cpu *Cpu)opCode98() byte {
	v1 := int16(cpu.reg.A)
	v2 := int16(cpu.reg.B)
	carry := int16(0)
	if cpu.getFlagC() {
		carry = int16(1)
	}
	res := v1 - v2 - carry
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	cpu.setFlagN()
	if (v1 & 0xF) - (v2 & 0xF) - (carry & 0xF) < 0 {
		cpu.setFlagH()
	}
	if res < 0 {
		cpu.setFlagC()
	}
	return 0
}

// SBC A,C
func (cpu *Cpu)opCode99() byte {
	v1 := int16(cpu.reg.A)
	v2 := int16(cpu.reg.C)
	carry := int16(0)
	if cpu.getFlagC() {
		carry = int16(1)
	}
	res := v1 - v2 - carry
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	cpu.setFlagN()
	if (v1 & 0xF) - (v2 & 0xF) - (carry & 0xF) < 0 {
		cpu.setFlagH()
	}
	if res < 0 {
		cpu.setFlagC()
	}
	return 0
}

// SBC A,D
func (cpu *Cpu)opCode9A() byte {
	v1 := int16(cpu.reg.A)
	v2 := int16(cpu.reg.D)
	carry := int16(0)
	if cpu.getFlagC() {
		carry = int16(1)
	}
	res := v1 - v2 - carry
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	cpu.setFlagN()
	if (v1 & 0xF) - (v2 & 0xF) - (carry & 0xF) < 0 {
		cpu.setFlagH()
	}
	if res < 0 {
		cpu.setFlagC()
	}
	return 0
}

// SBC A,E
func (cpu *Cpu)opCode9B() byte {
	v1 := int16(cpu.reg.A)
	v2 := int16(cpu.reg.E)
	carry := int16(0)
	if cpu.getFlagC() {
		carry = int16(1)
	}
	res := v1 - v2 - carry
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	cpu.setFlagN()
	if (v1 & 0xF) - (v2 & 0xF) - (carry & 0xF) < 0 {
		cpu.setFlagH()
	}
	if res < 0 {
		cpu.setFlagC()
	}
	return 0
}

// SBC A,H
func (cpu *Cpu)opCode9C() byte {
	v1 := int16(cpu.reg.A)
	v2 := int16(cpu.reg.H)
	carry := int16(0)
	if cpu.getFlagC() {
		carry = int16(1)
	}
	res := v1 - v2 - carry
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	cpu.setFlagN()
	if (v1 & 0xF) - (v2 & 0xF) - (carry & 0xF) < 0 {
		cpu.setFlagH()
	}
	if res < 0 {
		cpu.setFlagC()
	}
	return 0
}

// SBC A,L
func (cpu *Cpu)opCode9D() byte {
	v1 := int16(cpu.reg.A)
	v2 := int16(cpu.reg.L)
	carry := int16(0)
	if cpu.getFlagC() {
		carry = int16(1)
	}
	res := v1 - v2 - carry
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	cpu.setFlagN()
	if (v1 & 0xF) - (v2 & 0xF) - (carry & 0xF) < 0 {
		cpu.setFlagH()
	}
	if res < 0 {
		cpu.setFlagC()
	}
	return 0
}

// SBC A,(HL)
func (cpu *Cpu)opCode9E() byte {
	v1 := int16(cpu.reg.A)
	v2 := int16(cpu.gb.Mem.Read(cpu.reg.getHL()))
	carry := int16(0)
	if cpu.getFlagC() {
		carry = int16(1)
	}
	res := v1 - v2 - carry
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	cpu.setFlagN()
	if (v1 & 0xF) - (v2 & 0xF) - (carry & 0xF) < 0 {
		cpu.setFlagH()
	}
	if res < 0 {
		cpu.setFlagC()
	}
	return 0
}

// SBC A,A
func (cpu *Cpu)opCode9F() byte {
	v1 := int16(cpu.reg.A)
	v2 := int16(cpu.reg.A)
	carry := int16(0)
	if cpu.getFlagC() {
		carry = int16(1)
	}
	res := v1 - v2 - carry
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	cpu.setFlagN()
	if (v1 & 0xF) - (v2 & 0xF) - (carry & 0xF) < 0 {
		cpu.setFlagH()
	}
	if res < 0 {
		cpu.setFlagC()
	}
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
	cpu.reg.A = cpu.reg.A & cpu.gb.Mem.Read(cpu.reg.getHL())
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
	cpu.reg.A = cpu.reg.A ^ cpu.gb.Mem.Read(cpu.reg.getHL())
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
	cpu.reg.A = cpu.reg.A | cpu.gb.Mem.Read(cpu.reg.getHL())
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

// POP BC
func (cpu *Cpu)opCodeC1() byte {
	cpu.reg.setBC(cpu.gb.Mem.ReadUint16(cpu.reg.SP))
	cpu.reg.SP += 2
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

// PUSH BC
func (cpu *Cpu)opCodeC5() byte {
	cpu.reg.SP -= 2
	cpu.gb.Mem.WriteUint16(cpu.reg.SP, cpu.reg.getBC())
	return 0
}

// ADD A,d8
func (cpu *Cpu)opCodeC6() byte {
	a := uint16(cpu.reg.A)
	arg := uint16(cpu.gb.Mem.Read(cpu.reg.PC))
	cpu.reg.PC++
	res := a + arg
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if byte(res) == 0 {
		cpu.setFlagZ()
	}
	carry := a ^ arg ^ res
	if (carry & 0x10) > 0 {
		cpu.setFlagH()
	}
	if (carry & 0x100) > 0 {
		cpu.setFlagC()
	}
	return 0
}

// ADC A,d8
func (cpu *Cpu)opCodeCE() byte {
	v1 := uint16(cpu.reg.A)
	v2 := uint16(cpu.gb.Mem.Read(cpu.reg.PC))
	cpu.reg.PC++
	carry := uint16(0)
	if cpu.getFlagC() {
		carry = uint16(1)
	}
	res := v1 + v2 + carry
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if byte(res) == 0 {
		cpu.setFlagZ()
	}
	if (v1 & 0x000F) + (v2 & 0x000F) + carry > 0x000F {
		cpu.setFlagH()
	}
	if res > 0x00FF {
		cpu.setFlagC()
	}
	return 0
}

// POP DE
func (cpu *Cpu)opCodeD1() byte {
	cpu.reg.setDE(cpu.gb.Mem.ReadUint16(cpu.reg.SP))
	cpu.reg.SP += 2
	return 0
}

// PUSH DE
func (cpu *Cpu)opCodeD5() byte {
	cpu.reg.SP -= 2
	cpu.gb.Mem.WriteUint16(cpu.reg.SP, cpu.reg.getDE())
	return 0
}

// SUB d8
func (cpu *Cpu)opCodeD6() byte {
	v1 := int16(cpu.reg.A)
	v2 := int16(cpu.gb.Mem.Read(cpu.reg.PC))
	cpu.reg.PC++
	res := v1 - v2
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	cpu.setFlagN()
	if (v1 & 0xF) - (v2 & 0xF) < 0 {
		cpu.setFlagH()
	}
	if res < 0 {
		cpu.setFlagC()
	}
	return 0
}

// SBC A,d8
func (cpu *Cpu)opCodeDE() byte {
	v1 := int16(cpu.reg.A)
	v2 := int16(cpu.gb.Mem.Read(cpu.reg.PC))
	carry := int16(0)
	if cpu.getFlagC() {
		carry = int16(1)
	}
	cpu.reg.PC++
	res := v1 - v2 - carry
	cpu.reg.A = byte(res)
	cpu.resetFlagZNHC()
	if cpu.reg.A == 0 {
		cpu.setFlagZ()
	}
	cpu.setFlagN()
	if (v1 & 0xF) - (v2 & 0xF) - (carry & 0xF) < 0 {
		cpu.setFlagH()
	}
	if res < 0 {
		cpu.setFlagC()
	}
	return 0
}

// LDH (a8),A
func (cpu *Cpu)opCodeE0() byte {
	addr := cpu.gb.Mem.Read(cpu.reg.PC)
	cpu.reg.PC++
	cpu.gb.Mem.Write(0xFF00 + uint16(addr), cpu.reg.A)
	return 0
}

// POP HL
func (cpu *Cpu)opCodeE1() byte {
	cpu.reg.setHL(cpu.gb.Mem.ReadUint16(cpu.reg.SP))
	cpu.reg.SP += 2
	return 0
}

// LD (C),A
func (cpu *Cpu)opCodeE2() byte {
	cpu.gb.Mem.Write(0xFF00 + uint16(cpu.reg.C), cpu.reg.A)
	return 0
}

// PUSH HL
func (cpu *Cpu)opCodeE5() byte {
	cpu.reg.SP -= 2
	cpu.gb.Mem.WriteUint16(cpu.reg.SP, cpu.reg.getHL())
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

// LD (a16),A
func (cpu *Cpu)opCodeEA() byte {
	addr := cpu.gb.Mem.ReadUint16(cpu.reg.PC)
	cpu.reg.PC += 2
	cpu.gb.Mem.Write(addr, cpu.reg.A)
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
	arg := cpu.gb.Mem.Read(cpu.reg.PC)
	cpu.reg.PC++
	cpu.reg.A = cpu.gb.Mem.Read(0xFF00 + uint16(arg))
	cpu.reg.PC++
	return 0
}

// POP AF
func (cpu *Cpu)opCodeF1() byte {
	cpu.reg.setAF(cpu.gb.Mem.ReadUint16(cpu.reg.SP))
	cpu.reg.SP += 2
	return 0
}

// LD A,(C)
func (cpu *Cpu)opCodeF2() byte {
	cpu.reg.A = cpu.gb.Mem.Read(0xFF00 + uint16(cpu.reg.C))
	return 0
}

// DI
func (cpu *Cpu)opCodeF3() byte {
	cpu.setFlagIME(false)
	return 0
}

// PUSH AF
func (cpu *Cpu)opCodeF5() byte {
	cpu.reg.SP -= 2
	cpu.gb.Mem.WriteUint16(cpu.reg.SP, cpu.reg.getAF())
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

// LD HL,SP+r8
func (cpu *Cpu)opCodeF8() byte {
	arg := uint32(cpu.gb.Mem.Read(cpu.reg.PC))
	cpu.reg.PC++
	sp := uint32(cpu.reg.SP)
	ret := arg + sp
	cpu.reg.setHL(uint16(ret & 0x0000FFFF))
	cpu.resetFlagZNHC()
	carry := arg ^ sp ^ ret
	if (carry & 0x10) > 0 {
		cpu.setFlagH()
	}
	if (carry & 0x100) > 0 {
		cpu.setFlagC()
	}
	return 0
}

// LD SP,HL
func (cpu *Cpu)opCodeF9() byte {
	cpu.reg.SP = cpu.reg.getHL()
	return 0
}

// LD A,(a16)
func (cpu *Cpu)opCodeFA() byte {
	addr := cpu.gb.Mem.ReadUint16(cpu.reg.PC)
	cpu.reg.PC += 2
	cpu.reg.A = cpu.gb.Mem.Read(addr)
	return 0
}

// EI
func (cpu *Cpu)opCodeFB() byte {
	cpu.setFlagIME(true)
	return 0
}