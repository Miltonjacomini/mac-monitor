package collectors

/*
#include <mach/mach_host.h>
#include <mach/mach_init.h>
#include <mach/mach.h>
#include <mach/processor_info.h>
#include <sys/sysctl.h>
#include <stdlib.h>

static inline kern_return_t deallocate_processor_info(processor_info_array_t info, mach_msg_type_number_t count) {
    return vm_deallocate(mach_task_self(), (vm_address_t)info, count * sizeof(int));
}
*/
import "C"
import (
	"context"
	"fmt"
	"mac-monitor/pkg/models"
	"sync"
	"time"
	"unsafe"
)

type CPUCollector struct {
	mu            sync.Mutex
	prevTicks     []C.processor_cpu_load_info_data_t
	prevTimestamp time.Time
}

func NewCPUCollector() *CPUCollector {
	return &CPUCollector{}
}

func (cc *CPUCollector) Name() string {
	return "cpu"
}

func (cc *CPUCollector) Collect(ctx context.Context) (models.MetricsPayload, error) {
	var cpuCount C.natural_t
	var processorInfo C.processor_info_array_t
	var processorMsgCount C.mach_msg_type_number_t
	host := C.mach_host_self()

	ret := C.host_processor_info(C.host_t(host), C.PROCESSOR_CPU_LOAD_INFO, &cpuCount, &processorInfo, &processorMsgCount)
	if ret != C.KERN_SUCCESS {
		return models.MetricsPayload{}, fmt.Errorf("failed to call host_processor_info: %d", ret)
	}
	defer C.deallocate_processor_info(processorInfo, processorMsgCount)

	currentTicks := (*[1 << 30]C.processor_cpu_load_info_data_t)(unsafe.Pointer(processorInfo))[:cpuCount:cpuCount]
	now := time.Now()

	cc.mu.Lock()
	defer cc.mu.Unlock()

	var perCore []float64
	var totalUsage float64

	if len(cc.prevTicks) == int(cpuCount) {
		for i := 0; i < int(cpuCount); i++ {
			user := float64(currentTicks[i].cpu_ticks[C.CPU_STATE_USER] - cc.prevTicks[i].cpu_ticks[C.CPU_STATE_USER])
			system := float64(currentTicks[i].cpu_ticks[C.CPU_STATE_SYSTEM] - cc.prevTicks[i].cpu_ticks[C.CPU_STATE_SYSTEM])
			idle := float64(currentTicks[i].cpu_ticks[C.CPU_STATE_IDLE] - cc.prevTicks[i].cpu_ticks[C.CPU_STATE_IDLE])
			nice := float64(currentTicks[i].cpu_ticks[C.CPU_STATE_NICE] - cc.prevTicks[i].cpu_ticks[C.CPU_STATE_NICE])

			total := user + system + idle + nice
			if total > 0 {
				usage := (user + system + nice) / total * 100
				perCore = append(perCore, usage)
				totalUsage += usage
			} else {
				perCore = append(perCore, 0)
			}
		}
		totalUsage /= float64(cpuCount)
	} else {
		// First run, just return 0s
		for i := 0; i < int(cpuCount); i++ {
			perCore = append(perCore, 0)
		}
	}

	// Save ticks for next run
	cc.prevTicks = make([]C.processor_cpu_load_info_data_t, cpuCount)
	copy(cc.prevTicks, currentTicks)
	cc.prevTimestamp = now

	// Get Frequency (optional, using sysctl)
	var freq int64
	size := unsafe.Sizeof(freq)
	C.sysctlbyname(C.CString("hw.cpufrequency"), unsafe.Pointer(&freq), (*C.size_t)(unsafe.Pointer(&size)), nil, 0)

	metrics := models.CPUMetrics{
		TotalUsage: totalUsage,
		PerCore:    perCore,
		Frequency:  freq / 1000000, // Convert to MHz
		TopProcs:   []models.ProcessInfo{}, // To be implemented in a future task or subtask
	}

	return models.MetricsPayload{
		Timestamp: now,
		Data: map[string]interface{}{
			"cpu": metrics,
		},
	}, nil
}
