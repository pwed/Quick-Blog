package main

import (
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"sync"
)

var (
	wg   sync.WaitGroup
	args []string
)

func main() {
	mode := checkRunMode()

	paralellTask(ng, mode)

	paralellTask(gin, mode)

	wg.Wait()

	completeTask(goBin, mode)

	completeTask(goBuild, mode)
}

func paralellTask(fun func(string), m string) {
	wg.Add(1)
	go fun(m)
}

func completeTask(fun func(string), m string) {
	wg.Add(1)
	go fun(m)
	wg.Wait()
}

func checkRunMode() string {
	if len(args) == 2 {
		return args[1]
	}
	return "run"
}

func execWithStdout(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
}

func gin(m string) {
	defer wg.Done()
	if m != "run" {
		return
	}

	fmt.Println("Starting gin to build and watch for changes in the go project\n",
		"Angular build might take a while longer so routes will 404 for up to 30 seconds")

	os.Remove("bindata_assetfs.go")
	bd, _ := os.Create("bindata_assetfs.go")
	bd.WriteString(falseAssetFS)

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	ginPath := gopath + "/bin/gin"
	gin := exec.Command(ginPath, "-a", "8080", "-i", "-x", "_scripts", "dev")
	execWithStdout(gin)
	gin.Run()
}

func ng(m string) {
	defer wg.Done()

	arguments := []string{"build"}

	if m == "run" {
		arguments = append(arguments, "--watch")
		fmt.Println("Starting angular build with file watcher")
	}

	if m == "prod" {
		arguments = append(arguments, "--prod")
		fmt.Println("Starting angular build in production mode")
	}

	ng := exec.Command("ng", arguments...)
	ng.Dir = "gui"
	execWithStdout(ng)
	ng.Run()
}

func goBin(m string) {
	defer wg.Done()
	if m != "prod" {
		return
	}
	fmt.Println("Creating embedded static assets from angular build")

	goBin := exec.Command("go-bindata-assetfs", "static/...")
	execWithStdout(goBin)
	goBin.Run()
}

func goBuild(m string) {
	defer wg.Done()
	if m != "prod" {
		return
	}
	fmt.Println("Building production server with embeded assets")

	gb := exec.Command("go", "build")
	execWithStdout(gb)
	gb.Run()
}

func init() {
	args = os.Args
}

const falseAssetFS = `package main

import (
	assetfs "github.com/elazarl/go-bindata-assetfs"
)

func assetFS() *assetfs.AssetFS {
	return nil
}`
