package main

import (
    "fmt"
    "time"

    "github.com/zengzhiying/gomonchickey"
)

func main() {
    nowTimestamp := time.Now().Unix()
    nowTimeStr := monchickey.TimestampToString(nowTimestamp, "2006-01-02 15:04:05")
    fmt.Println(nowTimeStr)

    numSeq := []uint8{72, 101, 108, 108, 111, 32, 109, 111, 110, 99, 104, 105, 99, 107, 101, 121, 33}
    raw := monchickey.Uint8ToBytes(numSeq)
    fmt.Println(string(raw))
    fmt.Println(monchickey.Base64Encode(raw))
}
