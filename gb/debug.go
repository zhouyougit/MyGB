package gb

import (
	"bytes"
	"fmt"
	"github.com/c-bata/go-prompt"
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

func (d *Debuger) DebugOpcode() {
	if d.gb.Cpu.Halt {
		return
	}
	pc := d.gb.Cpu.reg.PC
	skipStep := d.skipStep()

	if d.showStep || !skipStep {
		d.printInstruction(pc, !skipStep)
	}

	if skipStep {
		return
	}

	d.doDebugInteract()
}

func (d *Debuger) SetStep() {
	d.step = true
	if d.gb.Cpu.Halt {
		d.doDebugInteract()
	}
}

func (d *Debuger) DebugInterrupt(i Interrupt) {
	skipStep := d.skipStep()

	if d.showStep || !skipStep {
		fmt.Printf("->$\tPush PC; JP 0x%04X;\t(Interrupted by %s)\n", i.Addr, i.Name)
	}

	if skipStep {
		return
	}
	//d.doDebugInteract()
}

func (d *Debuger) doDebugInteract() {
	c := true
	for c {
		text := d.p.Input()
		if len(text) == 0 {
			text = d.lastCmd
		}
		if len(text) == 0 {
			continue
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
	prefix := "  "
	if curr {
		prefix = "->"
	}
	opCodeByte := d.gb.Mem.Read(addr)
	var opCodeByte2 byte
	var opCode OPCode
	var argLength = 0
	argsAddr := addr + 1

	if opCodeByte == 0xCB {
		opCodeByte2 = d.gb.Mem.Read(addr+1)
		opCode = OPCodesMapCB[opCodeByte2]
		argsAddr++
	} else {
		opCode = OPCodesMap[opCodeByte]
		argLength = int(opCode.Length - 1)
	}

	if opCode.Func == nil {
		if opCodeByte == 0xCB {
			fmt.Printf("%s$0x%04X\tUnknown opCode [0x%02X%02X]\n", prefix, addr, opCodeByte, opCodeByte2)
		} else {
			fmt.Printf("%s$0x%04X\tUnknown opCode [0x%02X]\n", prefix, addr, opCodeByte)
		}
		return
	}

	var tab string
	if len(opCode.Mnemonic) < 8 {
		tab = "\t\t"
	} else {
		tab = "\t"
	}

	var argStr = ""
	if argLength > 0 {
		for i := uint16(0); i < uint16(argLength); i++ {
			argStr += fmt.Sprintf("%02X ", d.gb.Mem.Read(argsAddr + i))
		}
	}

	var argValue = d.generateOpCodeArgValueString(opCode, argsAddr)

	fmt.Printf("%s$0x%04X\t%s%s%s%s\n", prefix, addr, opCode.Mnemonic, tab, argStr, argValue)
}

func (d *Debuger) generateOpCodeArgValueString(code OPCode, addr uint16) string {
	out := bytes.Buffer{}
	out.WriteString("\t// ")
	ss := strings.Split(code.Mnemonic, " ")
	cmd := ss[0]

	out.WriteString(d.generateOpCodeArgValueByCmd(cmd))

	if len(ss) > 1 {
		argDefs := ss[1]
		argDefList := strings.Split(argDefs, ",")

		for idx, ad := range argDefList {
			out.WriteString(d.generateOneOpCodeArgValue(cmd, ad, addr, idx, len(argDefList)))
		}
	}
	if out.Len() == 4 {
		return ""
	}
	return out.String()
}

func (d *Debuger) generateOneOpCodeArgValue(cmd string, ad string, addr uint16, idx int, total int) string {
	switch ad {
	case "A":
		if idx != 0 || total == 1 {
			return fmt.Sprintf("A=%02X ", d.gb.Cpu.reg.A)
		}
	case "B":
		if idx != 0 || total == 1 {
			return fmt.Sprintf("B=%02X ", d.gb.Cpu.reg.B)
		}
	case "C":
		if idx == 0 {
			if cmd == "JR" || cmd == "CALL" || cmd == "JP" {
				return fmt.Sprintf("flagC=%t ", d.gb.Cpu.getFlagC())
			}
		} else if total == 1 {
			return fmt.Sprintf("C=%02X ", d.gb.Cpu.reg.C)
		}
	case "D":
		if idx != 0 || total == 1 {
			return fmt.Sprintf("D=%02X ", d.gb.Cpu.reg.D)
		}
	case "E":
		if idx != 0 || total == 1 {
			return fmt.Sprintf("E=%02X ", d.gb.Cpu.reg.E)
		}
	case "H":
		if idx != 0 || total == 1 {
			return fmt.Sprintf("H=%02X ", d.gb.Cpu.reg.H)
		}
	case "L":
		if idx != 0 || total == 1 {
			return fmt.Sprintf("L=%02X ", d.gb.Cpu.reg.L)
		}
	case "AF":
		if idx != 0 || total == 1 {
			return fmt.Sprintf("AF=%04X ", d.gb.Cpu.reg.getAF())
		}
	case "BC":
		if idx != 0 || total == 1 {
			return fmt.Sprintf("BC=%04X ", d.gb.Cpu.reg.getBC())
		}
	case "DE":
		if idx != 0 || total == 1 {
			return fmt.Sprintf("DE=%04X ", d.gb.Cpu.reg.getDE())
		}
	case "HL":
		if idx != 0 || total == 1 {
			return fmt.Sprintf("HL=%04X ", d.gb.Cpu.reg.getHL())
		}
	case "SP":
		if idx != 0 || total == 1 {
			return fmt.Sprintf("SP=%04X ", d.gb.Cpu.reg.SP)
		}
	case "PC":
		if idx != 0 || total == 1 {
			return fmt.Sprintf("PC=%04X ", d.gb.Cpu.reg.PC)
		}
	case "(a8)":
		a8 := d.gb.Mem.Read(addr)
		if idx == 0 {
			return fmt.Sprintf("0xFF00+a8=0xFF%02X ", a8)
		} else {
			return fmt.Sprintf("0xFF00+a8=0xFF%02X (0xFF+a8)=%02X ", a8, d.gb.Mem.Read(0xFF00 + uint16(a8)))
		}
	case "(a16)":
		a16 := d.gb.Mem.ReadUint16(addr)
		if idx == 0 {
			return fmt.Sprintf("a16=0x%04X ", a16)
		} else {
			return fmt.Sprintf("a16=0x%04X (a16)=%02X ", a16, d.gb.Mem.Read(a16))
		}
	case "a16":
		a16 := d.gb.Mem.ReadUint16(addr)
		return fmt.Sprintf("a16=0x%04X ", a16)
	case "d8":
		d8 := d.gb.Mem.Read(addr)
		return fmt.Sprintf("d8=%02X ", d8)
	case "d16":
		d16 := d.gb.Mem.ReadUint16(addr)
		return fmt.Sprintf("d16=%04X ", d16)
	case "Z","NZ":
		return fmt.Sprintf("flagZ=%t ", d.gb.Cpu.getFlagZ())
	case "NC":
		return fmt.Sprintf("flagC=%t ", d.gb.Cpu.getFlagC())
	case "r8":
		r8 := int8(d.gb.Mem.Read(addr))
		if r8 >= 0 {
			return fmt.Sprintf("r8=0x%02X ", r8)
		} else {
			return fmt.Sprintf("r8=-0x%02X ", -r8)
		}
	case "(HL)", "(HL-)", "(HL+)":
		if idx == 0 {
			return fmt.Sprintf("HL=0x%04X ", d.gb.Cpu.reg.getHL())
		} else {
			return fmt.Sprintf("HL=0x%04X (HL)=%02X", d.gb.Cpu.reg.getHL(), d.gb.Mem.Read(d.gb.Cpu.reg.getHL()))
		}
	}
	return ""
}

func (d *Debuger) generateOpCodeArgValueByCmd(cmd string) string {
	switch cmd {
	case "CP":
		return fmt.Sprintf("A=%02X ", d.gb.Cpu.reg.A)
	}
	return ""
}