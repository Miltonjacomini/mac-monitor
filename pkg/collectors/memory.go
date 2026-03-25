package collectors

/*
#include <mach/mach_host.h>
#include <mach/mach_init.h>
#include <mach/vm_statistics.h>
#include <sys/sysctl.h>
*/
import "C"
import (
	"context"
	"fmt"
	"mac-monitor/pkg/models"
	"time"
	"unsafe"
)

type MemoryCollector struct{}

func NewMemoryCollector() *MemoryCollector {
	return &MemoryCollector{}
}

func (mc *MemoryCollector) Name() string {
	return "memory"
}

func (mc *MemoryCollector) Collect(ctx context.Context) (models.MetricsPayload, error) {
	var vmStats C.vm_statistics64_data_t
	var count C.mach_msg_type_number_t = C.HOST_VM_INFO64_COUNT
	host := C.mach_host_self()

	ret := C.host_statistics64(C.host_t(host), C.HOST_VM_INFO64, C.host_info64_t(unsafe.Pointer(&vmStats)), &count)
	if ret != C.KERN_SUCCESS {
		return models.MetricsPayload{}, fmt.Errorf("failed to call host_statistics64: %d", ret)
	}

	pageSize := uint64(C.vm_kernel_page_size)
	
	used := (uint64(vmStats.active_count) + uint64(vmStats.inactive_count) + uint64(vmStats.wire_count)) * pageSize
	wired := uint64(vmStats.wire_count) * pageSize
	compressed := uint64(vmStats.compressor_page_count) * pageSize

	// Get memory pressure
	var pressure C.int
	size := unsafe.Sizeof(pressure)
	C.sysctlbyname(C.CString("kern.memo_pressure_level"), unsafe.Pointer(&pressure), (*C.size_t)(unsafe.Pointer(&size)), nil, 0)

	// Get Swap used
	var xsw_usage C.struct_xsw_usage
	xsw_size := unsafe.Sizeof(xsw_usage)
	C.sysctlbyname(C.CString("vm.swapusage"), unsafe.Pointer(&xsw_usage), (*C.size_t)(unsafe.Pointer(&xsw_size)), nil, 0)
	swapUsed := uint64(xsw_usage.xsu_used)

	metrics := models.MemoryMetrics{
		Used:       used,
		Wired:      wired,
		Compressed: compressed,
		Pressure:   float64(pressure), // Note: 1=normal, 2=warn, 4=critical in some macOS versions, or 0-100? Need to verify.
		SwapUsed:   swapUsed,
	}

	return models.MetricsPayload{
		Timestamp: time.Now(),
		Data: map[string]interface{}{
			"memory": metrics,
		},
	}, nil
}
