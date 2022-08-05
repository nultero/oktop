package state

import (
	"errors"
	"fmt"
)

// Adds a new process to the proctable.
func (pm PIDMap) NewProc(
	name string, pid uint64,
	utime, stime int64,
	cpuSum, cpuLast float64,
) {

	slicePc := float64(utime+stime) / cpuSum
	cpuPc := cpuLast * slicePc

	pm[pid] = proc_t{
		cpuPc: cpuPc,
		name:  name,
		utime: [2]int64{-1, utime},
		stime: [2]int64{-1, stime},
	}
}

func (pm PIDMap) UpdateProc(
	pid uint64,
	utime, stime int64,
	cpuSum, cpuLast float64,
) error {
	if proc, ok := pm[pid]; ok {
		proc.utime[0] = proc.utime[1]
		proc.utime[1] = utime

		proc.stime[0] = proc.stime[1]
		proc.stime[1] = stime

		diffU := utime - proc.utime[0]
		diffS := stime - proc.stime[0]

		slicePc := float64(diffU+diffS) / cpuSum
		proc.cpuPc = cpuLast * slicePc

	} else {
		return errors.New(
			fmt.Sprintf(
				"pid '%v' was supposed to be alive and in memory, but was not found",
				pid,
			),
		)
	}

	return nil
}
