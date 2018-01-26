package main

import (
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"sync"
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
	cmd.Stderr = os.Stderr
}

func gin() {
	defer wg.Done()
	fmt.Println("Starting gin to build and watch for changes in the go project\n",
		"Angular build might take a while longer so routes will 404 for up to 30 seconds")
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	ginPath := gopath + "/bin/gin"
	gin := exec.Command(ginPath, "-a", "8080", "-i", "-x", "_scripts")
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
