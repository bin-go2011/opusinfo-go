package main

import "golang.org/x/sys/windows"

var loadedDLL *windows.DLL

var (
	oggPageVersionProc *windows.Proc
)

func init() {
	loadedDLL = windows.MustLoadDLL("win32/ogg.dll")
	oggPageVersionProc = loadedDLL.MustFindProc("ogg_page_version")
}

func main() {

}
