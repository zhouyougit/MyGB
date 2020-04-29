package gb

const (
	TIMER_DIV_ADDR  uint16 = 0xFF04
	TIMER_TIMA_ADDR uint16 = 0xFF05
	TIMER_TMA_ADDR  uint16 = 0xFF06
	TIMER_TAC_ADDR  uint16 = 0xFF07

	TIMER_TAC_START_MASK byte = 0b100
	TIMER_TAC_FREQ_MASK byte = 0b11

)

var TimerFreqMap =  map[byte]int {
	0b00 : 4096,
	0b01 : 262144,
	0b10 : 65536,
	0b11 : 16384,
}
type Timer struct {
	gb *GameBoy
	cyclesCounter int
	dividerCounter int
	dividerMax int
}

func (t *Timer) Init(gb *GameBoy) error {
	t.gb = gb
	/**
	FF04 - DIV - Divider Register (R/W)
	This register is incremented at rate of 16384Hz (~16779Hz on SGB).
	In CGB Double Speed Mode it is incremented twice as fast, ie. at 32768Hz.
	Writing any value to this register resets it to 00h.
	 */
	t.dividerMax = gb.Cpu.frequency / 16384
	return nil
}
/**
FF07 - TAC - Timer Control (R/W)
  Bit 2    - Timer Stop  (0=Stop, 1=Start)
  Bits 1-0 - Input Clock Select
             00:   4096 Hz    (~4194 Hz SGB)
             01: 262144 Hz  (~268400 Hz SGB)
             10:  65536 Hz   (~67110 Hz SGB)
             11:  16384 Hz   (~16780 Hz SGB)

 */
func (t *Timer) isTimerStarted() bool {
	return t.gb.Mem.Read(TIMER_TAC_ADDR) & TIMER_TAC_START_MASK == TIMER_TAC_START_MASK
}

func (t *Timer) getTimerFreq() int {
	return TimerFreqMap[t.gb.Mem.Read(TIMER_TAC_ADDR) & TIMER_TAC_FREQ_MASK]
}

func (t *Timer) update(cycles int) {
	t.updateDivider(cycles)
	if !t.isTimerStarted() {
		return
	}

	t.cyclesCounter += cycles
	if t.cyclesCounter >= (t.gb.Cpu.frequency / t.getTimerFreq()) {
		t.cyclesCounter -= t.gb.Cpu.frequency / t.getTimerFreq()

		tima := t.gb.Mem.Read(TIMER_TIMA_ADDR)
		if tima == 0xFF {
			t.gb.Mem.Write(TIMER_TIMA_ADDR, t.gb.Mem.Read(TIMER_TMA_ADDR))
			t.gb.Cpu.RequestInterrupt(IntTimer)
		} else {
			t.gb.Mem.Write(TIMER_TIMA_ADDR, tima + 1)
		}
	}
}

func (t *Timer) updateDivider(cycles int) {
	t.dividerCounter += cycles
	if t.dividerCounter >= t.dividerMax {
		t.dividerCounter -= t.dividerMax
		t.gb.Mem.Write(TIMER_DIV_ADDR, t.gb.Mem.Read(TIMER_DIV_ADDR) + 1)
	}
}