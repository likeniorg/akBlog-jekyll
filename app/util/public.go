package util

import (
	"fmt"
	"os/exec"
)

// 命令快捷执行
func Command(shell string) {
	cmd := exec.Command("/bin/bash", "-c", shell)
	cmd.Run()
}

// 错误显示函数
func ErrprDisplay(err interface{}) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
