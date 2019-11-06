package monchickey_test

import (
    "fmt"
    "testing"

    "github.com/zengzhiying/gomonchickey"
)

// go test -v hash_util_test.go

func TestHashUtil(t *testing.T) {
    message := "This is message string."
    hash64Sum := monchickey.XXHashSum64([]byte(message))
    md5Bytes := monchickey.MD5Digest([]byte(message))
    md5Hex := monchickey.MD5HexDigest([]byte(message))
    fmt.Printf("%d, %q, %s\n", hash64Sum, md5Bytes, md5Hex)
    fmt.Println(fmt.Sprintf("%x", md5Bytes) == md5Hex)
}
