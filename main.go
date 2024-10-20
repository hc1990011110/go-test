package main

import (
	"fmt"

	"github.com/beego/beego/v2/core/logs"
)

func main() {
	Fuckyou()
}

func Fuckyou() {
	fmt.Println("Hello, World!")
	logs.Info("阿西吧")
	logs.Info("你会赖")
}
