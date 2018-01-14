// +build ignore

package main

import (
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go ng()

	go gin()

	wg.Wait()
}

func ExecWithStdout(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
}

func gin() {
	defer wg.Done()
	time.Sleep(time.Second * 10)
	fmt.Println("Starting gin to build and watch for changes in the go project")
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	fmt.Println(gopath)
	ginPath := gopath + "/bin/gin"
	gin := exec.Command(ginPath, "-a", "8080", "-i", "-x", "scripts")
	ExecWithStdout(gin)
	gin.Run()
}

func ng() {
	defer wg.Done()
	fmt.Println("Starting angular build with file watcher")

	ng := exec.Command("ng", "build", "--watch")
	ng.Dir = "gui"
	ExecWithStdout(ng)
	ng.Run()
}
