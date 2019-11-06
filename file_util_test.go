package monchickey_test

import (
    "fmt"
    "testing"

    "time"

    "monchickey"
)

// go test -v file_util_test.go

func TestPathStatus(t *testing.T) {
    fmt.Println(monchickey.PathIsExist("/usr/aaa"), monchickey.PathIsExist("/usr/bin"))
    fmt.Println(monchickey.PathIsFile("/etc/hosts"), monchickey.PathIsFile("/usr/local/"))
    fmt.Println(monchickey.PathIsDir("/usr/local"), monchickey.PathIsDir("/usr/bin/python"))
}

func TestGobSerialize(t *testing.T) {
    var a = 3
    fmt.Println(monchickey.GobSerialize("/root/a.gob", a))
    a += 2
    fmt.Println(monchickey.GobSerialize("/root/a.gob", a))

    type Aaa struct {
        A int
        B string
    }
    as := Aaa{A: 3, B:"hello"}
    fmt.Println(monchickey.GobSerialize("/root/aaa.gob", as))

    type Bbb struct {
        A time.Time
        B time.Duration
    }
    bb := Bbb{A: time.Now(), B: 10*time.Second}
    fmt.Println(monchickey.GobSerialize("/root/bbb.gob", bb))

    var b int
    err := monchickey.GobDeserialize("/root/a.gob", &b)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(b)
    }
    aaa := Aaa{}
    err = monchickey.GobDeserialize("/root/aaa.gob", &aaa)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(aaa)
    }

    bbb := Bbb{}
    err = monchickey.GobDeserialize("/root/bbb.gob", &bbb)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(bbb)
    }
}
