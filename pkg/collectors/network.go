package collectors

/*
#cgo LDFLAGS: -lproc
#include <sys/sysctl.h>
#include <sys/socket.h>
#include <net/if.h>
#include <net/if_dl.h>
#include <net/if_var.h>
#include <libproc.h>
#include <stdlib.h>
#include <arpa/inet.h>
#include <netinet/in.h>
#include <netinet/tcp.h>

static inline uint16_t my_ntohs(uint16_t n) {
    return ntohs(n);
}
*/
import "C"
import (
	"context"
	"mac-monitor/pkg/models"
	"strings"
	"time"
	"unsafe"
)

type NetworkCollector struct{}

func NewNetworkCollector() *NetworkCollector {
	return &NetworkCollector{}
}

func (nc *NetworkCollector) Name() string {
	return "network"
}

func (nc *NetworkCollector) Collect(ctx context.Context) (models.MetricsPayload, error) {
	ports := nc.getOpenPorts()
	ifaces := nc.getInterfaceStats()

	metrics := models.NetworkMetrics{
		OpenPorts:  ports,
		Interfaces: ifaces,
	}

	return models.MetricsPayload{
		Timestamp: time.Now(),
		Data: map[string]interface{}{
			"network": metrics,
		},
	}, nil
}

func (nc *NetworkCollector) getOpenPorts() []models.NetworkPort {
	var openPorts []models.NetworkPort
	
	pidsCount := C.proc_listallpids(nil, 0)
	if pidsCount <= 0 {
		return nil
	}
	
	pids := make([]C.int, pidsCount)
	pidsCount = C.proc_listallpids(unsafe.Pointer(&pids[0]), C.int(int(pidsCount)*int(unsafe.Sizeof(pids[0]))))

	for i := 0; i < int(pidsCount); i++ {
		pid := pids[i]
		if pid == 0 {
			continue
		}

		// List FDs for each PID
		fdsSize := C.proc_pidinfo(pid, C.PROC_PIDLISTFDS, 0, nil, 0)
		if fdsSize <= 0 {
			continue
		}

		numFds := int(fdsSize) / C.sizeof_struct_proc_fdinfo
		fds := make([]C.struct_proc_fdinfo, numFds)
		C.proc_pidinfo(pid, C.PROC_PIDLISTFDS, 0, unsafe.Pointer(&fds[0]), fdsSize)

		var procName [C.PROC_PIDPATHINFO_MAXSIZE]C.char
		C.proc_name(pid, unsafe.Pointer(&procName[0]), C.uint(C.PROC_PIDPATHINFO_MAXSIZE))
		name := C.GoString(&procName[0])

		for _, fd := range fds {
			if fd.proc_fdtype == C.PROX_FDTYPE_SOCKET {
				var sockInfo C.struct_socket_fdinfo
				ret := C.proc_pidfdinfo(pid, fd.proc_fd, C.PROC_PIDFDSOCKETINFO, unsafe.Pointer(&sockInfo), C.sizeof_struct_socket_fdinfo)
				if ret <= 0 {
					continue
				}

				if sockInfo.psi.soi_kind == C.SOCKINFO_TCP {
					tcpInfo := (*C.struct_tcp_sockinfo)(unsafe.Pointer(&sockInfo.psi.soi_proto[0]))
					state := tcpInfo.tcpsi_state
					
					// Only list listening or established for now
					if state == C.TSI_S_LISTEN || state == C.TSI_S_ESTABLISHED {
						var lport int
						// Handle IPv4 and IPv6
						if sockInfo.psi.soi_family == C.AF_INET {
							lport = int(C.my_ntohs(C.uint16_t((*C.struct_in_sockinfo)(unsafe.Pointer(&sockInfo.psi.soi_proto[0])).insi_lport)))
						} else if sockInfo.psi.soi_family == C.AF_INET6 {
							lport = int(C.my_ntohs(C.uint16_t((*C.struct_in_sockinfo)(unsafe.Pointer(&sockInfo.psi.soi_proto[0])).insi_lport)))
						}
						
						if lport > 0 {
							openPorts = append(openPorts, models.NetworkPort{
								Port:     lport,
								Protocol: "TCP",
								PID:      int(pid),
								Process:  name,
							})
						}
					}
				} else if sockInfo.psi.soi_kind == C.SOCKINFO_IN {
					// Handle UDP (generic INET)
					var lport int
					lport = int(C.my_ntohs(C.uint16_t((*C.struct_in_sockinfo)(unsafe.Pointer(&sockInfo.psi.soi_proto[0])).insi_lport)))
					if lport > 0 {
						openPorts = append(openPorts, models.NetworkPort{
							Port:     lport,
							Protocol: "UDP",
							PID:      int(pid),
							Process:  name,
						})
					}
				}
			}
		}
	}

	return openPorts
}

func (nc *NetworkCollector) getInterfaceStats() []models.InterfaceMetrics {
	var ifaces []models.InterfaceMetrics
	
	mib := []C.int{C.CTL_NET, C.PF_ROUTE, 0, 0, C.NET_RT_IFLIST2, 0}
	var size C.size_t
	if C.sysctl(&mib[0], 6, nil, &size, nil, 0) != 0 {
		return nil
	}

	buf := C.malloc(size)
	defer C.free(buf)

	if C.sysctl(&mib[0], 6, buf, &size, nil, 0) != 0 {
		return nil
	}

	ptr := uintptr(buf)
	end := ptr + uintptr(size)

	for ptr < end {
		ifm := (*C.struct_if_msghdr)(unsafe.Pointer(ptr))
		if ifm.ifm_type == C.RTM_IFINFO2 {
			ifm2 := (*C.struct_if_msghdr2)(unsafe.Pointer(ptr))
			
			// Get name
			sdl := (*C.struct_sockaddr_dl)(unsafe.Pointer(ptr + uintptr(ifm.ifm_msglen)))
			name := strings.ReplaceAll(C.GoString((*C.char)(unsafe.Pointer(&sdl.sdl_data[0]))), "\x00", "")
			name = name[:sdl.sdl_nlen]

			ifaces = append(ifaces, models.InterfaceMetrics{
				Name:       name,
				BytesIn:    uint64(ifm2.ifm_data.ifi_ibytes),
				BytesOut:   uint64(ifm2.ifm_data.ifi_obytes),
				PacketsIn:  uint64(ifm2.ifm_data.ifi_ipackets),
				PacketsOut: uint64(ifm2.ifm_data.ifi_opackets),
			})
		}
		ptr += uintptr(ifm.ifm_msglen)
	}

	return ifaces
}
