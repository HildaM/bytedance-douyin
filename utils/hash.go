package utils

import (
	"github.com/u2takey/go-utils/encrypt"
	"reflect"
	"unsafe"
)

func GetMd5(s string) string {
	return encrypt.Md5(zeroCopyString2Bytes(s))
}

func GetSHA256(s string) string {
	return encrypt.Sha256(zeroCopyString2Bytes(s))
}

func zeroCopyString2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))

	sliceHeader := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&sliceHeader))
}
