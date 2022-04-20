package main

import (
	"fmt"
	"os/exec"

	_ "gin-ojsys.cn/models"
	"gin-ojsys.cn/router"
)

func init() {
	cmd := exec.Command("swag", "init")
	cmd.Run()
	fmt.Println("swag init success")
}

func main() {
	r := router.Router()
	r.Run(":8080")
}
