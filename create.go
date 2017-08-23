package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
)

var echoRouter = `package main
	func initRouter(){
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
	}
`

func gocreate(dir, frame string) error {
	err := os.MkdirAll(path.Join(GOPATHSRC, dir), 0777)
	if err != nil {
		return err
	}
	fmt.Println(newMain(dir), newConfig(dir), newDB(dir), newRabbitmq(dir), newRedis(dir), newRouter(dir, frame), newGitgnore(dir))
	return nil
}

func newMain(dir string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, "main.go"))
	if err != nil {
		return err
	}
	f.WriteString(`package main

	func main(){

	}
	`)
	return nil
}
func newConfig(dir string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, "config.go"))
	if err != nil {
		return err
	}
	f.WriteString(`package main
    // Config struct
	type Config struct{
	}
	`)
	return nil
}
func newRouter(dir, frame string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, "router.go"))
	if err != nil {
		return err
	}
	switch frame {
	case "echo":
		f.WriteString(echoRouter)
	default:
		f.WriteString(`package main

	func initRouter(){

	}
	`)
	}

	return nil
}
func newDB(dir string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, "db.go"))
	if err != nil {
		return err
	}
	f.WriteString(`package main

	func initDB(){
	}
	`)
	return nil
}
func newRedis(dir string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, "redis.go"))
	if err != nil {
		return err
	}
	f.WriteString(`package main

	func initRedis(){
	}
	`)
	return nil
}
func newRabbitmq(dir string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, "rabbitmq.go"))
	if err != nil {
		return err
	}
	f.WriteString(`package main

	func initRabbitmq(){
	}
	`)
	return nil
}

func newGitgnore(dir string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, ".gitignore"))
	if err != nil {
		return err
	}
	f.WriteString(`
	# Binaries for programs and plugins
	*.exe
	*.dll
	*.so
	*.dylib

	# Test binary, build with 'go test -c'
	*.test

	# Output of the go coverage tool, specifically when used with LiteIDE
	*.out

	# Project-local glide cache, RE: https://github.com/Masterminds/glide/issues/736
	.glide/

	#jetbrains
	.idea/

	#.init
	.init/

	#vendor
	.vendor/

	#gitignore
	.gitignore

	`)
	return nil
}

func goinitfmt(dir string) {
	mycmd := "/d/gopath/src/goinit/fmt.sh" + " " + dir
	f := exec.Command("sh", "-c", mycmd)
	stdout, err := f.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Start()
	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Print(line)
	}
	f.Wait()

}
