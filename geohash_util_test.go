package monchickey_test

import (
    "fmt"
    "testing"

    "github.com/zengzhiying/gomonchickey"
)

// go test -v geohash_util_test.go

func TestGeoHash(t *testing.T) {

    // mergeValue, err := geohashEncode(116.38955, 39.9232, 6)
    // mergeValue >> 34
    // wx4g0e
    fmt.Println(monchickey.GeohashEncode(116.38955, 39.9232, 6))
    // wx4g0eb
    fmt.Println(monchickey.GeohashEncode(116.38955, 39.9232, 7))
    // wx4g0eb3
    fmt.Println(monchickey.GeohashEncode(116.38955, 39.9232, 8))
    // wx4g0eb33
    fmt.Println(monchickey.GeohashEncode(116.38955, 39.9232, 9))
    // wx4g0eb33p
    fmt.Println(monchickey.GeohashEncode(116.38955, 39.9232, 10))
    // wx4g0eb33pf
    fmt.Println(monchickey.GeohashEncode(116.389550, 39.9232, 11))
    // wx4g0eb33pfs
    fmt.Println(monchickey.GeohashEncode(116.38955, 39.9232, 12))

    encodedString, _ := monchickey.GeohashEncode(113.56291, 36.9271, 12)
    // ww8f04mgzw75
    fmt.Println(encodedString)
    // [28 28 8 14 0 4 19 15 31 28 7 5]
    // decodedBytes, _ := monchickey.base32Decode(encodedString)
    // fmt.Println(decodedBytes)

    // 113.56290997937322 36.92709996365011
    fmt.Println(monchickey.GeohashDecode(encodedString))
    // 113.55812072753906,38.33335876464844
    fmt.Println(monchickey.GeohashDecode("wwbf046"))
    // 113.55779886245728,38.33333730697632
    fmt.Println(monchickey.GeohashDecode("wwbf0467b"))
    // -112.5,67.5
    fmt.Println(monchickey.GeohashDecode("c"))
}
