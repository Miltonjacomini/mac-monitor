package collectors

/*
#include <sys/param.h>
#include <sys/mount.h>
#include <stdlib.h>
*/
import "C"
import (
	"context"
	"mac-monitor/pkg/models"
	"time"
	"unsafe"
)

type DiskCollector struct{}

func NewDiskCollector() *DiskCollector {
	return &DiskCollector{}
}

func (dc *DiskCollector) Name() string {
	return "disk"
}

func (dc *DiskCollector) Collect(ctx context.Context) (models.MetricsPayload, error) {
	var volumes []models.VolumeMetrics
	
	count := C.getfsstat(nil, 0, C.MNT_NOWAIT)
	if count <= 0 {
		return models.MetricsPayload{}, nil
	}

	bufSize := C.size_t(count) * C.sizeof_struct_statfs
	buf := (*C.struct_statfs)(C.malloc(bufSize))
	defer C.free(unsafe.Pointer(buf))

	count = C.getfsstat(buf, C.int(bufSize), C.MNT_NOWAIT)
	
	statfsList := (*[1 << 30]C.struct_statfs)(unsafe.Pointer(buf))[:count:count]

	for _, s := range statfsList {
		mountPoint := C.GoString(&s.f_mntonname[0])
		fileSystem := C.GoString(&s.f_fstypename[0])
		
		total := uint64(s.f_blocks) * uint64(s.f_bsize)
		free := uint64(s.f_bfree) * uint64(s.f_bsize)
		used := total - free

		volumes = append(volumes, models.VolumeMetrics{
			MountPoint: mountPoint,
			FileSystem: fileSystem,
			Total:      total,
			Used:       used,
			Free:       free,
		})
	}

	metrics := models.DiskMetrics{
		Volumes: volumes,
		// I/O rate is harder to get via statfs, set to 0 for now
		ReadBytesRate:  0,
		WriteBytesRate: 0,
	}

	return models.MetricsPayload{
		Timestamp: time.Now(),
		Data: map[string]interface{}{
			"disk": metrics,
		},
	}, nil
}
