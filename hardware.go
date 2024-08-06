package hardware

import (
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func GetSystemSection() (string, error) {
	runTimeOS := runtime.GOOS

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}
	hstat, err := host.Info()
	if err != nil {
		return "", err
	}
	output := fmt.Sprintf("Hostname: %s\nTotal Memory: %d\nUsed Memory: %d\nOS: %s",
		hstat.Hostname, vmStat.Total, vmStat.Used, runTimeOS)
	return output, nil
}

func GetCpuSection() (string, error) {
	cpuStat, err := cpu.Info()
	if err != nil {
		return "", err
	}
	output := fmt.Sprintf("CPU: %s\nCores: %d", cpuStat[0].ModelName, len(cpuStat))
	return output, nil
}

func GetDiskSection() (string, error) {
	diskStat, err := disk.Usage("/")
	if err != nil {
		return "", err
	}
	output := fmt.Sprintf("Total Disk Space: %d GB\nFree Disk Space: %d GB",
		diskStat.Total/1024/1024/1024, diskStat.Free/1024/1024/1024)
	return output, nil
}
