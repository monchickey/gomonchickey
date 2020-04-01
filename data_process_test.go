package monchickey_test

import (
    "fmt"
    "testing"

    "github.com/zengzhiying/gomonchickey"
)

// go test -v data_process_test.go

func TestDataProcess(t *testing.T) {
    arr := []int{1, 5, 7, 2, 3, 9}
    fmt.Println(monchickey.IntArrayContain(arr, 3), monchickey.IntArrayContain(arr, 0))
    t1 := int64(1572586810)
    fmt.Println(t1, " -> ", monchickey.TimestampToString(t1, "2006-01-02 15:04:05"))
    t2 := "2019-11-01 13:43:00"
    stamp, err := monchickey.StringToTimestamp(t2, "2006-01-02 15:04:05")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(t2, " -> ", stamp)
    }
    stamp, err = monchickey.TimeZoneStringToTimestamp(t2, "2006-01-02 15:04:05", "Asia/Shanghai")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(t2, " -> ", stamp)
    }



    high := byte(189)
    low := byte(82)

    var r uint16

    r = uint16(high) << 8 + uint16(low)
    fmt.Println(r)


    fmt.Println(monchickey.ByteToFloat16(88, 30))
    fmt.Println(monchickey.ByteToFloat16(53, 85))


    rawBytes := []byte("Hello!")
    encodedStr := monchickey.Base64Encode(rawBytes)
    fmt.Println(encodedStr)

    rawBytes, err = monchickey.Base64Decode(encodedStr)
    fmt.Println(string(rawBytes), err)
    rawBytes, err = monchickey.Base64Decode("Hello")
    fmt.Println(string(rawBytes), err)



    fmt.Println("------------------")
    fmt.Println(monchickey.Uint8Transform(182, 0), monchickey.Uint8Transform(263, 0))
    fmt.Println(monchickey.Uint16Transform(28272, 0), monchickey.Uint16Transform(78901, 0))
    srcBytes := []byte{23, 16, 78, 64, 128, 203}
    // 17104e4080cb
    fmt.Println(monchickey.EncodeToHex(srcBytes))
    fmt.Println(monchickey.HexDecode("17104e4080cb"))
    fmt.Println(monchickey.HexDecode("56ef1"))

    src16Seq := []uint16{28, 2819, 2901, 182}
    // 001c0b030b5500b6
    dstBytes := monchickey.Uint16ToBytesBigEndian(src16Seq)
    fmt.Printf("%q\n", monchickey.EncodeToHex(dstBytes))
    // 1c00030b550bb600
    dstBytes2 := monchickey.Uint16ToBytesLittleEndian(src16Seq)
    fmt.Printf("%q\n", monchickey.EncodeToHex(dstBytes2))

    src1, _ := monchickey.BytesToUint16BigEndian(dstBytes)
    for _, num := range src1 {
        fmt.Printf("%d ", num)
    }
    fmt.Println("")
    fmt.Println(monchickey.BytesToUint16BigEndian(dstBytes[:5]))
    src2, _ := monchickey.BytesToUint16LittleEndian(dstBytes2)
    for _, num := range src2 {
        fmt.Printf("%d ", num)
    }
    fmt.Println()
    fmt.Println(monchickey.BytesToUint16LittleEndian(dstBytes2[:5]))


    a := uint32(1)
    b := uint32(4294967295)
    fmt.Println(monchickey.SetUint32Bit(&a, 1, 1))
    fmt.Println(monchickey.SetUint32Bit(&b, 32, 0))
    // 2147483649 4294967294
    fmt.Println(a, b)
    // 1
    fmt.Println(monchickey.GetUint32Bit(b, 5))
    // 0
    fmt.Println(monchickey.GetUint32Bit(a, 31))

    c := uint64(1)
    d := uint64(18446744073709551615)
    // 9223372036854775809
    fmt.Println(monchickey.SetUint64Bit(&c, 1, 1))
    // 9223372036854775807
    fmt.Println(monchickey.SetUint64Bit(&d, 1, 0))
    fmt.Println(c, d)
}

func TestPack(t *testing.T) {
    a := uint64(18282918212901)
    b := uint32(999996)
    c := uint16(8912)

    bs := monchickey.Uint64PackLittleEndian(a)
    fmt.Println("uint64 little:", monchickey.Uint64UnpackLittleEndian(bs))
    bs = monchickey.Uint64PackBigEndian(a)
    fmt.Println("uint64 big:", monchickey.Uint64UnpackBigEndian(bs))

    bs = monchickey.Uint32PackLittleEndian(b)
    fmt.Println("uint32 little:", monchickey.Uint32UnpackLittleEndian(bs))
    bs = monchickey.Uint32PackBigEndian(b)
    fmt.Println("uint32 big:", monchickey.Uint32UnpackBigEndian(bs))

    bs = monchickey.Uint16PackLittleEndian(c)
    fmt.Println("uint16 little:", monchickey.Uint16UnpackLittleEndian(bs))
    bs = monchickey.Uint16PackBigEndian(c)
    fmt.Println("uint16 big:", monchickey.Uint16UnpackBigEndian(bs))
}
