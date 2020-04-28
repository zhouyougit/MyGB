package gb

import (
	"os"
)

type Cpu struct {
	frequency int
	reg Registers
	// IME - Interrupt Master Enable Flag(Write Only)
	ime bool

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
	return cpu.ime
}

func (cpu *Cpu) setFlagIME(val bool) {
	cpu.ime = val
}

func (cpu *Cpu) resetFlagIME() {
	cpu.ime = true
}

func (cpu *Cpu) executeNextOpcode() int {
	cpu.gb.Debuger.DebugStep()
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