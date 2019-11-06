package monchickey

import (
    "io/ioutil"

    "gopkg.in/yaml.v2"
)

// 加载yaml配置文件
// configFile: 配置文件路径; configs: 结构体指针, 注意是传地址
func GetYamlConfig(configFile string, configs interface{}) error {
    configContent, err := ioutil.ReadFile(configFile)
    if err != nil {
        return err
    }
    err = yaml.Unmarshal(configContent, configs)
    if err != nil {
        return err
    }
    return nil
}
