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

var (
	/**
	FF40 - LCDC - LCD Control (R/W)
	  Bit 7 - LCD Display Enable             (0=Off, 1=On)
	  Bit 6 - Window Tile Map Display Select (0=9800-9BFF, 1=9C00-9FFF)
	  Bit 5 - Window Display Enable          (0=Off, 1=On)
	  Bit 4 - BG & Window Tile Data Select   (0=8800-97FF, 1=8000-8FFF)
	  Bit 3 - BG Tile Map Display Select     (0=9800-9BFF, 1=9C00-9FFF)
	  Bit 2 - OBJ (Sprite) Size              (0=8x8, 1=8x16)
	  Bit 1 - OBJ (Sprite) Display Enable    (0=Off, 1=On)
	  Bit 0 - BG Display (for CGB see below) (0=Off, 1=On)
	 */
	LcdCtrlLCDEnable byte = 0x1 << 7
	LcdCtrlWinTileMapSelect byte = 0x1 << 6
	LcdCtrlWinEnable byte = 0x1 << 5
	LcdCtrlTileDataSelect byte = 0x1 << 4
	LcdCtrlBgTileMapSelect byte = 0x1 << 3
	LcdCtrlObjSize byte = 0x1 << 2
	LcdCtrlObjEnable byte = 0x1 << 1
	LcdCtrlBgEnable byte = 0x1 << 0
)

type Lcd struct {
	gb *GameBoy

	scanLineCounter int

	screen [160][144][3] uint8
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
	ctrl := l.gb.Mem.Read(LCD_CTRL_ADDR)

	if ctrl & LcdCtrlBgEnable > 0 {
		l.randerBg()
	}

	if ctrl & LcdCtrlWinEnable > 0 {
		l.randerWin()
	}

	if ctrl & LcdCtrlObjEnable > 0 {
		l.randerSprites()
	}
}

func (l *Lcd) randerBg() {
	m := l.gb.Mem
	ctrl := m.Read(LCD_CTRL_ADDR)

	scx := m.Read(LCD_SCX_ADDR)
	scy := m.Read(LCD_SCY_ADDR)

	//Bit 3 - BG Tile Map Display Select     (0=9800-9BFF, 1=9C00-9FFF)
	var bgMapBaseAddr uint16 = 0x9800
	if ctrl & LcdCtrlBgTileMapSelect > 0 {
		bgMapBaseAddr = 0x9C00
	}

	//Bit 4 - BG & Window Tile Data Select   (0=8800-97FF, 1=8000-8FFF)
	var tileDataBaseAddr uint16 = 0x8800
	tileDataSign := true
	if ctrl & LcdCtrlTileDataSelect > 0 {
		tileDataBaseAddr = 0x8000
		tileDataSign = false
	}

	ly := m.Read(LCD_LY_ADDR)
	curY := uint16(scy) + uint16(ly)

	//Each tile is sized 8x8 pixels
	tileRow := curY / 8

	for px := uint16(0); px < 160; px++ {
		curX := uint16(scx) + px

		//Each tile is sized 8x8 pixels
		tileCol := curX / 8

		//The gameboy contains two 32x32 tile background maps in VRAM at addresses 9800h-9BFFh and 9C00h-9FFFh
		tileNo := m.Read(bgMapBaseAddr + tileRow * 32 + tileCol)

		var tileDataAddr uint16
		if tileDataSign {
			//patterns have signed numbers from -128 to 127 (i.e. pattern #0 lies at address $9000)
			tileDataAddr = tileDataBaseAddr + uint16(int16(int8(tileNo)) + 128) * 16
		} else {
			tileDataAddr = tileDataBaseAddr + uint16(tileNo) * 16
		}

		tileLine := curY % 8

		loData := m.Read(tileDataAddr + tileLine * 2)
		hiData := m.Read(tileDataAddr + tileLine * 2 + 1)

		for bitPos := int16(7 - curX % 8); bitPos >= 0 && px < 160; bitPos--{

			//For each line, the first byte defines the least significant bits of the color numbers for each pixel,
			//and the second byte defines the upper bits of the color numbers.
			colorNum := (hiData >> bitPos) & 0x01
			colorNum = colorNum << 1
			colorNum = colorNum | ((loData >> bitPos) & 0x01)

			r, g, b := l.getMonochromeColor(colorNum, LCD_BGP_ADDR)

			l.screen[px][ly] = [3]byte{r,g,b}
			px++
		}
		px--
	}
}

func (l *Lcd) randerWin() {
	m := l.gb.Mem
	ctrl := m.Read(LCD_CTRL_ADDR)

	wx := m.Read(LCD_WX_ADDR)
	wy := m.Read(LCD_WY_ADDR) - 7

	//Bit 3 - BG Tile Map Display Select     (0=9800-9BFF, 1=9C00-9FFF)
	var winMapBaseAddr uint16 = 0x9800
	if ctrl & LcdCtrlWinTileMapSelect > 0 {
		winMapBaseAddr = 0x9C00
	}

	//Bit 4 - BG & Window Tile Data Select   (0=8800-97FF, 1=8000-8FFF)
	var tileDataBaseAddr uint16 = 0x8800
	tileDataSign := true
	if ctrl & LcdCtrlTileDataSelect > 0 {
		tileDataBaseAddr = 0x8000
		tileDataSign = false
	}

	ly := m.Read(LCD_LY_ADDR)
	curY := int16(ly) - int16(wy)
	if curY < 0 {
		return
	}

	//Each tile is sized 8x8 pixels
	tileRow := uint16(curY) / 8

	for px := uint16(0); px < 160; px++ {
		curX := int16(px) - int16(wx)
		if curX < 0 {
			continue
		}

		//Each tile is sized 8x8 pixels
		tileCol := uint16(curX) / 8

		//The gameboy contains two 32x32 tile background maps in VRAM at addresses 9800h-9BFFh and 9C00h-9FFFh
		tileNo := m.Read(winMapBaseAddr + tileRow * 32 + tileCol)

		var tileDataAddr uint16
		if tileDataSign {
			//patterns have signed numbers from -128 to 127 (i.e. pattern #0 lies at address $9000)
			tileDataAddr = tileDataBaseAddr + uint16(int16(int8(tileNo)) + 128) * 16
		} else {
			tileDataAddr = tileDataBaseAddr + uint16(tileNo) * 16
		}

		tileLine := uint16(curY) % 8

		loData := m.Read(tileDataAddr + tileLine * 2)
		hiData := m.Read(tileDataAddr + tileLine * 2 + 1)

		for bitPos := int16(7 - curX % 8); bitPos >= 0 && px < 160; bitPos--{

			//For each line, the first byte defines the least significant bits of the color numbers for each pixel,
			//and the second byte defines the upper bits of the color numbers.
			colorNum := (hiData >> bitPos) & 0x01
			colorNum = colorNum << 1
			colorNum = colorNum | ((loData >> bitPos) & 0x01)

			r, g, b := l.getMonochromeColor(colorNum, LCD_BGP_ADDR)

			l.screen[px][ly] = [3]byte{r,g,b}
			px++
		}
		px--
	}
}

func (l *Lcd) randerSprites() {

}

func (l *Lcd) getMonochromeColor(num byte, addr uint16) (r byte, g byte, b byte) {
	palettes := l.gb.Mem.Read(addr)

	color := (palettes >> (num * 2)) & 0x03

	switch color {
	case 0:
		//White
		return 0xFF, 0xFF, 0xFF
	case 1:
		//Light gray
		return 0xCC, 0xCC, 0xCC
	case 2:
		//Dark gray
		return 0x77, 0x77, 0x77
	case 3:
		//Dark
		return 0x00, 0x00, 0x00
	default:
		return 0xFF, 0xFF, 0xFF
	}
}