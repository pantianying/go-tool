package transtype

import "encoding/binary"

func Uint16ToByte(val uint16) []byte {
	var dst [2]byte
	binary.LittleEndian.PutUint16(dst[:], val)

	return dst[:]
}

func Uint32ToByte(val uint32) []byte {
	var dst [4]byte
	binary.LittleEndian.PutUint32(dst[:], val)

	return dst[:]
}

func Uint64ToByte(val uint64) []byte {
	var dst [8]byte
	binary.LittleEndian.PutUint64(dst[:], val)

	return dst[:]
}
