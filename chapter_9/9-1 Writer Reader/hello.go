package MarshalIndent

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// 创建一个Buffer 写入字符串
	var b bytes.Buffer

	b.Write([]byte("Hello "))

	// 使用 Fprintf 来将一个字符串拼接到 Buffer 里
	// 将 bytes.Buffer 的地址作为 io.Writer 类型值传入
	fmt.Fprintf(&b, "World!")

	// 将 Buffer 的内容输出到标准输出设备
	// 将 os.File 值的地址作为 io.Writer 类型值传入
	b.WriteTo(os.Stdout)
}
