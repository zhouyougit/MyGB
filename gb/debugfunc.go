package gb

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

type debugFunc func(*Debuger, []string) bool

var debugFuncMap = map[string]debugFunc {
	"help" : (* Debuger).debugFuncHelp,
	"h" : (* Debuger).debugFuncHelp,
	"next" : (* Debuger).debugFuncNext,
	"n" : (* Debuger).debugFuncNext,
	"step" : (* Debuger).debugFuncStep,
	"s" : (* Debuger).debugFuncStep,
	"finish" : (* Debuger).debugFuncFinish,
	"list" : (* Debuger).debugFuncList,
	"l" : (* Debuger).debugFuncList,
	"print" : (* Debuger).debugFuncPrint,
	"p" : (* Debuger).debugFuncPrint,
	"dump" : (* Debuger).debugFuncDump,
	"d" : (* Debuger).debugFuncDump,
	"continue" : (* Debuger).debugFuncContinue,
	"c" : (* Debuger).debugFuncContinue,
	"break" : (* Debuger).debugFuncBreak,
	"b" : (* Debuger).debugFuncBreak,
	"info" : (* Debuger).debugFuncInfo,
	"delete" : (* Debuger).debugFuncDelete,
	"set" : (* Debuger).debugFuncSet,
	"render" : (* Debuger).debugFuncRender,
	"stack" : (* Debuger).debugFuncStack,
	"save" : (* Debuger).debugFuncSave,
	"opcode" : (* Debuger).debugFuncOpCode,
	"exit" : (* Debuger).debugFuncExit,
}

var debugFuncInfoMap = map[string]debugFunc {
	"bp" : (* Debuger).debugFuncInfoBreakpoints,
	"cpu" : (* Debuger).debugFuncInfoCpu,
	"int" : (* Debuger).debugFuncInfoInt,
	"cart" : (* Debuger).debugFuncInfoCartridge,
	"timer" : (* Debuger).debugFuncInfoTimer,
	"cycle" : (* Debuger).debugFuncInfoCycle,
	"joypad" : (* Debuger).debugFuncInfoJoypad,
	"lcd" : (* Debuger).debugFuncInfoLcd,
}

var debugFuncDeleteMap = map[string]debugFunc {
	"bp" : (* Debuger).debugFuncDeleteBreakpoint,
}

var debugFuncSetMap = map[string]debugFunc {
	"step" : (* Debuger).debugFuncSetShowStep,
}

var debugFuncInfoLcdMap = map[string]debugFunc {
	"stat" : (* Debuger).debugFuncInfoLcdStat,
	"ctrl" : (* Debuger).debugFuncInfoLcdCtrl,
	"pal" : (* Debuger).debugFuncInfoLcdPal,
	"tileMap" : (* Debuger).debugFuncInfoLcdTileMap,
	"tileImg" : (* Debuger).debugFuncInfoLcdTileImg,
}

func (d *Debuger) debugFuncHelp(args []string) bool {
	fmt.Printf("===== Debuger Help =====\n" +
		"help,h\tPrint this Doc\n" +
		"next,n\tNext step\n" +
		"print,p\tPrint data from Memory address. eg. 'print 0x0100'\n" +
		"dump,d\tDump data from Memory address range. eg. 'dump 0x0100 0x0110'\n" +
		"continue,c\tContinue run\n" +
		"break,b\tCreate breakpoint at address. eg. 'break 0x0150'\n" +
		"stack\tdump all goroutine stack\n" +
		"save\tsave all opcode into file. eg. 'save dump.log'\n" +
		"info bp\tPrint all breakpoints\n" +
		"info cpu\tPrint CPU info\n" +
		"info cycle\tPrint Main Cycle info\n" +
		"info cart\tPrint Cartridge info\n" +
		"info int\tPrint Interrupt info\n" +
		"info timer\tPrint Timer info\n" +
		"info lcd stat\tPrint LCD stat\n" +
		"info lcd ctrl\tPrint LCD ctrl\n" +
		"info lcd pal\tPrint LCD palette\n" +
		"info lcd tileMap\tPrint tile map\n" +
		"info lcd tileImg\tPrint tile image\n" +
		"info joypad\tPrint Joypad info\n" +
		"set step\tSet showStep status.\n" +
		"delete bp\tbp\tDelete breakpoint at address. eg. 'delete bp 0x0150'\n" +
		"exit\t\tExit Program\n\n")
	return true
}

func (d *Debuger) debugFuncSubCmd(cmdMap map[string]debugFunc, args []string) bool {
	if len(args) < 1 {
		return true
	}
	subCmd := args[0]
	args = args[1:]

	df, exists := cmdMap[subCmd]
	if !exists {
		fmt.Printf("Unknown Sub Command : %s\n", subCmd)
		return true
	}
	return df(d, args)
}

func (d *Debuger) debugFuncInfo(args []string) bool {
	return d.debugFuncSubCmd(debugFuncInfoMap, args)
}

func (d *Debuger) debugFuncDelete(args []string) bool {
	return d.debugFuncSubCmd(debugFuncDeleteMap, args)
}

func (d *Debuger) debugFuncSet(args []string) bool {
	return d.debugFuncSubCmd(debugFuncSetMap, args)
}

func (d *Debuger) debugFuncInfoLcd(args []string) bool {
	return d.debugFuncSubCmd(debugFuncInfoLcdMap, args)
}

func (d *Debuger) debugFuncNext(args []string) bool {
	isCall, nextPC := d.checkCall(d.gb.Cpu.reg.PC)
	if isCall {
		d.breakpoints[nextPC] = BreakPoint{enable:true, temporary:true}
		d.step = false
	}
	return false
}

func (d *Debuger) debugFuncStep(args []string) bool {
	if len(args) > 0 {
		var err error
		d.stepSkipCount, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
			return true
		}
	}
	return false
}

func (d *Debuger) debugFuncFinish(args []string) bool {
	d.breakpoints[d.finishAddr] = BreakPoint{enable:true, temporary:true}
	d.step = false
	return false
}

func (d *Debuger) debugFuncList(args []string) bool {
	d.printInstruction(d.gb.Cpu.reg.PC, true)
	return true
}

func (d *Debuger) debugFuncPrint(args []string) bool {
	if len(args) < 1 {
		return true
	}
	addr, err := strconv.ParseUint(args[0], 0, 16)
	if err != nil {
		fmt.Println(err)
		return true
	}
	fmt.Printf("$%04X : ", addr)
	length := 1
	if len(args) > 1 {
		length, _ = strconv.Atoi(args[1])
	}
	for i := 0; i < length; i++ {
		fmt.Printf("%02X ", d.gb.Mem.Read(uint16(addr) + uint16(i)))
	}
	fmt.Printf("\n")
	return true
}

func (d *Debuger) debugFuncDump(args []string) bool {
	if len(args) != 2 {
		fmt.Println("Usage: dump [beginAddr] [endAddr]")
		return true
	}
	begin, err := strconv.ParseUint(args[0], 0, 16)
	if err != nil {
		fmt.Println(err)
		return true
	}
	end, err := strconv.ParseUint(args[1], 0, 16)
	if err != nil {
		fmt.Println(err)
		return true
	}
	d.gb.Mem.Dump(uint16(begin), uint16(end))
	return true
}

func (d *Debuger) debugFuncContinue(args []string) bool {
	d.step = false
	return false
}

func (d *Debuger) debugFuncBreak(args []string) bool {
	addr := d.gb.Cpu.reg.PC
	if len(args) > 0 {
		var err error
		addr, err = parseUint16(args[0])
		if err != nil {
			return true
		}
	}
	d.breakpoints[addr] = BreakPoint{enable:true, temporary:false}
	fmt.Printf("Create breakpoint at [0x%04X]\n", addr)
	return true
}

func (d *Debuger) debugFuncInfoBreakpoints(args []string) bool {
	fmt.Printf("name\t\taddr\tenabled\ttemporary\n")
	for addr, bp := range d.breakpoints {
		fmt.Printf("breakpoint\t%04X\t%v\t%v\n", addr, bp.enable, bp.temporary)
	}
	fmt.Println()
	return true
}
func (d *Debuger) debugFuncDeleteBreakpoint(args []string) bool {
	if len(args) < 1 {
		fmt.Println("Usage: delete bp [breakpoint addr]")
		return true
	}
	addr, err := parseUint16(args[0])
	if err != nil {
		fmt.Println("Usage: delete bp [breakpoint addr]")
		return true
	}
	_, exist := d.breakpoints[addr]
	if exist {
		delete(d.breakpoints, addr)
		fmt.Printf("Delete breakpoint at [0x%04X]\n", addr)
	} else {
		fmt.Printf("Breakpoint at [0x%04X] not found\n", addr)
	}
	return true
}

func (d *Debuger) debugFuncInfoCpu(args []string) bool {
	cpu := d.gb.Cpu
	btoi := func(a bool) int {
		if a {
			return 1
		} else {
			return 0
		}
	}
	fmt.Printf("== CPU Info ==\n" +
		"Reg  : AF=%04X BC=%04X DE=%04X HL=%04X SP=%04X PC=%04X\n" +
		"Flag : Z=%b N=%b H=%b C=%b IME=%b\n" +
		"Halt : %v\n\n",
		cpu.reg.getAF(), cpu.reg.getBC(), cpu.reg.getDE(), cpu.reg.getHL(), cpu.reg.SP, cpu.reg.PC,
		btoi(cpu.getFlagZ()), btoi(cpu.getFlagN()), btoi(cpu.getFlagH()),
		btoi(cpu.getFlagC()), btoi(cpu.getFlagIME()),
		cpu.Halt)
	return true
}

func (d *Debuger) debugFuncInfoInt(args []string) bool {
	enable := d.gb.Mem.Read(IE_ADDR)
	flag := d.gb.Mem.Read(IF_ADDR)

	fmt.Printf("== Interrupt Info ==\n" +
		"Name\t\tEnable\tFlag\n" +
		"IME\t\t%v\t-\n" +
		"V-Blank\t\t%t\t%t\n" +
		"LCD STAT\t%t\t%t\n" +
		"Timer\t\t%t\t%t\n" +
		"Serial\t\t%t\t%t\n" +
		"Joypad\t\t%t\t%t\n" +
		"\n",
		d.gb.Cpu.IME,
		enable & IntVBlank.Mask > 0, flag & IntVBlank.Mask > 0,
		enable & IntLcdStat.Mask > 0, flag & IntLcdStat.Mask > 0,
		enable & IntTimer.Mask > 0, flag & IntTimer.Mask > 0,
		enable & IntSerial.Mask > 0, flag & IntSerial.Mask > 0,
		enable & IntJoypad.Mask > 0, flag & IntJoypad.Mask > 0)
	return true
}

func (d *Debuger) debugFuncInfoCartridge(args []string) bool {
	//fmt.Println("next")
	return true
}

func (d *Debuger) debugFuncInfoCycle(args []string) bool {
	fmt.Printf("Main Cycles : %d/%d\n", d.gb.currCycles, d.gb.cyclesInFrame)
	return true
}

func (d *Debuger) debugFuncInfoLcdStat(args []string) bool {
	l := d.gb.Lcd
	fmt.Printf("== LCD Status Info ==\n" +
		"LCD Display Enabled\t: %t\n" +
		"LCD Cycles\t\t: %d/%d\n",
		l.isLCDDisplayEnabled(), l.scanLineCounter, LCD_FULL_CYCLES)
	ly := d.gb.Mem.Read(LCD_LY_ADDR)
	lyc := d.gb.Mem.Read(LCD_LYC_ADDR)
	fmt.Printf("Scan Line (LY)\t\t: %d (LYC=%d)\n", ly, lyc)

	stat := d.gb.Mem.Read(LCD_STAT_ADDR)

	mode := stat & 0b11
	switch mode {
	case 0:
		fmt.Printf("LCD Mode\t\t: 0 (H-Blank)\n")
	case 1:
		fmt.Printf("LCD Mode\t\t: 1 (V-Blank)\n")
	case 2:
		fmt.Printf("LCD Mode\t\t: 2 (Searching OAM)\n")
	case 3:
		fmt.Printf("LCD Mode\t\t: 3 (Transfering Data)\n")
	}

	fmt.Printf("Int Enabled\t\t: Mode0=%t Mode1=%t Mode2=%t LYC=%t\n",
		stat & 0b1000 == 0b1000, stat & 0b10000 == 0b10000,
		stat & 0b100000 == 0b100000, stat & 0b1000000 == 0b1000000)

	fmt.Println("")

	return true
}

func (d *Debuger) debugFuncInfoLcdCtrl(args []string) bool {
	bit2str := func(b byte, s1, s2 string) string {
		if b > 0 {
			return s1
		} else {
			return s2
		}
	}
	ctrl := d.gb.Mem.Read(LCD_CTRL_ADDR)
	fmt.Printf("== LCD Control Info ==\n" +
		"LCD Enable:\t\t%s\n" +
		"Win Enable:\t\t%s\n" +
		"Win Tile Map Select:\t%s\n" +
		"Bg Enable:\t\t%s\n" +
		"Bg Tile Map Select:\t%s\n" +
		"Tile Data Select:\t%s\n" +
		"Obj Enable:\t\t%s\n" +
		"Obj Size:\t\t%s\n\n",
		bit2str(ctrl & LcdCtrlLCDEnable, "On", "Off"),
		bit2str(ctrl & LcdCtrlWinEnable, "On", "Off"),
		bit2str(ctrl & LcdCtrlWinTileMapSelect, "0x9C00-0x9FFF", "0x9800-0x9BFF"),
		bit2str(ctrl & LcdCtrlBgEnable, "On", "Off"),
		bit2str(ctrl & LcdCtrlBgTileMapSelect, "0x9C00-0x9FFF", "0x9800-0x9BFF"),
		bit2str(ctrl & LcdCtrlTileDataSelect, "0x8000-0x8FFF", "0x8800-0x97FF"),
		bit2str(ctrl & LcdCtrlObjEnable, "On", "Off"),
		bit2str(ctrl & LcdCtrlObjSize, "8*16", "8*8"),
		)

	return true
}

func (d *Debuger) debugFuncInfoLcdPal(args []string) bool {
	printPal := func(addr uint16, name string) {
		pal := d.gb.Mem.Read(addr)
		fmt.Printf("%s:", name)

		colorMap := map[byte]string{
			0x00: "White",
			0x01: "Light gray",
			0x02: "Dark gray",
			0x03: "Black",
		}
		for i := 0; i < 4; i++ {
			color := (pal >> (i * 2)) & 0x03
			fmt.Printf("\t%d:%s", i, colorMap[color])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("== LCD Palettes Info ==\n")
	printPal(LCD_BGP_ADDR, "BGP")
	printPal(LCD_OBP0_ADDR, "OBP0")
	printPal(LCD_OBP1_ADDR, "OBP1")
	fmt.Printf("\n\n")
	return true
}

func (d *Debuger) debugFuncInfoLcdTileMap(args []string) bool {
	if len(args) < 1 {
		fmt.Printf("Usage: info lcd tileMap [lineAddr]")
		return true
	}

	baseAddr, err := strconv.ParseUint(args[0], 0, 16)
	if err != nil {
		fmt.Println(err)
		return true
	}

	lines := 1
	if len(args) > 1 {
		lines, _ = strconv.Atoi(args[1])
	}

	fmt.Printf("============ Dump Tile Map from 0x%04X =============\n" +
		"addr\t00 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19\n" +
		"---------------------------------------------------------------------", baseAddr)
	for line := 0; line < lines; line++ {
		begin := uint16(baseAddr) + uint16(line) * 32
		end := begin + 20
		for i := begin; i < end; i++ {
			if (i-begin)%20 == 0 {
				fmt.Printf("\n0x%04X\t", uint16(i))
			}
			fmt.Printf("%02X ", d.gb.Mem.Read(uint16(i)))
		}
	}
	fmt.Printf("\n")
	return true
}

func (d *Debuger) debugFuncInfoLcdTileImg(args []string) bool {
	if len(args) < 1 {
		fmt.Printf("Usage: info lcd tileImg [TileNo]")
		return true
	}
	tileNo, err := strconv.ParseUint(args[0], 0, 16)
	if err != nil {
		fmt.Println(err)
		return true
	}

	ctrl := d.gb.Mem.Read(LCD_CTRL_ADDR)
	var tileDataBaseAddr uint16 = 0x8800
	tileDataSign := true
	if ctrl & LcdCtrlTileDataSelect > 0 {
		tileDataBaseAddr = 0x8000
		tileDataSign = false
	}

	var tileDataAddr uint16
	if tileDataSign {
		//patterns have signed numbers from -128 to 127 (i.e. pattern #0 lies at address $9000)
		tileDataAddr = tileDataBaseAddr + uint16(int16(int8(tileNo)) + 128) * 16
	} else {
		tileDataAddr = tileDataBaseAddr + uint16(tileNo) * 16
	}

	var img [8][8]uint8
	colorMap := d.colorMap(LCD_BGP_ADDR)
	for row := uint16(0); row < 8; row++ {
		loData := d.gb.Mem.Read(tileDataAddr + row * 2)
		hiData := d.gb.Mem.Read(tileDataAddr + row * 2 + 1)

		for col := uint8(0); col < 8; col++ {
			bitPos := 7 - col

			//For each line, the first byte defines the least significant bits of the color numbers for each pixel,
			//and the second byte defines the upper bits of the color numbers.
			colorNum := (hiData >> bitPos) & 0x01
			colorNum = colorNum << 1
			colorNum = colorNum | ((loData >> bitPos) & 0x01)

			img[row][col] = colorMap[colorNum]
		}
	}

	tmap := map[byte]string {
		0: "\033[107m  ",
		1: "\033[47m  ",
		2: "\033[100m  ",
		3: "\033[40m  ",
	}
	fmt.Printf("== Tile [%02X] at 0x%04X == \n", tileNo, tileDataAddr)
	for row := byte(0); row < 8; row++ {
		for col := byte(0); col < 8; col++ {
			fmt.Print(tmap[img[row][col]])
		}
		fmt.Printf("\033[0m\n")
	}
	fmt.Printf("\n")
	return true
}

func (d *Debuger) colorMap(addr uint16) (r map[byte]byte) {
	pal := d.gb.Mem.Read(addr)
	r = map[byte]byte{}

	for i := byte(0); i < 4; i++ {
		color := (pal >> (i * 2)) & 0x03
		r[i] = color
	}
	return
}


func (d *Debuger) debugFuncInfoTimer(args []string) bool {
	t := d.gb.Timer
	fmt.Printf("== Timer Info ==\n")
	tima := d.gb.Mem.Read(TIMER_TIMA_ADDR)
	tma := d.gb.Mem.Read(TIMER_TMA_ADDR)
	div := d.gb.Mem.Read(TIMER_DIV_ADDR)

	fmt.Printf("Started\t\t: %t\n" +
		"Freq\t\t: %d\n" +
		"TIMA Cycles\t: %d/%d\n" +
		"TIMA/TMA\t: %02X/%02X\n" +
		"DIV Cycles\t: %d/%d\n" +
		"DIV\t\t: %02X\n\n",
		t.isTimerStarted(),
		t.getTimerFreq(),
		t.cyclesCounter, d.gb.Cpu.frequency / t.getTimerFreq(),
		tima, tma,
		t.dividerCounter, t.dividerMax,
		div)
	return true
}

func (d *Debuger) debugFuncInfoJoypad(args []string) bool {
	j := d.gb.Joypad
	testBtn := func(btn byte) string {
		if j.status & (0x1 << btn) == 0 {
			return "Pressed"
		} else {
			return "Released"
		}
	}
	selectBtn := "Direction"
	if j.selectButton {
		selectBtn = "Button"
	}
	fmt.Printf("== Joypad Info ==\n" +
		"Select %s\n" +
		"ButtonRight is %s\tButtonLeft is %s\n" +
		"ButtonUp is %s\tButtonDown is %s\n" +
		"ButtonA is %s\tButtonB is %s\n" +
		"ButtonSelect is %s\tButtonStart is %s\n",
		selectBtn,
		testBtn(ButtonRight), testBtn(ButtonLeft),
		testBtn(ButtonUp), testBtn(ButtonDown),
		testBtn(ButtonA), testBtn(ButtonB),
		testBtn(ButtonSelect), testBtn(ButtonStart))
	return true
}

func (d *Debuger) debugFuncStack(args []string) bool {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, true)]
	fmt.Printf("== Stack Info ==\n%s\n", buf)
	return true
}

func (d *Debuger) debugFuncSave(args []string) bool {
	if d.saveFile != nil {
		d.saveFile.Close()
		d.saveFile = nil
		d.savePath = ""

		if len(args) == 0 {
			fmt.Printf("Saved file %s\n", d.savePath)
			return true
		}
	}
	if len(args) == 0 {
		fmt.Println("Usage: save [fileName]")
		return true
	}
	d.savePath = args[0]

	var err error
	d.saveFile, err = os.OpenFile(d.savePath, os.O_WRONLY | os.O_CREATE, 0664)
	if err != nil {
		fmt.Println(err)
		return true
	}
	fmt.Printf("Save into file %s\n", d.savePath)
	return true
}

func (d *Debuger) debugFuncOpCode(args []string) bool {
	if len(args) == 0 {
		fmt.Println("Usage: opcode [hex]")
		return true
	}
	code := args[0]
	codeMap := OPCodesMap
	if len(code) == 2 {

	} else if code[:2] == "CB" {
		code = code[2:]
		codeMap = OPCodesMapCB
	} else {
		fmt.Printf("Unknown opCode [%X]\n", code)
		return true
	}
	code = "0x" + code

	codeNo, err := strconv.ParseUint(code, 0, 16)
	if err != nil {
		fmt.Println(err)
		return true
	}

	opCode := codeMap[codeNo]
	if opCode.Func == nil {
		fmt.Printf("Unknown opCode [%X]\n", code)
		return true
	}

	fmt.Printf("%s\n\n", opCode.Mnemonic)

	return true
}

func (d *Debuger) debugFuncExit(args []string) bool {
	os.Exit(0)
	return false
}

func (d *Debuger) debugFuncSetShowStep(args []string) bool {
	if len(args) < 1 {
		fmt.Printf("Usage : set step [true|false].\n")
	}
	d.showStep, _ = strconv.ParseBool(args[0])
	fmt.Printf("Show Step : %v\n", d.showStep)
	return true
}

func (d *Debuger) debugFuncRender(args []string) bool {
	d.gb.Monitor.RenderScreen()
	return true
}

func parseUint16(text string) (uint16, error) {
	val, err := strconv.ParseUint(text, 0, 16)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return uint16(val), nil
}