//go:build windows

// Code generated by 'go generate' using "github.com/Microsoft/go-winio/tools/mkwinsyscall"; DO NOT EDIT.

package computestorage

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	return e
}

var (
	modcomputestorage = windows.NewLazySystemDLL("computestorage.dll")

	procHcsAttachLayerStorageFilter = modcomputestorage.NewProc("HcsAttachLayerStorageFilter")
	procHcsAttachOverlayFilter      = modcomputestorage.NewProc("HcsAttachOverlayFilter")
	procHcsDestroyLayer             = modcomputestorage.NewProc("HcsDestroyLayer")
	procHcsDetachLayerStorageFilter = modcomputestorage.NewProc("HcsDetachLayerStorageFilter")
	procHcsDetachOverlayFilter      = modcomputestorage.NewProc("HcsDetachOverlayFilter")
	procHcsExportLayer              = modcomputestorage.NewProc("HcsExportLayer")
	procHcsFormatWritableLayerVhd   = modcomputestorage.NewProc("HcsFormatWritableLayerVhd")
	procHcsGetLayerVhdMountPath     = modcomputestorage.NewProc("HcsGetLayerVhdMountPath")
	procHcsImportLayer              = modcomputestorage.NewProc("HcsImportLayer")
	procHcsInitializeWritableLayer  = modcomputestorage.NewProc("HcsInitializeWritableLayer")
	procHcsSetupBaseOSLayer         = modcomputestorage.NewProc("HcsSetupBaseOSLayer")
	procHcsSetupBaseOSVolume        = modcomputestorage.NewProc("HcsSetupBaseOSVolume")
)

func hcsAttachLayerStorageFilter(layerPath string, layerData string) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(layerPath)
	if hr != nil {
		return
	}
	var _p1 *uint16
	_p1, hr = syscall.UTF16PtrFromString(layerData)
	if hr != nil {
		return
	}
	return _hcsAttachLayerStorageFilter(_p0, _p1)
}

func _hcsAttachLayerStorageFilter(layerPath *uint16, layerData *uint16) (hr error) {
	hr = procHcsAttachLayerStorageFilter.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsAttachLayerStorageFilter.Addr(), uintptr(unsafe.Pointer(layerPath)), uintptr(unsafe.Pointer(layerData)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsAttachOverlayFilter(volumePath string, layerData string) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(volumePath)
	if hr != nil {
		return
	}
	var _p1 *uint16
	_p1, hr = syscall.UTF16PtrFromString(layerData)
	if hr != nil {
		return
	}
	return _hcsAttachOverlayFilter(_p0, _p1)
}

func _hcsAttachOverlayFilter(volumePath *uint16, layerData *uint16) (hr error) {
	hr = procHcsAttachOverlayFilter.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsAttachOverlayFilter.Addr(), uintptr(unsafe.Pointer(volumePath)), uintptr(unsafe.Pointer(layerData)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsDestroyLayer(layerPath string) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(layerPath)
	if hr != nil {
		return
	}
	return _hcsDestroyLayer(_p0)
}

func _hcsDestroyLayer(layerPath *uint16) (hr error) {
	hr = procHcsDestroyLayer.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsDestroyLayer.Addr(), uintptr(unsafe.Pointer(layerPath)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsDetachLayerStorageFilter(layerPath string) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(layerPath)
	if hr != nil {
		return
	}
	return _hcsDetachLayerStorageFilter(_p0)
}

func _hcsDetachLayerStorageFilter(layerPath *uint16) (hr error) {
	hr = procHcsDetachLayerStorageFilter.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsDetachLayerStorageFilter.Addr(), uintptr(unsafe.Pointer(layerPath)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsDetachOverlayFilter(volumePath string, layerData string) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(volumePath)
	if hr != nil {
		return
	}
	var _p1 *uint16
	_p1, hr = syscall.UTF16PtrFromString(layerData)
	if hr != nil {
		return
	}
	return _hcsDetachOverlayFilter(_p0, _p1)
}

func _hcsDetachOverlayFilter(volumePath *uint16, layerData *uint16) (hr error) {
	hr = procHcsDetachOverlayFilter.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsDetachOverlayFilter.Addr(), uintptr(unsafe.Pointer(volumePath)), uintptr(unsafe.Pointer(layerData)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsExportLayer(layerPath string, exportFolderPath string, layerData string, options string) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(layerPath)
	if hr != nil {
		return
	}
	var _p1 *uint16
	_p1, hr = syscall.UTF16PtrFromString(exportFolderPath)
	if hr != nil {
		return
	}
	var _p2 *uint16
	_p2, hr = syscall.UTF16PtrFromString(layerData)
	if hr != nil {
		return
	}
	var _p3 *uint16
	_p3, hr = syscall.UTF16PtrFromString(options)
	if hr != nil {
		return
	}
	return _hcsExportLayer(_p0, _p1, _p2, _p3)
}

func _hcsExportLayer(layerPath *uint16, exportFolderPath *uint16, layerData *uint16, options *uint16) (hr error) {
	hr = procHcsExportLayer.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsExportLayer.Addr(), uintptr(unsafe.Pointer(layerPath)), uintptr(unsafe.Pointer(exportFolderPath)), uintptr(unsafe.Pointer(layerData)), uintptr(unsafe.Pointer(options)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsFormatWritableLayerVhd(handle windows.Handle) (hr error) {
	hr = procHcsFormatWritableLayerVhd.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsFormatWritableLayerVhd.Addr(), uintptr(handle))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsGetLayerVhdMountPath(vhdHandle windows.Handle, mountPath **uint16) (hr error) {
	hr = procHcsGetLayerVhdMountPath.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsGetLayerVhdMountPath.Addr(), uintptr(vhdHandle), uintptr(unsafe.Pointer(mountPath)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsImportLayer(layerPath string, sourceFolderPath string, layerData string) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(layerPath)
	if hr != nil {
		return
	}
	var _p1 *uint16
	_p1, hr = syscall.UTF16PtrFromString(sourceFolderPath)
	if hr != nil {
		return
	}
	var _p2 *uint16
	_p2, hr = syscall.UTF16PtrFromString(layerData)
	if hr != nil {
		return
	}
	return _hcsImportLayer(_p0, _p1, _p2)
}

func _hcsImportLayer(layerPath *uint16, sourceFolderPath *uint16, layerData *uint16) (hr error) {
	hr = procHcsImportLayer.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsImportLayer.Addr(), uintptr(unsafe.Pointer(layerPath)), uintptr(unsafe.Pointer(sourceFolderPath)), uintptr(unsafe.Pointer(layerData)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsInitializeWritableLayer(writableLayerPath string, layerData string, options string) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(writableLayerPath)
	if hr != nil {
		return
	}
	var _p1 *uint16
	_p1, hr = syscall.UTF16PtrFromString(layerData)
	if hr != nil {
		return
	}
	var _p2 *uint16
	_p2, hr = syscall.UTF16PtrFromString(options)
	if hr != nil {
		return
	}
	return _hcsInitializeWritableLayer(_p0, _p1, _p2)
}

func _hcsInitializeWritableLayer(writableLayerPath *uint16, layerData *uint16, options *uint16) (hr error) {
	hr = procHcsInitializeWritableLayer.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsInitializeWritableLayer.Addr(), uintptr(unsafe.Pointer(writableLayerPath)), uintptr(unsafe.Pointer(layerData)), uintptr(unsafe.Pointer(options)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsSetupBaseOSLayer(layerPath string, handle windows.Handle, options string) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(layerPath)
	if hr != nil {
		return
	}
	var _p1 *uint16
	_p1, hr = syscall.UTF16PtrFromString(options)
	if hr != nil {
		return
	}
	return _hcsSetupBaseOSLayer(_p0, handle, _p1)
}

func _hcsSetupBaseOSLayer(layerPath *uint16, handle windows.Handle, options *uint16) (hr error) {
	hr = procHcsSetupBaseOSLayer.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsSetupBaseOSLayer.Addr(), uintptr(unsafe.Pointer(layerPath)), uintptr(handle), uintptr(unsafe.Pointer(options)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsSetupBaseOSVolume(layerPath string, volumePath string, options string) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(layerPath)
	if hr != nil {
		return
	}
	var _p1 *uint16
	_p1, hr = syscall.UTF16PtrFromString(volumePath)
	if hr != nil {
		return
	}
	var _p2 *uint16
	_p2, hr = syscall.UTF16PtrFromString(options)
	if hr != nil {
		return
	}
	return _hcsSetupBaseOSVolume(_p0, _p1, _p2)
}

func _hcsSetupBaseOSVolume(layerPath *uint16, volumePath *uint16, options *uint16) (hr error) {
	hr = procHcsSetupBaseOSVolume.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsSetupBaseOSVolume.Addr(), uintptr(unsafe.Pointer(layerPath)), uintptr(unsafe.Pointer(volumePath)), uintptr(unsafe.Pointer(options)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}
