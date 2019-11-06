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
    dstBytes := monchickey.Uint16ToBytesBigend(src16Seq)
    fmt.Printf("%q\n", monchickey.EncodeToHex(dstBytes))
    // 1c00030b550bb600
    dstBytes2 := monchickey.Uint16ToBytesSmallend(src16Seq)
    fmt.Printf("%q\n", monchickey.EncodeToHex(dstBytes2))

    src1, _ := monchickey.BytesToUint16Bigend(dstBytes)
    for _, num := range src1 {
        fmt.Printf("%d ", num)
    }
    fmt.Println("")
    fmt.Println(monchickey.BytesToUint16Bigend(dstBytes[:5]))
    src2, _ := monchickey.BytesToUint16Smallend(dstBytes2)
    for _, num := range src2 {
        fmt.Printf("%d ", num)
    }
    fmt.Println()
    fmt.Println(monchickey.BytesToUint16Smallend(dstBytes2[:5]))
}
