package console

import (
	"fmt"
	"github.com/mgutz/ansi"
	"os"
)

// 打印成功信息 绿色
func Success(msg string) {
	colorOut(msg, "green")
}

// 打印错误信息 红色
func Error(msg string) {
	colorOut(msg, "red")
}

// 答应提示消息 黄色
func Warning(msg string) {
	colorOut(msg, "yellow")
}

// 打印一条错误消息 并退出
func Exit(msg string) {
	Error(msg)
	os.Exit(1)
}

// 语法糖 自带err!=nil判断
func ExitIf(err error) {
	if err != nil {
		Exit(err.Error())
	}
}

// colorOut 内部使用，设置高亮颜色
func colorOut(msg, color string) {
	fmt.Fprintln(os.Stdout, ansi.Color(msg, color))
}
