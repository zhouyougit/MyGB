package gb

import (
	"fmt"
	"os"
	"strconv"
)

type debugFunc func(*Debuger, []string) bool

var debugFuncMap = map[string]debugFunc {
	"help" : (* Debuger).debugFuncHelp,
	"h" : (* Debuger).debugFuncHelp,
	"next" : (* Debuger).debugFuncNext,
	"n" : (* Debuger).debugFuncNext,
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
	"exit" : (* Debuger).debugFuncExit,
}

var debugFuncInfoMap = map[string]debugFunc {
	"bp" : (* Debuger).debugFuncInfoBreakpoints,
	"cpu" : (* Debuger).debugFuncInfoCpu,
	"int" : (* Debuger).debugFuncInfoInt,
	"cart" : (* Debuger).debugFuncInfoCartridge,
	"timer" : (* Debuger).debugFuncInfoTimer,
	"lcd" : (* Debuger).debugFuncInfoLCD,
	"cycle" : (* Debuger).debugFuncInfoCycle,
}

var debugFuncDeleteMap = map[string]debugFunc {
	"bp" : (* Debuger).debugFuncDeleteBreakpoint,
}

var debugFuncSetMap = map[string]debugFunc {
	"step" : (* Debuger).debugFuncSetShowStep,
}

func (d *Debuger) debugFuncHelp(args []string) bool {
	fmt.Printf("===== Debuger Help =====\n" +
		"help,h\tPrint this Doc\n" +
		"next,n\tNext step\n" +
		"print,p\tPrint data from Memory address. eg. 'print 0x0100'\n" +
		"dump,d\tDump data from Memory address range. eg. 'dump 0x0100 0x0110'\n" +
		"continue,c\tContinue run\n" +
		"break,b\tCreate breakpoint at address. eg. 'break 0x0150'\n" +
		"info bp\tPrint all breakpoints\n" +
		"info cpu\tPrint CPU info\n" +
		"info cycle\tPrint Main Cycle info\n" +
		"info cart\tPrint Cartridge info\n" +
		"info int\tPrint Interrupt info\n" +
		"info timer\tPrint Timer info\n" +
		"info lcd\tPrint LCD info\n" +
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

func (d *Debuger) debugFuncNext(args []string) bool {
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
	d.breakpoints[addr] = true
	fmt.Printf("Create breakpoint at [0x%04X]\n", addr)
	return true
}

func (d *Debuger) debugFuncInfoBreakpoints(args []string) bool {
	fmt.Printf("name\t\taddr\tenabled\n")
	for addr, enabled := range d.breakpoints {
		fmt.Printf("breakpoint\t%04X\t%v\n", addr, enabled)
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
		"Flag : Z=%b N=%b H=%b C=%b IME=%b\n\n",
		cpu.reg.getAF(), cpu.reg.getBC(), cpu.reg.getDE(), cpu.reg.getHL(), cpu.reg.SP, cpu.reg.PC,
		btoi(cpu.getFlagZ()), btoi(cpu.getFlagN()), btoi(cpu.getFlagH()),
		btoi(cpu.getFlagC()), btoi(cpu.getFlagIME()))
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

func (d *Debuger) debugFuncInfoLCD(args []string) bool {
	l := d.gb.Lcd
	fmt.Printf("== LCD Info ==\n" +
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

func parseUint16(text string) (uint16, error) {
	val, err := strconv.ParseUint(text, 0, 16)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return uint16(val), nil
}