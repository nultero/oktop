package state

import "fmt"

type cpu_t struct {
	LCI     int // last CPU idle %, used for PollCPU
	Stamps  []float64
	Sum     int
	SumPrev int
}

func defaultCpu_t() cpu_t {
	return cpu_t{
		LCI:     0,
		Stamps:  []float64{},
		Sum:     0,
		SumPrev: 0,
	}
}

const multiDigit float64 = 10.0

// Last CPU percent.
func (c cpu_t) Last() float64 {
	return c.Stamps[len(c.Stamps)-1]
}

// Last CPU percent Sprintf'd to a string.
func (c cpu_t) LastToStr() string {
	pc := c.Stamps[len(c.Stamps)-1]
	if pc < multiDigit {
		return fmt.Sprintf(" %.2f", pc)
	}

	return fmt.Sprintf("%.2f", pc)
}
