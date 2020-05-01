package gb

const (
	LCD_CTRL_ADDR uint16 = 0xFF40
	LCD_STAT_ADDR uint16 = 0xFF41

	//LCD Position and Scrolling
	LCD_SCY_ADDR  uint16 = 0xFF42
	LCD_SCX_ADDR  uint16 = 0xFF43
	LCD_LY_ADDR   uint16 = 0xFF44
	LCD_LYC_ADDR  uint16 = 0xFF45
	LCD_WY_ADDR   uint16 = 0xFF4A
	LCD_WX_ADDR   uint16 = 0xFF4B

	//LCD Monochrome Palettes
	LCD_BGP_ADDR  uint16 = 0xFF47
	LCD_OBP0_ADDR uint16 = 0xFF48
	LCD_OBP1_ADDR uint16 = 0xFF49

	//LCD Color Palettes (CGB only)
	LCD_BGPI_ADDR uint16 = 0xFF68
	LCD_BGPD_ADDR uint16 = 0xFF69
	LCD_OBPI_ADDR uint16 = 0xFF6A
	LCD_OBPD_ADDR uint16 = 0xFF6B

	//LCD VRAM Bank (CGB only)
	LCD_VBK_ADDR  uint16 = 0xFF4F

	//LCD OAM DMA Transfers
	LCD_DMA_ADDR  uint16 = 0xFF46

	//LCD VRAM DMA Transfers (CGB only)
	LCD_HDMA1_ADDR uint16 = 0xFF51
	LCD_HDMA2_ADDR uint16 = 0xFF52
	LCD_HDMA3_ADDR uint16 = 0xFF53
	LCD_HDMA4_ADDR uint16 = 0xFF54
	LCD_HDMA5_ADDR uint16 = 0xFF55

	/**
	Mode 0 is present between 201-207 clks, 2 about 77-83 clks, and 3 about 169-175 clks.
	A complete cycle through these states takes 456 clks.
	VBlank lasts 4560 clks. A complete screen refresh occurs every 70224 clks.
	 */
	LCD_FULL_CYCLES int = 456
)

type Lcd struct {
	gb *GameBoy

	scanLineCounter int
}

func (l *Lcd) Init(gb *GameBoy) error {
	l.gb = gb
	return nil
}

func (l *Lcd) update(cycles int) {
	if !l.isLCDDisplayEnabled() {
		l.resetForLCDDisable()
		return
	}

	l.scanLineCounter += cycles
	if l.scanLineCounter < LCD_FULL_CYCLES {
		l.updateLCDStatusRegister()
		return
	}
	l.scanLineCounter -= LCD_FULL_CYCLES

	l.updateLCDStatusRegister()

	ly := l.gb.Mem.raw[LCD_LY_ADDR]
	ly++
	if ly > 153 {
		ly = 0
	}
	// Mem.Write() will reset the counter.
	l.gb.Mem.raw[LCD_LY_ADDR] = ly

	if ly == 144 {
		l.gb.Cpu.RequestInterrupt(IntVBlank)
	} else if ly < 144 {
		l.drawScanLine()
	}
}

/**
FF41 - STAT - LCDC Status (R/W)
  Bit 6 - LYC=LY Coincidence Interrupt (1=Enable) (Read/Write)
  Bit 5 - Mode 2 OAM Interrupt         (1=Enable) (Read/Write)
  Bit 4 - Mode 1 V-Blank Interrupt     (1=Enable) (Read/Write)
  Bit 3 - Mode 0 H-Blank Interrupt     (1=Enable) (Read/Write)
  Bit 2 - Coincidence Flag  (0:LYC<>LY, 1:LYC=LY) (Read Only)
  Bit 1-0 - Mode Flag       (Mode 0-3, see below) (Read Only)
            0: During H-Blank
            1: During V-Blank
            2: During Searching OAM-RAM
            3: During Transfering Data to LCD Driver
 */
func (l *Lcd) updateLCDStatusRegister() {
	stat := l.gb.Mem.Read(LCD_STAT_ADDR)

	ly := l.gb.Mem.Read(LCD_LY_ADDR)
	currMode := stat & 0b11
	var newMode byte = 0
	var needInt = false

	if ly >= 144 {
		// Mode V-Blank
		newMode = 1
		needInt = stat & 0b10000 == 0b10000
	} else {
		mode2end := 80
		mode3end := 80 + 172

		if l.scanLineCounter < mode2end {
			newMode = 2
			needInt = stat & 0b100000 == 0b100000
		} else if l.scanLineCounter < mode3end {
			newMode = 3
		} else {
			newMode = 0
			needInt = stat & 0b1000 == 0b1000
		}
	}
	if currMode != newMode {
		stat &= 0b11111100
		stat |= newMode
		if needInt {
			l.gb.Cpu.RequestInterrupt(IntLcdStat)
		}
	}

	lyc := l.gb.Mem.Read(LCD_LYC_ADDR)
	if ly == lyc {
		stat |= 0b0100
		if stat & 0b01000000 == 0b01000000 {
			l.gb.Cpu.RequestInterrupt(IntLcdStat)
		}
	} else {
		stat = stat &^ 0b0100
	}

	l.gb.Mem.Write(LCD_STAT_ADDR, stat)
}

func (l *Lcd) resetForLCDDisable() {
	stat := l.gb.Mem.Read(LCD_STAT_ADDR)
	l.scanLineCounter = 0
	l.gb.Mem.raw[LCD_LY_ADDR] = 0x00
	stat &= 0b11111100
	stat |= 0b01
	l.gb.Mem.Write(LCD_STAT_ADDR, stat)
}

func (l *Lcd) isLCDDisplayEnabled() bool {
	ctrl := l.gb.Mem.Read(LCD_CTRL_ADDR)
	return ctrl & 0x80 == 0x80
}

func (l *Lcd) drawScanLine() {

}