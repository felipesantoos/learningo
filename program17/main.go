package main

var a int   // 32 or 64 bits
var b int8  // 1 byte
var c int16 // 2 bytes
var d int32 // 4 bytes
var e int64 // 8 bytes

var f uint   // 32 or 64 bits
var g uint8  // 1 byte
var h uint16 // 2 bytes
var i uint32 // 4 bytes
var j uint64 // 8 bytes

var k byte // same as uint8
var l rune // same as int32

func main() {
	g = k
	d = l
}
