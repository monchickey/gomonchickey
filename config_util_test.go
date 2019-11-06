package monchickey_test

import (
    "fmt"
    "testing"

    "github.com/zengzhiying/gomonchickey"
)

// go test -v config_util_test.go

type Config struct {
    A struct {
        T1 int
        T2 string
    }
}

// go test -v config_util_test.go -test.run TestGetYamlConfig
func TestGetYamlConfig(t *testing.T) {
    config := Config{}
    err := monchickey.GetYamlConfig("resources/config.yaml", &config)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(config.A.T1, config.A.T2)
        fmt.Println(config.A.T1 == 3, config.A.T2 == "test")
    }
}
