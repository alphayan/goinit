package main

import (
	"fmt"
	"os/exec"
	"path"
)

func goremove(dir string) error {
	mycmd := path.Join(GOPATHSRC, dir)
	data, err := exec.Command("rm", "-rf", mycmd).Output()
	fmt.Println(data)
	return err
}
