package gb

import "fmt"

/**
General Memory Map
  0000-3FFF   16KB ROM Bank 00     (in cartridge, fixed at bank 00)
  4000-7FFF   16KB ROM Bank 01..NN (in cartridge, switchable bank number)
  8000-9FFF   8KB Video RAM (VRAM) (switchable bank 0-1 in CGB Mode)
  A000-BFFF   8KB External RAM     (in cartridge, switchable bank, if any)
  C000-CFFF   4KB Work RAM Bank 0 (WRAM)
  D000-DFFF   4KB Work RAM Bank 1 (WRAM)  (switchable bank 1-7 in CGB Mode)
  E000-FDFF   Same as C000-DDFF (ECHO)    (typically not used)
  FE00-FE9F   Sprite Attribute Table (OAM)
  FEA0-FEFF   Not Usable
  FF00-FF7F   I/O Ports
  FF80-FFFE   High RAM (HRAM)
  FFFF        Interrupt Enable Register
 */

type Memory struct {
	gb *GameBoy
	raw [0x10000]byte
	cartridge *Cartridge
}

/**
[$FF05] = $00   ; TIMA
[$FF06] = $00   ; TMA
[$FF07] = $00   ; TAC
[$FF10] = $80   ; NR10
[$FF11] = $BF   ; NR11
[$FF12] = $F3   ; NR12
[$FF14] = $BF   ; NR14
[$FF16] = $3F   ; NR21
[$FF17] = $00   ; NR22
[$FF19] = $BF   ; NR24
[$FF1A] = $7F   ; NR30
[$FF1B] = $FF   ; NR31
[$FF1C] = $9F   ; NR32
[$FF1E] = $BF   ; NR33
[$FF20] = $FF   ; NR41
[$FF21] = $00   ; NR42
[$FF22] = $00   ; NR43
[$FF23] = $BF   ; NR30
[$FF24] = $77   ; NR50
[$FF25] = $F3   ; NR51
[$FF26] = $F1-GB, $F0-SGB ; NR52
[$FF40] = $91   ; LCDC
[$FF42] = $00   ; SCY
[$FF43] = $00   ; SCX
[$FF45] = $00   ; LYC
[$FF47] = $FC   ; BGP
[$FF48] = $FF   ; OBP0
[$FF49] = $FF   ; OBP1
[$FF4A] = $00   ; WY
[$FF4B] = $00   ; WX
[$FFFF] = $00   ; IE
 */
func (mem *Memory) Init(gb *GameBoy) {
	mem.gb = gb
	mem.cartridge = &gb.Cartridge

	mem.raw[0xFF05] = 0x00
	mem.raw[0xFF06] = 0x00
	mem.raw[0xFF07] = 0x00
	mem.raw[0xFF10] = 0x80
	mem.raw[0xFF11] = 0xBF
	mem.raw[0xFF12] = 0xF3
	mem.raw[0xFF14] = 0xBF
	mem.raw[0xFF16] = 0x3F
	mem.raw[0xFF17] = 0x00
	mem.raw[0xFF19] = 0xBF
	mem.raw[0xFF1A] = 0x7F
	mem.raw[0xFF1B] = 0xFF
	mem.raw[0xFF1C] = 0x9F
	mem.raw[0xFF1E] = 0xBF
	mem.raw[0xFF20] = 0xFF
	mem.raw[0xFF21] = 0x00
	mem.raw[0xFF22] = 0x00
	mem.raw[0xFF23] = 0xBF
	mem.raw[0xFF24] = 0x77
	mem.raw[0xFF25] = 0xF3
	if mem.cartridge.isSGB() {
		mem.raw[0xFF26] = 0xF0
	} else {
		mem.raw[0xFF26] = 0xF1
	}
	mem.raw[0xFF40] = 0x91
	mem.raw[0xFF42] = 0x00
	mem.raw[0xFF43] = 0x00
	mem.raw[0xFF45] = 0x00
	mem.raw[0xFF47] = 0xFC
	mem.raw[0xFF48] = 0xFF
	mem.raw[0xFF49] = 0xFF
	mem.raw[0xFF4A] = 0x00
	mem.raw[0xFF4B] = 0x00
	mem.raw[0xFFFF] = 0x00
}

func (mem *Memory) Read(address uint16) byte {
	if address < 0x8000 {
		return mem.cartridge.MBC.Read(address)
	} else if (address >= 0xA000) && (address < 0xC000) {
		return mem.cartridge.MBC.Read(address)
	} else if address == 0xFF00 {
		return mem.gb.Joypad.readStatus()
	}
	return mem.raw[address]
}

func (mem *Memory) ReadUint16(address uint16) uint16 {
	byteLo := uint16(mem.Read(address))
	byteHi := uint16(mem.Read(address+1))
	return byteHi << 8 | byteLo
}

func (mem *Memory) Write(address uint16, data byte) {
	if address < 0x8000 {
		mem.cartridge.MBC.Write(address, data)
	} else if (address >= 0xA000) && (address < 0xC000) {
		mem.cartridge.MBC.Write(address, data)
	} else if (address >= 0xE000) && (address < 0xFE00) {
		// Echo of 8kB Internal RAM
		mem.raw[address] = data
		mem.raw[address - 0x2000] = data
	} else if address == 0xFF00 {
		mem.gb.Joypad.writeStatus(data)
	} else if address == 0xFF04 {
		//FF04 - DIV - Divider Register (R/W)
		mem.raw[address] = 0x00
	} else if address == 0xFF44 {
		//FF44 - LY - LCDC Y-Coordinate (R)
		mem.raw[address] = 0x00
	} else if address == 0xFF46 {
		//FF46 - DMA - DMA Transfer and Start Address (W)
		mem.doDMA(data)
	}
	mem.raw[address] = data
}

func (mem *Memory) WriteUint16(address uint16, data uint16) {
	byteHi := byte(data >> 8)
	byteLo := byte(data & 0xFF)
	mem.Write(address, byteLo)
	mem.Write(address+1, byteHi)
}

func (mem *Memory) Dump(begin, end uint16) {
	fmt.Printf("============ Dump Mem from 0x%04X to 0x%04X =============\n" +
		"addr\t00 01 02 03 04 05 06 07 08 09 0A 0B 0C 0D 0E 0F\n" +
		"-------------------------------------------------------", begin, end)
	for i := begin; i <= end; i++ {
		if (i - begin) % 16 == 0 {
			fmt.Printf("\n0x%04X\t", i)
		}
		fmt.Printf("%02X ", mem.Read(i))
	}
	fmt.Printf("\n\n")
}

func (mem *Memory) doDMA(data byte) {
	saddr := uint16(data) << 8
	for i := uint16(0); i < 0xA0; i++ {
		mem.Write(0xFE00 + i, mem.Read(saddr + i))
	}
}