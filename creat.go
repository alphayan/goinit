package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func gocreat(dir string) error {
	err := os.MkdirAll(path.Join(GOPATHSRC, dir), 0777)
	if err != nil {
		return err
	}
	fmt.Println(newMain(dir), newConfig(dir), newDB(dir), newRabbitmq(dir), newRedis(dir), newRouter(dir))

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

	type Config struct{
	}
	`)
	return nil
}
func newRouter(dir string) error {
	f, err := os.Create(path.Join(GOPATHSRC, dir, "router.go"))
	if err != nil {
		return err
	}
	f.WriteString(`package main

	func initRouter(){
	}
	`)
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
func goinitfmt(dir string) {
	mycmd := "/d/gopath/src/goinit/fmt.sh" + " " + dir
	f, err := exec.Command("sh", "-c", mycmd).Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(f))
}
