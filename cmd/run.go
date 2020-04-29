package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zhouyougit/MyGB/gb"
	"os"
	"os/signal"
	"syscall"
)

var (
	gameBoyArgs = gb.GameBoyArgs{}

	runCmd = &cobra.Command{
		Use:		"run",
		Aliases: 	[]string{"r"},
		Short:   	"run game locality",
		Long:    	"Load Game Boy game from file system and run locality\nUsage:\nmygb run romFile",
		Args:		cobra.ExactArgs(1),
		Run: 		func(cmd *cobra.Command, args []string) {
			gameBoyArgs.ROMPath = args[0]

			gameBoy, err := gb.NewGameBoy(&gameBoyArgs)
			if err != nil {
				fmt.Errorf("failed to Run GameBoy Game [%s] :%v", gameBoyArgs.ROMPath, err)
				return
			}
			//os.Exit(0)
			stop := make(chan struct{})
			go gameBoy.Run(stop)

			sigs := make(chan os.Signal, 1)
			signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
			for {
				sigIn := <- sigs
				if sigIn == syscall.SIGINT {
					gameBoy.Debuger.SetStep()
				} else if sigIn == syscall.SIGTERM {
					break
				}
			}
			close(stop)

			gameBoy.WaitFinish()
		},
	}
)

func init() {
	runCmd.PersistentFlags().BoolVarP(&gameBoyArgs.Debug, "debug", "d", true, "Use Debug Mode")
	runCmd.PersistentFlags().IntVarP(&gameBoyArgs.FPS, "fps", "f", 60, "Set FPS for GUI")
	runCmd.PersistentFlags().BoolVarP(&gameBoyArgs.SoundOn, "sound", "s", false, "Enable Sound for game")
	RootCmd.AddCommand(runCmd)
}