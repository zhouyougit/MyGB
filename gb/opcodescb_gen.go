package gb

var OPCodesMapCB = [0x100]OPCode{
	0x00: {
		Func: (*Cpu).opCodeCB00,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RLC B",
	},
	0x01: {
		Func: (*Cpu).opCodeCB01,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RLC C",
	},
	0x02: {
		Func: (*Cpu).opCodeCB02,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RLC D",
	},
	0x03: {
		Func: (*Cpu).opCodeCB03,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RLC E",
	},
	0x04: {
		Func: (*Cpu).opCodeCB04,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RLC H",
	},
	0x05: {
		Func: (*Cpu).opCodeCB05,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RLC L",
	},
	0x06: {
		Func: (*Cpu).opCodeCB06,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RLC (HL)",
	},
	0x07: {
		Func: (*Cpu).opCodeCB07,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RLC A",
	},
	0x08: {
		Func: (*Cpu).opCodeCB08,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RRC B",
	},
	0x09: {
		Func: (*Cpu).opCodeCB09,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RRC C",
	},
	0x0A: {
		Func: (*Cpu).opCodeCB0A,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RRC D",
	},
	0x0B: {
		Func: (*Cpu).opCodeCB0B,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RRC E",
	},
	0x0C: {
		Func: (*Cpu).opCodeCB0C,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RRC H",
	},
	0x0D: {
		Func: (*Cpu).opCodeCB0D,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RRC L",
	},
	0x0E: {
		Func: (*Cpu).opCodeCB0E,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RRC (HL)",
	},
	0x0F: {
		Func: (*Cpu).opCodeCB0F,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RRC A",
	},
	0x10: {
		Func: (*Cpu).opCodeCB10,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RL B",
	},
	0x11: {
		Func: (*Cpu).opCodeCB11,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RL C",
	},
	0x12: {
		Func: (*Cpu).opCodeCB12,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RL D",
	},
	0x13: {
		Func: (*Cpu).opCodeCB13,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RL E",
	},
	0x14: {
		Func: (*Cpu).opCodeCB14,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RL H",
	},
	0x15: {
		Func: (*Cpu).opCodeCB15,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RL L",
	},
	0x16: {
		Func: (*Cpu).opCodeCB16,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RL (HL)",
	},
	0x17: {
		Func: (*Cpu).opCodeCB17,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RL A",
	},
	0x18: {
		Func: (*Cpu).opCodeCB18,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RR B",
	},
	0x19: {
		Func: (*Cpu).opCodeCB19,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RR C",
	},
	0x1A: {
		Func: (*Cpu).opCodeCB1A,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RR D",
	},
	0x1B: {
		Func: (*Cpu).opCodeCB1B,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RR E",
	},
	0x1C: {
		Func: (*Cpu).opCodeCB1C,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RR H",
	},
	0x1D: {
		Func: (*Cpu).opCodeCB1D,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RR L",
	},
	0x1E: {
		Func: (*Cpu).opCodeCB1E,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RR (HL)",
	},
	0x1F: {
		Func: (*Cpu).opCodeCB1F,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RR A",
	},
	0x20: {
		Func: (*Cpu).opCodeCB20,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SLA B",
	},
	0x21: {
		Func: (*Cpu).opCodeCB21,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SLA C",
	},
	0x22: {
		Func: (*Cpu).opCodeCB22,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SLA D",
	},
	0x23: {
		Func: (*Cpu).opCodeCB23,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SLA E",
	},
	0x24: {
		Func: (*Cpu).opCodeCB24,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SLA H",
	},
	0x25: {
		Func: (*Cpu).opCodeCB25,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SLA L",
	},
	0x26: {
		Func: (*Cpu).opCodeCB26,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SLA (HL)",
	},
	0x27: {
		Func: (*Cpu).opCodeCB27,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SLA A",
	},
	0x28: {
		Func: (*Cpu).opCodeCB28,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SRA B",
	},
	0x29: {
		Func: (*Cpu).opCodeCB29,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SRA C",
	},
	0x2A: {
		Func: (*Cpu).opCodeCB2A,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SRA D",
	},
	0x2B: {
		Func: (*Cpu).opCodeCB2B,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SRA E",
	},
	0x2C: {
		Func: (*Cpu).opCodeCB2C,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SRA H",
	},
	0x2D: {
		Func: (*Cpu).opCodeCB2D,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SRA L",
	},
	0x2E: {
		Func: (*Cpu).opCodeCB2E,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SRA (HL)",
	},
	0x2F: {
		Func: (*Cpu).opCodeCB2F,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SRA A",
	},
	0x30: {
		Func: (*Cpu).opCodeCB30,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SWAP B",
	},
	0x31: {
		Func: (*Cpu).opCodeCB31,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SWAP C",
	},
	0x32: {
		Func: (*Cpu).opCodeCB32,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SWAP D",
	},
	0x33: {
		Func: (*Cpu).opCodeCB33,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SWAP E",
	},
	0x34: {
		Func: (*Cpu).opCodeCB34,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SWAP H",
	},
	0x35: {
		Func: (*Cpu).opCodeCB35,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SWAP L",
	},
	0x36: {
		Func: (*Cpu).opCodeCB36,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SWAP (HL)",
	},
	0x37: {
		Func: (*Cpu).opCodeCB37,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SWAP A",
	},
	0x38: {
		Func: (*Cpu).opCodeCB38,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SRL B",
	},
	0x39: {
		Func: (*Cpu).opCodeCB39,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SRL C",
	},
	0x3A: {
		Func: (*Cpu).opCodeCB3A,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SRL D",
	},
	0x3B: {
		Func: (*Cpu).opCodeCB3B,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SRL E",
	},
	0x3C: {
		Func: (*Cpu).opCodeCB3C,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SRL H",
	},
	0x3D: {
		Func: (*Cpu).opCodeCB3D,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SRL L",
	},
	0x3E: {
		Func: (*Cpu).opCodeCB3E,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SRL (HL)",
	},
	0x3F: {
		Func: (*Cpu).opCodeCB3F,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SRL A",
	},
	0x40: {
		Func: (*Cpu).opCodeCB40,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 0,B",
	},
	0x41: {
		Func: (*Cpu).opCodeCB41,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 0,C",
	},
	0x42: {
		Func: (*Cpu).opCodeCB42,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 0,D",
	},
	0x43: {
		Func: (*Cpu).opCodeCB43,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 0,E",
	},
	0x44: {
		Func: (*Cpu).opCodeCB44,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 0,H",
	},
	0x45: {
		Func: (*Cpu).opCodeCB45,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 0,L",
	},
	0x46: {
		Func: (*Cpu).opCodeCB46,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 0,(HL)",
	},
	0x47: {
		Func: (*Cpu).opCodeCB47,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 0,A",
	},
	0x48: {
		Func: (*Cpu).opCodeCB48,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 1,B",
	},
	0x49: {
		Func: (*Cpu).opCodeCB49,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 1,C",
	},
	0x4A: {
		Func: (*Cpu).opCodeCB4A,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 1,D",
	},
	0x4B: {
		Func: (*Cpu).opCodeCB4B,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 1,E",
	},
	0x4C: {
		Func: (*Cpu).opCodeCB4C,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 1,H",
	},
	0x4D: {
		Func: (*Cpu).opCodeCB4D,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 1,L",
	},
	0x4E: {
		Func: (*Cpu).opCodeCB4E,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 1,(HL)",
	},
	0x4F: {
		Func: (*Cpu).opCodeCB4F,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 1,A",
	},
	0x50: {
		Func: (*Cpu).opCodeCB50,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 2,B",
	},
	0x51: {
		Func: (*Cpu).opCodeCB51,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 2,C",
	},
	0x52: {
		Func: (*Cpu).opCodeCB52,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 2,D",
	},
	0x53: {
		Func: (*Cpu).opCodeCB53,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 2,E",
	},
	0x54: {
		Func: (*Cpu).opCodeCB54,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 2,H",
	},
	0x55: {
		Func: (*Cpu).opCodeCB55,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 2,L",
	},
	0x56: {
		Func: (*Cpu).opCodeCB56,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 2,(HL)",
	},
	0x57: {
		Func: (*Cpu).opCodeCB57,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 2,A",
	},
	0x58: {
		Func: (*Cpu).opCodeCB58,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 3,B",
	},
	0x59: {
		Func: (*Cpu).opCodeCB59,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 3,C",
	},
	0x5A: {
		Func: (*Cpu).opCodeCB5A,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 3,D",
	},
	0x5B: {
		Func: (*Cpu).opCodeCB5B,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 3,E",
	},
	0x5C: {
		Func: (*Cpu).opCodeCB5C,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 3,H",
	},
	0x5D: {
		Func: (*Cpu).opCodeCB5D,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 3,L",
	},
	0x5E: {
		Func: (*Cpu).opCodeCB5E,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 3,(HL)",
	},
	0x5F: {
		Func: (*Cpu).opCodeCB5F,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 3,A",
	},
	0x60: {
		Func: (*Cpu).opCodeCB60,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 4,B",
	},
	0x61: {
		Func: (*Cpu).opCodeCB61,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 4,C",
	},
	0x62: {
		Func: (*Cpu).opCodeCB62,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 4,D",
	},
	0x63: {
		Func: (*Cpu).opCodeCB63,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 4,E",
	},
	0x64: {
		Func: (*Cpu).opCodeCB64,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 4,H",
	},
	0x65: {
		Func: (*Cpu).opCodeCB65,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 4,L",
	},
	0x66: {
		Func: (*Cpu).opCodeCB66,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 4,(HL)",
	},
	0x67: {
		Func: (*Cpu).opCodeCB67,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 4,A",
	},
	0x68: {
		Func: (*Cpu).opCodeCB68,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 5,B",
	},
	0x69: {
		Func: (*Cpu).opCodeCB69,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 5,C",
	},
	0x6A: {
		Func: (*Cpu).opCodeCB6A,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 5,D",
	},
	0x6B: {
		Func: (*Cpu).opCodeCB6B,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 5,E",
	},
	0x6C: {
		Func: (*Cpu).opCodeCB6C,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 5,H",
	},
	0x6D: {
		Func: (*Cpu).opCodeCB6D,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 5,L",
	},
	0x6E: {
		Func: (*Cpu).opCodeCB6E,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 5,(HL)",
	},
	0x6F: {
		Func: (*Cpu).opCodeCB6F,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 5,A",
	},
	0x70: {
		Func: (*Cpu).opCodeCB70,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 6,B",
	},
	0x71: {
		Func: (*Cpu).opCodeCB71,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 6,C",
	},
	0x72: {
		Func: (*Cpu).opCodeCB72,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 6,D",
	},
	0x73: {
		Func: (*Cpu).opCodeCB73,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 6,E",
	},
	0x74: {
		Func: (*Cpu).opCodeCB74,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 6,H",
	},
	0x75: {
		Func: (*Cpu).opCodeCB75,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 6,L",
	},
	0x76: {
		Func: (*Cpu).opCodeCB76,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 6,(HL)",
	},
	0x77: {
		Func: (*Cpu).opCodeCB77,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 6,A",
	},
	0x78: {
		Func: (*Cpu).opCodeCB78,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 7,B",
	},
	0x79: {
		Func: (*Cpu).opCodeCB79,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 7,C",
	},
	0x7A: {
		Func: (*Cpu).opCodeCB7A,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 7,D",
	},
	0x7B: {
		Func: (*Cpu).opCodeCB7B,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 7,E",
	},
	0x7C: {
		Func: (*Cpu).opCodeCB7C,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 7,H",
	},
	0x7D: {
		Func: (*Cpu).opCodeCB7D,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 7,L",
	},
	0x7E: {
		Func: (*Cpu).opCodeCB7E,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 7,(HL)",
	},
	0x7F: {
		Func: (*Cpu).opCodeCB7F,
		Cycles: nil,
		Length: 2,
		Mnemonic: "BIT 7,A",
	},
	0x80: {
		Func: (*Cpu).opCodeCB80,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 0,B",
	},
	0x81: {
		Func: (*Cpu).opCodeCB81,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 0,C",
	},
	0x82: {
		Func: (*Cpu).opCodeCB82,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 0,D",
	},
	0x83: {
		Func: (*Cpu).opCodeCB83,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 0,E",
	},
	0x84: {
		Func: (*Cpu).opCodeCB84,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 0,H",
	},
	0x85: {
		Func: (*Cpu).opCodeCB85,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 0,L",
	},
	0x86: {
		Func: (*Cpu).opCodeCB86,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 0,(HL)",
	},
	0x87: {
		Func: (*Cpu).opCodeCB87,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 0,A",
	},
	0x88: {
		Func: (*Cpu).opCodeCB88,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 1,B",
	},
	0x89: {
		Func: (*Cpu).opCodeCB89,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 1,C",
	},
	0x8A: {
		Func: (*Cpu).opCodeCB8A,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 1,D",
	},
	0x8B: {
		Func: (*Cpu).opCodeCB8B,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 1,E",
	},
	0x8C: {
		Func: (*Cpu).opCodeCB8C,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 1,H",
	},
	0x8D: {
		Func: (*Cpu).opCodeCB8D,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 1,L",
	},
	0x8E: {
		Func: (*Cpu).opCodeCB8E,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 1,(HL)",
	},
	0x8F: {
		Func: (*Cpu).opCodeCB8F,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 1,A",
	},
	0x90: {
		Func: (*Cpu).opCodeCB90,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 2,B",
	},
	0x91: {
		Func: (*Cpu).opCodeCB91,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 2,C",
	},
	0x92: {
		Func: (*Cpu).opCodeCB92,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 2,D",
	},
	0x93: {
		Func: (*Cpu).opCodeCB93,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 2,E",
	},
	0x94: {
		Func: (*Cpu).opCodeCB94,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 2,H",
	},
	0x95: {
		Func: (*Cpu).opCodeCB95,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 2,L",
	},
	0x96: {
		Func: (*Cpu).opCodeCB96,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 2,(HL)",
	},
	0x97: {
		Func: (*Cpu).opCodeCB97,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 2,A",
	},
	0x98: {
		Func: (*Cpu).opCodeCB98,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 3,B",
	},
	0x99: {
		Func: (*Cpu).opCodeCB99,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 3,C",
	},
	0x9A: {
		Func: (*Cpu).opCodeCB9A,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 3,D",
	},
	0x9B: {
		Func: (*Cpu).opCodeCB9B,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 3,E",
	},
	0x9C: {
		Func: (*Cpu).opCodeCB9C,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 3,H",
	},
	0x9D: {
		Func: (*Cpu).opCodeCB9D,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 3,L",
	},
	0x9E: {
		Func: (*Cpu).opCodeCB9E,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 3,(HL)",
	},
	0x9F: {
		Func: (*Cpu).opCodeCB9F,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 3,A",
	},
	0xA0: {
		Func: (*Cpu).opCodeCBA0,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 4,B",
	},
	0xA1: {
		Func: (*Cpu).opCodeCBA1,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 4,C",
	},
	0xA2: {
		Func: (*Cpu).opCodeCBA2,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 4,D",
	},
	0xA3: {
		Func: (*Cpu).opCodeCBA3,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 4,E",
	},
	0xA4: {
		Func: (*Cpu).opCodeCBA4,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 4,H",
	},
	0xA5: {
		Func: (*Cpu).opCodeCBA5,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 4,L",
	},
	0xA6: {
		Func: (*Cpu).opCodeCBA6,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 4,(HL)",
	},
	0xA7: {
		Func: (*Cpu).opCodeCBA7,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 4,A",
	},
	0xA8: {
		Func: (*Cpu).opCodeCBA8,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 5,B",
	},
	0xA9: {
		Func: (*Cpu).opCodeCBA9,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 5,C",
	},
	0xAA: {
		Func: (*Cpu).opCodeCBAA,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 5,D",
	},
	0xAB: {
		Func: (*Cpu).opCodeCBAB,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 5,E",
	},
	0xAC: {
		Func: (*Cpu).opCodeCBAC,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 5,H",
	},
	0xAD: {
		Func: (*Cpu).opCodeCBAD,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 5,L",
	},
	0xAE: {
		Func: (*Cpu).opCodeCBAE,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 5,(HL)",
	},
	0xAF: {
		Func: (*Cpu).opCodeCBAF,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 5,A",
	},
	0xB0: {
		Func: (*Cpu).opCodeCBB0,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 6,B",
	},
	0xB1: {
		Func: (*Cpu).opCodeCBB1,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 6,C",
	},
	0xB2: {
		Func: (*Cpu).opCodeCBB2,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 6,D",
	},
	0xB3: {
		Func: (*Cpu).opCodeCBB3,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 6,E",
	},
	0xB4: {
		Func: (*Cpu).opCodeCBB4,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 6,H",
	},
	0xB5: {
		Func: (*Cpu).opCodeCBB5,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 6,L",
	},
	0xB6: {
		Func: (*Cpu).opCodeCBB6,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 6,(HL)",
	},
	0xB7: {
		Func: (*Cpu).opCodeCBB7,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 6,A",
	},
	0xB8: {
		Func: (*Cpu).opCodeCBB8,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 7,B",
	},
	0xB9: {
		Func: (*Cpu).opCodeCBB9,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 7,C",
	},
	0xBA: {
		Func: (*Cpu).opCodeCBBA,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 7,D",
	},
	0xBB: {
		Func: (*Cpu).opCodeCBBB,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 7,E",
	},
	0xBC: {
		Func: (*Cpu).opCodeCBBC,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 7,H",
	},
	0xBD: {
		Func: (*Cpu).opCodeCBBD,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 7,L",
	},
	0xBE: {
		Func: (*Cpu).opCodeCBBE,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 7,(HL)",
	},
	0xBF: {
		Func: (*Cpu).opCodeCBBF,
		Cycles: nil,
		Length: 2,
		Mnemonic: "RES 7,A",
	},
	0xC0: {
		Func: (*Cpu).opCodeCBC0,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 0,B",
	},
	0xC1: {
		Func: (*Cpu).opCodeCBC1,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 0,C",
	},
	0xC2: {
		Func: (*Cpu).opCodeCBC2,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 0,D",
	},
	0xC3: {
		Func: (*Cpu).opCodeCBC3,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 0,E",
	},
	0xC4: {
		Func: (*Cpu).opCodeCBC4,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 0,H",
	},
	0xC5: {
		Func: (*Cpu).opCodeCBC5,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 0,L",
	},
	0xC6: {
		Func: (*Cpu).opCodeCBC6,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 0,(HL)",
	},
	0xC7: {
		Func: (*Cpu).opCodeCBC7,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 0,A",
	},
	0xC8: {
		Func: (*Cpu).opCodeCBC8,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 1,B",
	},
	0xC9: {
		Func: (*Cpu).opCodeCBC9,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 1,C",
	},
	0xCA: {
		Func: (*Cpu).opCodeCBCA,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 1,D",
	},
	0xCB: {
		Func: (*Cpu).opCodeCBCB,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 1,E",
	},
	0xCC: {
		Func: (*Cpu).opCodeCBCC,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 1,H",
	},
	0xCD: {
		Func: (*Cpu).opCodeCBCD,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 1,L",
	},
	0xCE: {
		Func: (*Cpu).opCodeCBCE,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 1,(HL)",
	},
	0xCF: {
		Func: (*Cpu).opCodeCBCF,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 1,A",
	},
	0xD0: {
		Func: (*Cpu).opCodeCBD0,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 2,B",
	},
	0xD1: {
		Func: (*Cpu).opCodeCBD1,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 2,C",
	},
	0xD2: {
		Func: (*Cpu).opCodeCBD2,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 2,D",
	},
	0xD3: {
		Func: (*Cpu).opCodeCBD3,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 2,E",
	},
	0xD4: {
		Func: (*Cpu).opCodeCBD4,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 2,H",
	},
	0xD5: {
		Func: (*Cpu).opCodeCBD5,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 2,L",
	},
	0xD6: {
		Func: (*Cpu).opCodeCBD6,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 2,(HL)",
	},
	0xD7: {
		Func: (*Cpu).opCodeCBD7,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 2,A",
	},
	0xD8: {
		Func: (*Cpu).opCodeCBD8,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 3,B",
	},
	0xD9: {
		Func: (*Cpu).opCodeCBD9,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 3,C",
	},
	0xDA: {
		Func: (*Cpu).opCodeCBDA,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 3,D",
	},
	0xDB: {
		Func: (*Cpu).opCodeCBDB,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 3,E",
	},
	0xDC: {
		Func: (*Cpu).opCodeCBDC,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 3,H",
	},
	0xDD: {
		Func: (*Cpu).opCodeCBDD,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 3,L",
	},
	0xDE: {
		Func: (*Cpu).opCodeCBDE,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 3,(HL)",
	},
	0xDF: {
		Func: (*Cpu).opCodeCBDF,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 3,A",
	},
	0xE0: {
		Func: (*Cpu).opCodeCBE0,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 4,B",
	},
	0xE1: {
		Func: (*Cpu).opCodeCBE1,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 4,C",
	},
	0xE2: {
		Func: (*Cpu).opCodeCBE2,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 4,D",
	},
	0xE3: {
		Func: (*Cpu).opCodeCBE3,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 4,E",
	},
	0xE4: {
		Func: (*Cpu).opCodeCBE4,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 4,H",
	},
	0xE5: {
		Func: (*Cpu).opCodeCBE5,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 4,L",
	},
	0xE6: {
		Func: (*Cpu).opCodeCBE6,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 4,(HL)",
	},
	0xE7: {
		Func: (*Cpu).opCodeCBE7,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 4,A",
	},
	0xE8: {
		Func: (*Cpu).opCodeCBE8,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 5,B",
	},
	0xE9: {
		Func: (*Cpu).opCodeCBE9,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 5,C",
	},
	0xEA: {
		Func: (*Cpu).opCodeCBEA,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 5,D",
	},
	0xEB: {
		Func: (*Cpu).opCodeCBEB,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 5,E",
	},
	0xEC: {
		Func: (*Cpu).opCodeCBEC,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 5,H",
	},
	0xED: {
		Func: (*Cpu).opCodeCBED,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 5,L",
	},
	0xEE: {
		Func: (*Cpu).opCodeCBEE,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 5,(HL)",
	},
	0xEF: {
		Func: (*Cpu).opCodeCBEF,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 5,A",
	},
	0xF0: {
		Func: (*Cpu).opCodeCBF0,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 6,B",
	},
	0xF1: {
		Func: (*Cpu).opCodeCBF1,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 6,C",
	},
	0xF2: {
		Func: (*Cpu).opCodeCBF2,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 6,D",
	},
	0xF3: {
		Func: (*Cpu).opCodeCBF3,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 6,E",
	},
	0xF4: {
		Func: (*Cpu).opCodeCBF4,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 6,H",
	},
	0xF5: {
		Func: (*Cpu).opCodeCBF5,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 6,L",
	},
	0xF6: {
		Func: (*Cpu).opCodeCBF6,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 6,(HL)",
	},
	0xF7: {
		Func: (*Cpu).opCodeCBF7,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 6,A",
	},
	0xF8: {
		Func: (*Cpu).opCodeCBF8,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 7,B",
	},
	0xF9: {
		Func: (*Cpu).opCodeCBF9,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 7,C",
	},
	0xFA: {
		Func: (*Cpu).opCodeCBFA,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 7,D",
	},
	0xFB: {
		Func: (*Cpu).opCodeCBFB,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 7,E",
	},
	0xFC: {
		Func: (*Cpu).opCodeCBFC,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 7,H",
	},
	0xFD: {
		Func: (*Cpu).opCodeCBFD,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 7,L",
	},
	0xFE: {
		Func: (*Cpu).opCodeCBFE,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 7,(HL)",
	},
	0xFF: {
		Func: (*Cpu).opCodeCBFF,
		Cycles: nil,
		Length: 2,
		Mnemonic: "SET 7,A",
	},
}


// RLC B
func (cpu *Cpu)opCodeCB00() byte {
	data := cpu.reg.B
	data = data << 1 | data >> 7
	cpu.reg.B = data
	cpu.resetFlagZNHC()
	if data & 0x80 == 0x80 {
		cpu.setFlagC()
	}
	return 2
}
        

// RLC C
func (cpu *Cpu)opCodeCB01() byte {
	data := cpu.reg.C
	data = data << 1 | data >> 7
	cpu.reg.C = data
	cpu.resetFlagZNHC()
	if data & 0x80 == 0x80 {
		cpu.setFlagC()
	}
	return 2
}
        

// RLC D
func (cpu *Cpu)opCodeCB02() byte {
	data := cpu.reg.D
	data = data << 1 | data >> 7
	cpu.reg.D = data
	cpu.resetFlagZNHC()
	if data & 0x80 == 0x80 {
		cpu.setFlagC()
	}
	return 2
}
        

// RLC E
func (cpu *Cpu)opCodeCB03() byte {
	data := cpu.reg.E
	data = data << 1 | data >> 7
	cpu.reg.E = data
	cpu.resetFlagZNHC()
	if data & 0x80 == 0x80 {
		cpu.setFlagC()
	}
	return 2
}
        

// RLC H
func (cpu *Cpu)opCodeCB04() byte {
	data := cpu.reg.H
	data = data << 1 | data >> 7
	cpu.reg.H = data
	cpu.resetFlagZNHC()
	if data & 0x80 == 0x80 {
		cpu.setFlagC()
	}
	return 2
}
        

// RLC L
func (cpu *Cpu)opCodeCB05() byte {
	data := cpu.reg.L
	data = data << 1 | data >> 7
	cpu.reg.L = data
	cpu.resetFlagZNHC()
	if data & 0x80 == 0x80 {
		cpu.setFlagC()
	}
	return 2
}
        

// RLC (HL)
func (cpu *Cpu)opCodeCB06() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	data = data << 1 | data >> 7
	cpu.gb.Mem.Write(cpu.reg.getHL(), data)
	cpu.resetFlagZNHC()
	if data & 0x80 == 0x80 {
		cpu.setFlagC()
	}
	return 0
}
        

// RLC A
func (cpu *Cpu)opCodeCB07() byte {
	data := cpu.reg.A
	data = data << 1 | data >> 7
	cpu.reg.A = data
	cpu.resetFlagZNHC()
	if data & 0x80 == 0x80 {
		cpu.setFlagC()
	}
	return 2
}
        

// RRC B
func (cpu *Cpu)opCodeCB08() byte {
	data := cpu.reg.B
	carry := data & 0x01 == 0x01
	data = (data >> 1) | ((data & 0x01) << 7)
	cpu.reg.B = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}

        

// RRC C
func (cpu *Cpu)opCodeCB09() byte {
	data := cpu.reg.C
	carry := data & 0x01 == 0x01
	data = (data >> 1) | ((data & 0x01) << 7)
	cpu.reg.C = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}

        

// RRC D
func (cpu *Cpu)opCodeCB0A() byte {
	data := cpu.reg.D
	carry := data & 0x01 == 0x01
	data = (data >> 1) | ((data & 0x01) << 7)
	cpu.reg.D = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}

        

// RRC E
func (cpu *Cpu)opCodeCB0B() byte {
	data := cpu.reg.E
	carry := data & 0x01 == 0x01
	data = (data >> 1) | ((data & 0x01) << 7)
	cpu.reg.E = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}

        

// RRC H
func (cpu *Cpu)opCodeCB0C() byte {
	data := cpu.reg.H
	carry := data & 0x01 == 0x01
	data = (data >> 1) | ((data & 0x01) << 7)
	cpu.reg.H = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}

        

// RRC L
func (cpu *Cpu)opCodeCB0D() byte {
	data := cpu.reg.L
	carry := data & 0x01 == 0x01
	data = (data >> 1) | ((data & 0x01) << 7)
	cpu.reg.L = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}

        

// RRC (HL)
func (cpu *Cpu)opCodeCB0E() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	carry := data & 0x01 == 0x01
	data = (data >> 1) | ((data & 0x01) << 7)
	cpu.gb.Mem.Write(cpu.reg.getHL(), data)
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 0
}

        

// RRC A
func (cpu *Cpu)opCodeCB0F() byte {
	data := cpu.reg.A
	carry := data & 0x01 == 0x01
	data = (data >> 1) | ((data & 0x01) << 7)
	cpu.reg.A = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}

        

// RL B
func (cpu *Cpu)opCodeCB10() byte {
	data := cpu.reg.B
	carry := data & 0x80 == 0x80
	data = data << 1
	if cpu.getFlagC() {
		data |= 0x01
	}
	cpu.reg.B = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// RL C
func (cpu *Cpu)opCodeCB11() byte {
	data := cpu.reg.C
	carry := data & 0x80 == 0x80
	data = data << 1
	if cpu.getFlagC() {
		data |= 0x01
	}
	cpu.reg.C = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// RL D
func (cpu *Cpu)opCodeCB12() byte {
	data := cpu.reg.D
	carry := data & 0x80 == 0x80
	data = data << 1
	if cpu.getFlagC() {
		data |= 0x01
	}
	cpu.reg.D = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// RL E
func (cpu *Cpu)opCodeCB13() byte {
	data := cpu.reg.E
	carry := data & 0x80 == 0x80
	data = data << 1
	if cpu.getFlagC() {
		data |= 0x01
	}
	cpu.reg.E = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// RL H
func (cpu *Cpu)opCodeCB14() byte {
	data := cpu.reg.H
	carry := data & 0x80 == 0x80
	data = data << 1
	if cpu.getFlagC() {
		data |= 0x01
	}
	cpu.reg.H = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// RL L
func (cpu *Cpu)opCodeCB15() byte {
	data := cpu.reg.L
	carry := data & 0x80 == 0x80
	data = data << 1
	if cpu.getFlagC() {
		data |= 0x01
	}
	cpu.reg.L = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// RL (HL)
func (cpu *Cpu)opCodeCB16() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	carry := data & 0x80 == 0x80
	data = data << 1
	if cpu.getFlagC() {
		data |= 0x01
	}
	cpu.gb.Mem.Write(cpu.reg.getHL(), data)
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 0
}
        

// RL A
func (cpu *Cpu)opCodeCB17() byte {
	data := cpu.reg.A
	carry := data & 0x80 == 0x80
	data = data << 1
	if cpu.getFlagC() {
		data |= 0x01
	}
	cpu.reg.A = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// RR B
func (cpu *Cpu)opCodeCB18() byte {
	data := cpu.reg.B
	carry := data & 0x01 == 0x01
	data = data >> 1
	if cpu.getFlagC() {
		data |= 0x80
	}
	cpu.reg.B = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// RR C
func (cpu *Cpu)opCodeCB19() byte {
	data := cpu.reg.C
	carry := data & 0x01 == 0x01
	data = data >> 1
	if cpu.getFlagC() {
		data |= 0x80
	}
	cpu.reg.C = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// RR D
func (cpu *Cpu)opCodeCB1A() byte {
	data := cpu.reg.D
	carry := data & 0x01 == 0x01
	data = data >> 1
	if cpu.getFlagC() {
		data |= 0x80
	}
	cpu.reg.D = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// RR E
func (cpu *Cpu)opCodeCB1B() byte {
	data := cpu.reg.E
	carry := data & 0x01 == 0x01
	data = data >> 1
	if cpu.getFlagC() {
		data |= 0x80
	}
	cpu.reg.E = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// RR H
func (cpu *Cpu)opCodeCB1C() byte {
	data := cpu.reg.H
	carry := data & 0x01 == 0x01
	data = data >> 1
	if cpu.getFlagC() {
		data |= 0x80
	}
	cpu.reg.H = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// RR L
func (cpu *Cpu)opCodeCB1D() byte {
	data := cpu.reg.L
	carry := data & 0x01 == 0x01
	data = data >> 1
	if cpu.getFlagC() {
		data |= 0x80
	}
	cpu.reg.L = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// RR (HL)
func (cpu *Cpu)opCodeCB1E() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	carry := data & 0x01 == 0x01
	data = data >> 1
	if cpu.getFlagC() {
		data |= 0x80
	}
	cpu.gb.Mem.Write(cpu.reg.getHL(), data)
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 0
}
        

// RR A
func (cpu *Cpu)opCodeCB1F() byte {
	data := cpu.reg.A
	carry := data & 0x01 == 0x01
	data = data >> 1
	if cpu.getFlagC() {
		data |= 0x80
	}
	cpu.reg.A = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// SLA B
func (cpu *Cpu)opCodeCB20() byte {
	data := cpu.reg.B
	carry := data & 0x80 == 0x80
	data = data << 1
	cpu.reg.B = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// SLA C
func (cpu *Cpu)opCodeCB21() byte {
	data := cpu.reg.C
	carry := data & 0x80 == 0x80
	data = data << 1
	cpu.reg.C = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// SLA D
func (cpu *Cpu)opCodeCB22() byte {
	data := cpu.reg.D
	carry := data & 0x80 == 0x80
	data = data << 1
	cpu.reg.D = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// SLA E
func (cpu *Cpu)opCodeCB23() byte {
	data := cpu.reg.E
	carry := data & 0x80 == 0x80
	data = data << 1
	cpu.reg.E = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// SLA H
func (cpu *Cpu)opCodeCB24() byte {
	data := cpu.reg.H
	carry := data & 0x80 == 0x80
	data = data << 1
	cpu.reg.H = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// SLA L
func (cpu *Cpu)opCodeCB25() byte {
	data := cpu.reg.L
	carry := data & 0x80 == 0x80
	data = data << 1
	cpu.reg.L = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// SLA (HL)
func (cpu *Cpu)opCodeCB26() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	carry := data & 0x80 == 0x80
	data = data << 1
	cpu.gb.Mem.Write(cpu.reg.getHL(), data)
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 0
}
        

// SLA A
func (cpu *Cpu)opCodeCB27() byte {
	data := cpu.reg.A
	carry := data & 0x80 == 0x80
	data = data << 1
	cpu.reg.A = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// SRA B
func (cpu *Cpu)opCodeCB28() byte {
	data := cpu.reg.B
	carry := data & 0x01 == 0x01
	data = (data >> 1) | (data & 0x80)
	cpu.reg.B = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// SRA C
func (cpu *Cpu)opCodeCB29() byte {
	data := cpu.reg.C
	carry := data & 0x01 == 0x01
	data = (data >> 1) | (data & 0x80)
	cpu.reg.C = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// SRA D
func (cpu *Cpu)opCodeCB2A() byte {
	data := cpu.reg.D
	carry := data & 0x01 == 0x01
	data = (data >> 1) | (data & 0x80)
	cpu.reg.D = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// SRA E
func (cpu *Cpu)opCodeCB2B() byte {
	data := cpu.reg.E
	carry := data & 0x01 == 0x01
	data = (data >> 1) | (data & 0x80)
	cpu.reg.E = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// SRA H
func (cpu *Cpu)opCodeCB2C() byte {
	data := cpu.reg.H
	carry := data & 0x01 == 0x01
	data = (data >> 1) | (data & 0x80)
	cpu.reg.H = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// SRA L
func (cpu *Cpu)opCodeCB2D() byte {
	data := cpu.reg.L
	carry := data & 0x01 == 0x01
	data = (data >> 1) | (data & 0x80)
	cpu.reg.L = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// SRA (HL)
func (cpu *Cpu)opCodeCB2E() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	carry := data & 0x01 == 0x01
	data = (data >> 1) | (data & 0x80)
	cpu.gb.Mem.Write(cpu.reg.getHL(), data)
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 0
}
        

// SRA A
func (cpu *Cpu)opCodeCB2F() byte {
	data := cpu.reg.A
	carry := data & 0x01 == 0x01
	data = (data >> 1) | (data & 0x80)
	cpu.reg.A = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// SWAP B
func (cpu *Cpu)opCodeCB30() byte {
	data := cpu.reg.B
	data = data << 4 | data >> 4
	cpu.reg.B = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	return 2
}
        

// SWAP C
func (cpu *Cpu)opCodeCB31() byte {
	data := cpu.reg.C
	data = data << 4 | data >> 4
	cpu.reg.C = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	return 2
}
        

// SWAP D
func (cpu *Cpu)opCodeCB32() byte {
	data := cpu.reg.D
	data = data << 4 | data >> 4
	cpu.reg.D = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	return 2
}
        

// SWAP E
func (cpu *Cpu)opCodeCB33() byte {
	data := cpu.reg.E
	data = data << 4 | data >> 4
	cpu.reg.E = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	return 2
}
        

// SWAP H
func (cpu *Cpu)opCodeCB34() byte {
	data := cpu.reg.H
	data = data << 4 | data >> 4
	cpu.reg.H = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	return 2
}
        

// SWAP L
func (cpu *Cpu)opCodeCB35() byte {
	data := cpu.reg.L
	data = data << 4 | data >> 4
	cpu.reg.L = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	return 2
}
        

// SWAP (HL)
func (cpu *Cpu)opCodeCB36() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	data = data << 4 | data >> 4
	cpu.gb.Mem.Write(cpu.reg.getHL(), data)
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	return 0
}
        

// SWAP A
func (cpu *Cpu)opCodeCB37() byte {
	data := cpu.reg.A
	data = data << 4 | data >> 4
	cpu.reg.A = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	return 2
}
        

// SRL B
func (cpu *Cpu)opCodeCB38() byte {
	data := cpu.reg.B
	carry := data & 0x01 == 0x01
	data = data >> 1
	cpu.reg.B = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// SRL C
func (cpu *Cpu)opCodeCB39() byte {
	data := cpu.reg.C
	carry := data & 0x01 == 0x01
	data = data >> 1
	cpu.reg.C = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// SRL D
func (cpu *Cpu)opCodeCB3A() byte {
	data := cpu.reg.D
	carry := data & 0x01 == 0x01
	data = data >> 1
	cpu.reg.D = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// SRL E
func (cpu *Cpu)opCodeCB3B() byte {
	data := cpu.reg.E
	carry := data & 0x01 == 0x01
	data = data >> 1
	cpu.reg.E = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// SRL H
func (cpu *Cpu)opCodeCB3C() byte {
	data := cpu.reg.H
	carry := data & 0x01 == 0x01
	data = data >> 1
	cpu.reg.H = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// SRL L
func (cpu *Cpu)opCodeCB3D() byte {
	data := cpu.reg.L
	carry := data & 0x01 == 0x01
	data = data >> 1
	cpu.reg.L = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// SRL (HL)
func (cpu *Cpu)opCodeCB3E() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	carry := data & 0x01 == 0x01
	data = data >> 1
	cpu.gb.Mem.Write(cpu.reg.getHL(), data)
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 0
}
        

// SRL A
func (cpu *Cpu)opCodeCB3F() byte {
	data := cpu.reg.A
	carry := data & 0x01 == 0x01
	data = data >> 1
	cpu.reg.A = data
	cpu.resetFlagZNHC()
	if data == 0 {
		cpu.setFlagZ()
	}
	if carry {
		cpu.setFlagC()
	}
	return 2
}
        

// BIT 0,B
func (cpu *Cpu)opCodeCB40() byte {
	data := cpu.reg.B
	data = data & (0x1 << 0)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 0,C
func (cpu *Cpu)opCodeCB41() byte {
	data := cpu.reg.C
	data = data & (0x1 << 0)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 0,D
func (cpu *Cpu)opCodeCB42() byte {
	data := cpu.reg.D
	data = data & (0x1 << 0)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 0,E
func (cpu *Cpu)opCodeCB43() byte {
	data := cpu.reg.E
	data = data & (0x1 << 0)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 0,H
func (cpu *Cpu)opCodeCB44() byte {
	data := cpu.reg.H
	data = data & (0x1 << 0)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 0,L
func (cpu *Cpu)opCodeCB45() byte {
	data := cpu.reg.L
	data = data & (0x1 << 0)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 0,(HL)
func (cpu *Cpu)opCodeCB46() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	data = data & (0x1 << 0)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 1
}
        

// BIT 0,A
func (cpu *Cpu)opCodeCB47() byte {
	data := cpu.reg.A
	data = data & (0x1 << 0)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 1,B
func (cpu *Cpu)opCodeCB48() byte {
	data := cpu.reg.B
	data = data & (0x1 << 1)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 1,C
func (cpu *Cpu)opCodeCB49() byte {
	data := cpu.reg.C
	data = data & (0x1 << 1)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 1,D
func (cpu *Cpu)opCodeCB4A() byte {
	data := cpu.reg.D
	data = data & (0x1 << 1)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 1,E
func (cpu *Cpu)opCodeCB4B() byte {
	data := cpu.reg.E
	data = data & (0x1 << 1)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 1,H
func (cpu *Cpu)opCodeCB4C() byte {
	data := cpu.reg.H
	data = data & (0x1 << 1)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 1,L
func (cpu *Cpu)opCodeCB4D() byte {
	data := cpu.reg.L
	data = data & (0x1 << 1)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 1,(HL)
func (cpu *Cpu)opCodeCB4E() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	data = data & (0x1 << 1)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 1
}
        

// BIT 1,A
func (cpu *Cpu)opCodeCB4F() byte {
	data := cpu.reg.A
	data = data & (0x1 << 1)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 2,B
func (cpu *Cpu)opCodeCB50() byte {
	data := cpu.reg.B
	data = data & (0x1 << 2)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 2,C
func (cpu *Cpu)opCodeCB51() byte {
	data := cpu.reg.C
	data = data & (0x1 << 2)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 2,D
func (cpu *Cpu)opCodeCB52() byte {
	data := cpu.reg.D
	data = data & (0x1 << 2)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 2,E
func (cpu *Cpu)opCodeCB53() byte {
	data := cpu.reg.E
	data = data & (0x1 << 2)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 2,H
func (cpu *Cpu)opCodeCB54() byte {
	data := cpu.reg.H
	data = data & (0x1 << 2)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 2,L
func (cpu *Cpu)opCodeCB55() byte {
	data := cpu.reg.L
	data = data & (0x1 << 2)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 2,(HL)
func (cpu *Cpu)opCodeCB56() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	data = data & (0x1 << 2)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 1
}
        

// BIT 2,A
func (cpu *Cpu)opCodeCB57() byte {
	data := cpu.reg.A
	data = data & (0x1 << 2)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 3,B
func (cpu *Cpu)opCodeCB58() byte {
	data := cpu.reg.B
	data = data & (0x1 << 3)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 3,C
func (cpu *Cpu)opCodeCB59() byte {
	data := cpu.reg.C
	data = data & (0x1 << 3)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 3,D
func (cpu *Cpu)opCodeCB5A() byte {
	data := cpu.reg.D
	data = data & (0x1 << 3)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 3,E
func (cpu *Cpu)opCodeCB5B() byte {
	data := cpu.reg.E
	data = data & (0x1 << 3)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 3,H
func (cpu *Cpu)opCodeCB5C() byte {
	data := cpu.reg.H
	data = data & (0x1 << 3)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 3,L
func (cpu *Cpu)opCodeCB5D() byte {
	data := cpu.reg.L
	data = data & (0x1 << 3)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 3,(HL)
func (cpu *Cpu)opCodeCB5E() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	data = data & (0x1 << 3)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 1
}
        

// BIT 3,A
func (cpu *Cpu)opCodeCB5F() byte {
	data := cpu.reg.A
	data = data & (0x1 << 3)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 4,B
func (cpu *Cpu)opCodeCB60() byte {
	data := cpu.reg.B
	data = data & (0x1 << 4)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 4,C
func (cpu *Cpu)opCodeCB61() byte {
	data := cpu.reg.C
	data = data & (0x1 << 4)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 4,D
func (cpu *Cpu)opCodeCB62() byte {
	data := cpu.reg.D
	data = data & (0x1 << 4)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 4,E
func (cpu *Cpu)opCodeCB63() byte {
	data := cpu.reg.E
	data = data & (0x1 << 4)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 4,H
func (cpu *Cpu)opCodeCB64() byte {
	data := cpu.reg.H
	data = data & (0x1 << 4)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 4,L
func (cpu *Cpu)opCodeCB65() byte {
	data := cpu.reg.L
	data = data & (0x1 << 4)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 4,(HL)
func (cpu *Cpu)opCodeCB66() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	data = data & (0x1 << 4)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 1
}
        

// BIT 4,A
func (cpu *Cpu)opCodeCB67() byte {
	data := cpu.reg.A
	data = data & (0x1 << 4)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 5,B
func (cpu *Cpu)opCodeCB68() byte {
	data := cpu.reg.B
	data = data & (0x1 << 5)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 5,C
func (cpu *Cpu)opCodeCB69() byte {
	data := cpu.reg.C
	data = data & (0x1 << 5)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 5,D
func (cpu *Cpu)opCodeCB6A() byte {
	data := cpu.reg.D
	data = data & (0x1 << 5)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 5,E
func (cpu *Cpu)opCodeCB6B() byte {
	data := cpu.reg.E
	data = data & (0x1 << 5)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 5,H
func (cpu *Cpu)opCodeCB6C() byte {
	data := cpu.reg.H
	data = data & (0x1 << 5)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 5,L
func (cpu *Cpu)opCodeCB6D() byte {
	data := cpu.reg.L
	data = data & (0x1 << 5)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 5,(HL)
func (cpu *Cpu)opCodeCB6E() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	data = data & (0x1 << 5)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 1
}
        

// BIT 5,A
func (cpu *Cpu)opCodeCB6F() byte {
	data := cpu.reg.A
	data = data & (0x1 << 5)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 6,B
func (cpu *Cpu)opCodeCB70() byte {
	data := cpu.reg.B
	data = data & (0x1 << 6)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 6,C
func (cpu *Cpu)opCodeCB71() byte {
	data := cpu.reg.C
	data = data & (0x1 << 6)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 6,D
func (cpu *Cpu)opCodeCB72() byte {
	data := cpu.reg.D
	data = data & (0x1 << 6)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 6,E
func (cpu *Cpu)opCodeCB73() byte {
	data := cpu.reg.E
	data = data & (0x1 << 6)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 6,H
func (cpu *Cpu)opCodeCB74() byte {
	data := cpu.reg.H
	data = data & (0x1 << 6)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 6,L
func (cpu *Cpu)opCodeCB75() byte {
	data := cpu.reg.L
	data = data & (0x1 << 6)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 6,(HL)
func (cpu *Cpu)opCodeCB76() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	data = data & (0x1 << 6)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 1
}
        

// BIT 6,A
func (cpu *Cpu)opCodeCB77() byte {
	data := cpu.reg.A
	data = data & (0x1 << 6)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 7,B
func (cpu *Cpu)opCodeCB78() byte {
	data := cpu.reg.B
	data = data & (0x1 << 7)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 7,C
func (cpu *Cpu)opCodeCB79() byte {
	data := cpu.reg.C
	data = data & (0x1 << 7)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 7,D
func (cpu *Cpu)opCodeCB7A() byte {
	data := cpu.reg.D
	data = data & (0x1 << 7)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 7,E
func (cpu *Cpu)opCodeCB7B() byte {
	data := cpu.reg.E
	data = data & (0x1 << 7)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 7,H
func (cpu *Cpu)opCodeCB7C() byte {
	data := cpu.reg.H
	data = data & (0x1 << 7)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 7,L
func (cpu *Cpu)opCodeCB7D() byte {
	data := cpu.reg.L
	data = data & (0x1 << 7)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// BIT 7,(HL)
func (cpu *Cpu)opCodeCB7E() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	data = data & (0x1 << 7)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 1
}
        

// BIT 7,A
func (cpu *Cpu)opCodeCB7F() byte {
	data := cpu.reg.A
	data = data & (0x1 << 7)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 0,B
func (cpu *Cpu)opCodeCB80() byte {
	data := cpu.reg.B
	data = data &^ (0x1 << 0)
	cpu.reg.B = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 0,C
func (cpu *Cpu)opCodeCB81() byte {
	data := cpu.reg.C
	data = data &^ (0x1 << 0)
	cpu.reg.C = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 0,D
func (cpu *Cpu)opCodeCB82() byte {
	data := cpu.reg.D
	data = data &^ (0x1 << 0)
	cpu.reg.D = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 0,E
func (cpu *Cpu)opCodeCB83() byte {
	data := cpu.reg.E
	data = data &^ (0x1 << 0)
	cpu.reg.E = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 0,H
func (cpu *Cpu)opCodeCB84() byte {
	data := cpu.reg.H
	data = data &^ (0x1 << 0)
	cpu.reg.H = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 0,L
func (cpu *Cpu)opCodeCB85() byte {
	data := cpu.reg.L
	data = data &^ (0x1 << 0)
	cpu.reg.L = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 0,(HL)
func (cpu *Cpu)opCodeCB86() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	data = data &^ (0x1 << 0)
	cpu.gb.Mem.Write(cpu.reg.getHL(), data)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 0
}
        

// RES 0,A
func (cpu *Cpu)opCodeCB87() byte {
	data := cpu.reg.A
	data = data &^ (0x1 << 0)
	cpu.reg.A = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 1,B
func (cpu *Cpu)opCodeCB88() byte {
	data := cpu.reg.B
	data = data &^ (0x1 << 1)
	cpu.reg.B = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 1,C
func (cpu *Cpu)opCodeCB89() byte {
	data := cpu.reg.C
	data = data &^ (0x1 << 1)
	cpu.reg.C = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 1,D
func (cpu *Cpu)opCodeCB8A() byte {
	data := cpu.reg.D
	data = data &^ (0x1 << 1)
	cpu.reg.D = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 1,E
func (cpu *Cpu)opCodeCB8B() byte {
	data := cpu.reg.E
	data = data &^ (0x1 << 1)
	cpu.reg.E = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 1,H
func (cpu *Cpu)opCodeCB8C() byte {
	data := cpu.reg.H
	data = data &^ (0x1 << 1)
	cpu.reg.H = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 1,L
func (cpu *Cpu)opCodeCB8D() byte {
	data := cpu.reg.L
	data = data &^ (0x1 << 1)
	cpu.reg.L = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 1,(HL)
func (cpu *Cpu)opCodeCB8E() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	data = data &^ (0x1 << 1)
	cpu.gb.Mem.Write(cpu.reg.getHL(), data)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 0
}
        

// RES 1,A
func (cpu *Cpu)opCodeCB8F() byte {
	data := cpu.reg.A
	data = data &^ (0x1 << 1)
	cpu.reg.A = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 2,B
func (cpu *Cpu)opCodeCB90() byte {
	data := cpu.reg.B
	data = data &^ (0x1 << 2)
	cpu.reg.B = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 2,C
func (cpu *Cpu)opCodeCB91() byte {
	data := cpu.reg.C
	data = data &^ (0x1 << 2)
	cpu.reg.C = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 2,D
func (cpu *Cpu)opCodeCB92() byte {
	data := cpu.reg.D
	data = data &^ (0x1 << 2)
	cpu.reg.D = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 2,E
func (cpu *Cpu)opCodeCB93() byte {
	data := cpu.reg.E
	data = data &^ (0x1 << 2)
	cpu.reg.E = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 2,H
func (cpu *Cpu)opCodeCB94() byte {
	data := cpu.reg.H
	data = data &^ (0x1 << 2)
	cpu.reg.H = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 2,L
func (cpu *Cpu)opCodeCB95() byte {
	data := cpu.reg.L
	data = data &^ (0x1 << 2)
	cpu.reg.L = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 2,(HL)
func (cpu *Cpu)opCodeCB96() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	data = data &^ (0x1 << 2)
	cpu.gb.Mem.Write(cpu.reg.getHL(), data)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 0
}
        

// RES 2,A
func (cpu *Cpu)opCodeCB97() byte {
	data := cpu.reg.A
	data = data &^ (0x1 << 2)
	cpu.reg.A = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 3,B
func (cpu *Cpu)opCodeCB98() byte {
	data := cpu.reg.B
	data = data &^ (0x1 << 3)
	cpu.reg.B = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 3,C
func (cpu *Cpu)opCodeCB99() byte {
	data := cpu.reg.C
	data = data &^ (0x1 << 3)
	cpu.reg.C = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 3,D
func (cpu *Cpu)opCodeCB9A() byte {
	data := cpu.reg.D
	data = data &^ (0x1 << 3)
	cpu.reg.D = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 3,E
func (cpu *Cpu)opCodeCB9B() byte {
	data := cpu.reg.E
	data = data &^ (0x1 << 3)
	cpu.reg.E = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 3,H
func (cpu *Cpu)opCodeCB9C() byte {
	data := cpu.reg.H
	data = data &^ (0x1 << 3)
	cpu.reg.H = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 3,L
func (cpu *Cpu)opCodeCB9D() byte {
	data := cpu.reg.L
	data = data &^ (0x1 << 3)
	cpu.reg.L = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 3,(HL)
func (cpu *Cpu)opCodeCB9E() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	data = data &^ (0x1 << 3)
	cpu.gb.Mem.Write(cpu.reg.getHL(), data)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 0
}
        

// RES 3,A
func (cpu *Cpu)opCodeCB9F() byte {
	data := cpu.reg.A
	data = data &^ (0x1 << 3)
	cpu.reg.A = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 4,B
func (cpu *Cpu)opCodeCBA0() byte {
	data := cpu.reg.B
	data = data &^ (0x1 << 4)
	cpu.reg.B = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 4,C
func (cpu *Cpu)opCodeCBA1() byte {
	data := cpu.reg.C
	data = data &^ (0x1 << 4)
	cpu.reg.C = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 4,D
func (cpu *Cpu)opCodeCBA2() byte {
	data := cpu.reg.D
	data = data &^ (0x1 << 4)
	cpu.reg.D = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 4,E
func (cpu *Cpu)opCodeCBA3() byte {
	data := cpu.reg.E
	data = data &^ (0x1 << 4)
	cpu.reg.E = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 4,H
func (cpu *Cpu)opCodeCBA4() byte {
	data := cpu.reg.H
	data = data &^ (0x1 << 4)
	cpu.reg.H = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 4,L
func (cpu *Cpu)opCodeCBA5() byte {
	data := cpu.reg.L
	data = data &^ (0x1 << 4)
	cpu.reg.L = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 4,(HL)
func (cpu *Cpu)opCodeCBA6() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	data = data &^ (0x1 << 4)
	cpu.gb.Mem.Write(cpu.reg.getHL(), data)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 0
}
        

// RES 4,A
func (cpu *Cpu)opCodeCBA7() byte {
	data := cpu.reg.A
	data = data &^ (0x1 << 4)
	cpu.reg.A = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 5,B
func (cpu *Cpu)opCodeCBA8() byte {
	data := cpu.reg.B
	data = data &^ (0x1 << 5)
	cpu.reg.B = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 5,C
func (cpu *Cpu)opCodeCBA9() byte {
	data := cpu.reg.C
	data = data &^ (0x1 << 5)
	cpu.reg.C = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 5,D
func (cpu *Cpu)opCodeCBAA() byte {
	data := cpu.reg.D
	data = data &^ (0x1 << 5)
	cpu.reg.D = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 5,E
func (cpu *Cpu)opCodeCBAB() byte {
	data := cpu.reg.E
	data = data &^ (0x1 << 5)
	cpu.reg.E = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 5,H
func (cpu *Cpu)opCodeCBAC() byte {
	data := cpu.reg.H
	data = data &^ (0x1 << 5)
	cpu.reg.H = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 5,L
func (cpu *Cpu)opCodeCBAD() byte {
	data := cpu.reg.L
	data = data &^ (0x1 << 5)
	cpu.reg.L = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 5,(HL)
func (cpu *Cpu)opCodeCBAE() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	data = data &^ (0x1 << 5)
	cpu.gb.Mem.Write(cpu.reg.getHL(), data)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 0
}
        

// RES 5,A
func (cpu *Cpu)opCodeCBAF() byte {
	data := cpu.reg.A
	data = data &^ (0x1 << 5)
	cpu.reg.A = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 6,B
func (cpu *Cpu)opCodeCBB0() byte {
	data := cpu.reg.B
	data = data &^ (0x1 << 6)
	cpu.reg.B = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 6,C
func (cpu *Cpu)opCodeCBB1() byte {
	data := cpu.reg.C
	data = data &^ (0x1 << 6)
	cpu.reg.C = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 6,D
func (cpu *Cpu)opCodeCBB2() byte {
	data := cpu.reg.D
	data = data &^ (0x1 << 6)
	cpu.reg.D = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 6,E
func (cpu *Cpu)opCodeCBB3() byte {
	data := cpu.reg.E
	data = data &^ (0x1 << 6)
	cpu.reg.E = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 6,H
func (cpu *Cpu)opCodeCBB4() byte {
	data := cpu.reg.H
	data = data &^ (0x1 << 6)
	cpu.reg.H = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 6,L
func (cpu *Cpu)opCodeCBB5() byte {
	data := cpu.reg.L
	data = data &^ (0x1 << 6)
	cpu.reg.L = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 6,(HL)
func (cpu *Cpu)opCodeCBB6() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	data = data &^ (0x1 << 6)
	cpu.gb.Mem.Write(cpu.reg.getHL(), data)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 0
}
        

// RES 6,A
func (cpu *Cpu)opCodeCBB7() byte {
	data := cpu.reg.A
	data = data &^ (0x1 << 6)
	cpu.reg.A = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 7,B
func (cpu *Cpu)opCodeCBB8() byte {
	data := cpu.reg.B
	data = data &^ (0x1 << 7)
	cpu.reg.B = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 7,C
func (cpu *Cpu)opCodeCBB9() byte {
	data := cpu.reg.C
	data = data &^ (0x1 << 7)
	cpu.reg.C = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 7,D
func (cpu *Cpu)opCodeCBBA() byte {
	data := cpu.reg.D
	data = data &^ (0x1 << 7)
	cpu.reg.D = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 7,E
func (cpu *Cpu)opCodeCBBB() byte {
	data := cpu.reg.E
	data = data &^ (0x1 << 7)
	cpu.reg.E = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 7,H
func (cpu *Cpu)opCodeCBBC() byte {
	data := cpu.reg.H
	data = data &^ (0x1 << 7)
	cpu.reg.H = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 7,L
func (cpu *Cpu)opCodeCBBD() byte {
	data := cpu.reg.L
	data = data &^ (0x1 << 7)
	cpu.reg.L = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// RES 7,(HL)
func (cpu *Cpu)opCodeCBBE() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	data = data &^ (0x1 << 7)
	cpu.gb.Mem.Write(cpu.reg.getHL(), data)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 0
}
        

// RES 7,A
func (cpu *Cpu)opCodeCBBF() byte {
	data := cpu.reg.A
	data = data &^ (0x1 << 7)
	cpu.reg.A = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 0,B
func (cpu *Cpu)opCodeCBC0() byte {
	data := cpu.reg.B
	data = data | (0x1 << 0)
	cpu.reg.B = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 0,C
func (cpu *Cpu)opCodeCBC1() byte {
	data := cpu.reg.C
	data = data | (0x1 << 0)
	cpu.reg.C = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 0,D
func (cpu *Cpu)opCodeCBC2() byte {
	data := cpu.reg.D
	data = data | (0x1 << 0)
	cpu.reg.D = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 0,E
func (cpu *Cpu)opCodeCBC3() byte {
	data := cpu.reg.E
	data = data | (0x1 << 0)
	cpu.reg.E = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 0,H
func (cpu *Cpu)opCodeCBC4() byte {
	data := cpu.reg.H
	data = data | (0x1 << 0)
	cpu.reg.H = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 0,L
func (cpu *Cpu)opCodeCBC5() byte {
	data := cpu.reg.L
	data = data | (0x1 << 0)
	cpu.reg.L = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 0,(HL)
func (cpu *Cpu)opCodeCBC6() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	data = data | (0x1 << 0)
	cpu.gb.Mem.Write(cpu.reg.getHL(), data)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 0
}
        

// SET 0,A
func (cpu *Cpu)opCodeCBC7() byte {
	data := cpu.reg.A
	data = data | (0x1 << 0)
	cpu.reg.A = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 1,B
func (cpu *Cpu)opCodeCBC8() byte {
	data := cpu.reg.B
	data = data | (0x1 << 1)
	cpu.reg.B = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 1,C
func (cpu *Cpu)opCodeCBC9() byte {
	data := cpu.reg.C
	data = data | (0x1 << 1)
	cpu.reg.C = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 1,D
func (cpu *Cpu)opCodeCBCA() byte {
	data := cpu.reg.D
	data = data | (0x1 << 1)
	cpu.reg.D = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 1,E
func (cpu *Cpu)opCodeCBCB() byte {
	data := cpu.reg.E
	data = data | (0x1 << 1)
	cpu.reg.E = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 1,H
func (cpu *Cpu)opCodeCBCC() byte {
	data := cpu.reg.H
	data = data | (0x1 << 1)
	cpu.reg.H = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 1,L
func (cpu *Cpu)opCodeCBCD() byte {
	data := cpu.reg.L
	data = data | (0x1 << 1)
	cpu.reg.L = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 1,(HL)
func (cpu *Cpu)opCodeCBCE() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	data = data | (0x1 << 1)
	cpu.gb.Mem.Write(cpu.reg.getHL(), data)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 0
}
        

// SET 1,A
func (cpu *Cpu)opCodeCBCF() byte {
	data := cpu.reg.A
	data = data | (0x1 << 1)
	cpu.reg.A = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 2,B
func (cpu *Cpu)opCodeCBD0() byte {
	data := cpu.reg.B
	data = data | (0x1 << 2)
	cpu.reg.B = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 2,C
func (cpu *Cpu)opCodeCBD1() byte {
	data := cpu.reg.C
	data = data | (0x1 << 2)
	cpu.reg.C = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 2,D
func (cpu *Cpu)opCodeCBD2() byte {
	data := cpu.reg.D
	data = data | (0x1 << 2)
	cpu.reg.D = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 2,E
func (cpu *Cpu)opCodeCBD3() byte {
	data := cpu.reg.E
	data = data | (0x1 << 2)
	cpu.reg.E = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 2,H
func (cpu *Cpu)opCodeCBD4() byte {
	data := cpu.reg.H
	data = data | (0x1 << 2)
	cpu.reg.H = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 2,L
func (cpu *Cpu)opCodeCBD5() byte {
	data := cpu.reg.L
	data = data | (0x1 << 2)
	cpu.reg.L = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 2,(HL)
func (cpu *Cpu)opCodeCBD6() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	data = data | (0x1 << 2)
	cpu.gb.Mem.Write(cpu.reg.getHL(), data)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 0
}
        

// SET 2,A
func (cpu *Cpu)opCodeCBD7() byte {
	data := cpu.reg.A
	data = data | (0x1 << 2)
	cpu.reg.A = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 3,B
func (cpu *Cpu)opCodeCBD8() byte {
	data := cpu.reg.B
	data = data | (0x1 << 3)
	cpu.reg.B = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 3,C
func (cpu *Cpu)opCodeCBD9() byte {
	data := cpu.reg.C
	data = data | (0x1 << 3)
	cpu.reg.C = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 3,D
func (cpu *Cpu)opCodeCBDA() byte {
	data := cpu.reg.D
	data = data | (0x1 << 3)
	cpu.reg.D = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 3,E
func (cpu *Cpu)opCodeCBDB() byte {
	data := cpu.reg.E
	data = data | (0x1 << 3)
	cpu.reg.E = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 3,H
func (cpu *Cpu)opCodeCBDC() byte {
	data := cpu.reg.H
	data = data | (0x1 << 3)
	cpu.reg.H = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 3,L
func (cpu *Cpu)opCodeCBDD() byte {
	data := cpu.reg.L
	data = data | (0x1 << 3)
	cpu.reg.L = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 3,(HL)
func (cpu *Cpu)opCodeCBDE() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	data = data | (0x1 << 3)
	cpu.gb.Mem.Write(cpu.reg.getHL(), data)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 0
}
        

// SET 3,A
func (cpu *Cpu)opCodeCBDF() byte {
	data := cpu.reg.A
	data = data | (0x1 << 3)
	cpu.reg.A = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 4,B
func (cpu *Cpu)opCodeCBE0() byte {
	data := cpu.reg.B
	data = data | (0x1 << 4)
	cpu.reg.B = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 4,C
func (cpu *Cpu)opCodeCBE1() byte {
	data := cpu.reg.C
	data = data | (0x1 << 4)
	cpu.reg.C = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 4,D
func (cpu *Cpu)opCodeCBE2() byte {
	data := cpu.reg.D
	data = data | (0x1 << 4)
	cpu.reg.D = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 4,E
func (cpu *Cpu)opCodeCBE3() byte {
	data := cpu.reg.E
	data = data | (0x1 << 4)
	cpu.reg.E = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 4,H
func (cpu *Cpu)opCodeCBE4() byte {
	data := cpu.reg.H
	data = data | (0x1 << 4)
	cpu.reg.H = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 4,L
func (cpu *Cpu)opCodeCBE5() byte {
	data := cpu.reg.L
	data = data | (0x1 << 4)
	cpu.reg.L = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 4,(HL)
func (cpu *Cpu)opCodeCBE6() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	data = data | (0x1 << 4)
	cpu.gb.Mem.Write(cpu.reg.getHL(), data)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 0
}
        

// SET 4,A
func (cpu *Cpu)opCodeCBE7() byte {
	data := cpu.reg.A
	data = data | (0x1 << 4)
	cpu.reg.A = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 5,B
func (cpu *Cpu)opCodeCBE8() byte {
	data := cpu.reg.B
	data = data | (0x1 << 5)
	cpu.reg.B = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 5,C
func (cpu *Cpu)opCodeCBE9() byte {
	data := cpu.reg.C
	data = data | (0x1 << 5)
	cpu.reg.C = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 5,D
func (cpu *Cpu)opCodeCBEA() byte {
	data := cpu.reg.D
	data = data | (0x1 << 5)
	cpu.reg.D = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 5,E
func (cpu *Cpu)opCodeCBEB() byte {
	data := cpu.reg.E
	data = data | (0x1 << 5)
	cpu.reg.E = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 5,H
func (cpu *Cpu)opCodeCBEC() byte {
	data := cpu.reg.H
	data = data | (0x1 << 5)
	cpu.reg.H = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 5,L
func (cpu *Cpu)opCodeCBED() byte {
	data := cpu.reg.L
	data = data | (0x1 << 5)
	cpu.reg.L = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 5,(HL)
func (cpu *Cpu)opCodeCBEE() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	data = data | (0x1 << 5)
	cpu.gb.Mem.Write(cpu.reg.getHL(), data)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 0
}
        

// SET 5,A
func (cpu *Cpu)opCodeCBEF() byte {
	data := cpu.reg.A
	data = data | (0x1 << 5)
	cpu.reg.A = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 6,B
func (cpu *Cpu)opCodeCBF0() byte {
	data := cpu.reg.B
	data = data | (0x1 << 6)
	cpu.reg.B = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 6,C
func (cpu *Cpu)opCodeCBF1() byte {
	data := cpu.reg.C
	data = data | (0x1 << 6)
	cpu.reg.C = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 6,D
func (cpu *Cpu)opCodeCBF2() byte {
	data := cpu.reg.D
	data = data | (0x1 << 6)
	cpu.reg.D = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 6,E
func (cpu *Cpu)opCodeCBF3() byte {
	data := cpu.reg.E
	data = data | (0x1 << 6)
	cpu.reg.E = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 6,H
func (cpu *Cpu)opCodeCBF4() byte {
	data := cpu.reg.H
	data = data | (0x1 << 6)
	cpu.reg.H = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 6,L
func (cpu *Cpu)opCodeCBF5() byte {
	data := cpu.reg.L
	data = data | (0x1 << 6)
	cpu.reg.L = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 6,(HL)
func (cpu *Cpu)opCodeCBF6() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	data = data | (0x1 << 6)
	cpu.gb.Mem.Write(cpu.reg.getHL(), data)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 0
}
        

// SET 6,A
func (cpu *Cpu)opCodeCBF7() byte {
	data := cpu.reg.A
	data = data | (0x1 << 6)
	cpu.reg.A = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 7,B
func (cpu *Cpu)opCodeCBF8() byte {
	data := cpu.reg.B
	data = data | (0x1 << 7)
	cpu.reg.B = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 7,C
func (cpu *Cpu)opCodeCBF9() byte {
	data := cpu.reg.C
	data = data | (0x1 << 7)
	cpu.reg.C = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 7,D
func (cpu *Cpu)opCodeCBFA() byte {
	data := cpu.reg.D
	data = data | (0x1 << 7)
	cpu.reg.D = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 7,E
func (cpu *Cpu)opCodeCBFB() byte {
	data := cpu.reg.E
	data = data | (0x1 << 7)
	cpu.reg.E = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 7,H
func (cpu *Cpu)opCodeCBFC() byte {
	data := cpu.reg.H
	data = data | (0x1 << 7)
	cpu.reg.H = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 7,L
func (cpu *Cpu)opCodeCBFD() byte {
	data := cpu.reg.L
	data = data | (0x1 << 7)
	cpu.reg.L = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        

// SET 7,(HL)
func (cpu *Cpu)opCodeCBFE() byte {
	data := cpu.gb.Mem.Read(cpu.reg.getHL())
	data = data | (0x1 << 7)
	cpu.gb.Mem.Write(cpu.reg.getHL(), data)
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 0
}
        

// SET 7,A
func (cpu *Cpu)opCodeCBFF() byte {
	data := cpu.reg.A
	data = data | (0x1 << 7)
	cpu.reg.A = data
	if data == 0 {
		cpu.setFlagZ()
	}
	cpu.resetFlagN()
	cpu.setFlagH()
	return 2
}
        