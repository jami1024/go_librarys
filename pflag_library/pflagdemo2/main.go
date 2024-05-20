package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

func main() {
	// 创建一个新的FlagSet实例用于分组
	serverFlags := pflag.NewFlagSet("server", pflag.ExitOnError)

	// 定义属于"server"分组的变量
	var (
		serverPort int
		serverAddr string
	)

	// 为"server"分组添加标志
	serverFlags.IntVarP(&serverPort, "port", "p", 8080, "Server port")
	serverFlags.StringVarP(&serverAddr, "addr", "a", "localhost", "Server address")

	// 创建另一个FlagSet实例用于分组
	clientFlags := pflag.NewFlagSet("client", pflag.ExitOnError)

	// 定义属于"client"分组的变量
	var (
		clientServer string
		clientPath   string
	)

	// 为"client"分组添加标志
	clientFlags.StringVarP(&clientServer, "server", "s", "localhost:8080", "Server address")
	clientFlags.StringVarP(&clientPath, "path", "P", "/", "Request path")

	// 检查是否有任何参数传递
	if len(os.Args) == 1 {
		// 如果没有参数，显示帮助信息
		fmt.Println("Usage:")
		serverFlags.PrintDefaults()
		clientFlags.PrintDefaults()
		return
	}

	// 解析"server"分组的命令行参数
	if len(os.Args) > 1 && (os.Args[1] == "server" || os.Args[1] == "--server") {
		serverFlags.Parse(os.Args[2:])
		fmt.Printf("Starting server on %s:%d\n", serverAddr, serverPort)
	} else if len(os.Args) > 1 && (os.Args[1] == "client" || os.Args[1] == "--client") {
		// 解析"client"分组的命令行参数
		clientFlags.Parse(os.Args[2:])
		fmt.Printf("Making request to %s at %s\n", clientServer, clientPath)
	} else {
		// 如果参数不正确，显示帮助信息
		fmt.Println("Invalid command")
		fmt.Println("Usage:")
		serverFlags.PrintDefaults()
		clientFlags.PrintDefaults()
	}
}
