package base

import "encoding/binary"

//Int64 To Byte
func Int64ToBytes(src int64) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, uint64(src))
	return bytes
}

//Byte To Int64
func BytesToInt64(bytes []byte) int64 {
	return int64(binary.LittleEndian.Uint64(bytes))
}

//Int64 To Byte
func Int32ToBytes(src int32) []byte {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, uint32(src))
	return bytes
}

//Byte To Int32
func BytesToInt32(bytes []byte) int32 {
	return int32(binary.LittleEndian.Uint32(bytes))
}
