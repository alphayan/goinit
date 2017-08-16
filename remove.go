package main

import (
	"os/exec"
	"path"
)

func goremove(dir string) {
	mycmd := path.Join(GOPATHSRC, dir)
	exec.Command("rm", "-rf", mycmd).Output()
}
