package monchickey

import (
    "time"
    "math"
    "errors"
    "encoding/base64"
    "encoding/hex"
)

// 数据结构常用操作的封装

// 判断整数切片中是否包含元素
func IntArrayContain(arr []int, v int) bool {
    for _, x := range arr {
        if x == v {
            return true
        }
    }
    return false
}

// unix时间戳转时间字符串
// timestamp: 时间戳, 单位: s
// timeFormat: 格式化方式, 例如: "2006-01-02 15:04:05"
func TimestampToString(timestamp int64, timeFormat string) string {
    unixTime := time.Unix(timestamp, 0)
    return unixTime.Format(timeFormat)
}

// 字符串格式的时间转时间戳, 
// 注意: 按照本地时区进行转换
// 返回: 转换成功error为nil, 返回时间戳, 单位为s; 转换失败error不为nil
func StringToTimestamp(timeAsStr, timeFormat string) (int64, error) {
    // UTC
    // ts, err := time.Parse(timeFormat, timeAsStr)
    loc, err := time.LoadLocation("Local")
    if err != nil {
        return 0, err
    }
    ts, err := time.ParseInLocation(timeFormat, timeAsStr, loc)
    if err != nil {
        return 0, err
    }
    return ts.Unix(), nil
}

// 按照指定的时区转换时间戳, 比如: UTC - UTC时间, Asia/Shanghai - 东八区时间
func TimeZoneStringToTimestamp(timeAsStr, timeFormat, timeZone string) (int64, error) {
    loc, err := time.LoadLocation(timeZone)
    if err != nil {
        return 0, err
    }
    ts, err := time.ParseInLocation(timeFormat, timeAsStr, loc)
    if err != nil {
        return 0, err
    }
    return ts.Unix(), nil
}


// 两个字节转float16的数值表示
// 返回值: float64 对应float16表示的浮点数
func ByteToFloat16(high, low byte) float64 {
    s := 1.0
    if high >> 7 == 1 {
        s = -1.0
    }

    exp := int(high & 0x7c) >> 2
    E := exp - 15

    man := int(high & 0x03) << 8 + int(low)
    M := 1.0
    for i := 0; i < 10; i++ {
        v := man & (0x0200 >> uint(i))
        if v != 0 {
            M += 1 / math.Pow(2, float64(i + 1))
        }
    }

    return s * M * math.Pow(2, float64(E))
}

// base64编码
func Base64Encode(data []byte) string {
    return base64.StdEncoding.EncodeToString(data)
}

// base64解码
func Base64Decode(encodedData string) ([]byte, error) {
    return base64.StdEncoding.DecodeString(encodedData)
}

// 校验int是否在uint8范围内
// 校验通过返回value, 否则返回defaultValue默认值
func Uint8Transform(value, defaultValue int) int {
    if value >= 0 && value <= 255 {
        return value
    }
    return defaultValue
}

// 校验int是否在uint16范围内
// 校验通过返回value, 否则返回defaultValue默认值
func Uint16Transform(value, defaultValue int) int {
    if value >= 0 && value <= 65535 {
        return value
    }
    return defaultValue
}

// 编码二进制为hex格式的字符串
func EncodeToHex(src []byte) string {
    return hex.EncodeToString(src)
}

// 解码hex为二进制数组
func HexDecode(s string) ([]byte, error) {
    return hex.DecodeString(s)
}

// 打包uint8序列为二进制序列
func Uint8ToBytes(seq []uint8) []byte {
    dstBytes := make([]byte, len(seq))
    for i, value := range seq {
        dstBytes[i] = value
    }
    return dstBytes
}

// 打包uint16序列为二进制序列[大端]
func Uint16ToBytesBigend(seq []uint16) []byte {
    dstBytes := make([]byte, len(seq) * 2)
    for i, value := range seq {
        high := uint8(value >> 8)
        low := uint8(value & 0xff)
        dstBytes[2 * i] = high
        dstBytes[2 * i + 1] = low
    }
    return dstBytes
}

// 打包uint16序列为二进制序列[小端]
func Uint16ToBytesSmallend(seq []uint16) []byte {
    dstBytes := make([]byte, len(seq) * 2)
    for i , value := range seq {
        high := uint8(value >> 8)
        low := uint8(value & 0xff)
        dstBytes[2 * i] = low
        dstBytes[2 * i + 1] = high
    }
    return dstBytes
}

// 解包二进制数组为uint16切片[大端]
func BytesToUint16Bigend(bs []byte) ([]uint16, error) {
    if len(bs) % 2 != 0 {
        return nil, errors.New("Byte length must be a multiple of 2!")
    }
    src := make([]uint16, len(bs) / 2)
    for i := 0; i < len(bs); i += 2 {
        src[i / 2] = (uint16(bs[i]) << 8) + uint16(bs[i + 1])
    }
    return src, nil
}

// 解包二进制数组为uint16切片[小端]
func BytesToUint16Smallend(bs []byte) ([]uint16, error) {
    if len(bs) % 2 != 0 {
        return nil, errors.New("Byte length must be a multiple of 2!")
    }
    src := make([]uint16, len(bs) / 2)
    for i := 0; i < len(bs); i += 2 {
        src[i / 2] = (uint16(bs[i + 1]) << 8) + uint16(bs[i])
    }
    return src, nil
}
