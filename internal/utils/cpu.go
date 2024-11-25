package utils

import (
	"fmt"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

func UseCPUMinute() float64 {
	in := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(in)
	if err != nil {
		return 0
	}

	// Load averages are stored as fixed-point values: divide by 65536 to get the actual load
	return float64(in.Loads[0]) / 65536.0
}

func UseCPU(interval time.Duration) float64 {
	percentages, err := cpu.Percent(interval, false)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return 0
	}

	return percentages[0]
}
