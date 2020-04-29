package gb

import (
	"os"
)

type Interrupt struct {
	Name string
	Mask byte
	Pos	byte
	Addr uint16
}

const (
	/**
	FFFF - IE - Interrupt Enable (R/W)
	  Bit 0: V-Blank  Interrupt Enable  (INT 40h)  (1=Enable)
	  Bit 1: LCD STAT Interrupt Enable  (INT 48h)  (1=Enable)
	  Bit 2: Timer    Interrupt Enable  (INT 50h)  (1=Enable)
	  Bit 3: Serial   Interrupt Enable  (INT 58h)  (1=Enable)
	  Bit 4: Joypad   Interrupt Enable  (INT 60h)  (1=Enable)

	FF0F - IF - Interrupt Flag (R/W)
	  Bit 0: V-Blank  Interrupt Request (INT 40h)  (1=Request)
	  Bit 1: LCD STAT Interrupt Request (INT 48h)  (1=Request)
	  Bit 2: Timer    Interrupt Request (INT 50h)  (1=Request)
	  Bit 3: Serial   Interrupt Request (INT 58h)  (1=Request)
	  Bit 4: Joypad   Interrupt Request (INT 60h)  (1=Request)
	*/
	IF_ADDR uint16 = 0xFF0F
	IE_ADDR uint16 = 0xFFFF
)
var (
	IntVBlank  = Interrupt{"V-Blank", 0x1 << 0, 0, 0x40}
	IntLcdStat = Interrupt{"LCD STAT", 0x1 << 1, 1, 0x48}
	IntTimer   = Interrupt{"Timer", 0x1 << 2, 2, 0x50}
	IntSerial  = Interrupt{"Serial", 0x1 << 3, 3, 0x58}
	IntJoypad  = Interrupt{"Joypad", 0x1 << 4, 4, 0x60}
)

var IntPriority = []Interrupt {
	IntVBlank,
	IntLcdStat,
	IntTimer,
	IntSerial,
	IntJoypad,
}

type Cpu struct {
	frequency int
	reg Registers
	// IME - Interrupt Master Enable Flag(Write Only)
	IME  bool
	Halt bool

	gb *GameBoy

	showStep bool
}

/**
Registers
16bit	Hi	Lo	Name/Function
AF		A	-	Accumulator & Flags
BC		B	C	BC
DE		D	E	DE
HL		H	L	HL
SP		-	-	Stack Pointer
PC		-	-	Program Counter/Pointer

The Flag Register (lower 8bit of AF register)
Bit		Name	Set	Clr	Expl.
7		zf		Z	NZ	Zero Flag
6		n		-	-	Add/Sub-Flag (BCD)
5		h		-	-	Half Carry Flag (BCD)
4		cy		C	NC	Carry Flag
3-0		-		-	-	Not used (always zero)
 */
type Registers struct {
	A 	byte
	B	byte
	C	byte
	D	byte
	E	byte
	F	byte
	H	byte
	L	byte
	SP	uint16
	PC	uint16
}

/**
  AF=$01B0
  BC=$0013
  DE=$00D8
  HL=$014D
  Stack Pointer=$FFFE
 */
func (cpu *Cpu) Init(gb *GameBoy) {
	cpu.reg.setAF(0x01B0)
	cpu.reg.setBC(0x0013)
	cpu.reg.setDE(0x00D8)
	cpu.reg.setHL(0x014D)
	cpu.reg.SP = 0xFFFE
	cpu.reg.PC = 0x0100
	cpu.gb = gb
	cpu.showStep = gb.debug
}

func (reg *Registers) getAF() uint16 {
	return uint16(reg.A) << 8 | uint16(reg.F)
}

func (reg *Registers) setAF(data uint16) {
	reg.A = byte((data & 0xFF00) >> 8)
	reg.F = byte(data & 0x00FF)
}

func (reg *Registers) getBC() uint16 {
	return uint16(reg.B) << 8 | uint16(reg.C)
}

func (reg *Registers) setBC(data uint16) {
	reg.B = byte((data & 0xFF00) >> 8)
	reg.C = byte(data & 0x00FF)
}

func (reg *Registers) getDE() uint16 {
	return uint16(reg.D) << 8 | uint16(reg.E)
}

func (reg *Registers) setDE(data uint16) {
	reg.D = byte((data & 0xFF00) >> 8)
	reg.E = byte(data & 0x00FF)
}

func (reg *Registers) getHL() uint16 {
	return uint16(reg.H) << 8 | uint16(reg.L)
}

func (reg *Registers) setHL(data uint16) {
	reg.H = byte((data & 0xFF00) >> 8)
	reg.L = byte(data & 0x00FF)
}

func (cpu *Cpu) getFlagZ() bool{
	return cpu.reg.F & 0x80 > 0
}

func (cpu *Cpu) setFlagZ() {
	cpu.reg.F |= 0x80
}

func (cpu *Cpu) resetFlagZ() {
	cpu.reg.F &= 0x7F
}

func (cpu *Cpu) getFlagN() bool{
	return cpu.reg.F & 0x40 > 0
}

func (cpu *Cpu) setFlagN() {
	cpu.reg.F |= 0x40
}

func (cpu *Cpu) resetFlagN() {
	cpu.reg.F &= 0xBF
}

func (cpu *Cpu) getFlagH() bool{
	return cpu.reg.F & 0x20 > 0
}

func (cpu *Cpu) setFlagH() {
	cpu.reg.F |= 0x20
}

func (cpu *Cpu) resetFlagH() {
	cpu.reg.F &= 0xDF
}

func (cpu *Cpu) getFlagC() bool{
	return cpu.reg.F & 0x10 > 0
}

func (cpu *Cpu) setFlagC() {
	cpu.reg.F |= 0x10
}

func (cpu *Cpu) resetFlagC() {
	cpu.reg.F &= 0xEF
}

func (cpu *Cpu) resetFlagZNHC() {
	cpu.reg.F &= 0x0F
}

func (cpu *Cpu) getFlagIME() bool{
	return cpu.IME
}

func (cpu *Cpu) setFlagIME(val bool) {
	cpu.IME = val
}

func (cpu *Cpu) resetFlagIME() {
	cpu.IME = true
}

func (cpu *Cpu) executeNextOpcode() int {
	if cpu.Halt {
		return 4
	}
	var opCode OPCode
	opCodeByte := cpu.gb.Mem.Read(cpu.reg.PC)
	cpu.reg.PC++
	if opCodeByte == 0xCB {
		opCodeByteCB := cpu.gb.Mem.Read(cpu.reg.PC)
		cpu.reg.PC++
		opCode = OPCodesMapCB[opCodeByteCB]
	} else {
		opCode = OPCodesMap[opCodeByte]
	}
	if opCode.Func == nil {
		os.Exit(0)
		return 0
	}
	ret := opCode.Func(cpu)
	return opCode.Cycles[ret]
}

func (cpu *Cpu) RequestInterrupt(req Interrupt) {
	flag := cpu.gb.Mem.Read(IF_ADDR)
	flag |= req.Mask
	cpu.gb.Mem.Write(IF_ADDR, flag)
}

func (cpu *Cpu) checkInterrupt() int {
	if !cpu.IME {
		return 0
	}

	enabled := cpu.gb.Mem.Read(IE_ADDR)
	if enabled == 0 {
		return 0
	}

	flag := cpu.gb.Mem.Read(IF_ADDR)
	if flag == 0 {
		return 0
	}

	for _, i := range []Interrupt{} {
		if enabled & flag & i.Mask == 0 {
			continue
		}
		cpu.runInterrupt(i)
		return 16 + 4 // TODO: Interrupt cycles = PUSH PC + JP (HL)?
	}
	return 0
}

func (cpu *Cpu) runInterrupt(i Interrupt) {
	cpu.gb.Debuger.DebugInterrupt(i)
	cpu.IME = false
	cpu.Halt = false

	flag := cpu.gb.Mem.Read(IF_ADDR)
	flag = flag &^ i.Mask
	cpu.gb.Mem.Write(IF_ADDR, flag)

	cpu.reg.SP -= 2
	cpu.gb.Mem.WriteUint16(cpu.reg.SP, cpu.reg.PC)
	cpu.reg.PC = i.Addr
}