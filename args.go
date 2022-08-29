package main

import (
	"fmt"
	"ktop/calcs"
	"ktop/proc"
	"ktop/state"
	"os"
	"time"
)

// Checks for certain flags and otherwise prints things
// like help if those flags are present, but otherwise
// does nothing.
func parseArgs(args []string, stt *state.State) {

	//
	// TODOO obviously, set up a minimal help and a porcelain flag
	//

	funcQuits := false
	if len(args) >= 1 { // some crappy testing scaffold
		funcQuits = true
	} else {
		return
	}

	if args[0] == "t" {
		for {
			proc.Collect(stt)
			calcs.Aggregate(stt)
			s := fmt.Sprintf(
				"%.2f%%",
				stt.Cpu.LastCPUPercent,
			)
			bytes := []byte(s)
			for _, b := range bytes {
				fmt.Print(string(b), " ")
			}

			fmt.Print("\n")

			time.Sleep(stt.Time.PollRate)
		}
	}

	if funcQuits {
		os.Exit(0)
	}
}
