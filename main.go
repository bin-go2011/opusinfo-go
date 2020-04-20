package main

import (
	"fmt"
	"os"
	"reflect"
	"unsafe"

	"golang.org/x/sys/windows"
)

const CHUNK = 4500

var loadedDLL *windows.DLL

var (
	oggSyncInitProc,
	oggSyncPageSeekProc,
	oggSyncBufferProc,
	oggSyncWroteProc,
	oggPageSerialnoProc *windows.Proc
)

func init() {
	loadedDLL = windows.MustLoadDLL("win32/ogg.dll")
	oggSyncInitProc = loadedDLL.MustFindProc("ogg_sync_init")
	oggSyncPageSeekProc = loadedDLL.MustFindProc("ogg_sync_pageseek")
	oggSyncBufferProc = loadedDLL.MustFindProc("ogg_sync_buffer")
	oggSyncWroteProc = loadedDLL.MustFindProc("ogg_sync_wrote")
	oggPageSerialnoProc = loadedDLL.MustFindProc("ogg_page_serialno")
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var ogsync = oggSyncState{}
	var page = oggPage{}
	oggSyncInitProc.Call(uintptr(unsafe.Pointer(&ogsync)))

	oggSyncPageSeekProc.Call(uintptr(unsafe.Pointer(&ogsync)), uintptr(unsafe.Pointer(&page)))
	buffer, _, _ := oggSyncBufferProc.Call(uintptr(unsafe.Pointer(&ogsync)), CHUNK)

	var bytes []byte
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	sliceHeader.Cap = int(CHUNK)
	sliceHeader.Len = int(CHUNK)
	sliceHeader.Data = buffer

	n, err := file.Read(bytes)
	if err != nil {
		panic(err)
	}
	oggSyncWroteProc.Call(uintptr(unsafe.Pointer(&ogsync)), uintptr(n))

	serial, _, _ := oggPageSerialnoProc.Call(uintptr(unsafe.Pointer(&ogsync)), uintptr(unsafe.Pointer(&page)))
	fmt.Println(serial)
}
