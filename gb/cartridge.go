package gb

import (
	"fmt"
	"github.com/prometheus/common/log"
	"io/ioutil"
	"os"
)

type Cartridge struct {
	romPath string
	Header  CartridgeHeader
	MBC     MemoryBankController
}

type CartridgeHeader struct {
	raw []byte
	NintendoLogo [0x30]byte
	Title string
	ManufacturerCode [4]byte
	CGBFlag byte
	NewLicenseeCode [2]byte
	SGBFlag byte
	CartridgeType byte
	ROMSize byte
	RAMSize byte
	DestinationCode byte
	OldLicenseeCode byte
	MaskROMVersionNumber byte
	HeaderChecksum byte
	GlobalChecksum [2]byte
}

var cgbFlagDescMap = map[byte]string {
	0x80: "Game supports CGB functions, but works on old gameboys also.",
	0xC0: "Game works on CGB only (physically the same as 80h).",
}

var sgbFlagDescMap = map[byte]string {
	0x00: "No SGB functions (Normal Gameboy or CGB only game)",
	0x03: "Game supports SGB functions",
}

var cartridgeTypeMap = map[byte]string {
	0x00: "ROM ONLY",
	0x01: "MBC1",
	0x02: "MBC1+RAM",
	0x03: "MBC1+RAM+BATTERY",
	0x05: "MBC2",
	0x06: "MBC2+BATTERY",
	0x08: "ROM+RAM",
	0x09: "ROM+RAM+BATTERY",
	0x0B: "MMM01",
	0x0C: "MMM01+RAM",
	0x0D: "MMM01+RAM+BATTERY",
	0x0F: "MBC3+TIMER+BATTERY",
	0x10: "MBC3+TIMER+RAM+BATTERY",
	0x11: "MBC3",
	0x12: "MBC3+RAM",
	0x13: "MBC3+RAM+BATTERY",
	0x15: "MBC4",
	0x16: "MBC4+RAM",
	0x17: "MBC4+RAM+BATTERY",
	0x19: "MBC5",
	0x1A: "MBC5+RAM",
	0x1B: "MBC5+RAM+BATTERY",
	0x1C: "MBC5+RUMBLE",
	0x1D: "MBC5+RUMBLE+RAM",
	0x1E: "MBC5+RUMBLE+RAM+BATTERY",
	0xFC: "POCKET CAMERA",
	0xFD: "BANDAI TAMA5",
	0xFE: "HuC3",
	0xFF: "HuC1+RAM+BATTERY",
}

var romSizeNameMap = map[byte]string {
	0x00: "32KByte (no ROM banking)",
	0x01: "64KByte (4 banks)",
	0x02: "128KByte (8 banks)",
	0x03: "256KByte (16 banks)",
	0x04: "512KByte (32 banks)",
	0x05: "1MByte (64 banks) - only 63 banks used by MBC1",
	0x06: "2MByte (128 banks) - only 125 banks used by MBC1",
	0x07: "4MByte (256 banks)",
	0x52: "1.1MByte (72 banks)",
	0x53: "1.2MByte (80 banks)",
	0x54: "1.5MByte (96 banks)",
}

var romSizeBankMap = map[byte]uint32 {
	0x00: 2,
	0x01: 4,
	0x02: 8,
	0x03: 16,
	0x04: 32,
	0x05: 64,
	0x06: 128,
	0x07: 256,
	0x52: 72,
	0x53: 80,
	0x54: 96,
}

var ramSizeNameMap = map[byte]string {
	0x00: "None",
	0x01: "2 KBytes",
	0x02: "8 KBytes",
	0x03: "32 KBytes (4 banks of 8KBytes each",
}

var ramSizeBankMap = map[byte]uint32 {
	0x00: 0x00,
	0x01: 0x01,
	0x02: 0x01,
	0x03: 0x04,
}

var destinationCodeMap = map[byte]string {
	0x00: "Japanese",
	0x01: "Non-Japanese",
}

func (c *Cartridge) LoadRom() error {
	rom, err := ioutil.ReadFile(c.romPath)
	if err != nil {
		return err
	}
	ram, err := ioutil.ReadFile(c.romPath + ".dat")
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	c.Header.loadHeaders(rom[:0x014F+1])
	c.Header.dumpHeaders()

	if ram == nil {
		ram = make([]byte, uint32(RAMBankSize) * ramSizeBankMap[c.Header.RAMSize])
	}
	switch c.Header.CartridgeType {
	case 0x00, 0x08, 0x09:
		c.MBC = &MBCRom{
			rom: rom,
			ram: ram,
		}
		log.Debugf("Use MBCRom\n")
	case 0x01, 0x02, 0x03:
		c.MBC = &MBC1{
			rom: rom,
			ram: ram,
		}
		log.Debugf("Use MBC1\n")
	}
	return nil
}

func (c *Cartridge) isCGB() bool {
	return c.Header.CGBFlag & 0x80 > 0
}

func (c *Cartridge) isSGB() bool {
	return c.Header.SGBFlag == 0x03
}

func (h *CartridgeHeader) loadHeaders(rawHeader []byte) {
	h.raw = rawHeader
	//0100-0103 - Entry Point

	//0104-0133 - Nintendo Logo
	copy(h.NintendoLogo[:], rawHeader[0x0104:0x0133+1])

	//0134-0143 - Title
	h.Title = string(rawHeader[0x0134:0x0143 + 1])

	//013F-0142 - Manufacturer Code
	copy(h.ManufacturerCode[:], rawHeader[0x013F:0x0142 + 1])

	//0143 - CGB Flag
	h.CGBFlag = rawHeader[0x0143]

	if h.CGBFlag & 0x80 > 0 {
		h.Title = h.Title[:15]
	}

	//0144-0145 - New Licensee Code
	copy(h.NewLicenseeCode[:], rawHeader[0x0144:0x0145 + 1])

	//0146 - SGB Flag
	h.SGBFlag = rawHeader[0x0146]

	//0147 - Cartridge Type
	h.CartridgeType = rawHeader[0x0147]

	//0148 - ROM Size
	h.ROMSize = rawHeader[0x0148]

	//0149 - RAM Size
	h.RAMSize = rawHeader[0x0149]

	//014A - Destination Code
	h.DestinationCode = rawHeader[0x014A]

	//014B - Old Licensee Code
	h.OldLicenseeCode = rawHeader[0x014B]

	//014C - Mask ROM Version number
	h.MaskROMVersionNumber = rawHeader[0x014C]

	//014D - Header Checksum
	h.HeaderChecksum = rawHeader[0x014D]

	//014E-014F - Global Checksum
	copy(h.GlobalChecksum[:], rawHeader[0x014E:0x014F + 1])
}

func (h *CartridgeHeader) verifyChecksum() bool {
	var checksum byte = 0
	for i := 0x0134; i <= 0x014C; i++ {
		checksum = checksum - h.raw[i] - 1
	}
	return checksum == h.HeaderChecksum
}

func (h *CartridgeHeader) dumpHeaders() {
	fmt.Printf("====== Cartridge Headers ======\n")
	fmt.Printf("Title:\t%s\n", h.Title)
	fmt.Printf("Manufacturer Code:\t0x%X\n", h.ManufacturerCode)
	fmt.Printf("CGB Flag:\t0x%X (%v)\n", h.CGBFlag, h.CGBFlag & 0x80 > 0)
	fmt.Printf("New Licensee Code:\t0x%X\n", h.NewLicenseeCode)
	fmt.Printf("SGB Flag:\t0x%X (%v)\n", h.SGBFlag, h.SGBFlag == 0x03)
	fmt.Printf("Cartridge Type:\t0x%X (%s)\n", h.CartridgeType, cartridgeTypeMap[h.CartridgeType])
	fmt.Printf("ROM Size:\t0x%X (%s)\n", h.ROMSize, romSizeNameMap[h.ROMSize])
	fmt.Printf("RAM Size:\t0x%X (%s)\n", h.RAMSize, ramSizeNameMap[h.RAMSize])
	fmt.Printf("Destination Code:\t0x%X (%s)\n", h.DestinationCode, destinationCodeMap[h.DestinationCode])
	fmt.Printf("Old Licensee Code:\t0x%X\n", h.OldLicenseeCode)
	fmt.Printf("Mask ROM Version number:\t0x%X\n", h.MaskROMVersionNumber)

	verifyResult := "incorrect"
	if h.verifyChecksum() {
		verifyResult = "correct"
	}
	fmt.Printf("Header Checksum:\t0x%X (%s)\n", h.HeaderChecksum, verifyResult)
	fmt.Printf("Global Checksum:\t0x%X\n", h.GlobalChecksum)
	fmt.Printf("========== End Headers =========\n\n")
}