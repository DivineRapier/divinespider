package misc

import (
	"reflect"
	"strings"
	"unsafe"
)

// GetURLHost 获取 host 信息
func GetURLHost(url string) string {
	if strings.HasPrefix(url, "http://") {
		url = url[len("http://"):]
	} else if strings.HasPrefix(url, "https://") {
		url = url[len("https://"):]
	}

	index := strings.IndexByte(url, '/')
	if index == -1 {
		return url
	}
	return url[:index]
}

// StringToByteSlice low level function convert string to
// byte slice, result is immutable
func StringToByteSlice(s *string) []byte {
	var bytes []byte
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(s))
	bytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	bytesHeader.Data = stringHeader.Data
	bytesHeader.Len = stringHeader.Len
	bytesHeader.Cap = stringHeader.Len
	return bytes
}

// ByteSliceToString low level function convert byte slice
// to string, result is immutable
func ByteSliceToString(b []byte) string {
	s := ""
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	stringHeader.Data = sliceHeader.Data
	stringHeader.Len = sliceHeader.Len
	return s
}
