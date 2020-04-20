package main

type oggSyncState struct {
	data     uintptr
	storage  int32
	fill     int32
	returned int32

	unsynced    int32
	headerbytes int32
	bodybytes   int32
}

type oggPage struct {
	header     uintptr
	header_len int32
	body       uintptr
	body_len   int32
}
