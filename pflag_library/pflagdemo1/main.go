package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

var (
	strValue  string
	intValue  int
	boolValue bool
)

func init() {
	// 使用pflag而不是flag，因为pflag提供了更丰富的功能
	pflag.StringVarP(&strValue, "string", "s", "", "a string flag")
	pflag.IntVarP(&intValue, "int", "i", 0, "an int flag")
	pflag.BoolVarP(&boolValue, "bool", "b", false, "a bool flag")

	// 将pflag的标记与Go语言的flag包整合
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	// 将Go语言的flag包的解析设置为完成，防止flag.Parse()再次调用
	flag.CommandLine.Parse(nil)
}

func main() {

	// 检查是否有任何参数传递
	if len(os.Args) == 1 {
		// 如果没有参数，显示帮助信息
		fmt.Println("Usage:")
		pflag.PrintDefaults()
		return
	}
	// 打印出命令行参数的值
	fmt.Println("String flag (-s, --string):", strValue)
	fmt.Println("Int flag (-i, --int):", intValue)
	fmt.Println("Bool flag (-b, --bool):", boolValue)

}
