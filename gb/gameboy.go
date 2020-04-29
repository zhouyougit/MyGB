package gb

import (
	"sync"
	"time"
)
const (
	CPU_MAIN_FREQUENCY = 4194304
)

type GameBoyArgs struct {
	Debug bool
	FPS int
	SoundOn bool
	ROMPath string
}

type GameBoy struct {
	//
	Cartridge Cartridge
	Cpu	Cpu
	Mem Memory
	Timer Timer
	Lcd Lcd

	//Video
	//Sound
	//Joypad
	Debuger *Debuger

	//Args
	debug bool
	fps int
	cyclesInFrame int
	currCycles int

	finish sync.WaitGroup
}

func NewGameBoy(args *GameBoyArgs) (*GameBoy, error) {
	gb := &GameBoy{}

	gb.debug = args.Debug
	gb.fps = args.FPS
	gb.Cartridge.romPath = args.ROMPath
	gb.Cpu.frequency = CPU_MAIN_FREQUENCY
	gb.cyclesInFrame = gb.Cpu.frequency / gb.fps
	gb.currCycles = 0

	gb.init()
	return gb, nil
}

func (gb *GameBoy) Run(stop <-chan struct{}) error {
	ticker := time.NewTicker(time.Second / time.Duration(gb.fps))
	gb.finish.Add(1)
	for {
		select {
		case <-ticker.C:
			gb.updateFrame()
		case <-stop:
			gb.finish.Done()
			println("finish")
			return nil
		}
	}
}

func (gb *GameBoy) WaitFinish() {
	gb.finish.Wait()
}

func (gb *GameBoy) init() error {
	err := gb.Cartridge.LoadRom()
	if err != nil {
		return err
	}
	gb.Mem.Init(&gb.Cartridge)
	//gb.Mem.Dump(0x0185, 0x0187)
	gb.Cpu.Init(gb)

	gb.Timer.Init(gb)

	gb.Lcd.Init(gb)

	if gb.debug {
		gb.Debuger = NewDebuger(gb)
	}
	return nil
}

func (gb *GameBoy) updateFrame() {
	for gb.currCycles < gb.cyclesInFrame {
		gb.Debuger.DebugOpcode()
		cycles := gb.Cpu.checkInterrupt()
		if cycles == 0 {
			cycles = gb.Cpu.executeNextOpcode()
		}

		gb.currCycles += cycles
		gb.Timer.update(cycles)
		gb.Lcd.update(cycles)
	}
	gb.currCycles -= gb.cyclesInFrame
}