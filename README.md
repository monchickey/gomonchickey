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

type Coordinate monchickey.Coordinate

func PolygonContain(pointSet []Coordinate, p Coordinate) (int, error) {
    newPointSet := make([]monchickey.Coordinate, len(pointSet))
    for i, c := range pointSet {
        newPointSet[i] = monchickey.Coordinate(c)
    }
    return monchickey.PolygonContain(newPointSet, monchickey.Coordinate(p))
}

func main() {
    nowTimestamp := time.Now().Unix()
    nowTimeStr := monchickey.TimestampToString(nowTimestamp, "2006-01-02 15:04:05")
    fmt.Println(nowTimeStr)

    numSeq := []uint8{72, 101, 108, 108, 111, 32, 109, 111, 110, 99, 104, 105, 99, 107, 101, 121, 33}
    raw := monchickey.Uint8ToBytes(numSeq)
    fmt.Println(string(raw))
    fmt.Println(monchickey.Base64Encode(raw))

    geoHash, err := monchickey.GeohashEncode(113.56291, 36.9271, 12)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("(113.56291, 36.9271) geohash Encode:", geoHash)
    }
    longitude, latitude, err := monchickey.GeohashDecode(geoHash)
    if err == nil {
        fmt.Println(geoHash, "Decode:(", longitude, latitude, ")")
    }

    pointSet := []Coordinate{
        Coordinate{1, 1},
        Coordinate{1, 4},
        Coordinate{4, 4},
        Coordinate{4, 1},
    }

    fmt.Println("Polygon: (1,1)-(1,4)-(4,4)-(4,1): ")
    v, _ := PolygonContain(pointSet, Coordinate{1, 1})
    fmt.Println("  (1, 1) in", v)  // 边上
    v, _ = PolygonContain(pointSet, Coordinate{2, 2})
    fmt.Println("  (2, 2) in", v)  // 内部
    v, _ = PolygonContain(pointSet, Coordinate{5, 1})
    fmt.Println("  (5, 1) in", v)  // 外部
}


```
> output:  
2020-02-23 15:05:55
Hello monchickey!
SGVsbG8gbW9uY2hpY2tleSE=
(113.56291, 36.9271) geohash Encode: ww8f04mgzw75
ww8f04mgzw75 Decode:( 113.56290997937322 36.92709996365011 )
Polygon: (1,1)-(1,4)-(4,4)-(4,1): 
  (1, 1) in 0
  (2, 2) in 1
  (5, 1) in -1
