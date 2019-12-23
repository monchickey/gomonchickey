# gomonchickey - go语言实现的工具库
### go version: 1.13

### example code

```go
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

    geoHash, err := monchickey.GeohashEncode(113.56291, 36.9271, 12)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("(113.56291, 36.9271) geohash Encode:", geoHash)
    }
}

```
> output:  
2019-11-06 21:13:17  
Hello monchickey!
(113.56291, 36.9271) geohash Encode: ww8f04mgzw75
