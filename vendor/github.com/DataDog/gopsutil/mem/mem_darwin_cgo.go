// +build darwin
// +build cgo

package mem

/*
#ifndef TEST_H_
#define TEST_H_
#include <mach/mach_host.h>
#endif
*/
import "C"

import (
	"fmt"
	"syscall"
	"unsafe"
)

// VirtualMemory returns VirtualmemoryStat.
func VirtualMemory() (*VirtualMemoryStat, error) {
	count := C.mach_msg_type_number_t(C.HOST_VM_INFO_COUNT)
	var vmstat C.vm_statistics_data_t

	status := C.host_statistics(C.host_t(C.mach_host_self()),
		C.HOST_VM_INFO,
		C.host_info_t(unsafe.Pointer(&vmstat)),
		&count)

	if status != C.KERN_SUCCESS {
		return nil, fmt.Errorf("host_statistics error=%d", status)
	}

	pageSize := uint64(syscall.Getpagesize())
	total, err := getHwMemsize()
	if err != nil {
		return nil, err
	}
	totalCount := C.natural_t(total / pageSize)

	availableCount := vmstat.inactive_count + vmstat.free_count
	usedPercent := 100 * float64(totalCount-availableCount) / float64(totalCount)

	usedCount := totalCount - availableCount

	return &VirtualMemoryStat{
		Total:       total,
		Available:   pageSize * uint64(availableCount),
		Used:        pageSize * uint64(usedCount),
		UsedPercent: usedPercent,
		Free:        pageSize * uint64(vmstat.free_count),
		Active:      pageSize * uint64(vmstat.active_count),
		Inactive:    pageSize * uint64(vmstat.inactive_count),
		Wired:       pageSize * uint64(vmstat.wire_count),
	}, nil
}
