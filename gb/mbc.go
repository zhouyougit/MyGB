package gb

import "log"

var (
	ROMBankSize uint16 = 0x4000
	RAMBankSize uint16 = 0x2000
)


type MemoryBankController interface {
	Read(uint16) byte
	Write(uint16, byte)
}

type MBCRom struct {
	rom []byte
	ram []byte
}

func (mbc *MBCRom) Read(address uint16) byte {
	if address < 0x8000 {
		return mbc.rom[address]
	} else if (address >= 0xA000) && (address < 0xC000) {
		return mbc.ram[address - 0xA000]
	}
	log.Fatalf("cartridge memory access out of range, address = 0x%X", address)
	return 0
}

func (mbc *MBCRom) Write(address uint16, data byte) {
	if (address >= 0xA000) && (address < 0xC000) {
		mbc.ram[address - 0xA000] = data
	}
}

type MBC1 struct {
	rom []byte
	currentRomBank byte
	ram []byte
	currentRamBank byte
	enableRam bool
	romBankingMode bool
}

func (mbc *MBC1) Read(address uint16) byte {
	if address < 0x4000 {
		//0000-3FFF - ROM Bank 00 (Read Only)
		return mbc.rom[address]
	} else if address < 0x8000 {
		//4000-7FFF - ROM Bank 01-7F (Read Only)
		return mbc.rom[uint32(address - 0x4000) + uint32(mbc.currentRomBank) * 0x4000]
	} else if (address >= 0xA000) && (address < 0xC000) {
		//A000-BFFF - RAM Bank 00-03, if any (Read/Write)
		if !mbc.enableRam {
			return byte(0)
		}
		return mbc.ram[uint32(address - 0xA000) + uint32(mbc.currentRamBank) * 0x2000]
	}
	log.Fatalf("cartridge memory access out of range, address = 0x%X", address)
	return 0
}

func (mbc *MBC1) Write(address uint16, data byte) {
	if address <= 0x1FFF {
		//0000-1FFF - RAM Enable (Write Only)
		val := data & 0x0F
		mbc.enableRam = val == 0x0A
	} else if address < 0x4000 {
		//2000-3FFF - ROM Bank Number (Write Only)
		val := data & 0x1F
		if val == 0 {
			val++
		}
		mbc.currentRomBank &= 0xE0
		mbc.currentRomBank |= val
	} else if address < 0x6000 {
		//4000-5FFF - RAM Bank Number - or - Upper Bits of ROM Bank Number (Write Only)
		if mbc.romBankingMode {
			val := data & 0xE0
			mbc.currentRomBank &= 0x1F
			mbc.currentRomBank |= val
		} else {
			mbc.currentRamBank = data & 0x3
		}
	} else if address < 0x8000 {
		val := data & 0x01
		mbc.romBankingMode = val == 0
	} else if (address >= 0xA000) && (address < 0xC000) {
		if mbc.enableRam {
			mbc.ram[uint32(address - 0xA000) + uint32(mbc.currentRamBank) * 0x2000] = data
		}
	}
}

type MBC2 struct {
	rom []byte
}

type MBC3 struct {
	rom []byte
}

type MBC4 struct {
	rom []byte
}

type MBC5 struct {
	rom []byte
}