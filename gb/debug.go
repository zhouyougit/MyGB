package gb

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"os"
	"strconv"
	"strings"
)

type Debuger struct {
	gb *GameBoy
	p *prompt.Prompt
	lastCmd string
	breakpoints map[uint16]bool
	step bool
	showStep bool
	stepSkipCount int
}

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
	"cart" : (* Debuger).debugFuncCartridge,
	"exit" : (* Debuger).debugFuncExit,
}

var debugFuncInfoMap = map[string]debugFunc {
	"bp" : (* Debuger).debugFuncInfoBreakpoints,
	"cpu" : (* Debuger).debugFuncCpu,
}

var debugFuncDeleteMap = map[string]debugFunc {
	"bp" : (* Debuger).debugFuncDeleteBreakpoint,
}

func NewDebuger(gb *GameBoy) *Debuger {
	d := &Debuger{
		gb: gb,
		step: true,
		showStep: gb.debug,
		breakpoints: make(map[uint16]bool),
	}
	d.p = prompt.New(d.debugParse, d.debugComplate, prompt.OptionPrefix("debug> "))
	return d
}

func (d *Debuger) skipStep() bool {
	if !d.step {
		bp, exist := d.breakpoints[d.gb.Cpu.reg.PC]
		if exist && bp {
			d.step = true
		} else {
			return true
		}
	}
	if d.stepSkipCount > 0 {
		d.stepSkipCount--
		return true
	}
	return false
}

func (d *Debuger) DebugStep() {
	pc := d.gb.Cpu.reg.PC
	skipStep := d.skipStep()

	if d.showStep || !skipStep {
		d.printInstruction(pc, !skipStep)
	}

	if skipStep {
		return
	}

	c := true
	for c {
		text := d.p.Input()
		if len(text) == 0 {
			text = d.lastCmd
		}
		args := strings.Split(text, " ")
		cmd := args[0]
		args = args[1:]

		df, exists := debugFuncMap[cmd]
		if !exists {
			fmt.Printf("Unknown Command : %s\n", cmd)
			continue
		}
		c = df(d, args)

		d.lastCmd = text
	}
}

func (d *Debuger) debugParse(cmd string) {
	println(cmd)
}

func (d *Debuger) debugComplate(doc prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

func (d *Debuger) printInstruction(addr uint16, curr bool) {
	opCodeByte := d.gb.Mem.Read(addr)
	opCode := OPCodesMap[opCodeByte]
	var args = ""
	for i := uint16(1); i < opCode.Length; i++ {
		args += fmt.Sprintf("%02X ", d.gb.Mem.Read(d.gb.Cpu.reg.PC + i))
	}

	if curr {
		fmt.Print("->")
	} else {
		fmt.Print("  ")
	}
	if opCode.Func != nil {
		fmt.Printf("$%04X\t%s\t%s\n", addr, opCode.Mnemonic, args)
	} else {
		fmt.Printf("$%04X\tUnknown opCode [%02X]\n", addr, opCodeByte)
	}
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
		"delete bp\tbp\tDelete breakpoint at address. eg. 'delete bp 0x0150'\n" +
		"cart\t\tPrint Cartridge info\n" +
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

func (d *Debuger) debugFuncCpu(args []string) bool {
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
		cpu.getAF(), cpu.getBC(), cpu.getDE(), cpu.getHL(), cpu.reg.SP, cpu.reg.PC,
		btoi(cpu.getFlagZ()), btoi(cpu.getFlagN()), btoi(cpu.getFlagH()),
		btoi(cpu.getFlagC()), btoi(cpu.getFlagIME()))
	return true
}

func (d *Debuger) debugFuncCartridge(args []string) bool {
	//fmt.Println("next")
	return true
}

func (d *Debuger) debugFuncExit(args []string) bool {
	os.Exit(0)
	return false
}


func parseUint16(text string) (uint16, error) {
	val, err := strconv.ParseUint(text, 0, 16)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return uint16(val), nil
}