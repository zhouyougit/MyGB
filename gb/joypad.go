package gb

const (
	ButtonRight byte = 0
	ButtonLeft byte = 1
	ButtonUp byte = 2
	ButtonDown byte = 3
	ButtonA byte = 4
	ButtonB byte = 5
	ButtonSelect byte = 6
	ButtonStart byte = 7
)
type Joypad struct {
	gb *GameBoy
	status byte
	selectButton bool // true: button / false: direction
}

func (j *Joypad) Init(gb *GameBoy) error {
	j.gb = gb
	j.status = 0xFF
	return nil
}

func (j *Joypad) PressButton(button byte) {
	op := byte(0x1) << button
	oldSt := j.status & op
	j.status = j.status &^ op
	if oldSt != (j.status & op) && ((j.selectButton && button > 3) || (!j.selectButton && button < 4)){
		j.gb.Cpu.RequestInterrupt(IntJoypad)
	}
}

func (j *Joypad) ReleaseButton(button byte) {
	op := byte(0x1) << button
	oldSt := j.status & op
	j.status = j.status | op
	if oldSt != j.status & op {
		//
	}
}

func (j *Joypad) readStatus() byte {
	if j.selectButton {
		return 0x10 | (j.status >> 4)
	} else {
		return 0x20 | (j.status & 0x0F)
	}
}

func (j *Joypad) writeStatus(stat byte) {
	if stat & 0x20 == 0 {
		j.selectButton = true
	} else {
		j.selectButton = false
	}
}